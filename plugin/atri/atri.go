/*
Package atri 本文件基于 https://github.com/Kyomotoi/ATRI
为 Golang 移植版，语料、素材均来自上述项目
本项目遵守 AGPL v3 协议进行开源
*/
package atri

import (
	"encoding/base64"
	"math/rand"
	"time"

	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/message"

	ctrl "github.com/FloatTech/zbpctrl"
	"github.com/FloatTech/zbputils/control"
)

type datagetter func(string, bool) ([]byte, error)

func (dgtr datagetter) randImage(file ...string) message.MessageSegment {
	data, err := dgtr(file[rand.Intn(len(file))], true)
	if err != nil {
		return message.Text("ERROR: ", err)
	}
	return message.ImageBytes(data)
}

func (dgtr datagetter) randRecord(file ...string) message.MessageSegment {
	data, err := dgtr(file[rand.Intn(len(file))], true)
	if err != nil {
		return message.Text("ERROR: ", err)
	}
	return message.Record("base64://" + base64.StdEncoding.EncodeToString(data))
}

func randText(text ...string) message.MessageSegment {
	return message.Text(text[rand.Intn(len(text))])
}

// isAtriSleeping 凌晨0点到6点，ATRI 在睡觉，不回应任何请求
func isAtriSleeping(*zero.Ctx) bool {
	if now := time.Now().Hour(); now >= 1 && now < 6 {
		return false
	}
	return true
}

