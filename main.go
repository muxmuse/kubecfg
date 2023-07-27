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

package main

import (
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/kubecfg/kubecfg/cmd"
	"github.com/kubecfg/kubecfg/pkg/kubecfg"
	pkgversion "github.com/kubecfg/kubecfg/pkg/version"
)

// Version is overridden using `-X main.version` during release builds
var version = pkgversion.DevVersion

func main() {
	pkgversion.Version = version

	if err := cmd.RootCmd.Execute(); err != nil {
		// PersistentPreRunE may not have been run for early
		// errors, like invalid command line flags.
		logFmt := cmd.NewLogFormatter(log.StandardLogger().Out)
		log.SetFormatter(logFmt)
		log.Error(err.Error())

		switch err {
		case kubecfg.ErrDiffFound:
			os.Exit(10)
		default:
			os.Exit(1)
		}
	}
}
