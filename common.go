package ontap

import (
	log "github.com/sirupsen/logrus"
	"os"
	"strings"
)

type HttpError struct {
	Error struct {
		Message string `json:"message"`
		Code    string `json:"code"`
	} `json:"error"`
}

func ConfigLog() {
	level := os.Getenv("ONTAP")
	level = strings.ToUpper(level)

	if level == "DEBUG" {
		log.SetFormatter(&log.JSONFormatter{})

	} else {
		log.SetFormatter(&log.TextFormatter{})
	}
}
