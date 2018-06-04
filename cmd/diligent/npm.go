package main

import (
	"github.com/senseyeio/diligent/npm"
	"github.com/spf13/cobra"
)

var (
	npmDevDeps bool
)

const npmAPIURL = "https://registry.npmjs.org"

// npmCmd represents the npm command
var npmCmd = &cobra.Command{
	Use:   "npm [filePath]",
	Short: "Exposes NPM specific options",
	Long: `The NPM command is the same as the diligent command, but it exposes additional NPM options.
Can only be used with package.json files`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		npmDeper := npm.NewWithOptions(npmAPIURL, npm.Config{
			DevDependencies: npmDevDeps,
		})
		runDep(npmDeper, getReporter(), args[0])
	},
}

func init() {
	RootCmd.AddCommand(npmCmd)
	npmCmd.Flags().BoolVarP(&npmDevDeps, "devDeps", "d", false, "Include developer dependencies")
}