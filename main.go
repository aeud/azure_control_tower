package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

func slack() error {
	slackEndpoint := os.Getenv("ACT_SLACK_ENDPOINT")
	portalURL := os.Getenv("ACT_PORTAL_URL")
	vmName := os.Getenv("ACT_VM_NAME")
	vmCost := os.Getenv("ACT_VM_COST")
	mentions := strings.Split(os.Getenv("ACT_SLACK_MENTIONS"), ",")
	for i, m := range mentions {
		mentions[i] = fmt.Sprintf("<@%s>", m)
	}
	_, err := http.Post(
		slackEndpoint,
		"application/json",
		strings.NewReader(fmt.Sprintf(`{
			"username": "Azure Control Tower",
			"text": "%s, your VM _%s_ is still up and running. This machine will cost *%s€* if it runs the whole night.",
			"icon_emoji": ":azure:",
			"attachments": [
				{
					"text": "Would you like to stop it?",
					"color": "#3AA3E3",
					"attachment_type": "default",
					"actions": [
						{
							"text": "Go to the Portal",
							"type": "button",
							"url": "%s",
							"style": "default"
						}
					]
				}
			]
		}`, strings.Join(mentions, ", "), vmName, vmCost, portalURL)),
	)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	if err := slack(); err != nil {
		panic(err)
	}
}
