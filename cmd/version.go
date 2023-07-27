// Copyright 2017 The kubecfg authors
//
//
//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.

package cmd

import (
	"fmt"

	jsonnet "github.com/google/go-jsonnet"
	kubecfgVersion "github.com/kubecfg/kubecfg/pkg/version"
	"github.com/spf13/cobra"
	k8sVersion "k8s.io/client-go/pkg/version"
)

func init() {
	RootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version information",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		out := cmd.OutOrStdout()
		fmt.Fprintln(out, "kubecfg version:", kubecfgVersion.Get())
		fmt.Fprintln(out, "jsonnet version:", jsonnet.Version())
		fmt.Fprintln(out, "client-go version:", k8sVersion.Get())
	},
}
