package spider

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/go-kratos/kratos/v2/log"
)

func newLogger() log.Logger {
	return log.With(log.NewStdLogger(os.Stdout),
		"ts", log.DefaultTimestamp,
		"caller", log.Caller(4),
	)
}

func compareAnimeMagnets(a, b []*AnimeMagnet) *AnimeMagnet {
	for i, _ := range a {
		if a[i].MagnetLink != b[i].MagnetLink {
			return a[i]
		}

		if a[i].Name != b[i].Name {
			fmt.Printf("a=%v\n", a[i])
			fmt.Printf("b=%v\n", b[i])
			return a[i]
		}
	}
	return nil
}

func TestMioBtSpider(t *testing.T) {
	miobtSpider := MiobtSpider{
		BaseSpider: BaseSpider{
			log: log.NewHelper(newLogger()),
		},
	}
	animeMagnetsBeCompare := []*AnimeMagnet{
		{
			Name:       "[orion origin发布组] 更衣人偶坠入爱河/恋上换装娃娃 Sono Bisque Doll wa Koi wo Suru [07] [1440p] [GB] [网盘] [2022年1月番]",
			MagnetLink: "magnet:?xt=urn:btih:e83f7a6092c662faa16003769114abf0a5ef2148&tr=http://open.acgtracker.com:1096/announce",
		},
		{
			Name:       "[orion origin发布组] 更衣人偶坠入爱河/恋上换装娃娃 Sono Bisque Doll wa Koi wo Suru [07] [1080p] [GB] [网盘] [2022年1月番]",
			MagnetLink: "magnet:?xt=urn:btih:0c692371a9e9992c3c4e3d04d2555d39a34424b3&tr=http://open.acgtracker.com:1096/announce",
		},
		{
			Name:       "非凡公主希瑞.She-Ra.and.the.Princesses.of.Power.S05E12.中英字幕.WEB.720P.甜饼字幕组.mp4",
			MagnetLink: "magnet:?xt=urn:btih:b85f1de7a776504ac7063c0aab216507a1a2f5a0&tr=http://open.acgtracker.com:1096/announce",
		},
		{
			Name:       "[霜庭云花Sub&森之屋动画组][平凡職業造就世界最強 第二季 / Arifureta Shokugyou de Sekai Saikyou S2][06][1080P][AVC][繁日内嵌]",
			MagnetLink: "magnet:?xt=urn:btih:4e86ad12b659826d1248a7bd5bfb0dfe2787b280&tr=http://open.acgtracker.com:1096/announce",
		},
		{
			Name:       "[霜庭云花Sub&森之屋动画组]平凡职业成就世界最强 第二季 / Arifureta Shokugyou de Sekai Saikyou S2 [06][WebRip 1080P][简繁日内封]",
			MagnetLink: "magnet:?xt=urn:btih:96c18b8f26ff079793c6baca87c4d98379b24d5b&tr=http://open.acgtracker.com:1096/announce",
		},
		{
			Name:       "[MagicStar] 真凶标签 / 真犯人フラグ EP17 [WEBDL] [1080p] [HULU]【生】【附日字】",
			MagnetLink: "magnet:?xt=urn:btih:746238e27a5b10e789ce1ad1d9d48eba6cb5dc8a&tr=http://open.acgtracker.com:1096/announce",
		},
		{
			Name:       "[ARIA吧汉化组]Aria Special Navigation AQUARIA III(Aria The Benedizione)",
			MagnetLink: "magnet:?xt=urn:btih:d8d1801b19c169258a5bdd6a037c3a38b7507607&tr=http://open.acgtracker.com:1096/announce",
		},
		{
			Name:       "[NC-Raws] FUTSAL BOYS!! - 07 (B-Global 1920x1080 HEVC AAC MKV)",
			MagnetLink: "magnet:?xt=urn:btih:630c55583b4343cbc61332104748a8866fc40a64&tr=http://open.acgtracker.com:1096/announce",
		},
		{
			Name:       "[NC-Raws] Futsal Boys！！！！！ - 07 (Baha 1920x1080 AVC AAC MP4)",
			MagnetLink: "magnet:?xt=urn:btih:44a3a78e0d2497e6d0b2e4ad311f6426875efe86&tr=http://open.acgtracker.com:1096/announce",
		},
		{
			Name:       "[Lilith-Raws] 薔薇王的葬列 / Baraou no Souretsu - 07 [Baha][WEB-DL][1080p][AVC AAC][CHT][MP4]",
			MagnetLink: "magnet:?xt=urn:btih:a78ffdc7a571e8ccaeecf893f098f34cd1600c22&tr=http://open.acgtracker.com:1096/announce",
		},
		{
			Name:       "自由之缰.Free.Rein.S01E09.中英字幕.WEB.720P.甜饼字幕组.mp4",
			MagnetLink: "magnet:?xt=urn:btih:c5f0f2c3456183eb52b933cf93e0cd460025cc14&tr=http://open.acgtracker.com:1096/announce",
		},
		{
			Name:       "自由之缰.Free.Rein.S01E08.中英字幕.WEB.720P.甜饼字幕组.mp4",
			MagnetLink: "magnet:?xt=urn:btih:5d2a9f09737dcf0e8a91c19fd4ffcb2af7878e22&tr=http://open.acgtracker.com:1096/announce",
		},
		{
			Name:       "[c.c動漫][1月新番][Ryman's Club / 白領羽球部][04][簡繁內掛][AVC_AAC][1080P][MKV]",
			MagnetLink: "magnet:?xt=urn:btih:c7135a681323e2a1e037431e1165c9f3422c762b&tr=http://open.acgtracker.com:1096/announce",
		},
		{
			Name:       "[c.c動漫][4月番][慕留人 -火影忍者新世代-][Boruto -Naruto Next Generations-][237][BIG5][720P][MP4]",
			MagnetLink: "magnet:?xt=urn:btih:9cf2df4024bd17a09386c0a0010471f2bdd59f4f&tr=http://open.acgtracker.com:1096/announce",
		},
		{
			Name:       "[MagicStar] DCU ~持有手铐的潜水员~ / DCU ~手錠を持ったダイバー~ EP05 [WEBDL] [1080p]【生】【附日字】",
			MagnetLink: "magnet:?xt=urn:btih:8b71338f0727200da91ba75b2b0b68b962ff0d35&tr=http://open.acgtracker.com:1096/announce",
		},
		{
			Name:       "[MagicStar] 逃亡医生F / 逃亡医F EP06 [WEBDL] [1080p] [HULU]【生】【附日字】",
			MagnetLink: "magnet:?xt=urn:btih:7b824ad2a7ef3a23906fdd2b2faf6bd3f55de96c&tr=http://open.acgtracker.com:1096/announce",
		},
		{
			Name:       "[NC-Raws] 锈色铠甲 黎明 / Sabiiro no Armor: Reimei - 07 (B-Global 1920x1080 HEVC AAC MKV)",
			MagnetLink: "magnet:?xt=urn:btih:70502f5f4138b7b2a92d0b18413802c158a6c414&tr=http://open.acgtracker.com:1096/announce",
		},
		{
			Name:       "[NC-Raws] 薔薇王的葬列 / Baraou no Souretsu - 07 (B-Global 3840x2160 HEVC AAC MKV)",
			MagnetLink: "magnet:?xt=urn:btih:6cb876fa76ab853b630856ba34402765cb0ecbb1&tr=http://open.acgtracker.com:1096/announce",
		},
		{
			Name:       "[NC-Raws] 薔薇王的葬列 / Baraou no Souretsu - 07 (Baha 1920x1080 AVC AAC MP4)",
			MagnetLink: "magnet:?xt=urn:btih:f0ba2fe747e6edbafc15e07ae969228a0ce4244e&tr=http://open.acgtracker.com:1096/announce",
		},
		{
			Name:       "【幻月字幕组】【22年日剧】【如果 有一所只有帅哥的高中】【06】【1080P】【中日双语】",
			MagnetLink: "magnet:?xt=urn:btih:e5a83a99e00f3fd03565dabcc89df75fab481397&tr=http://open.acgtracker.com:1096/announce",
		},
		{
			Name:       "[OPFans楓雪動漫][ONE PIECE 海賊王][第1011話][1080p][周日版][MKV]",
			MagnetLink: "magnet:?xt=urn:btih:79166482e2c94b941dda49ef2e84abc5139471fe&tr=http://open.acgtracker.com:1096/announce",
		},
		{
			Name:       "[orion origin发布组] 明日酱的水手服/明日同学的水手服 Akebi-chan no Sailor-fuku [07] [1440p] [GB] [mkv] [网盘] [2022年1月番]",
			MagnetLink: "magnet:?xt=urn:btih:a28bd20be4cc930341698d8d63f035baecf4e0a5&tr=http://open.acgtracker.com:1096/announce",
		},
		{
			Name:       "[orion origin发布组] 明日酱的水手服/明日同学的水手服 Akebi-chan no Sailor-fuku [07] [1080p] [GB] [mkv] [网盘] [2022年1月番]",
			MagnetLink: "magnet:?xt=urn:btih:59e4041f2a47cdc611cc0ff541d3246e0c54823d&tr=http://open.acgtracker.com:1096/announce",
		},
		{
			Name:       "[Skymoon-Raws][One Piece 海賊王][1011][ViuTV][WEB-RIP][CHT][SRTx2][1080p][MKV](先行版本) V2",
			MagnetLink: "magnet:?xt=urn:btih:4c6d1f5ad211685cface420a178b9969dd7a62d8&tr=http://open.acgtracker.com:1096/announce",
		},
		{
			Name:       "[OPFans枫雪动漫][ONE PIECE 海贼王][第1011话][720p][周日版][MP4]",
			MagnetLink: "magnet:?xt=urn:btih:6d0148813b079b3bc54d65c1782bee9de4720242&tr=http://open.acgtracker.com:1096/announce",
		},
		{
			Name:       "[Skymoon-Raws][One Piece 海賊王][1011][ViuTV][WEB-DL][1080p][AVC AAC][CHT][MP4+ASS](正式版本) V2",
			MagnetLink: "magnet:?xt=urn:btih:6e7a4a22779aba647523da36bd1cdf6eb75a668e&tr=http://open.acgtracker.com:1096/announce",
		},
		{
			Name:       "[NC-Raws] 龙蛇演义 / Dragon's Disciple - 06 (B-Global Donghua 1920x1080 HEVC AAC MKV)",
			MagnetLink: "magnet:?xt=urn:btih:1b01890c4acc775f907162062c1bf49b0a8c342d&tr=http://open.acgtracker.com:1096/announce",
		},
		{
			Name:       "[神滨市立大学勇者部&LoliHouse] 魔法纪录 魔法少女小圆外传 第二季 -觉醒前夜- / Magia Record S2 [06-08][BDRip 1080p HEVC-10bit FLAC][简繁内封]",
			MagnetLink: "magnet:?xt=urn:btih:f660492b2d1970aafd1783c0f41fb8347fa16f57&tr=http://open.acgtracker.com:1096/announce",
		},
		{
			Name:       "[离谱Sub] 杀爱 / 殺し愛 / Koroshi Ai [06][HEVC AAC][1080p][简繁内封]",
			MagnetLink: "magnet:?xt=urn:btih:691f4874590565226e5bb37f7bac6bbd08a254e9&tr=http://open.acgtracker.com:1096/announce",
		},
		{
			Name:       "[离谱Sub] 相愛相殺 / 殺し愛 / Koroshi Ai [03][AVC AAC][1080p][繁體內嵌]",
			MagnetLink: "magnet:?xt=urn:btih:1dc4f1ffe83027e348339aa5e78ac8f81bc21dee&tr=http://open.acgtracker.com:1096/announce",
		},
		{
			Name:       "[离谱Sub] 杀爱 / 殺し愛 / Koroshi Ai [06][AVC AAC][1080p][简体内嵌]",
			MagnetLink: "magnet:?xt=urn:btih:b5889e97bd6dfd8015f04c3356badcb8c9860c72&tr=http://open.acgtracker.com:1096/announce",
		},
		{
			Name:       "[桜都字幕组] 终末的后宫 / Shuumatsu no Harem [07][1080P][简体内嵌]",
			MagnetLink: "magnet:?xt=urn:btih:5de921441b85392adebc2a605ab4fa7f04803e6c&tr=http://open.acgtracker.com:1096/announce",
		},
		{
			Name:       "[桜都字幕组] 终末的后宫 / Shuumatsu no Harem [07][1080P][简体内嵌]",
			MagnetLink: "magnet:?xt=urn:btih:9fae0c993cea0ff60f8773deb8d897a5109daecf&tr=http://open.acgtracker.com:1096/announce",
		},
		{
			Name:       "[桜都字幕组] 终末的后宫 / Shuumatsu no Harem [07][1080P][简繁内封]",
			MagnetLink: "magnet:?xt=urn:btih:92ef70346cbde06655e5b849e18a8a2d884172ab&tr=http://open.acgtracker.com:1096/announce",
		},
		{
			Name:       "[神滨市立大学勇者部&LoliHouse] 魔法纪录 魔法少女小圆外传 第二季 -觉醒前夜- / Magia Record S2 [01-03][BDRip 1080p HEVC-10bit FLAC][简繁内封] v2",
			MagnetLink: "magnet:?xt=urn:btih:b9a2ba27123fd0c8a394cfb8774089c8920bc9b0&tr=http://open.acgtracker.com:1096/announce",
		},
		{
			Name:       "[雪飘工作室][ワッチャプリマジ！/Whatcha _Pri-maji!/哇恰美妙魔法！][19][简](检索:/美妙旋律/美妙天堂/美妙频道) 附外挂字幕",
			MagnetLink: "magnet:?xt=urn:btih:68322434dea2b9c09c05bedcfda1992403356e8e&tr=http://open.acgtracker.com:1096/announce",
		},
		{
			Name:       "【幻月字幕组】【22年日剧】【鹿枫堂四色日和】【06】【1080P】【中日双语】",
			MagnetLink: "magnet:?xt=urn:btih:e2c1d538a83ff1b0348f0a4df5624ec34fab52b5&tr=http://open.acgtracker.com:1096/announce",
		},
		{
			Name:       "【飞沐team】大头儿子和小头爸爸 156集全(1995)[国语无字][DVD修复版](1080P)[MKV]",
			MagnetLink: "magnet:?xt=urn:btih:7babb62f3af830c6bac61ff0d465bbfe2fea0088&tr=http://open.acgtracker.com:1096/announce",
		},
		{
			Name:       "[星空字幕組][白領羽球部 / Ryman's Club][02][繁日雙語][1080P][WEBrip][MP4]（急招翻譯、校對）",
			MagnetLink: "magnet:?xt=urn:btih:51f6ca83bd8af986618e22348f781909f2c94f5b&tr=http://open.acgtracker.com:1096/announce",
		},
		{
			Name:       "[星空字幕组][白领羽球部 / Ryman's Club][02][简日双语][1080P][WEBrip][MP4]（急招翻译、校对）",
			MagnetLink: "magnet:?xt=urn:btih:8756a6497f5961a27a28aebec83a8ce952fa78d5&tr=http://open.acgtracker.com:1096/announce",
		},
		{
			Name:       "[NC-Raws] 博人传之火影忍者次世代 / Boruto - Naruto Next Generations - 237 (B-Global 1920x1080 HEVC AAC MKV)",
			MagnetLink: "magnet:?xt=urn:btih:c4f911fbaf872be72f7bec7eb38132b8ccaba6d2&tr=http://open.acgtracker.com:1096/announce",
		},
		{
			Name:       "[神滨市立大学勇者部&LoliHouse] 魔法纪录 魔法少女小圆外传 第二季 -觉醒前夜- / Magia Record S2 [04-05][BDRip 1080p HEVC-10bit FLAC][简繁内封]",
			MagnetLink: "magnet:?xt=urn:btih:b36b5ae62040a621049e3fb762b09fcfae25ce61&tr=http://open.acgtracker.com:1096/announce",
		},
		{
			Name:       "[220216][少女与战车]『ガールズ&パンツァー 劇場版』ボーカルミニアルバム「音楽道、まだまだ邁進中です!!!」ChouCho×佐咲紗花、渕上舞、田中理恵、喜多村英梨、川澄綾子、吉岡麻耶、金元寿子、瀬戸麻沙美、能登麻美子、竹達彩奈[FLAC]",
			MagnetLink: "magnet:?xt=urn:btih:a4e944be5f22c97e983990c3b73215d9f283e139&tr=http://open.acgtracker.com:1096/announce",
		},
		{
			Name:       "[220216][少女与战车]『ガールズ&パンツァー 劇場版』ボーカルミニアルバム「音楽道、まだまだ邁進中です!!!」ChouCho×佐咲紗花、渕上舞、田中理恵、喜多村英梨、川澄綾子、吉岡麻耶、金元寿子、瀬戸麻沙美、能登麻美子、竹達彩奈[320K]",
			MagnetLink: "magnet:?xt=urn:btih:87d2fccb09c49ab6abf1c9926c23d452aaa89bc3&tr=http://open.acgtracker.com:1096/announce",
		},
		{
			Name:       "[DBD-Raws][枪神斯托拉塔斯/Gunslinger Stratos The Animation/ガンスリンガー ストラトス THE ANIMATION][01-12TV全集+SP+特典映像][1080P][BDRip][HEVC-10bit][FLAC][MKV]",
			MagnetLink: "magnet:?xt=urn:btih:cca8da9e9d7f2b7f0f8696797d54de884677b069&tr=http://open.acgtracker.com:1096/announce",
		},
		{
			Name:       "[Skymoon-Raws][One Piece 海賊王][1011][ViuTV][WEB-DL][1080p][AVC AAC][CHT][MP4+ASS](正式版本)",
			MagnetLink: "magnet:?xt=urn:btih:828c0da93a199526e35a0a8d518d805221f683e1&tr=http://open.acgtracker.com:1096/announce",
		},
		{
			Name:       "[Skymoon-Raws][One Piece 海賊王][1011][ViuTV][WEB-RIP][CHT][SRT][1080p][MKV](先行版本)",
			MagnetLink: "magnet:?xt=urn:btih:b1e762663a3257189775aeb2001331ea31e894df&tr=http://open.acgtracker.com:1096/announce",
		},
		{
			Name:       "[桜都字幕组] 自称贤者弟子的贤者 / Kenja no Deshi o Nanoru Kenja [06][1080p][简体内嵌]",
			MagnetLink: "magnet:?xt=urn:btih:bc4c90d11999aaa50e2674a308685dc8aa1b97aa&tr=http://open.acgtracker.com:1096/announce",
		},
		{
			Name:       "[桜都字幕組] 自稱賢者弟子的賢者 / Kenja no Deshi o Nanoru Kenja [06][1080p][繁體內嵌]",
			MagnetLink: "magnet:?xt=urn:btih:98d9dec39cb71fa949d26e139f159e3fb92a1489&tr=http://open.acgtracker.com:1096/announce",
		},
		{
			Name:       "[桜都字幕组] 自称贤者弟子的贤者 / Kenja no Deshi o Nanoru Kenja [06][1080p][简繁内封]",
			MagnetLink: "magnet:?xt=urn:btih:abcd915c85f5dc4af6a60124adebf6b2eaceb1eb&tr=http://open.acgtracker.com:1096/announce",
		},
	}
	animeMagnets, err := miobtSpider.ExtractData(context.Background(), miobtHtml)
	if err != nil {
		t.Fatalf("miobtSpider.ExtractData error %v", err)
	}

	animeManget := compareAnimeMagnets(animeMagnets, animeMagnetsBeCompare)
	if animeManget != nil {
		t.Fatalf("miobtSpider.ExtractData animeMagnet=%v not equal the result ", animeManget)
	}
}

func TestDMHYSpider(t *testing.T) {
	dmhySpider := DmhySpider{
		BaseSpider: BaseSpider{
			log: log.NewHelper(newLogger()),
		},
	}

	animeMagnets, err := dmhySpider.ExtractData(context.Background(), dmhyHtml)
	if err != nil {
		t.Fatalf("DmhySpider.ExtractData error %v", err)
	}

	if len(animeMagnets) != 80 {
		t.Fatalf("DmhySpider.ExtractData length not equal 80 %v", len(animeMagnets))
	}
}

func TestNyaaSpider(t *testing.T) {
	nyaaSpider := NyaaSpider{
		BaseSpider: BaseSpider{
			log: log.NewHelper(newLogger()),
		},
	}

	animeMagnets, err := nyaaSpider.ExtractData(context.Background(), nyaaHtml)
	if err != nil {
		t.Fatalf("DmhySpider.ExtractData error %v", err)
	}

	if len(animeMagnets) != 75 {
		t.Fatalf("DmhySpider.ExtractData length not equal 75 %v", len(animeMagnets))
	}
}
