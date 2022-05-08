package sensor

import (
	"encoding/json"
	"fmt"
	"net/http"

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
}

func Init(c *SensorConfig) *AwairElementSensor {
	sensor := &AwairElementSensor{
		host:              c.Endpoint,
		settingsEndpoint:  "settings/config/data",
		telemetryEndpoint: "air-data/latest",
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
func (s AwairElementSensor) SaveTelemetry(*AwairElementTelemetry) error {
	return nil
}
func (s AwairElementSensor) SaveSettings(*AwairElementSettings) error {

	return nil
}
