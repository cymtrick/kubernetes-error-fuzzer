// Copyright © 2013 Steve Francia <spf@spf13.com>.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package cobra is a commander providing a simple interface to create powerful modern CLI interfaces.
// In addition to providing an interface, Cobra simultaneously provides a controller to organize your application code.
package cobra

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"

	flag "github.com/spf13/pflag"
)

// FParseErrWhitelist configures Flag parse errors to be ignored
type FParseErrWhitelist flag.ParseErrorsWhitelist

// Command is just that, a command for your application.
// E.g.  'go run ...' - 'run' is the command. Cobra requires
// you to define the usage and description as part of your command
// definition to ensure usability.
type Command struct {
	// Use is the one-line usage message.
	// Recommended syntax is as follow:
	//   [ ] identifies an optional argument. Arguments that are not enclosed in brackets are required.
	//   ... indicates that you can specify multiple values for the previous argument.
	//   |   indicates mutually exclusive information. You can use the argument to the left of the separator or the
	//       argument to the right of the separator. You cannot use both arguments in a single use of the command.
	//   { } delimits a set of mutually exclusive arguments when one of the arguments is required. If the arguments are
	//       optional, they are enclosed in brackets ([ ]).
	// Example: add [-F file | -D dir]... [-f format] profile
	Use string

	// Aliases is an array of aliases that can be used instead of the first word in Use.
	Aliases []string

	// SuggestFor is an array of command names for which this command will be suggested -
	// similar to aliases but only suggests.
	SuggestFor []string

	// Short is the short description shown in the 'help' output.
	Short string

	// Long is the long message shown in the 'help <this-command>' output.
	Long string

	// Example is examples of how to use the command.
	Example string

	// ValidArgs is list of all valid non-flag arguments that are accepted in bash completions
	ValidArgs []string
	// ValidArgsFunction is an optional function that provides valid non-flag arguments for bash completion.
	// It is a dynamic version of using ValidArgs.
	// Only one of ValidArgs and ValidArgsFunction can be used for a command.
	ValidArgsFunction func(cmd *Command, args []string, toComplete string) ([]string, ShellCompDirective)

	// Expected arguments
	Args PositionalArgs

	// ArgAliases is List of aliases for ValidArgs.
	// These are not suggested to the user in the bash completion,
	// but accepted if entered manually.
	ArgAliases []string

	// BashCompletionFunction is custom functions used by the bash autocompletion generator.
	BashCompletionFunction string

	// Deprecated defines, if this command is deprecated and should print this string when used.
	Deprecated string

	// Hidden defines, if this command is hidden and should NOT show up in the list of available commands.
	Hidden bool

	// Annotations are key/value pairs that can be used by applications to identify or
	// group commands.
	Annotations map[string]string

	// Version defines the version for this command. If this value is non-empty and the command does not
	// define a "version" flag, a "version" boolean flag will be added to the command and, if specified,
	// will print content of the "Version" variable. A shorthand "v" flag will also be added if the
	// command does not define one.
	Version string

	// The *Run functions are executed in the following order:
	//   * PersistentPreRun()
	//   * PreRun()
	//   * Run()
	//   * PostRun()
	//   * PersistentPostRun()
	// All functions get the same args, the arguments after the command name.
	//
	// PersistentPreRun: children of this command will inherit and execute.
	PersistentPreRun func(cmd *Command, args []string)
	// PersistentPreRunE: PersistentPreRun but returns an error.
	PersistentPreRunE func(cmd *Command, args []string) error
	// PreRun: children of this command will not inherit.
	PreRun func(cmd *Command, args []string)
	// PreRunE: PreRun but returns an error.
	PreRunE func(cmd *Command, args []string) error
	// Run: Typically the actual work function. Most commands will only implement this.
	Run func(cmd *Command, args []string)
	// RunE: Run but returns an error.
	RunE func(cmd *Command, args []string) error
	// PostRun: run after the Run command.
	PostRun func(cmd *Command, args []string)
	// PostRunE: PostRun but returns an error.
	PostRunE func(cmd *Command, args []string) error
	// PersistentPostRun: children of this command will inherit and execute after PostRun.
	PersistentPostRun func(cmd *Command, args []string)
	// PersistentPostRunE: PersistentPostRun but returns an error.
	PersistentPostRunE func(cmd *Command, args []string) error

	// SilenceErrors is an option to quiet errors down stream.
	SilenceErrors bool

	// SilenceUsage is an option to silence usage when an error occurs.
	SilenceUsage bool

	// DisableFlagParsing disables the flag parsing.
	// If this is true all flags will be passed to the command as arguments.
	DisableFlagParsing bool

	// DisableAutoGenTag defines, if gen tag ("Auto generated by spf13/cobra...")
	// will be printed by generating docs for this command.
	DisableAutoGenTag bool

	// DisableFlagsInUseLine will disable the addition of [flags] to the usage
	// line of a command when printing help or generating docs
	DisableFlagsInUseLine bool

	// DisableSuggestions disables the suggestions based on Levenshtein distance
	// that go along with 'unknown command' messages.
	DisableSuggestions bool
	// SuggestionsMinimumDistance defines minimum levenshtein distance to display suggestions.
	// Must be > 0.
	SuggestionsMinimumDistance int

	// TraverseChildren parses flags on all parents before executing child command.
	TraverseChildren bool

	// FParseErrWhitelist flag parse errors to be ignored
	FParseErrWhitelist FParseErrWhitelist

	ctx context.Context

	// commands is the list of commands supported by this program.
	commands []*Command
	// parent is a parent command for this command.
	parent *Command
	// Max lengths of commands' string lengths for use in padding.
	commandsMaxUseLen         int
	commandsMaxCommandPathLen int
	commandsMaxNameLen        int
	// commandsAreSorted defines, if command slice are sorted or not.
	commandsAreSorted bool
	// commandCalledAs is the name or alias value used to call this command.
	commandCalledAs struct {
		name   string
		called bool
	}

	// args is actual args parsed from flags.
	args []string
	// flagErrorBuf contains all error messages from pflag.
	flagErrorBuf *bytes.Buffer
	// flags is full set of flags.
	flags *flag.FlagSet
	// pflags contains persistent flags.
	pflags *flag.FlagSet
	// lflags contains local flags.
	lflags *flag.FlagSet
	// iflags contains inherited flags.
	iflags *flag.FlagSet
	// parentsPflags is all persistent flags of cmd's parents.
	parentsPflags *flag.FlagSet
	// globNormFunc is the global normalization function
	// that we can use on every pflag set and children commands
	globNormFunc func(f *flag.FlagSet, name string) flag.NormalizedName

	// usageFunc is usage func defined by user.
	usageFunc func(*Command) error
	// usageTemplate is usage template defined by user.
	usageTemplate string
	// flagErrorFunc is func defined by user and it's called when the parsing of
	// flags returns an error.
	flagErrorFunc func(*Command, error) error
	// helpTemplate is help template defined by user.
	helpTemplate string
	// helpFunc is help func defined by user.
	helpFunc func(*Command, []string)
	// helpCommand is command with usage 'help'. If it's not defined by user,
	// cobra uses default help command.
	helpCommand *Command
	// versionTemplate is the version template defined by user.
	versionTemplate string

	// inReader is a reader defined by the user that replaces stdin
	inReader io.Reader
	// outWriter is a writer defined by the user that replaces stdout
	outWriter io.Writer
	// errWriter is a writer defined by the user that replaces stderr
	errWriter io.Writer
}

