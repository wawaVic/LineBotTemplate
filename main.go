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
	"strconv"
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
	var BotMessage string
	var GoogleKey string = "google"
	var YoutubeKey string = "youtube"
	var SearchLength int
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

	var Who string = "誰"


	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
            	UserMessage = message.Text
			}
		}

		ifsearch := strings.Split(UserMessage, " ")
		SearchLength = len(ifsearch)
  		if SearchLength > 1 {
			engine, SearchFor := ifsearch[0], ifsearch[1]
			engine = strings.ToLower(engine)
			if reflect.DeepEqual(engine, GoogleKey) {
				BotMessage = "https://www.google.com.tw/#q=" + SearchFor
			}
			if reflect.DeepEqual(engine, YoutubeKey) {
				BotMessage = "https://www.youtube.com/results?search_query=" + SearchFor
			}
		}

		if reflect.DeepEqual(UserMessage, Hello) {
			BotMessage = "大家好阿，我是民間小精靈，把我搞出來的是一個帥哥"
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

        //亂數區
		if reflect.DeepEqual(UserMessage, Who) {
			BotMessage = randomcase
		}


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
	answers := []string{
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
	rand.Seed(time.Now().UnixNano()) // Try changing this number!
	var who string = answers[rand.Intn(len(answers))]
	
	return who
}