package config

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestOpenAndReadFile(t *testing.T) {
	tests := map[string]struct {
		input_file string
		want       *Config
	}{
		"name": {input_file: "testdata/awair-postgres.yaml", want: &Config{
			Sensors: []SensorConfig{{Name: "sala", Model: "awair_element", Endpoint: "10.0.1.20"}},
			Store: Store{
				Type:     "postgres",
				InfluxDb: &Influxdb{Host: "localhost:8086", Token: "8vq-gjtKAPfrXCwaLPN3EeabGjUUCZFTKjumX7t1IkiNtGq_d6I-XqY6wm4iozMVM5qBgCbbK0UQNFyQJBYFDw==", Org: "admin", Bucket: "sensor"},
			},
		}},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, _ := openAndReadFile(tc.input_file)
			diff := cmp.Diff(tc.want, got)
			if diff != "" {
				t.Fatalf(diff)
			}
		})
	}
}
