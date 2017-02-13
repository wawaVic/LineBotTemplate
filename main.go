// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	//"strconv"
	"reflect"
	"strings"
	"math/rand"
    "time"

	"github.com/line/line-bot-sdk-go/linebot"
)

var bot *linebot.Client

func main() {
	var err error
	bot, err = linebot.New(os.Getenv("ChannelSecret"), os.Getenv("ChannelAccessToken"))
	log.Println("Bot:", bot, " err:", err)
	http.HandleFunc("/callback", callbackHandler)
	port := os.Getenv("PORT")
	addr := fmt.Sprintf(":%s", port)
	http.ListenAndServe(addr, nil)
}

func callbackHandler(w http.ResponseWriter, r *http.Request) {
	events, err := bot.ParseRequest(r)

	if err != nil {
		if err == linebot.ErrInvalidSignature {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(500)
		}
		return
	}

	var UserMessage string
	var MultiMessage string
	var BotMessage, BotMessage_contain  string
	var GoogleKey string = "google"
	var YoutubeKey string = "youtube"
	var SingerKey string = "歌手"
	var AlbumeKey string = "專輯"
	var SongTitleKey string = "歌名"
	var LyricsKey string = "歌詞"
	var SearchLength int
	var RandomLength int
	//var SpaceSplitLength int
	var Hello string = "打招呼"
	var Weak string = "嫩啦"
	var Wantpapapa string = "我像大砲"
	var Baby string = "媽寶"
	var ChineseNewYear string = "拜年"
	var Crazyck string = "ck101"
	var LINEBotName string = "精靈"
	var TheCopperKey string = "頭目萬歲"
	var UpdateInfo string = "/?"

	var Vic string = "殺蛙"
	var Benson string = "陳冠宇"
	var Drew string = "彥竹"
	var Ziv string = "七七"
	var Kai string = "啾啾"
	var Lee string = "李志乾"
	var Hector string = "頂超"
	var Neal string = "賴柏采"
	var Cat string = "小貓咪"
	var Alvin string = "張銘仁"

	//var inputOne string = ""
	//var inputTwo string = ""
	//var inputThree string = ""
	var WhoKey string = "誰"
	var WhereKey string = "哪裡"
	var DoingKey string = "在幹嘛"
	PlayFingers := []string{"剪刀","石頭","布"}

	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
            	UserMessage = message.Text
            	MultiMessage = message.Text
			}
		}
		/*好難弄
		//輸入以空白隔開字串
		spacesplit := strings.Split(MultiMessage, " ")
		SpaceSplitLength = len(spacesplit)
  		if SpaceSplitLength == 2 {
			//inputOne, inputTwo, inputThree := spacesplit[0], spacesplit[1], spacesplit[2]
			inputOne = spacesplit[0]
			inputTwo = spacesplit[1]
		}
  		if SpaceSplitLength == 3 {
			inputOne = spacesplit[0]
			inputTwo = spacesplit[1]
			inputThree = spacesplit[2]
		}
		*/

        //搜尋區
		ifsearch := strings.Split(MultiMessage, " ")
		SearchLength = len(ifsearch)
  		if SearchLength == 2 {
			engine, SearchFor := ifsearch[0], ifsearch[1]
			engine = strings.ToLower(engine)
			if reflect.DeepEqual(engine, GoogleKey) {
				BotMessage = "Google搜尋結果：\n" + "https://www.google.com.tw/#q=" + SearchFor
			}
			if reflect.DeepEqual(engine, YoutubeKey) {
				BotMessage = "Youtube搜尋結果：\n https://www.youtube.com/results?search_query=" + SearchFor
			}
			if reflect.DeepEqual(engine, SingerKey) {
				BotMessage = "魔鏡歌詞網搜尋結果：\n https://mojim.com/" + SearchFor + ".html?t1"
			}
			if reflect.DeepEqual(engine, AlbumeKey) {
				BotMessage = "魔鏡歌詞網搜尋結果：\n https://mojim.com/" + SearchFor + ".html?t2"
			}
			if reflect.DeepEqual(engine, SongTitleKey) {
				BotMessage = "魔鏡歌詞網搜尋結果：\n https://mojim.com/" + SearchFor + ".html?t3"
			}
			if reflect.DeepEqual(engine, LyricsKey) {
				BotMessage = "魔鏡歌詞網搜尋結果：\n https://mojim.com/" + SearchFor + ".html?t4"
			}
		}
        //搜尋區end

        //亂數區
        /*
		if reflect.DeepEqual(UserMessage, WhoKey) {
			BotMessage = randomcase()
		}
		*/
		ifrandom := strings.Split(MultiMessage, " ")
		RandomLength = len(ifrandom)
  		if RandomLength == 3 {
		who, where, doing := ifrandom[0], ifrandom[1], ifrandom[2]
//			if reflect.DeepEqual(who, WhoKey); reflect.DeepEqual(where, WhereKey); reflect.DeepEqual(doing, DoingKey) {
//				BotMessage = randomcase()
//			}
			if reflect.DeepEqual(who, WhoKey) {
				if reflect.DeepEqual(where, WhereKey) {
					if reflect.DeepEqual(doing, DoingKey) {
    					BotMessage = randomcase()
					}
				}
			}
		}
        //亂數區end

		//猜拳區
		for _, Finger := range PlayFingers {
			if reflect.DeepEqual(UserMessage, Finger) {
			BotMessage = FingersGame(Finger)
			}
		}
		//猜拳區end

		//單詞回覆區
		if reflect.DeepEqual(UserMessage, UpdateInfo) {
			BotMessage = "目前版本：V1.2.2 \n" +
						 "更新日期：2017.02.13 \n" +
						 "更新內容： \n" +
						 "1.可以搜尋歌詞，用法：\n" +
						 "    歌手 蕭敬騰\n" +
						 "    專輯 五月天\n" +
						 "    歌名 易燃易爆炸\n" +
						 "    歌詞 打開門就見山"
						 //"1.小精靈銅鑰匙開賣啦！"
						 //"1.可以與民間小精靈猜拳了！"
						 //"1.可以與民間小精靈猜拳了！ \n" +
						 //"2.嫩啦與媽寶擁有了更多詞彙 \n" +
						 //"3.加入2個秘密關鍵字"
		}
		if reflect.DeepEqual(UserMessage, Hello) {
			BotMessage = "大家好阿，我是民間小精靈，你們這些小GG"
		}
		if reflect.DeepEqual(UserMessage, Weak) {
			BotMessage = randomSingle("weak")
		}
		if reflect.DeepEqual(UserMessage, Wantpapapa) {
			BotMessage = "他說OK沒問題"
		}
		if reflect.DeepEqual(UserMessage, Baby) {
			BotMessage = randomSingle("baby")
		}
		if reflect.DeepEqual(UserMessage, ChineseNewYear) {
			BotMessage = "祝大家狗年行大運，旺旺旺旺！"
		}
		if reflect.DeepEqual(UserMessage, Vic) {
			BotMessage = "樹林彭于晏!"
		}
		if reflect.DeepEqual(UserMessage, Benson) {
			BotMessage = "天母國文自以為小老師"
		}
		if reflect.DeepEqual(UserMessage, Drew) {
			BotMessage = "臺中送粽雞"
		}
		if reflect.DeepEqual(UserMessage, Ziv) {
			BotMessage = "龜山嫩啦NPC"
		}
		if reflect.DeepEqual(UserMessage, Kai) {
			BotMessage = "南港欺負粉粉大哥"
		}
		if reflect.DeepEqual(UserMessage, Lee) {
			BotMessage = "天母大白鯊"
		}
		if reflect.DeepEqual(UserMessage, Hector) {
			BotMessage = "北投癡漢"
		}
		if reflect.DeepEqual(UserMessage, Neal) {
			BotMessage = "新竹波波熊股東"
		}
		if reflect.DeepEqual(UserMessage, Cat) {
			BotMessage = "基隆廟口鼎邊銼難吃"
		}
		if reflect.DeepEqual(UserMessage, Alvin) {
			BotMessage = "永和有永和路，中和也有永和路，中和有中和路，永和也有中和路；中和的中和路有接永和的中和路，永和的永和路沒接中和的永和路；永和的中和路有接永和的永和路，中和的永和路沒接中和的中和路。"
		}
		//單詞回覆區end

		//關鍵字包含區
		if strings.Contains(UserMessage, Crazyck) {
			BotMessage_contain = "卡提諾狂新聞上線囉！"
		}
		if strings.Contains(UserMessage, LINEBotName) {
			BotMessage_contain = "誰cue偶？"
		}
		if strings.Contains(UserMessage, TheCopperKey) {
			BotMessage_contain = "獲得銅鑰匙(10)"
		}
		//關鍵字包含區end

