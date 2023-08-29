package sensor

import (
	"encoding/json"
	"fmt"
	"net/http"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	. "github.com/rafaribe/sensor-manager/config"
)

// Awair element payload
type AwairElementTelemetry struct {
	Timestamp      string  `json:"timestamp"`
	Score          int64   `json:"score"`
	DewPoint       float64 `json:"dew_point"`
	Temp           float64 `json:"temp"`
	Humid          float64 `json:"humid"`
	AbsHumid       float64 `json:"abs_humid"`
	Co2            int64   `json:"co2"`
	Co2Est         int64   `json:"co2_est"`
	Co2EstBaseline int64   `json:"co2_est_baseline"`
	Voc            int64   `json:"voc"`
	VocBaseline    int64   `json:"voc_baseline"`
	VocH2Raw       int64   `json:"voc_h2_raw"`
	VocEthanolRaw  int64   `json:"voc_ethanol_raw"`
	Pm25           int64   `json:"pm25"`
	Pm10Est        int64   `json:"pm10_est"`
}

// Awair element settings
type AwairElementSettings struct {
	DeviceUUID string `json:"device_uuid"`
	WifiMac    string `json:"wifi_mac"`
	Ssid       string `json:"ssid"`
	IP         string `json:"ip"`
	Netmask    string `json:"netmask"`
	Gateway    string `json:"gateway"`
	FwVersion  string `json:"fw_version"`
	Timezone   string `json:"timezone"`
	Display    string `json:"display"`
	Led        struct {
		Mode       string `json:"mode"`
		Brightness int    `json:"brightness"`
	} `json:"led"`
	VocFeatureSet int `json:"voc_feature_set"`
}

type AwairElementSensor struct {
	host              string
	settingsEndpoint  string
	telemetryEndpoint string
	storageClient     influxdb2.Client
	bucket            string
	org               string
}

func Init(c *SensorConfig, s *Store) *AwairElementSensor {
	sensor := &AwairElementSensor{
		host:              c.Endpoint,
		settingsEndpoint:  "settings/config/data",
		telemetryEndpoint: "air-data/latest",
		storageClient:     influxdb2.NewClient(s.InfluxDb.Host, s.InfluxDb.Token),
		bucket:            s.InfluxDb.Bucket,
		org:               s.InfluxDb.Org,
	}
	return sensor
}
func (s *AwairElementSensor) GetSettings() (*AwairElementSettings, error) {
	res, err := http.Get(fmt.Sprintf("http://%s/%s", s.host, s.settingsEndpoint))
	if err != nil {
		return nil, err
	}
	set := new(AwairElementSettings)
	if err := json.NewDecoder(res.Body).Decode(set); err != nil {
		return nil, err
	}
	return set, nil
}
func (s AwairElementSensor) GetTelemetry() (*AwairElementTelemetry, error) {
	res, err := http.Get(fmt.Sprintf("http://%s/%s", s.host, s.telemetryEndpoint))
	if err != nil {
		return nil, err
	}
	telemetry := new(AwairElementTelemetry)
	if err := json.NewDecoder(res.Body).Decode(telemetry); err != nil {
		return nil, err
	}
	return telemetry, nil
}
func (s AwairElementSensor) SaveTelemetry(telemetry *AwairElementTelemetry) {
	writeAPI := s.storageClient.WriteAPI(s.org, s.bucket)
	defer writeAPI.Flush()

	point := influxdb2.NewPointWithMeasurement("awair_telemetry").
		AddField("score", telemetry.Score).
		AddField("dew_point", telemetry.DewPoint).
		AddField("temp", telemetry.Temp).
		AddField("humid", telemetry.Humid).
		AddField("abs_humid", telemetry.AbsHumid).
		AddField("co2", telemetry.Co2).
		AddField("co2_est", telemetry.Co2Est).
		AddField("co2_est_baseline", telemetry.Co2EstBaseline).
		AddField("voc", telemetry.Voc).
		AddField("voc_baseline", telemetry.VocBaseline).
		AddField("voc_h2_raw", telemetry.VocH2Raw).
		AddField("voc_ethanol_raw", telemetry.VocEthanolRaw).
		AddField("pm25", telemetry.Pm25).
		AddField("pm10_est", telemetry.Pm10Est)

	// Write the point to the database
	writeAPI.WritePoint(point)
}
func (s AwairElementSensor) SaveSettings(settings *AwairElementSettings) {

	writeAPI := s.storageClient.WriteAPI(s.org, s.bucket)
	defer writeAPI.Flush()

	point := influxdb2.NewPointWithMeasurement("awair_settings").
		AddTag("device_uuid", settings.DeviceUUID).
		AddTag("wifi_mac", settings.WifiMac).
		AddField("ssid", settings.Ssid).
		AddField("ip", settings.IP).
		AddField("netmask", settings.Netmask).
		AddField("gateway", settings.Gateway).
		AddField("fw_version", settings.FwVersion).
		AddField("timezone", settings.Timezone).
		AddField("display", settings.Display).
		AddField("led_mode", settings.Led.Mode).
		AddField("led_brightness", settings.Led.Brightness).
		AddField("voc_feature_set", settings.VocFeatureSet)

	writeAPI.WritePoint(point)
}
