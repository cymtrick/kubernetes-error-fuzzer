/*
Copyright 2017 The Kubernetes Authors.

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
	"fmt"
	"io"
	"os"

	"github.com/golang/glog"
	"github.com/spf13/cobra"
	"k8s.io/kubernetes/pkg/kubectl/cmd/templates"
	cmdutil "k8s.io/kubernetes/pkg/kubectl/cmd/util"
	"k8s.io/kubernetes/pkg/kubectl/plugins"
	"k8s.io/kubernetes/pkg/util/i18n"
)

var (
	plugin_long = templates.LongDesc(`
		Runs a command-line plugin.

		Plugins are subcommands that are not part of the major command-line distribution 
		and can even be provided by third-parties. Please refer to the documentation and 
		examples for more information about how to install and write your own plugins.`)
)

// NewCmdPlugin creates the command that is the top-level for plugin commands.
func NewCmdPlugin(f cmdutil.Factory, in io.Reader, out, err io.Writer) *cobra.Command {
	// Loads plugins and create commands for each plugin identified
	loadedPlugins, loadErr := f.PluginLoader().Load()
	if loadErr != nil {
		glog.V(1).Infof("Unable to load plugins: %v", loadErr)
	}

	cmd := &cobra.Command{
		Use:   "plugin NAME",
		Short: i18n.T("Runs a command-line plugin"),
		Long:  plugin_long,
		Run: func(cmd *cobra.Command, args []string) {
			if len(loadedPlugins) == 0 {
				cmdutil.CheckErr(fmt.Errorf("no plugins installed."))
			}
			cmdutil.DefaultSubCommandRun(err)(cmd, args)
		},
	}

	if len(loadedPlugins) > 0 {
		pluginRunner := f.PluginRunner()
		for _, p := range loadedPlugins {
			cmd.AddCommand(NewCmdForPlugin(p, pluginRunner, in, out, err))
		}
	}

	return cmd
}

// NewCmdForPlugin creates a command capable of running the provided plugin.
func NewCmdForPlugin(plugin *plugins.Plugin, runner plugins.PluginRunner, in io.Reader, out, errout io.Writer) *cobra.Command {
	if !plugin.IsValid() {
		return nil
	}

	return &cobra.Command{
		Use:     plugin.Name,
		Short:   plugin.ShortDesc,
		Long:    templates.LongDesc(plugin.LongDesc),
		Example: templates.Examples(plugin.Example),
		Run: func(cmd *cobra.Command, args []string) {
			ctx := plugins.RunningContext{
				In:         in,
				Out:        out,
				ErrOut:     errout,
				Args:       args,
				Env:        os.Environ(),
				WorkingDir: plugin.Dir,
			}
			if err := runner.Run(plugin, ctx); err != nil {
				cmdutil.CheckErr(err)
			}
		},
	}
}
