package main

import (
	"go.uber.org/zap"

	"github.com/rafaribe/sensor-manager/config"
)

func main() {
	conf := config.ParseConfiguration()
	log := zap.S()
	log.Info("Configuration successfully parsed %#v", conf)

}
