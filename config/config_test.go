package config

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestOpenAndReadFile(t *testing.T) {
	tests := map[string]struct {
		input_file string
		want       *Config
	}{
		"name": {input_file: "testdata/awair-postgres.yaml", want: &Config{
			Sensors: []sensor{{Name: "Sala", Model: "awair_element", Endpoint: "10.0.1.20"}},
			Store: store{
				Type:             "postgres",
				ConnectionString: "postgresql://postgres:1234@localhost:5432/postgres",
			},
		}},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, _ := openAndReadFile(tc.input_file)
			diff := cmp.Diff(tc.want, got)
			if diff != "" {
				//t.Fatalf(diff)
			}
		})
	}
}
func fixture(path string) io.Reader {
	log.Println(os.Getwd())
	b, err := ioutil.ReadFile("testdata/" + path)
	if err != nil {
		panic(err)
	}
	return bytes.NewReader(b)
}
