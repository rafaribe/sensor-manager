package sensor

type SensorSettings interface {
}
type Sensor interface {
	Init() any
	GetSettings() (any, error)
	GetTelemetry() (any, error)
	SaveTelemetry() error
}
