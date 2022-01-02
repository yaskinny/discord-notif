package script

import (
	"fmt"
	"os"
	"strings"

	"github.com/yaskinny/discord-notif/pkg/models"
)

type Script struct {
	//
}

// goes through NOTIF_FIELDS environment variable and set the ones user set for template
func (s Script) Setter() ([]models.Field, error) {
	var fields []models.Field
	scriptEnvs := os.Getenv(`NOTIF_FIELDS`)
	if scriptEnvs == "" {
		fmt.Fprintf(os.Stderr, "`NOTIF_FIELDS` is empty, default template used.\n")
		fields = append(fields, models.Field{
			Name:   "USER",
			Value:  os.Getenv(`USER`),
			Inline: "false",
		})
		fields = append(fields, models.Field{
			Name:   "SHELL",
			Value:  os.Getenv(`SHELL`),
			Inline: "false",
		})
		return fields, nil
	}
	scriptTags := strings.Split(scriptEnvs, ",")
	for _, v := range scriptTags {
		this := os.Getenv(v)
		if this != "" {
			fields = append(fields, models.Field{
				Name:   v,
				Value:  this,
				Inline: "false",
			})
		}
	}
	return fields, nil
}
