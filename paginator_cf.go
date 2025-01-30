package cf

// 翻页器

import (
	"math"
	"regexp"
	"strings"

	"github.com/jianka/cf/cast"
)

// 生成分页数据
// totalCount 数据总数量
// page 当前页
// pageSize 每页数据数量
// url 当前Url
// param 跳转地址携带的参数
func PaginatorData(totalCount int64, page int, pageSize int, url string, param ...map[string]interface{}) string {
	// 如果总数量为0返回空字符串
	if totalCount < 1 {
		return ""
	}
	// 初始化数据
	params := "?"
	if len(param) == 1 {
		str := ""
		for k, v := range param[0] {
			str += k + "=" + cast.ToString(v) + "&"
		}
		if str != "" {
			params = "?" + SubStrLast(str)
		}
	}
	if params == "?" {
		params += "page="
	} else {
		params += "&page="
	}
	// 页面地址加参数
	urlParam := url + params
	// 当前页面
	currentPage := page
	// 总数量
	total := totalCount
	// 每个页面数量
	listRows := pageSize
	// 是否有下一页
	hasMore := false
	// 是否有上一页
	hasLess := false
	// 分页元素切片
	var htmlSlice []string
	// 最大页数
	lastPage := int(math.Ceil(float64(total) / float64(listRows)))
	// 如果总数量小于等于每页数据数量 只返回数量信息
	if total <= int64(listRows) {
		GetTotalCount(&htmlSlice, total)
	} else {
		var pageSlice []map[string]string
		// 是否有下一页
		if lastPage > currentPage {
			hasMore = true
		}
		// 是否有上一页
		if currentPage > 1 {
			hasLess = true
		}
		// 上一页
		GetPreviousButton(&htmlSlice, urlParam+cast.ToString(currentPage-1), hasLess)
		// 组装数据
		if lastPage <= 10 {
			// 页数小于10 一次性显示全
			for i := 1; i <= lastPage; i++ {
				if i == currentPage {
					pageSlice = append(pageSlice, map[string]string{
						"index": cast.ToString(i),
						"state": "active", //click active disabled
						"url":   "",
					})
				} else {
					pageSlice = append(pageSlice, map[string]string{
						"index": cast.ToString(i),
						"state": "click", //click active disabled
						"url":   urlParam + cast.ToString(i),
					})
				}
			}
		} else {
			// 页数大于10时处理
			var qtemp []interface{}
			var htemp []interface{}
			var ztemp []interface{}
			qPage := currentPage
			// 当前页面向前取三
			for i := 3; i > 0; i-- {
				if qPage-i > 0 {
					qtemp = append(qtemp, qPage-i)
				}
			}
			// 如果还未到第一页 再补... 之后再向前补两位
			if len(qtemp) > 0 && cast.ToInt(qtemp[0]) > 1 {
				qPage = cast.ToInt(qtemp[0])
				var btemp []interface{}
				// 补入 1 2
				for i := 1; i < 3; i++ {
					if i < qPage {
						btemp = append(btemp, i)
					}
				}
				// 只有向前补位为两个，且向前取三的页面大于3 才会有三个点的补位符
				if len(btemp) == 2 && qPage > 3 {
					btemp = append(btemp, "...")
				}
				// 合并向前取页切片
				btemp = append(btemp, qtemp...)
				qtemp = btemp
			}

			hPage := currentPage
			// 向后取三
			for i := 1; i < 4; i++ {
				if hPage+i <= lastPage {
					htemp = append(htemp, hPage+i)
				}
			}
			// 如果还未到最后一页 再补... 之后再向后补两位
			if len(htemp) > 0 {
				hPage = cast.ToInt(htemp[len(htemp)-1])
				// 总页数 - 2 大于最后补入的页数时才会有三个点的补位
				if lastPage-2 > hPage {
					htemp = append(htemp, "...")
				}
				for i := 1; i > -1; i-- {
					if lastPage-i > hPage {
						htemp = append(htemp, lastPage-i)
					}
				}
			}
			ztemp = append(ztemp, qtemp...)
			ztemp = append(ztemp, currentPage)
			ztemp = append(ztemp, htemp...)
			for _, v := range ztemp {
				i := cast.ToInt(v)
				if i > 0 {
					// 大于0为一个有效的点击链接
					if i == currentPage {
						pageSlice = append(pageSlice, map[string]string{
							"index": cast.ToString(i),
							"state": "active", //click active disabled
							"url":   "",
						})
					} else {
						pageSlice = append(pageSlice, map[string]string{
							"index": cast.ToString(i),
							"state": "click", //click active disabled
							"url":   urlParam + cast.ToString(i),
						})
					}
				} else {
					pageSlice = append(pageSlice, map[string]string{
						"index": cast.ToString(v),
						"state": "disabled", //click active disabled
						"url":   "",
					})
				}
			}
		}
		GetPageButton(&htmlSlice, pageSlice)
		// 下一页
		GetNextButton(&htmlSlice, urlParam+cast.ToString(currentPage+1), hasMore)
		GetTotalCount(&htmlSlice, total)
	}
	return render(htmlSlice)
}

