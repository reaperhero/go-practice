package goquery

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"net/http"
	"regexp"
	"strings"
	"testing"
)

func TestZhiHu(t *testing.T) {
	url := "https://www.zhihu.com/question/59139092/answer/3494924065"

	// 创建一个请求对象
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
	}

	// 设置请求头，模拟浏览器
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, as Is) Chrome/91.0.4472.124 Safari/537.36")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/svg+xml;q=0.9,*/*;q=0.9")
	req.Header.Set("Accept-Language", "en-US,en;q=0.8")
	// 创建一个HTTP客户端
	client := &http.Client{}
	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	all, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	fmt.Println(string(all))
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	doc.Find("li").Each(func(i int, s *goquery.Selection) {
		if s.Text() == "夏天的傍晚去西湖边散步纳凉吹风" {
			fmt.Println("找到匹配文本，链接为:", url)
		}
	})
}

func TestName(t *testing.T) {
	s := ` <li data-pid="EWJ6wEDE">夏天的傍晚去西湖边散步纳凉吹风</li>
                                                                                <li data-pid="E_fCFy9G">去良渚古城遗址公园看小鹿 </li>
                                                                                <li data-pid="QrP-Wgdg">去虎跑路吸氧，直到太阳下山才回家</li>
                                                                                <li data-pid="h2TR_JTp">去灵隐寺拜佛，那边绿化超舒服</li>
                                                                                <li data-pid="JZGPg_D_">去法喜寺求姻缘，黄墙拍照好出片</li>
                                                                                <li data-pid="zAz3IViF">去西湖边的长椅上等一场落日（看不腻）</li>
                                                                                <li data-pid="IcjYq-rj">去西湖景区的大树旁喂松鼠 ️</li>
                                                                                <li data-pid="7gcaPbJW">去雷峰塔下的博物馆探寻许仙和白娘子的爱情故事踪迹</li>
                                                                                <li data-pid="a9lxryLJ">去茅家埠寻找绿野仙踪</li>
                                                                                <li data-pid="up9H3M4o">去免门票的中国湿地博物馆拍好看的照片，去沉浸场馆看超大屏4D记录片</li>
                                                                                <li data-pid="PZuJatOE">云松书舍感受江南中式园林</li>
                                                                                <li data-pid="1xLs8pby">丝绸博物馆打卡看展</li>
                                                                                <li data-pid="WCLSLhdU">在晓书馆落地窗前看一本喜欢的书</li>
                                                                                <li data-pid="gOWv45pz">去满觉陇赏桂花雨</li>
                                                                                <li data-pid="pdwCWQQu">去高高的马家坞空中观景台看杭州全貌，顺便看一场夕阳</li>
                                                                                <li data-pid="_TqN01KD">去白塔公园坐绿皮火车马</li>
                                                                                <li data-pid="He2IGWXq">去钱塘江边看落日飞车</li>
                                                                                <li data-pid="kVZI3kr0">去香积寺吃一次素斋饭</li>
                                                                                <li data-pid="-SWGDtCv">去香积寺对面的杭州英蓝中心看一场浪漫的落日展</li>
                                                                                <li data-pid="bBUpDrWX">去码头花2块钱坐一次京杭大运河轮渡⛴️</li>
                                                                                <li data-pid="Wt-FYwJR">去杭州钱江世纪公园的江边看电影感日落字幕</li>
                                                                                <li data-pid="oAR9g2Kw">去天都城感受一下杭州小巴黎的埃菲尔铁塔 </li>
                                                                                <li data-pid="KwM_8vrF">爬上宝石山看日出，还能看到西湖全景</li>
                                                                                <li data-pid="PKCydRNX">去九溪烟村徒步</li>
                                                                                <li data-pid="meg4gXuE">去龙井茶园呼吸新鲜空气</li>
                                                                                <li data-pid="yE9ro8-q">去玉鸟集感受一下杭州阿那亚</li>
                                                                                <li data-pid="HGzM5hA2">去单创空间书店点一杯咖啡看一下午书</li>
                                                                                <li data-pid="kpQjR-Xj">去天目里逛逛小众设计师的各种店</li>
                                                                                <li data-pid="9M95SBtn">去良渚遗址公园等粮仓打卡拍照看落日</li>
                                                                                <li data-pid="fcf6Qdb4">坐在西湖边的长椅上感受四季，西湖环湖骑行</li>
                                                                                <li data-pid="oJ7vHFnQ">去057 壹码头（南星桥地铁站）打卡“杭州文和友”</li>
                                                                                <li data-pid="N-Ncc1Tx">去小河直街感受淳朴舒服的江南水乡</li>
                                                                                <li data-pid="YWi3XZ0X">去国家版本馆看一场建筑之美 </li>
                                                                                <li data-pid="KZa0D4pA">去浙江展览馆看展</li>
                                                                                <li data-pid="79t_Y1AQ">在苏堤漫步，下雨天更美</li>
                                                                                <li data-pid="rYJTrIXm">在金沙湖的沙滩上等一场夕阳</li>
                                                                                <li data-pid="J67S00x_">打卡逼格满满的中国美院象山，红墙超出片</li>
                                                                                <li data-pid="rEZnCPVk">去杭州最美大学浙大之江校区</li>
                                                                                <li data-pid="G7_XPPIj">去青芝坞浙大玉泉校区重温青春校园氛围</li>
                                                                                <li data-pid="5AbOz0lQ">在太子湾野餐</li>
                                                                                <li data-pid="KPsNjexk">在下沙青年林野餐露营骑行散步</li>
                                                                                <li data-pid="EdlOIkh1">在白蓝地文创园打卡“纪念碑谷”</li>
                                                                                <li data-pid="G9FaM_jr">在长桥溪生态公园看一次睡莲</li>
                                                                                <li data-pid="Aj7Gv5uY">去钱江世界城亚运公园感受亚运氛围</li>
                                                                                <li data-pid="YiU_9GGC">去植物园拥抱自然，和大佛像合照</li>
                                                                                <li data-pid="5dO7ATYR">去钱塘江观潮</li>
                                                                                <li data-pid="c8FyetP0">打卡杭州新地标粮仓艺术公园</li>
                                                                                <li data-pid="35EZDLlo">夜游西湖打卡范仲淹笔下的“浮光跃金”</li>
                                                                                <li data-pid="F1CZt1nE">去动漫博物馆“重返童年”</li>
                                                                                <li data-pid="Gua35Jn5">打卡高楼大厦间的小桥流水西兴古镇</li>
                                                                                <li data-pid="qqWpKf5F">去杭州花圃看夏色渐郁的江南水乡</li>
                                                                                <li data-pid="mBz2QgRl">去中国伞博物馆打卡戴望舒的雨巷</li>
                                                                                <li data-pid="CieLFSG_">打卡西湖边的绝美博物馆杭博</li>
                                                                                <li data-pid="4a70ShSg">遇上雨天就去感受江南烟雨</li>
                                                                                <li data-pid="o7fDBzLm">打卡杭州新展“时间上的家”</li>
                                                                                <li data-pid="lnCLDLrG">去看各种盛放的花</li>
                                                                                <li data-pid="zbO6rqTx">去东巢艺术公园打卡白敬亭去过的公园</li>
                                                                                <li data-pid="TBiiKIXB">去运河广场啃葱包烩</li>
                                                                                <li data-pid="EtIo3T-J">去杭州的各个网红书店看一下午书</li>
                                                                                <li data-pid="56zBnbOz">去杭州的城隍庙早点摊吃一次早餐</li>
                                                                                <li data-pid="HcekkzGQ">去下沙大学城高沙社区的淄博烧烤店排队吃烧烤</li>
                                                                                <li data-pid="opMjlClu">去十里琅珰的北高峰拜财神庙</li>
                                                                                <li data-pid="2VimggvK">去市民中心的图书馆找本书消磨时间</li>
                                                                                <li data-pid="QEjjtUh0">去图书馆睡觉，毕竟图书馆跟教室一样好睡</li>
                                                                                <li data-pid="pL2UnF4Y">去杨公堤骑自行车 ，骑行吹风烦恼全消散了</li>`
	split := strings.Split(s, "\n")
	re := regexp.MustCompile(`>(.*?)<`)
	for _, s2 := range split {
		// 使用FindStringSubmatch查找匹配项
		match := re.FindStringSubmatch(s2)
		if len(match) > 1 {
			// 输出匹配到的内容（去掉<p>和</p>标签）
			fmt.Printf("%s\n", strings.Trim(match[1], ""))
		}
	}

}