/*關閉
		//模式區
        mode, e := strconv.Atoi(UserMessage)
	    if e != nil {
    	    fmt.Println(e)
	    }
	    //modestr := strconv.Itoa(mode)

		switch mode {
		case 1:
			BotMessage = "中文"
		case 2:
			BotMessage = "中文二"
		}
		//模式區end
*/

		//reply
		if BotMessage != "" {
			bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(BotMessage)).Do()
		}
		if BotMessage_contain != "" {
			bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(BotMessage_contain)).Do()
		}
		//reply end
//			bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("Debug: " + UserMessage)).Do()
  	}

/*原版
	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.ID+":"+ message.Text +" OK!")).Do()
				err != nil {
					log.Print(err)
				}
			}
		}
	}
*/
}

func randomcase() string {
	var who string = ""
	var where string = ""
	var doing string = ""

	name := []string{
		"殺蛙",
		"陳冠宇",
		"彥竹",
		"七七",
		"啾啾",
		"李志乾",
		"頂超",
		"賴博彩",
		"小貓咪",
		"張銘仁",
	}
	place := []string{
		"在大庭廣眾下",
		"在陳冠宇家",
		"在空無一人的浴室",
		"在安靜的圖書館",
		"在一萬英呎的高空",
		"在家裡的地板",
		"在公司開重要會議時",
		"在暗戀的女生的房間",
		"在丁丁房間",
		"在召換峽谷",
	}
	thing := []string{
		"放鞭炮",
		"發呆",
		"開噴",
		"大喊：我要當老師",
		"癡漢笑",
		"放屁",
		"吃很燙的拉麵",
		"切生魚片",
		"脫褲子",
		"打手槍",
		"比賽挖鼻屎",
	}
	rand.Seed(time.Now().UnixNano())
	who = name[rand.Intn(len(name))]
	where = place[rand.Intn(len(place))]
	doing = thing[rand.Intn(len(thing))]

	var result string =	strings.Join([]string{who, where, doing}, " ")

	return result
}

