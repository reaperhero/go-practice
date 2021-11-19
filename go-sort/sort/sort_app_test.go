package sort

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
)

type DownloadItem struct {
	AppId         int // appID
	DownloadTimes int // 下载次数
}

func (d DownloadItem) String() string {
	return fmt.Sprintf("AppId:%d,DownloadTimes:%d", d.AppId, d.DownloadTimes)
}

type DownloadCollection []*DownloadItem

func (d DownloadCollection) Len() int {
	return len(d)
}
func (d DownloadCollection) Swap(i int, j int) {
	d[i], d[j] = d[j], d[i]
}

// 根据app下载量降序排列
func (d DownloadCollection) Less(i int, j int) bool {
	return d[i].DownloadTimes > d[j].DownloadTimes
}

func Test_sort_01(t *testing.T) {
	a := make(DownloadCollection, 5)
	for i := 0; i < len(a); i++ {
		a[i] = &DownloadItem{i + 1, rand.Intn(1000)}
	}
	fmt.Println(a) // [AppId:1,DownloadTimes:81 AppId:2,DownloadTimes:887 AppId:3,DownloadTimes:847 AppId:4,DownloadTimes:59 AppId:5,DownloadTimes:81]
	sort.Sort(a)
	fmt.Println(a) //[AppId:2,DownloadTimes:887 AppId:3,DownloadTimes:847 AppId:1,DownloadTimes:81 AppId:5,DownloadTimes:81 AppId:4,DownloadTimes:59]
}