// Context returns underlying command context. If command wasn't
// executed with ExecuteContext Context returns Background context.
func (c *Command) Context() context.Context {
	return c.ctx
}

// SetArgs sets arguments for the command. It is set to os.Args[1:] by default, if desired, can be overridden
// particularly useful when testing.
func (c *Command) SetArgs(a []string) {
	c.args = a
}

// SetOutput sets the destination for usage and error messages.
// If output is nil, os.Stderr is used.
// Deprecated: Use SetOut and/or SetErr instead
func (c *Command) SetOutput(output io.Writer) {
	c.outWriter = output
	c.errWriter = output
}

// SetOut sets the destination for usage messages.
// If newOut is nil, os.Stdout is used.
func (c *Command) SetOut(newOut io.Writer) {
	c.outWriter = newOut
}

// SetErr sets the destination for error messages.
// If newErr is nil, os.Stderr is used.
func (c *Command) SetErr(newErr io.Writer) {
	c.errWriter = newErr
}

// SetIn sets the source for input data
// If newIn is nil, os.Stdin is used.
func (c *Command) SetIn(newIn io.Reader) {
	c.inReader = newIn
}

// SetUsageFunc sets usage function. Usage can be defined by application.
func (c *Command) SetUsageFunc(f func(*Command) error) {
	c.usageFunc = f
}

// SetUsageTemplate sets usage template. Can be defined by Application.
func (c *Command) SetUsageTemplate(s string) {
	c.usageTemplate = s
}

// SetFlagErrorFunc sets a function to generate an error when flag parsing
// fails.
func (c *Command) SetFlagErrorFunc(f func(*Command, error) error) {
	c.flagErrorFunc = f
}

// SetHelpFunc sets help function. Can be defined by Application.
func (c *Command) SetHelpFunc(f func(*Command, []string)) {
	c.helpFunc = f
}

// SetHelpCommand sets help command.
func (c *Command) SetHelpCommand(cmd *Command) {
	c.helpCommand = cmd
}

// SetHelpTemplate sets help template to be used. Application can use it to set custom template.
func (c *Command) SetHelpTemplate(s string) {
	c.helpTemplate = s
}

// SetVersionTemplate sets version template to be used. Application can use it to set custom template.
func (c *Command) SetVersionTemplate(s string) {
	c.versionTemplate = s
}

// SetGlobalNormalizationFunc sets a normalization function to all flag sets and also to child commands.
// The user should not have a cyclic dependency on commands.
func (c *Command) SetGlobalNormalizationFunc(n func(f *flag.FlagSet, name string) flag.NormalizedName) {
	c.Flags().SetNormalizeFunc(n)
	c.PersistentFlags().SetNormalizeFunc(n)
	c.globNormFunc = n

	for _, command := range c.commands {
		command.SetGlobalNormalizationFunc(n)
	}
}

// OutOrStdout returns output to stdout.
func (c *Command) OutOrStdout() io.Writer {
	return c.getOut(os.Stdout)
}

// OutOrStderr returns output to stderr
func (c *Command) OutOrStderr() io.Writer {
	return c.getOut(os.Stderr)
}

// ErrOrStderr returns output to stderr
func (c *Command) ErrOrStderr() io.Writer {
	return c.getErr(os.Stderr)
}

// InOrStdin returns input to stdin
func (c *Command) InOrStdin() io.Reader {
	return c.getIn(os.Stdin)
}

func (c *Command) getOut(def io.Writer) io.Writer {
	if c.outWriter != nil {
		return c.outWriter
	}
	if c.HasParent() {
		return c.parent.getOut(def)
	}
	return def
}

func (c *Command) getErr(def io.Writer) io.Writer {
	if c.errWriter != nil {
		return c.errWriter
	}
	if c.HasParent() {
		return c.parent.getErr(def)
	}
	return def
}

func (c *Command) getIn(def io.Reader) io.Reader {
	if c.inReader != nil {
		return c.inReader
	}
	if c.HasParent() {
		return c.parent.getIn(def)
	}
	return def
}

// UsageFunc returns either the function set by SetUsageFunc for this command
// or a parent, or it returns a default usage function.
func (c *Command) UsageFunc() (f func(*Command) error) {
	if c.usageFunc != nil {
		return c.usageFunc
	}
	if c.HasParent() {
		return c.Parent().UsageFunc()
	}
	return func(c *Command) error {
		c.mergePersistentFlags()
		err := tmpl(c.OutOrStderr(), c.UsageTemplate(), c)
		if err != nil {
			c.PrintErrln(err)
		}
		return err
	}
}

// Usage puts out the usage for the command.
// Used when a user provides invalid input.
// Can be defined by user by overriding UsageFunc.
func (c *Command) Usage() error {
	return c.UsageFunc()(c)
}

// HelpFunc returns either the function set by SetHelpFunc for this command
// or a parent, or it returns a function with default help behavior.
func (c *Command) HelpFunc() func(*Command, []string) {
	if c.helpFunc != nil {
		return c.helpFunc
	}
	if c.HasParent() {
		return c.Parent().HelpFunc()
	}
	return func(c *Command, a []string) {
		c.mergePersistentFlags()
		// The help should be sent to stdout
		// See https://github.com/spf13/cobra/issues/1002
		err := tmpl(c.OutOrStdout(), c.HelpTemplate(), c)
		if err != nil {
			c.PrintErrln(err)
		}
	}
}

