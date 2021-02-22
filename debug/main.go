package main

import (
	"fmt"

	"github.com/88250/lute"
)

func main() {
	luteEngine := lute.New() // 默认已经启用 GFM 支持以及中文语境优化
	// html := luteEngine.MarkdownStr("demo", "**Lute** - A structured markdown engine.")
	markdown, err := luteEngine.HTML2Markdown(`
		<meta charset='utf-8'><div>
		<ol style="margin:0px;" yne-block-type="list"><li style="white-space:pre-wrap;line-height:1.75;font-size:14px;text-align:left;list-style-position:inside;word-break:break-word;color:rgb(0, 0, 0);font-weight:normal;font-style:normal;text-decoration:none;background-color:rgba(0, 0, 0, 0);font-family:&quot;Microsoft YaHei&quot;, STXihei;list-style-type:decimal;">的序号是10万名），绝大部分用户可获得的9折优惠券（这个更重要）；</li><li style="white-space:pre-wrap;line-height:1.75;font-size:14px;text-align:left;list-style-position:inside;word-break:break-word;color:rgb(0, 0, 0);font-weight:normal;font-style:normal;text-decoration:none;background-color:rgba(0, 0, 0, 0);font-family:&quot;Microsoft YaHei&quot;, STXihei;list-style-type:decimal;">社交效应，俗称凑热闹：越多的人、好友参与，自己受到的影响越大，活动越火热，被激发兴趣参与活动的可能性就越大。</li></ol></div><div yne-bulb-block="paragraph" style="white-space:pre-wrap;line-height:1.75;font-size:14px;text-align:left;"><span style="color: rgb(51, 51, 51);">商业目标和用户想要参与的目标其实是密切相关的。</span></div><div yne-bulb-block="paragraph" style="white-space:pre-wrap;line-height:1.75;font-size:14px;text-align:left;">
		<span style="color: rgb(51, 51, 51);">例如：社交效应能够让品牌得到快速传播，优惠券的派发则引发用户去使用从而提升优衣库的营收，降低的参与门槛则不仅让客群扩大，更能提升用户的品牌认知。</span>
		</div>
	
			<img
				alt="0"
				data-media-type="image"
				src="https://note.youdao.com/yws/public/resource/4debda872da1c21da335e8d16a7a548e/xmlnote/46B942211458481A99CC540DD899A8EE/15154"
				style="width: 620px; height: 188px;"
			/>
	
	<div yne-bulb-block="paragraph" style="white-space:pre-wrap;line-height:1.75;font-size:14px;text-align:left;"><span style="color: rgb(51, 51, 51);">在Luckyline线上排队形式的营销传播活动，称之为游戏。那么，它有哪些游戏元素呢？</span></div><div yne-bulb-block="paragraph" style="white-space:pre-wrap;line-height:1.75;font-size:14px;text-align:left;"><span style="color: rgb(51, 51, 51);">虚拟排队并非必须是游戏，例如：我们在银行里排队的感受，如果把这样的排队迁移到网上，也需要很长的时间拿到号，也会有索然无味之感。</span></div><div yne-bulb-block="paragraph" style="white-space:pre-wrap;line-height:1.75;font-size:14px;text-align:left;"><span style="color: rgb(51, 51, 51);">而在Luckyline线上虚拟排队，其具备了游戏的一些主要元素，它就是一种游戏，把排队变成了一项娱乐休闲的互动游戏。</span></div><div yne-bulb-block="paragraph" style="white-space:pre-wrap;line-height:1.75;font-size:14px;text-align:left;"><span style="color: rgb(51, 51, 51);">我认为，luckyline的虚拟排队，主要表现以下主要游戏元素：明确的目标、随机性（或者叫偶然性）、奖励行为、社交化、角色特征、限时性、低门槛。</span></div><div yne-bulb-block="paragraph" style="white-space:pre-wrap;line-height:1.75;font-size:14px;text-align:left;"><span style="color: rgb(51, 51, 51);">下面我们逐一进行解释：</span></div><div yne-bulb-block="paragraph" style="white-space:pre-wrap;line-height:1.75;font-size:14px;text-align:left;"><span style="font-size: 20px;">1. 明确的目标</span></div>
	`)
	if err == nil {
		fmt.Println(markdown)
	}

	html := luteEngine.MarkdownStr("demo", markdown)
	fmt.Println(html)

}
