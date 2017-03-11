package phones

import (
	"strconv"

	"github.com/henrylee2cn/pholcus/app/downloader/request"
	"github.com/henrylee2cn/pholcus/common/goquery"

	"fmt"
	. "github.com/henrylee2cn/pholcus/app/spider"
	"strings"
)

func init() {
	GanJi_Ershouche.Register()
}

var GanJi_Ershouche = &Spider{
	Name:        "SH_GANJI_ERSHOUCHE_USER_INFO",
	Description: "取二手车卖家与mobile_phone [http://sh.ganji.com/ershouche/]",
	Pausetime:   300,
	// Keyin:   KEYIN,
	// Limit:        LIMIT,
	EnableCookie: false,
	RuleTree: &RuleTree{
		Root: func(ctx *Context) {
			ctx.AddQueue(&request.Request{
				Url:  "http://sh.ganji.com/ershouche/",
				Rule: "PAGES",
				Temp: map[string]interface{}{"p": 1},
			})
		},

		Trunk: map[string]*Rule{
			"PAGES": {
				ParseFunc: func(ctx *Context) {
					var curr = ctx.GetTemp("p", 0).(int)
					c := ctx.GetDom().Find("ul.pageLink.clearfix>li>a.linkOn").Text()

					if c != strconv.Itoa(curr) {
						fmt.Printf("当前列表页不存在 %v", c)
						return
					}
					ctx.AddQueue(&request.Request{
						Url:  "http://sh.ganji.com/ershouche/o" + strconv.Itoa(curr+1),
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
							url, _ := s.Find(".infor-titbox > a").Attr("href")
							ctx.AddQueue(&request.Request{
								Url:      url,
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
					user_name = query.Find(".veh-tel-phone > .v-p2").First().Text()
					mobile_phone = query.Find(".telephone").First().Text()

					user_name = strings.Replace(user_name, "GJ.use(\"tool/webim/js/webim.js\");", "", -1)
					user_name = strings.Replace(user_name, "联系人：", "", -1)
					user_name = strings.Replace(user_name, " ", "", -1)
					user_name = strings.Replace(user_name, "\n", "", -1)

					mobile_phone = strings.Replace(mobile_phone, " ", "", -1)
					if len(mobile_phone) == 0 {
						return
					}

					fmt.Println("user_name：" + user_name + "，mobile_phone： " + mobile_phone)

					ctx.Output(map[int]interface{}{
						0: user_name,
						1: mobile_phone,
					})
				},
			},
		},
	},
}