// Help puts out the help for the command.
// Used when a user calls help [command].
// Can be defined by user by overriding HelpFunc.
func (c *Command) Help() error {
	c.HelpFunc()(c, []string{})
	return nil
}

// UsageString returns usage string.
func (c *Command) UsageString() string {
	// Storing normal writers
	tmpOutput := c.outWriter
	tmpErr := c.errWriter

	bb := new(bytes.Buffer)
	c.outWriter = bb
	c.errWriter = bb

	c.Usage()

	// Setting things back to normal
	c.outWriter = tmpOutput
	c.errWriter = tmpErr

	return bb.String()
}

// FlagErrorFunc returns either the function set by SetFlagErrorFunc for this
// command or a parent, or it returns a function which returns the original
// error.
func (c *Command) FlagErrorFunc() (f func(*Command, error) error) {
	if c.flagErrorFunc != nil {
		return c.flagErrorFunc
	}

	if c.HasParent() {
		return c.parent.FlagErrorFunc()
	}
	return func(c *Command, err error) error {
		return err
	}
}

var minUsagePadding = 25

// UsagePadding return padding for the usage.
func (c *Command) UsagePadding() int {
	if c.parent == nil || minUsagePadding > c.parent.commandsMaxUseLen {
		return minUsagePadding
	}
	return c.parent.commandsMaxUseLen
}

var minCommandPathPadding = 11

// CommandPathPadding return padding for the command path.
func (c *Command) CommandPathPadding() int {
	if c.parent == nil || minCommandPathPadding > c.parent.commandsMaxCommandPathLen {
		return minCommandPathPadding
	}
	return c.parent.commandsMaxCommandPathLen
}

var minNamePadding = 11

// NamePadding returns padding for the name.
func (c *Command) NamePadding() int {
	if c.parent == nil || minNamePadding > c.parent.commandsMaxNameLen {
		return minNamePadding
	}
	return c.parent.commandsMaxNameLen
}

// UsageTemplate returns usage template for the command.
func (c *Command) UsageTemplate() string {
	if c.usageTemplate != "" {
		return c.usageTemplate
	}

	if c.HasParent() {
		return c.parent.UsageTemplate()
	}
	return `Usage:{{if .Runnable}}
  {{.UseLine}}{{end}}{{if .HasAvailableSubCommands}}
  {{.CommandPath}} [command]{{end}}{{if gt (len .Aliases) 0}}

Aliases:
  {{.NameAndAliases}}{{end}}{{if .HasExample}}

Examples:
{{.Example}}{{end}}{{if .HasAvailableSubCommands}}

Available Commands:{{range .Commands}}{{if (or .IsAvailableCommand (eq .Name "help"))}}
  {{rpad .Name .NamePadding }} {{.Short}}{{end}}{{end}}{{end}}{{if .HasAvailableLocalFlags}}

Flags:
{{.LocalFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}{{if .HasAvailableInheritedFlags}}

Global Flags:
{{.InheritedFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}{{if .HasHelpSubCommands}}

Additional help topics:{{range .Commands}}{{if .IsAdditionalHelpTopicCommand}}
  {{rpad .CommandPath .CommandPathPadding}} {{.Short}}{{end}}{{end}}{{end}}{{if .HasAvailableSubCommands}}

Use "{{.CommandPath}} [command] --help" for more information about a command.{{end}}
`
}

// HelpTemplate return help template for the command.
func (c *Command) HelpTemplate() string {
	if c.helpTemplate != "" {
		return c.helpTemplate
	}

	if c.HasParent() {
		return c.parent.HelpTemplate()
	}
	return `{{with (or .Long .Short)}}{{. | trimTrailingWhitespaces}}

{{end}}{{if or .Runnable .HasSubCommands}}{{.UsageString}}{{end}}`
}

// VersionTemplate return version template for the command.
func (c *Command) VersionTemplate() string {
	if c.versionTemplate != "" {
		return c.versionTemplate
	}

	if c.HasParent() {
		return c.parent.VersionTemplate()
	}
	return `{{with .Name}}{{printf "%s " .}}{{end}}{{printf "version %s" .Version}}
`
}

func hasNoOptDefVal(name string, fs *flag.FlagSet) bool {
	flag := fs.Lookup(name)
	if flag == nil {
		return false
	}
	return flag.NoOptDefVal != ""
}

func shortHasNoOptDefVal(name string, fs *flag.FlagSet) bool {
	if len(name) == 0 {
		return false
	}

	flag := fs.ShorthandLookup(name[:1])
	if flag == nil {
		return false
	}
	return flag.NoOptDefVal != ""
}

func stripFlags(args []string, c *Command) []string {
	if len(args) == 0 {
		return args
	}
	c.mergePersistentFlags()

	commands := []string{}
	flags := c.Flags()

Loop:
	for len(args) > 0 {
		s := args[0]
		args = args[1:]
		switch {
		case s == "--":
			// "--" terminates the flags
			break Loop
		case strings.HasPrefix(s, "--") && !strings.Contains(s, "=") && !hasNoOptDefVal(s[2:], flags):
			// If '--flag arg' then
			// delete arg from args.
			fallthrough // (do the same as below)
		case strings.HasPrefix(s, "-") && !strings.Contains(s, "=") && len(s) == 2 && !shortHasNoOptDefVal(s[1:], flags):
			// If '-f arg' then
			// delete 'arg' from args or break the loop if len(args) <= 1.
			if len(args) <= 1 {
				break Loop
			} else {
				args = args[1:]
				continue
			}
		case s != "" && !strings.HasPrefix(s, "-"):
			commands = append(commands, s)
		}
	}

	return commands
}

// argsMinusFirstX removes only the first x from args.  Otherwise, commands that look like
// openshift admin policy add-role-to-user admin my-user, lose the admin argument (arg[4]).
func argsMinusFirstX(args []string, x string) []string {
	for i, y := range args {
		if x == y {
			ret := []string{}
			ret = append(ret, args[:i]...)
			ret = append(ret, args[i+1:]...)
			return ret
		}
	}
	return args
}

func isFlagArg(arg string) bool {
	return ((len(arg) >= 3 && arg[1] == '-') ||
		(len(arg) >= 2 && arg[0] == '-' && arg[1] != '-'))
}

