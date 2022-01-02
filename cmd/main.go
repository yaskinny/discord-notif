package main

import (
	"os"

	"github.com/yaskinny/discord-notif/pkg/functions"
	"github.com/yaskinny/discord-notif/pkg/models"
)

func init() {
	functions.PreCheck()
}
func main() {
	var cli models.Cli
	cli.Name = os.Args[1]
	cli.Kind = os.Args[2]
	cli.State = os.Args[3]
	functions.SendMsg(cli)
}
