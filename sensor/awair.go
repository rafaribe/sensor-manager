package sensor

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
	
}