// Find the target command given the args and command tree
// Meant to be run on the highest node. Only searches down.
func (c *Command) Find(args []string) (*Command, []string, error) {
	var innerfind func(*Command, []string) (*Command, []string)

	innerfind = func(c *Command, innerArgs []string) (*Command, []string) {
		argsWOflags := stripFlags(innerArgs, c)
		if len(argsWOflags) == 0 {
			return c, innerArgs
		}
		nextSubCmd := argsWOflags[0]

		cmd := c.findNext(nextSubCmd)
		if cmd != nil {
			return innerfind(cmd, argsMinusFirstX(innerArgs, nextSubCmd))
		}
		return c, innerArgs
	}

	commandFound, a := innerfind(c, args)
	if commandFound.Args == nil {
		return commandFound, a, legacyArgs(commandFound, stripFlags(a, commandFound))
	}
	return commandFound, a, nil
}

func (c *Command) findSuggestions(arg string) string {
	if c.DisableSuggestions {
		return ""
	}
	if c.SuggestionsMinimumDistance <= 0 {
		c.SuggestionsMinimumDistance = 2
	}
	suggestionsString := ""
	if suggestions := c.SuggestionsFor(arg); len(suggestions) > 0 {
		suggestionsString += "\n\nDid you mean this?\n"
		for _, s := range suggestions {
			suggestionsString += fmt.Sprintf("\t%v\n", s)
		}
	}
	return suggestionsString
}

func (c *Command) findNext(next string) *Command {
	matches := make([]*Command, 0)
	for _, cmd := range c.commands {
		if cmd.Name() == next || cmd.HasAlias(next) {
			cmd.commandCalledAs.name = next
			return cmd
		}
		if EnablePrefixMatching && cmd.hasNameOrAliasPrefix(next) {
			matches = append(matches, cmd)
		}
	}

	if len(matches) == 1 {
		return matches[0]
	}

	return nil
}

// Traverse the command tree to find the command, and parse args for
// each parent.
func (c *Command) Traverse(args []string) (*Command, []string, error) {
	flags := []string{}
	inFlag := false

	for i, arg := range args {
		switch {
		// A long flag with a space separated value
		case strings.HasPrefix(arg, "--") && !strings.Contains(arg, "="):
			// TODO: this isn't quite right, we should really check ahead for 'true' or 'false'
			inFlag = !hasNoOptDefVal(arg[2:], c.Flags())
			flags = append(flags, arg)
			continue
		// A short flag with a space separated value
		case strings.HasPrefix(arg, "-") && !strings.Contains(arg, "=") && len(arg) == 2 && !shortHasNoOptDefVal(arg[1:], c.Flags()):
			inFlag = true
			flags = append(flags, arg)
			continue
		// The value for a flag
		case inFlag:
			inFlag = false
			flags = append(flags, arg)
			continue
		// A flag without a value, or with an `=` separated value
		case isFlagArg(arg):
			flags = append(flags, arg)
			continue
		}

		cmd := c.findNext(arg)
		if cmd == nil {
			return c, args, nil
		}

		if err := c.ParseFlags(flags); err != nil {
			return nil, args, err
		}
		return cmd.Traverse(args[i+1:])
	}
	return c, args, nil
}

// SuggestionsFor provides suggestions for the typedName.
func (c *Command) SuggestionsFor(typedName string) []string {
	suggestions := []string{}
	for _, cmd := range c.commands {
		if cmd.IsAvailableCommand() {
			levenshteinDistance := ld(typedName, cmd.Name(), true)
			suggestByLevenshtein := levenshteinDistance <= c.SuggestionsMinimumDistance
			suggestByPrefix := strings.HasPrefix(strings.ToLower(cmd.Name()), strings.ToLower(typedName))
			if suggestByLevenshtein || suggestByPrefix {
				suggestions = append(suggestions, cmd.Name())
			}
			for _, explicitSuggestion := range cmd.SuggestFor {
				if strings.EqualFold(typedName, explicitSuggestion) {
					suggestions = append(suggestions, cmd.Name())
				}
			}
		}
	}
	return suggestions
}

// VisitParents visits all parents of the command and invokes fn on each parent.
func (c *Command) VisitParents(fn func(*Command)) {
	if c.HasParent() {
		fn(c.Parent())
		c.Parent().VisitParents(fn)
	}
}

// Root finds root command.
func (c *Command) Root() *Command {
	if c.HasParent() {
		return c.Parent().Root()
	}
	return c
}

// ArgsLenAtDash will return the length of c.Flags().Args at the moment
// when a -- was found during args parsing.
func (c *Command) ArgsLenAtDash() int {
	return c.Flags().ArgsLenAtDash()
}

func (c *Command) execute(a []string) (err error) {
	if c == nil {
		return fmt.Errorf("Called Execute() on a nil Command")
	}

	if len(c.Deprecated) > 0 {
		c.Printf("Command %q is deprecated, %s\n", c.Name(), c.Deprecated)
	}

	// initialize help and version flag at the last point possible to allow for user
	// overriding
	c.InitDefaultHelpFlag()
	c.InitDefaultVersionFlag()

	err = c.ParseFlags(a)
	if err != nil {
		return c.FlagErrorFunc()(c, err)
	}

	// If help is called, regardless of other flags, return we want help.
	// Also say we need help if the command isn't runnable.
	helpVal, err := c.Flags().GetBool("help")
	if err != nil {
		// should be impossible to get here as we always declare a help
		// flag in InitDefaultHelpFlag()
		c.Println("\"help\" flag declared as non-bool. Please correct your code")
		return err
	}

	if helpVal {
		return flag.ErrHelp
	}

	// for back-compat, only add version flag behavior if version is defined
	if c.Version != "" {
		versionVal, err := c.Flags().GetBool("version")
		if err != nil {
			c.Println("\"version\" flag declared as non-bool. Please correct your code")
			return err
		}
		if versionVal {
			err := tmpl(c.OutOrStdout(), c.VersionTemplate(), c)
			if err != nil {
				c.Println(err)
			}
			return err
		}
	}

	if !c.Runnable() {
		return flag.ErrHelp
	}

	c.preRun()

	argWoFlags := c.Flags().Args()
	if c.DisableFlagParsing {
		argWoFlags = a
	}

	if err := c.ValidateArgs(argWoFlags); err != nil {
		return err
	}

	for p := c; p != nil; p = p.Parent() {
		if p.PersistentPreRunE != nil {
			if err := p.PersistentPreRunE(c, argWoFlags); err != nil {
				return err
			}
			break
		} else if p.PersistentPreRun != nil {
			p.PersistentPreRun(c, argWoFlags)
			break
		}
	}
	if c.PreRunE != nil {
		if err := c.PreRunE(c, argWoFlags); err != nil {
			return err
		}
	} else if c.PreRun != nil {
		c.PreRun(c, argWoFlags)
	}

	if err := c.validateRequiredFlags(); err != nil {
		return err
	}
	if c.RunE != nil {
		if err := c.RunE(c, argWoFlags); err != nil {
			return err
		}
	} else {
		c.Run(c, argWoFlags)
	}
	if c.PostRunE != nil {
		if err := c.PostRunE(c, argWoFlags); err != nil {
			return err
		}
	} else if c.PostRun != nil {
		c.PostRun(c, argWoFlags)
	}
	for p := c; p != nil; p = p.Parent() {
		if p.PersistentPostRunE != nil {
			if err := p.PersistentPostRunE(c, argWoFlags); err != nil {
				return err
			}
			break
		} else if p.PersistentPostRun != nil {
			p.PersistentPostRun(c, argWoFlags)
			break
		}
	}

	return nil
}

