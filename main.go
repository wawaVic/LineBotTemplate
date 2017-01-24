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
//	var s string = "2"


	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
            	UserMessage = message.Text
			}
		}

//	mode := strconv.ParseInt(s, 10, 64)

        mode, e := strconv.Atoi(UserMessage)
	    if e != nil {
    	    fmt.Println(e)
	    }
	    modestr := strconv.Itoa(mode)
/*
		if mode = 2 {
			BotMessage = "222"
		}
*/


		switch mode {
		case 1:
			BotMessage = "中文"
		case 2:
			BotMessage = "中文二"
		}

//		if BotMessage != "" {
//			bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("mode: " + mode + " Message: " + BotMessage + " UserMessage: " + UserMessage)).Do()
			bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("mode: " + modestr + " Message: " + BotMessage + " UserMessage: " + UserMessage)).Do()
//		}
  	}

/*
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