// 生成一个可点击的按钮
// url
// page
func GetAvailablePageWrapper(url string, page string) string {
	return `<li><a href="` + HTML2str(url) + `">` + page + `</a></li>`
}

// 生成一个禁用的按钮
// text
func GetDisabledTextWrapper(text string) string {
	return `<li class="disabled"><span>` + text + `</span></li>`
}

// 生成一个激活的按钮
func GetActivePageWrapper(text string) string {
	return `<li class="active"><span>` + text + `</span></li>`
}

// 生成分页按钮数据
// s *[]string 分页数据
// p []map[string]string  index 索引 state 状态 url 链接
// state click可点击 active激活 disabled禁用
func GetPageButton(s *[]string, p []map[string]string) {
	for _, v := range p {
		switch v["state"] {
		case "click":
			*s = append(*s, GetAvailablePageWrapper(v["url"], v["index"]))
		case "active":
			*s = append(*s, GetActivePageWrapper(v["index"]))
		case "disabled":
			*s = append(*s, GetDisabledTextWrapper(v["index"]))
		}
	}
}

// 生成上一页按钮
func GetPreviousButton(s *[]string, url string, b bool) {
	if b {
		*s = append(*s, GetAvailablePageWrapper(url, "上一页"))
	} else {
		_ = 1
		//*s = append(*s, GetDisabledTextWrapper("上一页"))
	}
}

// 生成下一页按钮
func GetNextButton(s *[]string, url string, b bool) {
	if b {
		*s = append(*s, GetAvailablePageWrapper(url, "下一页"))
	} else {
		_ = 1
		// *s = append(*s, GetDisabledTextWrapper("下一页"))
	}
}

// 生成数据总数量
func GetTotalCount(s *[]string, i int64) {
	*s = append(*s, `<li class="layui-laypage-count"> 共 `+cast.ToString(i)+` 条 </li>`)
}

// 渲染分页html
func render(s []string) string {
	str := `<ul class="pagination">`
	for _, v := range s {
		str += v
	}
	str += `</ul>`
	return str
}

// HTML2str returns escaping text convert from html.
func HTML2str(html string) string {

	re := regexp.MustCompile(`\<[\S\s]+?\>`)
	html = re.ReplaceAllStringFunc(html, strings.ToLower)

	//remove STYLE
	re = regexp.MustCompile(`\<style[\S\s]+?\</style\>`)
	html = re.ReplaceAllString(html, "")

	//remove SCRIPT
	re = regexp.MustCompile(`\<script[\S\s]+?\</script\>`)
	html = re.ReplaceAllString(html, "")

	re = regexp.MustCompile(`\<[\S\s]+?\>`)
	html = re.ReplaceAllString(html, "\n")

	re = regexp.MustCompile(`\s{2,}`)
	html = re.ReplaceAllString(html, "\n")

	return strings.TrimSpace(html)
}
