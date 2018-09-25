package cmd

import (
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Removes all the unnecessary files",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		cleanAll()
	},
}

func cleanAll() {
	cwd, _ := os.Getwd()

	migration := filepath.Join(cwd, "migrations")
	os.RemoveAll(migration)

	routesGo := filepath.Join(cwd, "acl", "routes.go")
	os.Remove(routesGo)

	bindataGo := filepath.Join(cwd, "migrate")
	os.RemoveAll(bindataGo)

	tempDir := filepath.Join(cwd, "tmp")
	os.RemoveAll(tempDir)
}

func init() {
	rootCmd.AddCommand(cleanCmd)

}
