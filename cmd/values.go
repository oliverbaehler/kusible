/*
Copyright © 2021 Bedag Informatik AG

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

package cmd

import (
	"github.com/bedag/kusible/pkg/printer"
	"github.com/bedag/kusible/pkg/values"
	"github.com/spf13/cobra"

	log "github.com/sirupsen/logrus"
)

func newValuesCmd(c *Cli) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "values GROUP ...",
		Short: "Compile values for a list of groups",
		Long: `Use the given groups to compile a single values yaml file.
	The groups are priorized from least to most specific.
	Values of groups of higher priorities override values
	of groups with lower priorities.`,
		Args:                  cobra.MinimumNArgs(1),
		TraverseChildren:      true,
		DisableFlagsInUseLine: true,
		RunE:                  c.wrap(runValues),
	}
	addEjsonFlags(cmd)
	addEvalFlags(cmd)
	addGroupsFlags(cmd)
	addOutputFlags(cmd)

	return cmd
}

func runValues(c *Cli, cmd *cobra.Command, args []string) error {
	groups := args
	groupVarsDir := c.viper.GetString("group-vars-dir")
	skipEval := c.viper.GetBool("skip-eval")

	ejsonSettings := getEjsonSettings(c)

	values, err := values.New(groupVarsDir, groups, skipEval, ejsonSettings)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err.Error(),
		}).Error("Failed to compile group vars.")
		return err
	}

	printFn := func(fields []string) map[string]interface{} {
		all := values.Map()
		if len(fields) < 1 {
			return all
		}

		result := make(map[string]interface{}, len(fields))

		for _, field := range fields {
			if val, ok := all[field]; ok {
				result[field] = val
			}
		}
		return result
	}

	printerQueue := printer.Queue{printer.NewJob(printFn)}

	return c.output(printerQueue)
}
