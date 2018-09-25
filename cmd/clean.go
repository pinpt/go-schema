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
		cwd, _ := os.Getwd()

		migration := filepath.Join(cwd, "schema", "migrations")
		os.RemoveAll(migration)

		routesGo := filepath.Join(cwd, "schema", "acl", "routes.go")
		os.Remove(routesGo)

		bindataGo := filepath.Join(cwd, "schema", "migrate", "bindata.go")
		os.Remove(bindataGo)

		tempDir := filepath.Join(cwd, "tmp")
		os.RemoveAll(tempDir)
	},
}

func init() {
	rootCmd.AddCommand(cleanCmd)

}