func (c *Command) preRun() {
	for _, x := range initializers {
		x()
	}
}

// ExecuteContext is the same as Execute(), but sets the ctx on the command.
// Retrieve ctx by calling cmd.Context() inside your *Run lifecycle functions.
func (c *Command) ExecuteContext(ctx context.Context) error {
	c.ctx = ctx
	return c.Execute()
}

// Execute uses the args (os.Args[1:] by default)
// and run through the command tree finding appropriate matches
// for commands and then corresponding flags.
func (c *Command) Execute() error {
	_, err := c.ExecuteC()
	return err
}

// ExecuteC executes the command.
func (c *Command) ExecuteC() (cmd *Command, err error) {
	if c.ctx == nil {
		c.ctx = context.Background()
	}

	// Regardless of what command execute is called on, run on Root only
	if c.HasParent() {
		return c.Root().ExecuteC()
	}

	// windows hook
	if preExecHookFn != nil {
		preExecHookFn(c)
	}

	// initialize help as the last point possible to allow for user
	// overriding
	c.InitDefaultHelpCmd()

	args := c.args

	// Workaround FAIL with "go test -v" or "cobra.test -test.v", see #155
	if c.args == nil && filepath.Base(os.Args[0]) != "cobra.test" {
		args = os.Args[1:]
	}

	// initialize the hidden command to be used for bash completion
	c.initCompleteCmd(args)

	var flags []string
	if c.TraverseChildren {
		cmd, flags, err = c.Traverse(args)
	} else {
		cmd, flags, err = c.Find(args)
	}
	if err != nil {
		// If found parse to a subcommand and then failed, talk about the subcommand
		if cmd != nil {
			c = cmd
		}
		if !c.SilenceErrors {
			c.PrintErrln("Error:", err.Error())
			c.PrintErrf("Run '%v --help' for usage.\n", c.CommandPath())
		}
		return c, err
	}

	cmd.commandCalledAs.called = true
	if cmd.commandCalledAs.name == "" {
		cmd.commandCalledAs.name = cmd.Name()
	}

	// We have to pass global context to children command
	// if context is present on the parent command.
	if cmd.ctx == nil {
		cmd.ctx = c.ctx
	}

	err = cmd.execute(flags)
	if err != nil {
		// Always show help if requested, even if SilenceErrors is in
		// effect
		if err == flag.ErrHelp {
			cmd.HelpFunc()(cmd, args)
			return cmd, nil
		}

		// If root command has SilentErrors flagged,
		// all subcommands should respect it
		if !cmd.SilenceErrors && !c.SilenceErrors {
			c.PrintErrln("Error:", err.Error())
		}

		// If root command has SilentUsage flagged,
		// all subcommands should respect it
		if !cmd.SilenceUsage && !c.SilenceUsage {
			c.Println(cmd.UsageString())
		}
	}
	return cmd, err
}

func (c *Command) ValidateArgs(args []string) error {
	if c.Args == nil {
		return nil
	}
	return c.Args(c, args)
}

func (c *Command) validateRequiredFlags() error {
	if c.DisableFlagParsing {
		return nil
	}

	flags := c.Flags()
	missingFlagNames := []string{}
	flags.VisitAll(func(pflag *flag.Flag) {
		requiredAnnotation, found := pflag.Annotations[BashCompOneRequiredFlag]
		if !found {
			return
		}
		if (requiredAnnotation[0] == "true") && !pflag.Changed {
			missingFlagNames = append(missingFlagNames, pflag.Name)
		}
	})

	if len(missingFlagNames) > 0 {
		return fmt.Errorf(`required flag(s) "%s" not set`, strings.Join(missingFlagNames, `", "`))
	}
	return nil
}

// InitDefaultHelpFlag adds default help flag to c.
// It is called automatically by executing the c or by calling help and usage.
// If c already has help flag, it will do nothing.
func (c *Command) InitDefaultHelpFlag() {
	c.mergePersistentFlags()
	if c.Flags().Lookup("help") == nil {
		usage := "help for "
		if c.Name() == "" {
			usage += "this command"
		} else {
			usage += c.Name()
		}
		c.Flags().BoolP("help", "h", false, usage)
	}
}

// InitDefaultVersionFlag adds default version flag to c.
// It is called automatically by executing the c.
// If c already has a version flag, it will do nothing.
// If c.Version is empty, it will do nothing.
func (c *Command) InitDefaultVersionFlag() {
	if c.Version == "" {
		return
	}

	c.mergePersistentFlags()
	if c.Flags().Lookup("version") == nil {
		usage := "version for "
		if c.Name() == "" {
			usage += "this command"
		} else {
			usage += c.Name()
		}
		if c.Flags().ShorthandLookup("v") == nil {
			c.Flags().BoolP("version", "v", false, usage)
		} else {
			c.Flags().Bool("version", false, usage)
		}
	}
}

