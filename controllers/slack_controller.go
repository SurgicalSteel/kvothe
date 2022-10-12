package controllers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/SurgicalSteel/kvothe/models"
	"github.com/gin-gonic/gin"
)

var subCommands = map[string]int{
	"person-get": 1,
	"person-add": 2,
	"quote-get":  3,
	"quote-add":  4,
}

func (kc *KvotheController) HandleSlackSlashCommand(c *gin.Context) {
	c.Request.ParseForm()

	var (
		ssc     models.SlashCommand
		message string
	)
	ssc.Token = c.Request.PostForm.Get("token")
	ssc.TeamID = c.Request.PostForm.Get("team_id")
	ssc.TeamDomain = c.Request.PostForm.Get("team_domain")
	ssc.EnterpriseID = c.Request.PostForm.Get("enterprise_id")
	ssc.EnterpriseName = c.Request.PostForm.Get("enterprise_name")
	ssc.ChannelID = c.Request.PostForm.Get("channel_id")
	ssc.ChannelName = c.Request.PostForm.Get("channel_name")
	ssc.UserID = c.Request.PostForm.Get("user_id")
	ssc.UserName = c.Request.PostForm.Get("user_name")
	ssc.Command = c.Request.PostForm.Get("command")
	ssc.Text = c.Request.PostForm.Get("text")
	ssc.ResponseURL = c.Request.PostForm.Get("response_url")
	ssc.TriggerID = c.Request.PostForm.Get("trigger_id")
	ssc.APIAppID = c.Request.PostForm.Get("api_app_id")

	verificationToken := kc.Configurations.Core.Kvothe.Slack.VerificationToken
	if verificationToken != ssc.Token {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	subCommandFound := false
	subCommand := ""
	for sc, _ := range subCommands {
		if strings.HasPrefix(ssc.Text, sc) {
			subCommandFound = true
			subCommand = sc
		}
	}

	if !subCommandFound {
		message = "The sub command that you entered is not registered in our system"
		c.Writer.Write([]byte(message))
		return
	}

	params := strings.TrimPrefix(ssc.Text, subCommand)
	message, err := kc.processMessage(subCommand, params)
	if err != nil {
		c.Writer.Write([]byte(err.Error()))
		return
	}

	c.Writer.Write([]byte(message))
}

func (kc *KvotheController) processMessage(subCommand, params string) (string, error) {

	switch subCommand {
	case "person-add":
		return kc.Services.AddNewPersonMockData(params)
	case "quote-add":
		return kc.Services.AddNewQuoteMockData(params)
	case "person-get":
		return kc.Services.GetPersonByIDMockData(params)
	case "quote-get":
		return kc.Services.GetQuoteByIDMockData(params)
	default:
		return "", fmt.Errorf("sub command %s is not registered in our system", subCommand)
	}
}
