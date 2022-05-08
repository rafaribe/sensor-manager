package sensor

import (
	"time"

	"go.uber.org/zap"
)

func SettingsRoutine(sensor *AwairElementSensor, interval int32) {
	for {
		log := zap.S()
		settings, err := sensor.GetSettings()
		if err != nil {
			log.Errorf("Error getting settings for sensor %#v", sensor)
		}
		err = sensor.SaveSettings(settings)
		if err != nil {
			log.Errorf("Cannot save settings for sensor %#v", sensor)
		}
		log.Infof("Settings fetched at %s", time.Now())
		time.Sleep(time.Duration(interval) * time.Second)
	}
}
func TelemetryRoutine(sensor *AwairElementSensor, interval int32) {
	for {
		log := zap.S()
		telemetry, err := sensor.GetTelemetry()
		//log.Infow("Telemetry: %#v", telemetry)
		if err != nil {
			log.Errorf("Error getting telemetry for sensor %#v", sensor)
		}
		err = sensor.SaveTelemetry(telemetry)
		if err != nil {
			log.Infof("Cannot save settings for sensor %#v", sensor)
		}
		log.Debugf("Telemetry fetched at %s", time.Now())
		time.Sleep(time.Duration(interval) * time.Second)
	}
}