// InitDefaultHelpCmd adds default help command to c.
// It is called automatically by executing the c or by calling help and usage.
// If c already has help command or c has no subcommands, it will do nothing.
func (c *Command) InitDefaultHelpCmd() {
	if !c.HasSubCommands() {
		return
	}

	if c.helpCommand == nil {
		c.helpCommand = &Command{
			Use:   "help [command]",
			Short: "Help about any command",
			Long: `Help provides help for any command in the application.
Simply type ` + c.Name() + ` help [path to command] for full details.`,
			ValidArgsFunction: func(c *Command, args []string, toComplete string) ([]string, ShellCompDirective) {
				var completions []string
				cmd, _, e := c.Root().Find(args)
				if e != nil {
					return nil, ShellCompDirectiveNoFileComp
				}
				if cmd == nil {
					// Root help command.
					cmd = c.Root()
				}
				for _, subCmd := range cmd.Commands() {
					if subCmd.IsAvailableCommand() || subCmd == cmd.helpCommand {
						if strings.HasPrefix(subCmd.Name(), toComplete) {
							completions = append(completions, fmt.Sprintf("%s\t%s", subCmd.Name(), subCmd.Short))
						}
					}
				}
				return completions, ShellCompDirectiveNoFileComp
			},
			Run: func(c *Command, args []string) {
				cmd, _, e := c.Root().Find(args)
				if cmd == nil || e != nil {
					c.Printf("Unknown help topic %#q\n", args)
					c.Root().Usage()
				} else {
					cmd.InitDefaultHelpFlag() // make possible 'help' flag to be shown
					cmd.Help()
				}
			},
		}
	}
	c.RemoveCommand(c.helpCommand)
	c.AddCommand(c.helpCommand)
}

// ResetCommands delete parent, subcommand and help command from c.
func (c *Command) ResetCommands() {
	c.parent = nil
	c.commands = nil
	c.helpCommand = nil
	c.parentsPflags = nil
}

// Sorts commands by their names.
type commandSorterByName []*Command

func (c commandSorterByName) Len() int           { return len(c) }
func (c commandSorterByName) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }
func (c commandSorterByName) Less(i, j int) bool { return c[i].Name() < c[j].Name() }

// Commands returns a sorted slice of child commands.
func (c *Command) Commands() []*Command {
	// do not sort commands if it already sorted or sorting was disabled
	if EnableCommandSorting && !c.commandsAreSorted {
		sort.Sort(commandSorterByName(c.commands))
		c.commandsAreSorted = true
	}
	return c.commands
}

// AddCommand adds one or more commands to this parent command.
func (c *Command) AddCommand(cmds ...*Command) {
	for i, x := range cmds {
		if cmds[i] == c {
			panic("Command can't be a child of itself")
		}
		cmds[i].parent = c
		// update max lengths
		usageLen := len(x.Use)
		if usageLen > c.commandsMaxUseLen {
			c.commandsMaxUseLen = usageLen
		}
		commandPathLen := len(x.CommandPath())
		if commandPathLen > c.commandsMaxCommandPathLen {
			c.commandsMaxCommandPathLen = commandPathLen
		}
		nameLen := len(x.Name())
		if nameLen > c.commandsMaxNameLen {
			c.commandsMaxNameLen = nameLen
		}
		// If global normalization function exists, update all children
		if c.globNormFunc != nil {
			x.SetGlobalNormalizationFunc(c.globNormFunc)
		}
		c.commands = append(c.commands, x)
		c.commandsAreSorted = false
	}
}

// RemoveCommand removes one or more commands from a parent command.
func (c *Command) RemoveCommand(cmds ...*Command) {
	commands := []*Command{}
main:
	for _, command := range c.commands {
		for _, cmd := range cmds {
			if command == cmd {
				command.parent = nil
				continue main
			}
		}
		commands = append(commands, command)
	}
	c.commands = commands
	// recompute all lengths
	c.commandsMaxUseLen = 0
	c.commandsMaxCommandPathLen = 0
	c.commandsMaxNameLen = 0
	for _, command := range c.commands {
		usageLen := len(command.Use)
		if usageLen > c.commandsMaxUseLen {
			c.commandsMaxUseLen = usageLen
		}
		commandPathLen := len(command.CommandPath())
		if commandPathLen > c.commandsMaxCommandPathLen {
			c.commandsMaxCommandPathLen = commandPathLen
		}
		nameLen := len(command.Name())
		if nameLen > c.commandsMaxNameLen {
			c.commandsMaxNameLen = nameLen
		}
	}
}

// Print is a convenience method to Print to the defined output, fallback to Stderr if not set.
func (c *Command) Print(i ...interface{}) {
	fmt.Fprint(c.OutOrStderr(), i...)
}

// Println is a convenience method to Println to the defined output, fallback to Stderr if not set.
func (c *Command) Println(i ...interface{}) {
	c.Print(fmt.Sprintln(i...))
}

// Printf is a convenience method to Printf to the defined output, fallback to Stderr if not set.
func (c *Command) Printf(format string, i ...interface{}) {
	c.Print(fmt.Sprintf(format, i...))
}

// PrintErr is a convenience method to Print to the defined Err output, fallback to Stderr if not set.
func (c *Command) PrintErr(i ...interface{}) {
	fmt.Fprint(c.ErrOrStderr(), i...)
}

// PrintErrln is a convenience method to Println to the defined Err output, fallback to Stderr if not set.
func (c *Command) PrintErrln(i ...interface{}) {
	c.PrintErr(fmt.Sprintln(i...))
}

// PrintErrf is a convenience method to Printf to the defined Err output, fallback to Stderr if not set.
func (c *Command) PrintErrf(format string, i ...interface{}) {
	c.PrintErr(fmt.Sprintf(format, i...))
}

// CommandPath returns the full path to this command.
func (c *Command) CommandPath() string {
	if c.HasParent() {
		return c.Parent().CommandPath() + " " + c.Name()
	}
	return c.Name()
}

// UseLine puts out the full usage for a given command (including parents).
func (c *Command) UseLine() string {
	var useline string
	if c.HasParent() {
		useline = c.parent.CommandPath() + " " + c.Use
	} else {
		useline = c.Use
	}
	if c.DisableFlagsInUseLine {
		return useline
	}
	if c.HasAvailableFlags() && !strings.Contains(useline, "[flags]") {
		useline += " [flags]"
	}
	return useline
}

// DebugFlags used to determine which flags have been assigned to which commands
// and which persist.
func (c *Command) DebugFlags() {
	c.Println("DebugFlags called on", c.Name())
	var debugflags func(*Command)

	debugflags = func(x *Command) {
		if x.HasFlags() || x.HasPersistentFlags() {
			c.Println(x.Name())
		}
		if x.HasFlags() {
			x.flags.VisitAll(func(f *flag.Flag) {
				if x.HasPersistentFlags() && x.persistentFlag(f.Name) != nil {
					c.Println("  -"+f.Shorthand+",", "--"+f.Name, "["+f.DefValue+"]", "", f.Value, "  [LP]")
				} else {
					c.Println("  -"+f.Shorthand+",", "--"+f.Name, "["+f.DefValue+"]", "", f.Value, "  [L]")
				}
			})
		}
		if x.HasPersistentFlags() {
			x.pflags.VisitAll(func(f *flag.Flag) {
				if x.HasFlags() {
					if x.flags.Lookup(f.Name) == nil {
						c.Println("  -"+f.Shorthand+",", "--"+f.Name, "["+f.DefValue+"]", "", f.Value, "  [P]")
					}
				} else {
					c.Println("  -"+f.Shorthand+",", "--"+f.Name, "["+f.DefValue+"]", "", f.Value, "  [P]")
				}
			})
		}
		c.Println(x.flagErrorBuf)
		if x.HasSubCommands() {
			for _, y := range x.commands {
				debugflags(y)
			}
		}
	}

	debugflags(c)
}

