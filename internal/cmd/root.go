package cmd

import (
	"os"
	"path/filepath"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/yendo/rdme/internal/parser"
)

var (
	chdir    string
	fileName string
)

func Root() *cobra.Command {
	cmd := cobra.Command{
		Use:           "rdme",
		Short:         "Execute commands directly from a README",
		Long:          "Parses commands directly from a README (best-effort) to make them executable under a unique name.",
		SilenceErrors: true,
		SilenceUsage:  true,
		ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			return nil, cobra.ShellCompDirectiveNoFileComp
		},
	}

	pflags := cmd.PersistentFlags()

	pflags.StringVar(&chdir, "chdir", ".", "Switch to a different working directory before exeucing the command.")
	pflags.StringVar(&fileName, "filename", "README.md", "A name of the README file.")

	cmd.AddCommand(runCmd())
	cmd.AddCommand(listCmd())
	cmd.AddCommand(printCmd())
	cmd.AddCommand(tasksCmd())

	return &cmd
}

func newParser() (*parser.Parser, error) {
	source, err := os.ReadFile(filepath.Join(chdir, fileName))
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return parser.New(source), nil
}

func lookup(snippets parser.Snippets, name string) (*parser.Snippet, error) {
	snippet, found := snippets.Lookup(name)
	if !found {
		return nil, errors.Errorf("command %q not found; known command names: %s", name, snippets.Names())
	}
	return snippet, nil
}
