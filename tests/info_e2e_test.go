package tests_test

import (
	"os/exec"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Info E2E", func() {
	It("prints e2e config overrides", func() {
		cmd := exec.Command("task", "run", "--", "info", "--env", "e2e")
		outputBytes, err := cmd.CombinedOutput()
		output := string(outputBytes)
		Expect(err).To(BeNil(), "should run without error, got: %s", output)
		Expect(output).To(ContainSubstring("App Version: 1.0.0-e2e"))
		Expect(output).To(ContainSubstring("Cloud Provider: fakecloud"))
		Expect(output).To(ContainSubstring("Cloud Region: eu-central-1"))
	})
})
