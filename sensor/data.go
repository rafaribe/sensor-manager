package sensor

type SensorSettings interface {
}
type Sensor interface {
	GetSettings() (any, error)
	GetTelemetry() (any, error)
	SaveTelemetry() error
}
