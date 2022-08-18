package functions

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"os"
	"regexp"

	"github.com/yaskinny/discord-notif/pkg/discord"
	"github.com/yaskinny/discord-notif/pkg/models"
)

func help() {
	fmt.Fprintf(os.Stderr, `Usage:
discord-notif <name> <kind> <state>
state:
        error   It will exit with exit status code 1.
        start   just notify on starting
        succeed just notify on succeed
kind:
        pipeline  Use within your Gitlab CI or Drone pipeline
        script    Use in Your scripts
environment variables:
        DISCORD_URL       Your discord bot API URL
        DISCORD_TAGS      Comma seperated list of user/role ids to tag in notifications
        DISCORD_USERNAME  Username for Bot(Optional, default is 'Notification')
        DISCORD_AVATAR    Your discord bot avatar URL(Optional, default is a 'Sad PEPE')
        NOTIF_TAGS        Comma seperated list of environments to use in templating notifications
                          For Drone CI: https://docs.drone.io/pipeline/environment/reference
                            if the variable name is DRONE_COMMIT_AUTHOR, use it like DroneCommitAuthor
                          For Gitlab CI: https://docs.gitlab.com/ee/ci/variables/predefined_variables.html
                            if the variable name is GITLAB_USER_ID, use it like GitlabUserId
                          For scripts: only works when kind is 'script'
                            list what You want to see in template like NOTIF_TAGS='USER,HOSTNAME'
                          If This variable is not set, for each mode There is a default template and it will be used
examples:
pipeline -> NOTIF_FIELDS="DroneCommitAuthor,DroneCommit" discord-notif project1 pipeline start
script -> DISCORD_TAGS='11111111111111,22222222222222,&333333333333' NOTIF_FIELDS="USER,HOSTNAME,SHELL" discord-notif backup-mysql script succeed
`)
}

func isValidApi(url string) bool {
	regex := regexp.MustCompile(`https:\/\/(discord|discordapp)\.com\/api\/webhooks\/[0-9]{18,20}\/[a-zA-Z0-9_\-]+`)
	return regex.Match([]byte(url))
}

var (
	ErrParamCount        = "Wrong Number of parameters."
	ErrInvalidState      = "State is not supported."
	ErrInvalidDiscordApi = "Discord URL is not valid."
	ErrEmptyDiscordApi   = "Discord URL is empty, please set `DISCORD_URL`"
)
var DiscordApi string = os.Getenv(`DISCORD_URL`)

func PreCheck() {
	if DiscordApi == "" {
		fmt.Fprintf(os.Stderr, "err ~> \t %v\n", ErrEmptyDiscordApi)
		help()
		os.Exit(1)
	}

	switch {
	// check number of params are correct
	case len(os.Args) != 4:
		fmt.Fprintf(os.Stderr, "err ~> \t %v \n", ErrParamCount)
		help()
		os.Exit(1)
		// check if state is correct
	case !(os.Args[3] == "start" || os.Args[3] == "error" || os.Args[3] == "succeed"):
		fmt.Fprintf(os.Stderr, "err ~> \t %v\n", ErrInvalidState)
		help()
		os.Exit(1)
	// validate discord url
	case isValidApi(DiscordApi) != true:
		fmt.Fprintf(os.Stderr, "err ~> \t %v\n", ErrInvalidDiscordApi)
		help()
		os.Exit(1)
	}
}

func sendHttpRequest(c models.Cli, p []byte) (int, error) {
	req, err := http.NewRequest(http.MethodPost, DiscordApi, bytes.NewBuffer(p))
	if err != nil {
		return 999, err
	}
	req.Header.Add("content-type", "application/json")
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return 999, err
	}
	defer resp.Body.Close()
	if resp.StatusCode > 299 || resp.StatusCode < 200 {
		return resp.StatusCode, errors.New(fmt.Sprintf("Something is wrong with request -> %v\n", resp.Status))
	}
	return resp.StatusCode, nil
}
func SendMsg(c models.Cli) {
	var m discord.Message
	payload, err := m.NotifCreator(c)
	if err != nil {
		fmt.Fprintf(os.Stderr, "err ~> \t %v\n", err)
		help()
		os.Exit(1)
	}
	r, err := sendHttpRequest(c, payload)
	if err != nil {
		fmt.Fprintf(os.Stderr, "err ~> \t %v\n", err)
		os.Exit(1)
	}
	fmt.Fprintf(os.Stdout, "Notification sent successfully with response code: %d\n", r)
}
