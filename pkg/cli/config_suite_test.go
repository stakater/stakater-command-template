package cli

import (
	"os"
	"path/filepath"
	"stakater-cmd/internal/config"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/spf13/pflag"
)

func TestCLI(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "CLI Suite")
}

var _ = Describe("Config", func() {
	var (
		tempDir   string
		configDir string
		base      string
	)

	writeTempConfig := func(content, filename string) {
		f, err := os.Create(filename)
		Expect(err).To(BeNil())
		_, err = f.WriteString(content)
		Expect(err).To(BeNil())
		Expect(f.Close()).To(BeNil())
		DeferCleanup(func() { os.Remove(filename) })
	}

	BeforeEach(func() {
		base = `
app:
  name:    base
  version: 1.0.0
cloud:
  provider: aws
  region:   us-west-2
`
		tempDir = GinkgoT().TempDir()
		configDir = filepath.Join(tempDir, "configs")
		Expect(os.MkdirAll(configDir, 0755)).To(BeNil())
		writeTempConfig(base, filepath.Join(configDir, "config.yaml"))
		oldWd, _ := os.Getwd()
		DeferCleanup(func() { os.Chdir(oldWd) })
		os.Chdir(tempDir)
	})

	It("should override config with environment-specific file", func() {
		override := `
cloud:
  region: eu-central-1
`
		writeTempConfig(override, filepath.Join(configDir, "config.test.yaml"))
		os.Args = []string{"cmd", "--env", "test"}

		v, err := config.InitConfig()
		Expect(err).To(BeNil())
		cfg, err := config.GetConfig(v)
		Expect(err).To(BeNil())

		Expect(cfg.Cloud.Region).To(Equal("eu-central-1"))
		Expect(cfg.Cloud.Provider).To(Equal("aws"))
		Expect(v.GetString("cloud.region")).To(Equal("eu-central-1"))
		Expect(v.GetString("cloud.provider")).To(Equal("aws"))
	})

	It("should override config with environment variables", func() {
		os.Setenv("CLOUD_REGION", "env-region")
		DeferCleanup(func() { os.Unsetenv("CLOUD_REGION") })
		os.Args = []string{"cmd"}

		v, err := config.InitConfig()
		Expect(err).To(BeNil())
		cfg, err := config.GetConfig(v)
		Expect(err).To(BeNil())

		Expect(cfg.Cloud.Region).To(Equal("env-region"))
		Expect(v.GetString("cloud.region")).To(Equal("env-region"))
	})

	It("should override config with command line arguments", func() {
		fs := pflag.NewFlagSet("test", pflag.ContinueOnError)
		fs.String("region", "", "")
		fs.Parse([]string{"--region", "arg-region"})
		os.Args = []string{"cmd"}

		v, err := config.InitConfig()
		Expect(err).To(BeNil())

		v.BindPFlag("cloud.region", fs.Lookup("region"))

		cfg, err := config.GetConfig(v)
		Expect(err).To(BeNil())

		Expect(cfg.Cloud.Region).To(Equal("arg-region"))
		Expect(v.GetString("cloud.region")).To(Equal("arg-region"))
	})
})
