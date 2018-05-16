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
	"os"
	"os/exec"
	"syscall"

	"github.com/golang/glog"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"k8s.io/kubernetes/pkg/kubectl/cmd/templates"
	cmdutil "k8s.io/kubernetes/pkg/kubectl/cmd/util"
	"k8s.io/kubernetes/pkg/kubectl/genericclioptions"
	"k8s.io/kubernetes/pkg/kubectl/plugins"
	"k8s.io/kubernetes/pkg/kubectl/util/i18n"
)

var (
	plugin_long = templates.LongDesc(`
		Runs a command-line plugin.

		Plugins are subcommands that are not part of the major command-line distribution 
		and can even be provided by third-parties. Please refer to the documentation and 
		examples for more information about how to install and write your own plugins.`)
)

// NewCmdPlugin creates the command that is the top-level for plugin commands.
func NewCmdPlugin(f cmdutil.Factory, streams genericclioptions.IOStreams) *cobra.Command {
	// Loads plugins and create commands for each plugin identified
	loadedPlugins, loadErr := pluginLoader().Load()
	if loadErr != nil {
		glog.V(1).Infof("Unable to load plugins: %v", loadErr)
	}

	cmd := &cobra.Command{
		Use: "plugin NAME",
		DisableFlagsInUseLine: true,
		Short: i18n.T("Runs a command-line plugin"),
		Long:  plugin_long,
		Run: func(cmd *cobra.Command, args []string) {
			if len(loadedPlugins) == 0 {
				cmdutil.CheckErr(fmt.Errorf("no plugins installed."))
			}
			cmdutil.DefaultSubCommandRun(streams.ErrOut)(cmd, args)
		},
	}

	if len(loadedPlugins) > 0 {
		pluginRunner := pluginRunner()
		for _, p := range loadedPlugins {
			cmd.AddCommand(NewCmdForPlugin(f, p, pluginRunner, streams))
		}
	}

	return cmd
}

// NewCmdForPlugin creates a command capable of running the provided plugin.
func NewCmdForPlugin(f cmdutil.Factory, plugin *plugins.Plugin, runner plugins.PluginRunner, streams genericclioptions.IOStreams) *cobra.Command {
	if !plugin.IsValid() {
		return nil
	}

	cmd := &cobra.Command{
		Use:     plugin.Name,
		Short:   plugin.ShortDesc,
		Long:    templates.LongDesc(plugin.LongDesc),
		Example: templates.Examples(plugin.Example),
		Run: func(cmd *cobra.Command, args []string) {
			if len(plugin.Command) == 0 {
				cmdutil.DefaultSubCommandRun(streams.ErrOut)(cmd, args)
				return
			}

			envProvider := &plugins.MultiEnvProvider{
				&plugins.PluginCallerEnvProvider{},
				&plugins.OSEnvProvider{},
				&plugins.PluginDescriptorEnvProvider{
					Plugin: plugin,
				},
				&flagsPluginEnvProvider{
					cmd: cmd,
				},
				&factoryAttrsPluginEnvProvider{
					factory: f,
				},
			}

			runningContext := plugins.RunningContext{
				IOStreams:   streams,
				Args:        args,
				EnvProvider: envProvider,
				WorkingDir:  plugin.Dir,
			}

			if err := runner.Run(plugin, runningContext); err != nil {
				if exiterr, ok := err.(*exec.ExitError); ok {
					// check for (and exit with) the correct exit code
					// from a failed plugin command execution
					if status, ok := exiterr.Sys().(syscall.WaitStatus); ok {
						fmt.Fprintf(streams.ErrOut, "error: %v\n", err)
						os.Exit(status.ExitStatus())
					}
				}

				cmdutil.CheckErr(err)
			}
		},
	}

	for _, flag := range plugin.Flags {
		cmd.Flags().StringP(flag.Name, flag.Shorthand, flag.DefValue, flag.Desc)
	}

	for _, childPlugin := range plugin.Tree {
		cmd.AddCommand(NewCmdForPlugin(f, childPlugin, runner, streams))
	}

	return cmd
}

type flagsPluginEnvProvider struct {
	cmd *cobra.Command
}

func (p *flagsPluginEnvProvider) Env() (plugins.EnvList, error) {
	globalPrefix := "KUBECTL_PLUGINS_GLOBAL_FLAG_"
	env := plugins.EnvList{}
	p.cmd.InheritedFlags().VisitAll(func(flag *pflag.Flag) {
		env = append(env, plugins.FlagToEnv(flag, globalPrefix))
	})
	localPrefix := "KUBECTL_PLUGINS_LOCAL_FLAG_"
	p.cmd.LocalFlags().VisitAll(func(flag *pflag.Flag) {
		env = append(env, plugins.FlagToEnv(flag, localPrefix))
	})
	return env, nil
}

type factoryAttrsPluginEnvProvider struct {
	factory cmdutil.Factory
}

func (p *factoryAttrsPluginEnvProvider) Env() (plugins.EnvList, error) {
	cmdNamespace, _, err := p.factory.DefaultNamespace()
	if err != nil {
		return plugins.EnvList{}, err
	}
	return plugins.EnvList{
		plugins.Env{N: "KUBECTL_PLUGINS_CURRENT_NAMESPACE", V: cmdNamespace},
	}, nil
}

// pluginLoader loads plugins from a path set by the KUBECTL_PLUGINS_PATH env var.
// If this env var is not set, it defaults to
//   "~/.kube/plugins", plus
//  "./kubectl/plugins" directory under the "data dir" directory specified by the XDG
// system directory structure spec for the given platform.
func pluginLoader() plugins.PluginLoader {
	if len(os.Getenv("KUBECTL_PLUGINS_PATH")) > 0 {
		return plugins.KubectlPluginsPathPluginLoader()
	}
	return plugins.TolerantMultiPluginLoader{
		plugins.XDGDataDirsPluginLoader(),
		plugins.UserDirPluginLoader(),
	}
}

func pluginRunner() plugins.PluginRunner {
	return &plugins.ExecPluginRunner{}
}
