package discord

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/yaskinny/discord-notif/pkg/drone"
	"github.com/yaskinny/discord-notif/pkg/models"
)

type Notif interface {
	Setter() ([]models.Field, error)
}

func SetMessage(n Notif) ([]models.Field, error) {
	f, err := n.Setter()
	return f, err
}

var N Notif

var Colors = map[string]string{
	"Red":    "16723200",
	"Green":  "65295",
	"Purple": "16711892",
}

type Message struct {
	Username string         `json:"username"`
	Avatar   string         `json:"avatar_url"`
	Content  string         `json:"content"`
	Embeds   []models.Embed `json:"embeds"`
}

func (m *Message) EmbedsSetter(c models.Cli) {
	var (
		color  string
		fields []models.Field
		err    error
	)
	switch c.Kind {
	case "pipeline":
		this := os.Getenv(`DRONE`)
		if this == "true" {
			var d drone.Drone
			fields, err = SetMessage(d)
			if err != nil {
				fmt.Fprintf(os.Stderr, "err ~> \t %v\n", err)
				os.Exit(1)
			}
		}
		// this = os.Getenv(`GITLAB_CI`)

	}

	switch c.State {
	case "start":
		color = Colors["Purple"]
	case "succeed":
		color = Colors["Green"]
	case "error":
		color = Colors["Red"]
	}
	e := models.Embed{
		Title:       c.Kind,
		Description: c.Name,
		Color:       color,
		Fields:      fields,
	}
	m.Embeds = append(m.Embeds, e)
}

// create json body for notification
func (m *Message) NotifCreator(c models.Cli) ([]byte, error) {
	m.EmbedsSetter(c)
	m.Content = SetContent(c)
	m.Avatar = os.Getenv(`DISCORD_AVATAR`)
	m.Username = os.Getenv(`DISCORD_USERNAME`)
	if m.Avatar == "" {
		m.Avatar = "https://i.imgur.com/8nLFCVP.png"
	}
	if m.Username == "" {
		m.Username = "Notification"
	}

	msg, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	return msg, nil
}

func SetContent(c models.Cli) string {
	date := time.Now()
	var this string
	switch c.State {
	case "start":
		this = "**__started__**"
	case "error":
		this = "**__failed__**"
	case "succeed":
		this = "**__succeeded__**"
	}
	return fmt.Sprintf("%v %v %v at %v!", c.Kind, c.Name, this, date.Format("2006.01.02 15:04:05"))
}