// Name returns the command's name: the first word in the use line.
func (c *Command) Name() string {
	name := c.Use
	i := strings.Index(name, " ")
	if i >= 0 {
		name = name[:i]
	}
	return name
}

// HasAlias determines if a given string is an alias of the command.
func (c *Command) HasAlias(s string) bool {
	for _, a := range c.Aliases {
		if a == s {
			return true
		}
	}
	return false
}

// CalledAs returns the command name or alias that was used to invoke
// this command or an empty string if the command has not been called.
func (c *Command) CalledAs() string {
	if c.commandCalledAs.called {
		return c.commandCalledAs.name
	}
	return ""
}

// hasNameOrAliasPrefix returns true if the Name or any of aliases start
// with prefix
func (c *Command) hasNameOrAliasPrefix(prefix string) bool {
	if strings.HasPrefix(c.Name(), prefix) {
		c.commandCalledAs.name = c.Name()
		return true
	}
	for _, alias := range c.Aliases {
		if strings.HasPrefix(alias, prefix) {
			c.commandCalledAs.name = alias
			return true
		}
	}
	return false
}

// NameAndAliases returns a list of the command name and all aliases
func (c *Command) NameAndAliases() string {
	return strings.Join(append([]string{c.Name()}, c.Aliases...), ", ")
}

// HasExample determines if the command has example.
func (c *Command) HasExample() bool {
	return len(c.Example) > 0
}

// Runnable determines if the command is itself runnable.
func (c *Command) Runnable() bool {
	return c.Run != nil || c.RunE != nil
}

// HasSubCommands determines if the command has children commands.
func (c *Command) HasSubCommands() bool {
	return len(c.commands) > 0
}

// IsAvailableCommand determines if a command is available as a non-help command
// (this includes all non deprecated/hidden commands).
func (c *Command) IsAvailableCommand() bool {
	if len(c.Deprecated) != 0 || c.Hidden {
		return false
	}

	if c.HasParent() && c.Parent().helpCommand == c {
		return false
	}

	if c.Runnable() || c.HasAvailableSubCommands() {
		return true
	}

	return false
}

// IsAdditionalHelpTopicCommand determines if a command is an additional
// help topic command; additional help topic command is determined by the
// fact that it is NOT runnable/hidden/deprecated, and has no sub commands that
// are runnable/hidden/deprecated.
// Concrete example: https://github.com/spf13/cobra/issues/393#issuecomment-282741924.
func (c *Command) IsAdditionalHelpTopicCommand() bool {
	// if a command is runnable, deprecated, or hidden it is not a 'help' command
	if c.Runnable() || len(c.Deprecated) != 0 || c.Hidden {
		return false
	}

	// if any non-help sub commands are found, the command is not a 'help' command
	for _, sub := range c.commands {
		if !sub.IsAdditionalHelpTopicCommand() {
			return false
		}
	}

	// the command either has no sub commands, or no non-help sub commands
	return true
}

// HasHelpSubCommands determines if a command has any available 'help' sub commands
// that need to be shown in the usage/help default template under 'additional help
// topics'.
func (c *Command) HasHelpSubCommands() bool {
	// return true on the first found available 'help' sub command
	for _, sub := range c.commands {
		if sub.IsAdditionalHelpTopicCommand() {
			return true
		}
	}

	// the command either has no sub commands, or no available 'help' sub commands
	return false
}

// HasAvailableSubCommands determines if a command has available sub commands that
// need to be shown in the usage/help default template under 'available commands'.
func (c *Command) HasAvailableSubCommands() bool {
	// return true on the first found available (non deprecated/help/hidden)
	// sub command
	for _, sub := range c.commands {
		if sub.IsAvailableCommand() {
			return true
		}
	}

	// the command either has no sub commands, or no available (non deprecated/help/hidden)
	// sub commands
	return false
}

// HasParent determines if the command is a child command.
func (c *Command) HasParent() bool {
	return c.parent != nil
}

// GlobalNormalizationFunc returns the global normalization function or nil if it doesn't exist.
func (c *Command) GlobalNormalizationFunc() func(f *flag.FlagSet, name string) flag.NormalizedName {
	return c.globNormFunc
}

// Flags returns the complete FlagSet that applies
// to this command (local and persistent declared here and by all parents).
func (c *Command) Flags() *flag.FlagSet {
	if c.flags == nil {
		c.flags = flag.NewFlagSet(c.Name(), flag.ContinueOnError)
		if c.flagErrorBuf == nil {
			c.flagErrorBuf = new(bytes.Buffer)
		}
		c.flags.SetOutput(c.flagErrorBuf)
	}

	return c.flags
}

// LocalNonPersistentFlags are flags specific to this command which will NOT persist to subcommands.
func (c *Command) LocalNonPersistentFlags() *flag.FlagSet {
	persistentFlags := c.PersistentFlags()

	out := flag.NewFlagSet(c.Name(), flag.ContinueOnError)
	c.LocalFlags().VisitAll(func(f *flag.Flag) {
		if persistentFlags.Lookup(f.Name) == nil {
			out.AddFlag(f)
		}
	})
	return out
}

// LocalFlags returns the local FlagSet specifically set in the current command.
func (c *Command) LocalFlags() *flag.FlagSet {
	c.mergePersistentFlags()

	if c.lflags == nil {
		c.lflags = flag.NewFlagSet(c.Name(), flag.ContinueOnError)
		if c.flagErrorBuf == nil {
			c.flagErrorBuf = new(bytes.Buffer)
		}
		c.lflags.SetOutput(c.flagErrorBuf)
	}
	c.lflags.SortFlags = c.Flags().SortFlags
	if c.globNormFunc != nil {
		c.lflags.SetNormalizeFunc(c.globNormFunc)
	}

	addToLocal := func(f *flag.Flag) {
		if c.lflags.Lookup(f.Name) == nil && c.parentsPflags.Lookup(f.Name) == nil {
			c.lflags.AddFlag(f)
		}
	}
	c.Flags().VisitAll(addToLocal)
	c.PersistentFlags().VisitAll(addToLocal)
	return c.lflags
}

