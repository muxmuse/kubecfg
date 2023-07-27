// +build integration

package integration

import (
	"bytes"
	"fmt"
	"os/exec"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("completion", func() {
	var shellArg string
	var env []string
	var output string

	JustBeforeEach(func() {
		outbuf := bytes.Buffer{}

		args := []string{"completion"}
		if shellArg != "" {
			args = append(args, "--shell", shellArg)
		}
		fmt.Fprintf(GinkgoWriter, "Running %q %q\n", *kubecfgBin, args)
		cmd := exec.Command(*kubecfgBin, args...)
		cmd.Env = env
		cmd.Stdout = &outbuf
		cmd.Stderr = GinkgoWriter

		err := cmd.Run()
		Expect(err).NotTo(HaveOccurred())

		output = outbuf.String()
	})

	Context("With --shell=bash", func() {
		BeforeEach(func() {
			shellArg = "bash"
			env = []string{}
		})
		It("should produce bash completion", func() {
			Expect(output).To(MatchRegexp("complete .* kubecfg"))
		})
	})

	Context("With --shell=zsh", func() {
		BeforeEach(func() {
			shellArg = "zsh"
			env = []string{}
		})
		It("should produce zsh completion", func() {
			Expect(output).To(HavePrefix("#compdef kubecfg\n#compdef _kubecfg kubecfg\n"))
		})
	})

	Context("With SHELL=/bin/bash4", func() {
		BeforeEach(func() {
			shellArg = ""
			env = append(env, "SHELL=/bin/bash4")
		})
		It("should produce bash completion", func() {
			Expect(output).To(MatchRegexp("complete .* kubecfg"))
		})
	})
})
