package page

import "math"

type Page struct {
}

// 参数：页数，每页数，总数
// 返回分页的内容 7个页数 前 1，2，3，4，5 后 的格式返回,小于5页返回具体页数
func (*Page) Paginator(page, prePage int, nums int64) map[string]interface{} {

	var beforePage int                                             //前一页地址
	var AfterPage int                                              //后一页地址
	totalPages := int(math.Ceil(float64(nums) / float64(prePage))) //page总数
	if page > totalPages {
		page = totalPages
	}
	if page <= 0 {
		page = 1
	}
	var pages []int
	switch {
	case page >= totalPages-5 && totalPages > 5: //最后5页
		start := totalPages - 5 + 1
		beforePage = page - 1
		AfterPage = int(math.Min(float64(totalPages), float64(page+1)))
		pages = make([]int, 5)
		for i, _ := range pages {
			pages[i] = start + i
		}
	case page >= 3 && totalPages > 5:
		start := page - 3 + 1
		pages = make([]int, 5)
		beforePage = page - 3
		for i, _ := range pages {
			pages[i] = start + i
		}
		beforePage = page - 1
		AfterPage = page + 1
	default:
		pages = make([]int, int(math.Min(5, float64(totalPages))))
		for i, _ := range pages {
			pages[i] = i + 1
		}
		beforePage = int(math.Max(float64(1), float64(page-1)))
		AfterPage = page + 1
	}
	paginatorMap := make(map[string]interface{})
	paginatorMap["pages"] = pages
	paginatorMap["totalPages"] = totalPages
	paginatorMap["beforePage"] = beforePage
	paginatorMap["AfterPage"] = AfterPage
	paginatorMap["currPage"] = page
	return paginatorMap
}
