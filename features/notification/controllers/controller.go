package controllers

import (
	"errors"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

type Controller struct{}

func (u Controller) SendNotification(c *gin.Context) {
	sendNotification(c)
}

//*** API endpoint usage ***
// notification?channels=<'line' | 'telegram' | 'line,telegram'>&message=<message>
func sendNotification(c *gin.Context) {

	message, err2 := url.PathUnescape(c.Query("message"))
	if err2 != nil || message == "" {
		c.String(http.StatusBadRequest, "Error")
	}

	channelString, err3 := url.PathUnescape(c.Query("channels"))
	if err3 != nil || channelString == "" {
		c.String(http.StatusBadRequest, "Error")
	}
	channels := strings.Split(channelString, ",")

	chatIDs := getChatRoomIDs()

	if message == "" || len(channels) == 0 {
		c.String(http.StatusBadRequest, "Error")
	}

	if len(chatIDs) == 0 {
		err := sendMessage("", message, channels)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
		}
	} else {
		for _, chatID := range chatIDs {
			err := sendMessage(chatID, message, channels)
			if err != nil {
				c.String(http.StatusInternalServerError, err.Error())
			}
		}

	}

	c.String(http.StatusOK, "Error")
}

//TODO: Database request for 'botID', 'chatID'
func getChatRoomIDs() []string {
	return []string{}
}

func sendMessage(chatID string, message string, channels []string) error {
	for _, channel := range channels {
		switch channel {
		case "telegram":
			telegramError := sendTelegramMessage(chatID, message)
			if telegramError != nil {
				return telegramError
			}

		case "line":
			lineError := sendLineMessage(chatID, message)
			if lineError != nil {
				return lineError
			}
		}
	}
	return nil
}

func sendTelegramMessage(chatID string, message string) error {
	form := url.Values{}
	form.Add("key", os.Getenv("CHAT_BOT_SECRET_KEY"))
	form.Add("template", "false")
	if len(chatID) > 0 {
		form.Add("chat_id", chatID)
	}
	form.Add("message", message)

	request, error := http.NewRequest("POST", os.Getenv("CHAT_BOT_URI"), strings.NewReader(form.Encode()))
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	if error != nil {
		return errors.New("cannot create reqeust telegram")
	}
	client := &http.Client{}
	_, err2 := client.Do(request)
	if err2 != nil {
		return errors.New("cannot call telegram")
	}
	return nil
}

//TODO: Send request to line
func sendLineMessage(chatID string, message string) error {
	return nil
}
