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
			Sensors: []Sensor{{Name: "Sala", Model: "awair_element", Endpoint: "10.0.1.20"}},
			Store: Store{
				Type:     "postgres",
				Postgres: &Postgres{ConnectionString: "host=localhost port=5432 user=postgres password=postgres dbname=postgres sslmode=disable"},
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
