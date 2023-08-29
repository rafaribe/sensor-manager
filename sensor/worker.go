package sensor

import (
	"time"

	"go.uber.org/zap"
)

func SettingsRoutine(sensor *AwairElementSensor, interval int32) {
	log := zap.S()
	for {
		settings, err := sensor.GetSettings()
		if err != nil {
			log.Errorf("Error getting settings for sensor %#v: %v", sensor, err)
		} else {
			sensor.SaveSettings(settings)
			log.Infof("Settings fetched and saved at %s", time.Now())
		}
		time.Sleep(5 * time.Minute)
	}
}

func TelemetryRoutine(sensor *AwairElementSensor, interval int32) {
	log := zap.S()
	for {
		telemetry, err := sensor.GetTelemetry()
		if err != nil {
			log.Warnf("Error getting telemetry for sensor %#v: %v", sensor, err)
		} else {
			sensor.SaveTelemetry(telemetry)
			log.Infof("Saved telemetry for sensor %#v at %s", sensor, time.Now())
		}
		time.Sleep(time.Duration(interval) * time.Second)
	}
}
