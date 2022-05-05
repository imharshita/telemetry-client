/*
Copyright 2020 The Telemetry Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package reporter

import (
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

func NewReporterCommand(log *zap.SugaredLogger) *cobra.Command {
	cmd := &cobra.Command{
		Args:  cobra.NoArgs,
		Use:   "reporter",
		Short: "Telemetry reporter",
	}

	cmd.AddCommand(
		newStdoutReporterCommand(),
		newHTTPReporterCommand(log),
	)
	return cmd
}
