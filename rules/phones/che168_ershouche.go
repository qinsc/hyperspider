package phones

import (
	"strconv"

	"github.com/henrylee2cn/pholcus/app/downloader/request"
	"github.com/henrylee2cn/pholcus/common/goquery"

	"fmt"
	. "github.com/henrylee2cn/pholcus/app/spider"
	//"strings"
	//"encoding/hex"
)

func init() {
	Che168_Ershouche.Register()
}

var Che168_Ershouche = &Spider{
	Name:        "SH_CHE168_OLDCAR",
	Description: "上海汽车之家二手车 [http://sh.ganji.com/ershouche/]",
	Pausetime:   300,
	// Keyin:   KEYIN,
	// Limit:        LIMIT,
	EnableCookie: false,
	RuleTree: &RuleTree{
		Root: func(ctx *Context) {
			ctx.AddQueue(&request.Request{
				Url:  "http://www.che168.com/shanghai/list/",
				Rule: "PAGES",
				Temp: map[string]interface{}{"p": 1},
			})
		},

		Trunk: map[string]*Rule{
			"PAGES": {
				ParseFunc: func(ctx *Context) {
					var curr = ctx.GetTemp("p", 0).(int)
					c := ctx.GetDom().Find("#listpagination > a.current").Text()

					if c != strconv.Itoa(curr) {
						fmt.Printf("当前列表页不存在 %v", c)
						return
					}

					href, _ := ctx.GetDom().Find("#listpagination > a.current").Attr("href")
					fmt.Println("href = " + href)

					hrefHead := "http://www.che168.com/shanghai/shanghai/a0_0msdgscncgpi1lto7csp"
					fmt.Println("hrefHead = " + hrefHead)

					hrefTail := href[len("/shanghai/shanghai/a0_0msdgscncgpi1ltocsp")+1:]
					fmt.Println("hrefTail = " + hrefTail)

					ctx.AddQueue(&request.Request{
						Url:  hrefHead + strconv.Itoa(curr+1) + hrefTail,
						Rule: "PAGE",
						Temp: map[string]interface{}{"p": curr + 1},
					})
					ctx.Parse("PAGE")
				},
			},

			"PAGE": {
				ParseFunc: func(ctx *Context) {
					ctx.GetDom().
						Find(".infor").
						Each(func(i int, s *goquery.Selection) {
							url, _ := s.Find(".viewlist_ul > li > a.carinfo").Attr("href")
							ctx.AddQueue(&request.Request{
								Url:      "http://www.che168.com/" + url,
								Rule:     "USER_PHONES",
								Priority: 1,
							})
						})
				},
			},

			"USER_PHONES": {
				ItemFields: []string{
					"user_name",
					"mobile_phone",
				},
				ParseFunc: func(ctx *Context) {
					query := ctx.GetDom()

					var user_name, mobile_phone string
					user_name = query.Find(".car-address").First().Text()
					mobile_phone = query.Find(".btn-iphone3").First().Text()
					fmt.Println(" 联系人 = " + user_name)
					fmt.Println(" 电话 = " + mobile_phone)

					//user_name = strings.Replace(user_name, "GJ.use(\"tool/webim/js/webim.js\");", "", -1)
					//user_name = strings.Replace(user_name, "联系人：", "", -1)
					//user_name = strings.Replace(user_name, " ", "", -1)
					//user_name = strings.Replace(user_name, "\n", "", -1)
					//
					//mobile_phone = strings.Replace(mobile_phone, " ", "", -1)
					//if len(mobile_phone) == 0 {
					//	return
					//}
					//
					//fmt.Println("user_name：" + user_name + "，mobile_phone： " + mobile_phone)
					//
					//ctx.Output(map[int]interface{}{
					//	0: user_name,
					//	1: mobile_phone,
					//})
				},
			},
		},
	},
}
