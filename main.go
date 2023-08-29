package main

import (
	"sync"

	"go.uber.org/zap"

	"github.com/rafaribe/sensor-manager/config"
	sensor "github.com/rafaribe/sensor-manager/sensor"
)

func main() {
	conf := config.ParseConfiguration()
	log := zap.S()
	log.Infof("Configuration successfully parsed %#v", conf)
	var wg sync.WaitGroup
	for _, sensorConfig := range conf.Sensors {
		wg.Add(2)
		log.Info("Sensor %#v", sensorConfig)
		if sensorConfig.Model == "awair_element" {
			awair := sensor.Init(&sensorConfig, &conf.Store)
			go func() {
				defer wg.Done()
				go sensor.SettingsRoutine(awair, conf.ScrapeInterval*5)
				go sensor.TelemetryRoutine(awair, conf.ScrapeInterval)
			}()
		}
	}
	wg.Wait()
}
