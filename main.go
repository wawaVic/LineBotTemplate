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
	var Hello string = "安安"
	var Vic string = "殺蛙"
	var Benson string = "陳冠宇"
	var Drew string = "彥竹"

	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
            	UserMessage = message.Text
			}
		}

		ifgoogle := strings.Split(UserMessage, " ")

		google, SearchFoor := ifgoogle[0], ifgoogle[1]
		google = strings.ToLower(google)
		if reflect.DeepEqual(google, GoogleKey) {
			BotMessage = "https://www.google.com.tw/#q=" + SearchFoor
		}
		bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("Debug: " + UserMessage)).Do()
		
		if reflect.DeepEqual(UserMessage, Hello) {
			BotMessage = "你好阿"
		}
		if reflect.DeepEqual(UserMessage, Vic) {
			BotMessage = "新北彭于晏!"
		}
		if reflect.DeepEqual(UserMessage, Benson) {
			BotMessage = "天母烤秋勤大師"
		}
		if reflect.DeepEqual(UserMessage, Drew) {
			BotMessage = "! 臺中馮迪索"
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
			bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(BotMessage + " Debug" + UserMessage)).Do()
		}
		//reply end
			bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("Debug: " + UserMessage)).Do()
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