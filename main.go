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
	var BotMessage string
	var GoogleKey string = "google"
	var YoutubeKey string = "youtube"
	//var SearchLength int
	//var RandomLength int
	var SpaceSplitLength int
	var Hello string = "打招呼"
	var Weak string = "嫩啦"
	var Wantpapapa string = "我像大砲"
	var Baby string = "媽寶"
	
	var Vic string = "殺蛙"
	var Benson string = "陳冠宇"
	var Drew string = "彥竹"
	var Ziv string = "七七"
	//var Kai string = "啾啾"
	//var Lee string = "李志乾"
	//var Hector string = "頂超"
	//var Neal string = "賴博彩"
	//var Cat string = "小貓咪"
	//var Alvin string = "張銘仁"

	var WhoKey string = "誰"
	var WhereKey string = "哪裡"
	var DoingKey string = "在幹嘛"


	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
            	UserMessage = message.Text
            	MultiMessage = message.Text
			}
		}

		//輸入以空白隔開字串
		spacesplit := strings.Split(MultiMessage, " ")
		SpaceSplitLength = len(spacesplit)
		inputOne, inputTwo, inputThree := spacesplit[0], spacesplit[1], spacesplit[2]



        //搜尋區
		//ifsearch := strings.Split(MultiMessage, " ")
		//SearchLength = len(ifsearch)
  		//if SearchLength == 2 {
		//	engine, SearchFor := ifsearch[0], ifsearch[1]
			inputOne = strings.ToLower(inputOne)
			if reflect.DeepEqual(inputOne, GoogleKey) {
				BotMessage = "Google搜尋結果：\n https://www.google.com.tw/#q=" + inputTwo
			}
			if reflect.DeepEqual(inputOne, YoutubeKey) {
				BotMessage = "Youtube搜尋結果：\n https://www.youtube.com/results?search_query=" + inputTwo
			}
		//}
        //搜尋區end

        //亂數區
        /*
		if reflect.DeepEqual(UserMessage, WhoKey) {
			BotMessage = randomcase()
		}
		*/
		//ifrandom := strings.Split(MultiMessage, " ")
		//RandomLength = len(ifrandom)
  		//if RandomLength == 3 {
		//	who, where, doing := ifrandom[0], ifrandom[1], ifrandom[2]
//			if reflect.DeepEqual(who, WhoKey); reflect.DeepEqual(where, WhereKey); reflect.DeepEqual(doing, DoingKey) {
//				BotMessage = randomcase()
//			}
			if reflect.DeepEqual(inputOne, WhoKey) {
				if reflect.DeepEqual(inputTwo, WhereKey) {
					if reflect.DeepEqual(inputThree, DoingKey) {
    					BotMessage = randomcase()
					}
				}
			}
		//}
        //亂數區end



		if reflect.DeepEqual(UserMessage, Hello) {
			BotMessage = "大家好阿，我是民間小精靈，你們這些小GG"
		}
		if reflect.DeepEqual(UserMessage, Weak) {
			BotMessage = "嫩"
		}
		if reflect.DeepEqual(UserMessage, Wantpapapa) {
			BotMessage = "他說OK沒問題"
		}
		if reflect.DeepEqual(UserMessage, Baby) {
			BotMessage = "哭哭喔"
		}
		if reflect.DeepEqual(UserMessage, Vic) {
			BotMessage = "樹林彭于晏!"
		}
		if reflect.DeepEqual(UserMessage, Benson) {
			BotMessage = "天母烤秋勤大師"
		}
		if reflect.DeepEqual(UserMessage, Drew) {
			BotMessage = "臺中送粽雞"
		}
		if reflect.DeepEqual(UserMessage, Ziv) {
			BotMessage = "龜山嫩啦NPC"
		}

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
//			bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("mode: " + modestr + " Message: " + BotMessage + " UserMessage: " + UserMessage)).Do()
			bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(BotMessage)).Do()
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
	var who string = name[rand.Intn(len(name))]
	var where string = place[rand.Intn(len(place))]
	var doing string = thing[rand.Intn(len(thing))]

	var result string =	strings.Join([]string{who, where, doing}, " ")

	return result
}