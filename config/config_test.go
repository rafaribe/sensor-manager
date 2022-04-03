package config_test

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"os"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/rafaribe/planetwatch-awair-uploader/config"
)

var _ = Describe("Parse the configuration file contents into a struct", func() {
	When("The file contains a firebase configuration and local awair sensor", func() {
		It("Should parse the file correctly", func() {
			yaml := fixture("firebase-local-awair.yaml")
			config, err := config.UnmarshalConfig(yaml)
			Expect(err).To(BeNil())
			Expect(config.Store.Firebase.SecureCredentialsFile).To(Equal("/tmp/firebase-credentials.json"))
			Expect(len(config.Inputs.LocalAwair)).To(Equal(1))
			Expect(config.Outputs.PlanetWatch).ToNot(BeNil())
			Expect(config.Outputs.PlanetWatch.Endpoint).To(Equal("https://wearableapi.planetwatch.io/api/data/devicedata"))
			Expect(len(config.Outputs.PlanetWatch.AccessToken)).To(BeNumerically(">", 200))
			Expect(len(config.Outputs.PlanetWatch.AccessToken)).To(BeNumerically("==", 240))
			Expect(config.Outputs.PlanetWatch.AccessToken).Should(HavePrefix("eyJ"))
			Expect(len(config.Outputs.PlanetWatch.RefreshToken)).To(BeNumerically(">", 200))
			Expect(len(config.Outputs.PlanetWatch.RefreshToken)).To(BeNumerically("==", 241))
			Expect(config.Outputs.PlanetWatch.RefreshToken).Should(HavePrefix("eyJ"))

		})
	})
})

func fixture(path string) io.Reader {
	log.Println(os.Getwd())
	b, err := ioutil.ReadFile("testdata/" + path)
	if err != nil {
		panic(err)
	}
	return bytes.NewReader(b)
}
