package main

import (
	"github.com/rafaribe/planetwatch-awair-uploader/config"
	"go.uber.org/zap"
)

func main() {
	conf := config.ParseConfiguration()
	log := zap.S()
	log.Info("Configuration successfully parsed %#v", conf)
}
