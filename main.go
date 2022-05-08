package main

import (
	"go.uber.org/zap"

	"github.com/rafaribe/sensor-manager/config"
	sensor "github.com/rafaribe/sensor-manager/sensor"
)

func main() {
	conf := config.ParseConfiguration()
	log := zap.S()
	log.Info("Configuration successfully parsed %#v", conf)
	for _, sensorConfig := range conf.Sensors {
		log.Info("Sensor %#v", sensorConfig)
		if sensorConfig.Model == "awair_element" {
			sensor := sensor.Init(&sensorConfig)
			settings, err := sensor.GetSettings()
			if err != nil {
				log.Error("Error getting settings for sensor %#v", sensorConfig)
			}
			log.Debugw("Sensor Settings %#v", settings)
		}

	}
}