func init() { // 插件主体
	engine := control.Register("atri", &ctrl.Options[*zero.Ctx]{
		DisableOnDefault: false,
		Brief:            "atri人格文本回复",
		Help: "本插件基于 ATRI ，为 Golang 移植版\n" +
			"- ATRI醒醒\n- ATRI睡吧\n- 萝卜子\n- 喜欢 | 爱你 | 爱 | suki | daisuki | すき | 好き | 贴贴 | 老婆 | 亲一个 | mua\n" +
			"- 草你妈 | 操你妈 | 脑瘫 | 废柴 | fw | 废物 | 战斗 | 爬 | 爪巴 | sb | SB | 傻B\n- 早安 | 早哇 | 早上好 | ohayo | 哦哈哟 | お早う | 早好 | 早 | 早早早\n" +
			"- 中午好 | 午安 | 午好\n- 晚安 | oyasuminasai | おやすみなさい | 晚好 | 晚上好\n- 高性能 | 太棒了 | すごい | sugoi | 斯国一 | よかった\n" +
			"- 没事 | 没关系 | 大丈夫 | 还好 | 不要紧 | 没出大问题 | 没伤到哪\n- 好吗 | 是吗 | 行不行 | 能不能 | 可不可以\n- 啊这\n- 我好了\n- ？ | ? | ¿\n" +
			"- 离谱\n- 答应我",
		PublicDataFolder: "Atri",
		OnEnable: func(ctx *zero.Ctx) {
			ctx.SendChain(message.Text("嗯呜呜……夏生先生……？"))
		},
		OnDisable: func(ctx *zero.Ctx) {
			ctx.SendChain(message.Text("Zzz……Zzz……"))
		},
	})
	engine.UsePreHandler(isAtriSleeping)
	var dgtr datagetter = engine.GetLazyData
	engine.OnFullMatch("蛇", isAtriSleeping).SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			switch rand.Intn(2) {
			case 0:
				ctx.SendChain(randText("【蛇】在哦", "【蛇】盯上你了哦", "是想来找我玩吗~小白鼠？"))
			case 1:
				ctx.SendChain(randText("抓住你了哦~小白鼠~"))
			}
		})
	engine.OnFullMatchGroup([]string{"蛇~蛇~", "梅比乌斯"}, isAtriSleeping).SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			ctx.SendChain(randText("【蛇】在哦", "【蛇】盯上你了哦", "是想来找我玩吗~小白鼠？", "抓住你了哦~小白鼠~"))
		})
	engine.OnFullMatchGroup([]string{"喜欢", "爱你", "爱", "suki", "daisuki", "すき", "好き", "贴贴", "老婆", "亲一个", "mua"}, isAtriSleeping, zero.OnlyToMe).SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			ctx.SendChain(randText(
				"是要隔着衣服贴，还是从领口伸进去贴呀?小~白~鼠~",
				"小~白~鼠~？",
				"贴这么近，是对我有什么想法吗？小白鼠？",
				"来吧小白鼠，牵起我的手，加入这进化的路途吧~",
				"可以哟小白鼠，来和我做点有意思的事吧~",
				"看来我们都很闲呢，要去我的实验室里坐坐吗~？",
				"这是...表白吗？真是意外呢，我的小白鼠~",
				"你是喜欢我这副躯体呢？还是...（笑~",
				"想让我也喜欢你？你知道该怎么做~ ",
			))
		})
	engine.OnKeywordGroup([]string{"草你妈", "操你妈", "脑瘫", "废柴", "fw", "five", "废物", "战斗", "爬", "爪巴", "sb", "SB", "傻B"}, isAtriSleeping, zero.OnlyToMe).SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			ctx.SendChain(randText(
				"既然说出了这样的话~ 那你应该已经做好觉悟了吧，呵呵呵~",
				"做好准备哦，小白鼠~接下来，可是会很痛的~",
				"把你做成标本，怎么样~",
				"呵呵呵~ 可不要~逃走哦~！",
				"哎呀，生命可真是脆弱呢~ 你觉得呢？我的小白鼠~？",
			))
		})
	engine.OnFullMatchGroup([]string{"早安", "早哇", "早上好", "ohayo", "哦哈哟", "お早う", "早好", "早", "早早早"}).SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			now := time.Now().Hour()
			switch {
			case now < 6: // 凌晨
				ctx.SendChain(message.Reply(ctx.Event.MessageID), randText(
					"zzzz......",
					"zzzzzzzz......",
					"克莱因...来..zzz..帮人家..zzz..",
					"如果是要找梅比乌斯博士的话...博士还在休息",
					"有什么我可以帮忙的吗",
				))
			case now >= 6 && now < 12:
				ctx.SendChain(message.Reply(ctx.Event.MessageID), randText(
					"啊......早上好...克莱因(哈欠)",
					"唔...哈啊啊~~~克莱因？......不是啊~",
					"早上好......无聊的早晨呢~陪我玩玩吧，小白鼠？",
					"早上好...睡觉？博士的工作...还没有做完，我还能...工作...",
				))
			case now >= 12 && now < 18:
				ctx.SendChain(message.Reply(ctx.Event.MessageID), randText(
					"现在可不是早上好的时间哦~ ",
					"难道你昨天晚上做了什么吗？我的小白鼠~？",
					"繁衍，也是生命延续的一种形式...没有？呵呵~",
					"这个时间...小白鼠~？来陪我做点有意思的事吧~",
				))
			case now >= 18 && now < 24:
				ctx.SendChain(message.Reply(ctx.Event.MessageID), randText(
					"即使是【蛇】...这个时间也该睡觉了呢~",
					"啊，早上...哦不对，晚上好",
					"早上好？难不成，小白鼠~ 你是昼伏夜出？",
				))
			}
		})
	engine.OnFullMatchGroup([]string{"中午好", "午安", "午好"}).SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			now := time.Now().Hour()
			if now > 11 && now < 15 { // 中午
				ctx.SendChain(message.Reply(ctx.Event.MessageID), randText(
					"午安哦~ 我的小白鼠~ ",
					"午安，小白鼠，做个好梦哦~ 呵呵~",
				))
			}
		})
	engine.OnFullMatchGroup([]string{"晚安", "oyasuminasai", "おやすみなさい", "晚好", "晚上好"}).SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			now := time.Now().Hour()
			switch {
			case now < 6: // 凌晨
				ctx.SendChain(message.Reply(ctx.Event.MessageID), randText(
					"zzzz......",
					"zzzzzzzz......",
					"梅比乌斯博士已经休息了，有什么事情找我就行...",
					"不早了舰长，请注意休息...不然会影响实验结果",
				))
			case now >= 6 && now < 11:
				ctx.SendChain(message.Reply(ctx.Event.MessageID), randText(
					"晚安？是我睡过头了吗？还是小白鼠你睡过头了呢~",
					"晚上好？难不成，小白鼠~ 你是昼伏夜出吗？呵呵~",
					"【蛇】要冬眠了哦~ 呵呵~",
				))
			case now >= 11 && now < 19:
				ctx.SendChain(message.Reply(ctx.Event.MessageID), randText(
					"纠正，应该是午安……舰长",
					"这个时间...小白鼠~？来陪我做点有意思的事吧~",
				))
			case now >= 19 && now < 24:
				ctx.SendChain(message.Reply(ctx.Event.MessageID), randText(
					"晚安，我的小白鼠，做个好梦~",
					"呵呵~ 小白鼠~ 明天见~",
					"小白鼠~猜猜我会不会趁你睡着的时候………… 呵呵~这就怕了吗~",
					"克莱因还需要继续完成博士的工作，舰长请先去休息",
				))
			}
		})
	engine.OnKeywordGroup([]string{"高性能", "太棒了", "すごい", "sugoi", "斯国一", "よかった"}, isAtriSleeping, zero.OnlyToMe).SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			ctx.SendChain(randText(
				"太棒了？你是说哪里呢？我的小白鼠~",
				"觉得自己运气不错？要不要去我的实验室里试试呀~",
			))
		})
	engine.OnKeywordGroup([]string{"好吗", "是吗", "行不行", "能不能", "可不可以"}, isAtriSleeping).SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			if rand.Intn(2) == 0 {
				ctx.SendChain(randImage("YES.png", "NO.jpg"))
			}
		})
	engine.OnKeywordGroup([]string{"我好了"}, isAtriSleeping).SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			ctx.SendChain(message.Reply(ctx.Event.MessageID), randText(
				"你好了？你是说哪里呢？我的小白鼠~",
				"觉得自己运气不错？要不要去我的实验室里试试呀~",
			))
		})
	engine.OnFullMatchGroup([]string{"来点刀子", "来份刀子"}, isAtriSleeping, zero.OnlyToMe).SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			ctx.SendChain(randText(
				"琪亚娜，抬起头，继续前进吧。去吧这个不完美的故事，变成你所期望的样子。",
				"姬子老师，我答应过你的。我会把这个不完美的故事，变成我们期望的样子。",
				"那一天，你向我伸出了手。从你抓住我的那一刻起，我的命运就被你改变了。你是我生命中最重要的人，如果拯救你是一种罪，那就让我来当这个罪人。",
				"一个人，要犯下多少恶行，才能在地狱的尽头，将她带回黎明。一个人，要走多远的距离，才能在时光的尽头，追回最初的自己。",
				"“姬子温柔的注视着你，不再言语”",
				"我可是最擅长逃跑了啊……这种事情……怎么可能呢？芽衣姐……我……不想死……",
				"好吧……好吧，既然只有我才能做到了……就这么一次！最后一次！等你回来之后，记得一定要……再夸夸我啊",
				"你所选择的，名为【生存】，而人类所选择的……名为【文明】。既然这会成为他们的选择，那么……我的选择，当然也会一样",
				"羽渡尘，最后再帮我一次吧，然后我就会明白……自己能够去做些什么",
				"多说无益，律者，这里不会成为你的猎场。无论如何……我看起来才更像是反派……对吧？",
				"未来，就是我要为你创造的最后一种无限。再见了 我的【理解者】",
				"我真正想要的……我知道我没资格拥有，但这一次，我有没有【的确】保护了什么……",
				"这一次，我将自己的的生命压进枪膛……只为，拯救【一人】",
				"所以，请告诉我，在这个故事的最后……我……成为人了吗？",
			))
		})
	engine.OnKeyword("答应我", isAtriSleeping, zero.OnlyToMe).SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			ctx.SendChain(randText("和蛇立约定？小心会被吞掉哦~"))
		})
	engine.OnKeywordGroup([]string{"是时候了", "到时间了"}, isAtriSleeping).SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			ctx.SendChain(message.Reply(ctx.Event.MessageID), randText(
				"今天也很准时呢~ 我可爱的小白鼠哟~",
				"来吧，我可爱的小白鼠。我们的时间……还很长哦~",
			))
		})
	engine.OnKeywordGroup([]string{"启动", "起动", "启洞", "起洞"}, isAtriSleeping).SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			ctx.SendChain(message.Reply(ctx.Event.MessageID), randText(
				"这么迫不及待啊，我的小白鼠？",
				"不考虑再和我玩一会儿吗？",
				"有干这种事情的时间的话……不如来和我玩玩？我可爱的小白鼠~",
			))
		})
	engine.OnKeywordGroup([]string{"难绷", "绷"}, isAtriSleeping).SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			ctx.SendChain(message.Reply(ctx.Event.MessageID), randText(
				"哦？是遇到了什么有趣的事情吗？",
				"呵呵，我可爱的小白鼠……你这副样子，还真是可爱呢~",
			))
		})
	engine.OnKeywordGroup([]string{"信丰饶", "永生", "长生", "药王", "丰饶"}, isAtriSleeping).SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			ctx.SendChain(message.Reply(ctx.Event.MessageID), randText(
				"永生吗……呵呵~",
				"呵呵，我的小白鼠……这种事你不会相信的对吗？",
				"哎呀哎呀，我可不喜欢把这种事情挂在嘴边哦~ 你说对吗？",
				"寿命不过是思想的枷锁，冲破他，你就能看到更广阔的世界~",
			))
		})
	engine.OnKeywordGroup([]string{"粉色妖精小姐", "爱莉希雅", "爱门", "爱莉", "喇叭"}, isAtriSleeping).SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			ctx.SendChain(message.Reply(ctx.Event.MessageID), randText(
				"哎呀呀，明明有我在这里，竟然还不能让你满足吗？小白鼠~",
				"呵呵，我喜欢你【现在】的样子哦~",
				"哎呀呀，这究竟是怎么回事呢？",
			))
		})
	engine.OnKeywordGroup([]string{"千劫", "千师傅", "劫哥"}, isAtriSleeping).SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			ctx.SendChain(message.Reply(ctx.Event.MessageID), randText(
				"一定是千劫干的哦……呵呵~",
			))
		})
	engine.OnKeywordGroup([]string{"出了", "好耶", "出货了"}, isAtriSleeping).SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			ctx.SendChain(message.Reply(ctx.Event.MessageID), randText(
				"觉得自己运气不错？要不要来我的实验室试试呢？",
				"哎呀调皮的小白鼠？是想让我陪你好好玩一会儿吗？",
				"你是在渴求我的祝福吗？别担心，我会把它给你的",
			))
		})
	engine.OnFullMatchGroup([]string{"希儿"}, isAtriSleeping).SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			ctx.SendChain(message.Reply(ctx.Event.MessageID), randText(
				"希儿……是个可爱的孩子呢",
				"哦？呵呵……真是可爱呢~",
			))
		})
}