// InheritedFlags returns all flags which were inherited from parent commands.
func (c *Command) InheritedFlags() *flag.FlagSet {
	c.mergePersistentFlags()

	if c.iflags == nil {
		c.iflags = flag.NewFlagSet(c.Name(), flag.ContinueOnError)
		if c.flagErrorBuf == nil {
			c.flagErrorBuf = new(bytes.Buffer)
		}
		c.iflags.SetOutput(c.flagErrorBuf)
	}

	local := c.LocalFlags()
	if c.globNormFunc != nil {
		c.iflags.SetNormalizeFunc(c.globNormFunc)
	}

	c.parentsPflags.VisitAll(func(f *flag.Flag) {
		if c.iflags.Lookup(f.Name) == nil && local.Lookup(f.Name) == nil {
			c.iflags.AddFlag(f)
		}
	})
	return c.iflags
}

// NonInheritedFlags returns all flags which were not inherited from parent commands.
func (c *Command) NonInheritedFlags() *flag.FlagSet {
	return c.LocalFlags()
}

// PersistentFlags returns the persistent FlagSet specifically set in the current command.
func (c *Command) PersistentFlags() *flag.FlagSet {
	if c.pflags == nil {
		c.pflags = flag.NewFlagSet(c.Name(), flag.ContinueOnError)
		if c.flagErrorBuf == nil {
			c.flagErrorBuf = new(bytes.Buffer)
		}
		c.pflags.SetOutput(c.flagErrorBuf)
	}
	return c.pflags
}

// ResetFlags deletes all flags from command.
func (c *Command) ResetFlags() {
	c.flagErrorBuf = new(bytes.Buffer)
	c.flagErrorBuf.Reset()
	c.flags = flag.NewFlagSet(c.Name(), flag.ContinueOnError)
	c.flags.SetOutput(c.flagErrorBuf)
	c.pflags = flag.NewFlagSet(c.Name(), flag.ContinueOnError)
	c.pflags.SetOutput(c.flagErrorBuf)

	c.lflags = nil
	c.iflags = nil
	c.parentsPflags = nil
}

// HasFlags checks if the command contains any flags (local plus persistent from the entire structure).
func (c *Command) HasFlags() bool {
	return c.Flags().HasFlags()
}

// HasPersistentFlags checks if the command contains persistent flags.
func (c *Command) HasPersistentFlags() bool {
	return c.PersistentFlags().HasFlags()
}

// HasLocalFlags checks if the command has flags specifically declared locally.
func (c *Command) HasLocalFlags() bool {
	return c.LocalFlags().HasFlags()
}

// HasInheritedFlags checks if the command has flags inherited from its parent command.
func (c *Command) HasInheritedFlags() bool {
	return c.InheritedFlags().HasFlags()
}

// HasAvailableFlags checks if the command contains any flags (local plus persistent from the entire
// structure) which are not hidden or deprecated.
func (c *Command) HasAvailableFlags() bool {
	return c.Flags().HasAvailableFlags()
}

// HasAvailablePersistentFlags checks if the command contains persistent flags which are not hidden or deprecated.
func (c *Command) HasAvailablePersistentFlags() bool {
	return c.PersistentFlags().HasAvailableFlags()
}

// HasAvailableLocalFlags checks if the command has flags specifically declared locally which are not hidden
// or deprecated.
func (c *Command) HasAvailableLocalFlags() bool {
	return c.LocalFlags().HasAvailableFlags()
}

// HasAvailableInheritedFlags checks if the command has flags inherited from its parent command which are
// not hidden or deprecated.
func (c *Command) HasAvailableInheritedFlags() bool {
	return c.InheritedFlags().HasAvailableFlags()
}

// Flag climbs up the command tree looking for matching flag.
func (c *Command) Flag(name string) (flag *flag.Flag) {
	flag = c.Flags().Lookup(name)

	if flag == nil {
		flag = c.persistentFlag(name)
	}

	return
}

// Recursively find matching persistent flag.
func (c *Command) persistentFlag(name string) (flag *flag.Flag) {
	if c.HasPersistentFlags() {
		flag = c.PersistentFlags().Lookup(name)
	}

	if flag == nil {
		c.updateParentsPflags()
		flag = c.parentsPflags.Lookup(name)
	}
	return
}

// ParseFlags parses persistent flag tree and local flags.
func (c *Command) ParseFlags(args []string) error {
	if c.DisableFlagParsing {
		return nil
	}

	if c.flagErrorBuf == nil {
		c.flagErrorBuf = new(bytes.Buffer)
	}
	beforeErrorBufLen := c.flagErrorBuf.Len()
	c.mergePersistentFlags()

	// do it here after merging all flags and just before parse
	c.Flags().ParseErrorsWhitelist = flag.ParseErrorsWhitelist(c.FParseErrWhitelist)

	err := c.Flags().Parse(args)
	// Print warnings if they occurred (e.g. deprecated flag messages).
	if c.flagErrorBuf.Len()-beforeErrorBufLen > 0 && err == nil {
		c.Print(c.flagErrorBuf.String())
	}

	return err
}

// Parent returns a commands parent command.
func (c *Command) Parent() *Command {
	return c.parent
}

// mergePersistentFlags merges c.PersistentFlags() to c.Flags()
// and adds missing persistent flags of all parents.
func (c *Command) mergePersistentFlags() {
	c.updateParentsPflags()
	c.Flags().AddFlagSet(c.PersistentFlags())
	c.Flags().AddFlagSet(c.parentsPflags)
}

// updateParentsPflags updates c.parentsPflags by adding
// new persistent flags of all parents.
// If c.parentsPflags == nil, it makes new.
func (c *Command) updateParentsPflags() {
	if c.parentsPflags == nil {
		c.parentsPflags = flag.NewFlagSet(c.Name(), flag.ContinueOnError)
		c.parentsPflags.SetOutput(c.flagErrorBuf)
		c.parentsPflags.SortFlags = false
	}

	if c.globNormFunc != nil {
		c.parentsPflags.SetNormalizeFunc(c.globNormFunc)
	}

	c.Root().PersistentFlags().AddFlagSet(flag.CommandLine)

	c.VisitParents(func(parent *Command) {
		c.parentsPflags.AddFlagSet(parent.PersistentFlags())
	})
}