func randomSingle(Keyword string) string {
	var result string =	""

	weak := []string{
		"嫩",
		"嫩啦",
		"女束ㄆ",
		"嗨，肉腳，你們好阿",
		"連我阿罵都比你還行",
		"B嘴，廢物",
	}
	baby := []string{
		"哭哭喔",
		"只會叫JG",
		"讀私立喔",
		"媽寶喔",
	}

	rand.Seed(time.Now().UnixNano())

	if reflect.DeepEqual(Keyword, "weak") {
		result = weak[rand.Intn(len(weak))]
	}
	if reflect.DeepEqual(Keyword, "baby") {
		result = baby[rand.Intn(len(baby))]
	}

	return result
}

//猜拳遊戲 石頭3 剪刀2 布1
func FingersGame(Challenger string) string {
	PlayFingers := []string{"剪刀","石頭","布"}
	var Stone string = "石頭"
	var Scissor string = "剪刀"
	var Paper string = "布"
	var ChallengerValue int
	var Bot string
	var BotValue int
	var result string =	""

	//定義使用者出拳數值
	if reflect.DeepEqual(Challenger, Stone) {
		ChallengerValue = 3
	}
	if reflect.DeepEqual(Challenger, Scissor) {
		ChallengerValue = 2
	}
	if reflect.DeepEqual(Challenger, Paper) {
		ChallengerValue = 1
	}

	//BOT出拳
    rand.Seed(time.Now().UnixNano())
	Bot = PlayFingers[rand.Intn(len(PlayFingers))]

	//定義BOT出拳數值
	if reflect.DeepEqual(Bot, Stone) {
		BotValue = 3
	}
	if reflect.DeepEqual(Bot, Scissor) {
		BotValue = 2
	}
	if reflect.DeepEqual(Bot, Paper) {
		BotValue = 1
	}

	if (BotValue == 3 && ChallengerValue == 2) || (BotValue == 2 && ChallengerValue == 1) || (BotValue == 1 && ChallengerValue == 3){
		result = "我出" + Bot + "，嫩"
	} else if (ChallengerValue == 3 &&  BotValue == 2) || (ChallengerValue == 2 &&  BotValue == 1) || (ChallengerValue == 1 &&  BotValue == 3){
		result = "我出" + Bot + "，你厲害，最棒最棒"
	} else if BotValue == ChallengerValue {
		result = "我也出" + Bot + "，再來啊！"
	} else {
		result = "猜拳機壞掉了，KUKU"
	}

	return result
}