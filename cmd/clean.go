package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Removes all the generated files",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		cleanAll()
	},
}

func cleanAll() {
	rootDir, _ := os.Getwd()
	deleteMigrationsDir(rootDir)
	deleteRoutesGoFile(rootDir)
	deleteMigrateDir(rootDir)
	deleteTempDir(rootDir)
	deleteAllGoFiles(rootDir)
}

func deleteRoutesGoFile(rootDir string) {
	routesGo := filepath.Join(rootDir, "acl", "routes.go")
	os.Remove(routesGo)
}

func deleteMigrationsDir(rootDir string) {
	migration := filepath.Join(rootDir, "migrations")
	os.RemoveAll(migration)
}

func deleteMigrateDir(rootDir string) {
	bindataGo := filepath.Join(rootDir, "migrate")
	os.RemoveAll(bindataGo)
}

func deleteTempDir(rootDir string) {
	tempDir := filepath.Join(rootDir, "tmp")
	os.RemoveAll(tempDir)
}

func deleteAllGoFiles(rootDir string) error {
	src := filepath.Join(rootDir, "schema")
	files, err := ioutil.ReadDir(src)
	if err != nil {
		return err
	}
	for _, f := range files {
		n := f.Name()
		if strings.HasSuffix(n, ".go") {
			fmt.Println(fmt.Sprintf("deleting %v", n))
			os.Remove(filepath.Join(src, n))
		}
	}
	return nil
}

func init() {
	rootCmd.AddCommand(cleanCmd)

}
