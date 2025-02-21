package handlers

import (
	"fmt"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
)

/*	This handler handles the text sent from telex, processed and then sent a formatted response back
 */
func HandleGenerate(ctx *gin.Context) {
	var msgReq MsgRequest
	// var data map[string]interface{}

	err := ctx.Bind(&msgReq)
	if err != nil {
		log.Println("error bindind post data: ", err)
		ctx.JSON(400, gin.H{
			"error": " invalid JSON payload",
		})
		return
	}

	log.Printf("All request data from telex: %+v\n", msgReq)
	fmt.Println("*******************")
	log.Printf("Channel ID is: %s\n", msgReq.ChannelID)
	log.Printf("Message is: %s\n", msgReq.Message)
	log.Printf("Settings are: %+v\n", msgReq.Settings)
	fmt.Println("*******************")

	// if msg is as a result of webhook

	if strings.HasPrefix(msgReq.Message, "*****") {
		ctx.JSON(200, gin.H{
			"event_name": "Webhook Form Msg",
			"message":    msgReq.Message,
			"status":     "success",
			"username":   "formify-bot",
		})
		// ctx.JSON(400, gin.H{
		// 	"status":  "error",
		// 	"message": "invalid command",
		// })
		return
	}
	// *******************

	text := ExtractText(msgReq.Message)
	// text := msgReq.Message

	if text != "/generate_url" {
		ctx.JSON(200, gin.H{
			"event_name": "Invalid Command",
			"message":    "type '/generate_url' to get the unique url for your html forms",
			"status":     "success",
			"username":   "formify-bot",
		})
		// ctx.JSON(400, gin.H{
		// 	"status":  "error",
		// 	"message": "invalid command",
		// })
		return
	}

	form_name := msgReq.Settings[0].Default
	url := GenerateUniqueURL(msgReq.Settings)
	message := FormatMSG(form_name, url)

	// returns the message as a json response to telex
	ctx.JSON(200, gin.H{
		"event_name": "Unique URL Generated",
		"message":    message,
		"status":     "success",
		"username":   "formify-bot",
	})
}
