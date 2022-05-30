package sqlx

import (
	"fmt"
	"math"
	"testing"
)

// select * from user order by id asc limit 10
// select * from user order by id asc limit 10000,10  // 第10000行记录开始取，第二个参数表示总共取10行记
// select * from user where id >=10000 order by id asc limit 10
// select * from login LIMIT [数量] OFFSET [起始位]
func TestMysqlPageOffset(t *testing.T) {
	//query := ""
}

func TestMysqlPageSize(t *testing.T) {
	param := NewPaginatorParam().DefaultPage(0).DefaultLimit(10)
	var (
		selectQuery = `select * from user where id >= 10000 order by id asc limit ? offset ?`
		countQuery  = "select count(*) from user where id >= 10000"
		Count       = 0
		result      []string
	)
	Db.Get(&Count, countQuery)
	Db.Select(&result, selectQuery, (param.Page-1)*param.Limit,param.Limit)
	fmt.Println(Count, result)
}

type paginatorParam struct {
	Page  int
	Limit int
}

func NewPaginatorParam() *paginatorParam {
	return &paginatorParam{}
}

func (p *paginatorParam) DefaultPage(page int) *paginatorParam {
	p.Page = page
	return p
}

func (p *paginatorParam) DefaultLimit(limit int) *paginatorParam {
	p.Limit = limit
	return p
}

func Paginator(page, prepage int, nums int64) map[string]interface{} {

	var firstpage int //前一页地址
	var lastpage int  //后一页地址
	//根据nums总数，和prepage每页数量 生成分页总数
	totalpages := int(math.Ceil(float64(nums) / float64(prepage))) //page总数
	if page > totalpages {
		page = totalpages
	}
	if page <= 0 {
		page = 1
	}
	var pages []int
	switch {
	case page >= totalpages-5 && totalpages > 5: //最后5页
		start := totalpages - 5 + 1
		firstpage = page - 1
		lastpage = int(math.Min(float64(totalpages), float64(page+1)))
		pages = make([]int, 5)
		for i, _ := range pages {
			pages[i] = start + i
		}
	case page >= 3 && totalpages > 5:
		start := page - 3 + 1
		pages = make([]int, 5)
		firstpage = page - 3
		for i, _ := range pages {
			pages[i] = start + i
		}
		firstpage = page - 1
		lastpage = page + 1
	default:
		pages = make([]int, int(math.Min(5, float64(totalpages))))
		for i, _ := range pages {
			pages[i] = i + 1
		}
		firstpage = int(math.Max(float64(1), float64(page-1)))
		lastpage = page + 1
	}
	paginatorMap := make(map[string]interface{})
	paginatorMap["pages"] = pages
	paginatorMap["totalpages"] = totalpages
	paginatorMap["firstpage"] = firstpage
	paginatorMap["lastpage"] = lastpage
	paginatorMap["currpage"] = page
	return paginatorMap
}
