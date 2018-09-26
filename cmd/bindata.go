package cmd

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/pinpt/go-common/fileutil"
	"github.com/spf13/cobra"
)

var bindataCmd = &cobra.Command{
	Use:   "bindata",
	Short: "Generate static data into go source files",
	Run: func(cmd *cobra.Command, args []string) {
		if err := runBindataCmd(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func runBindataCmd() error {

	rootDir, _ := os.Getwd()
	migrationsDir := filepath.Join(rootDir, "migrations")
	if !fileutil.FileExists(migrationsDir) {
		return fmt.Errorf("`./migrations` directory not found")
	}
	files, err := ioutil.ReadDir(migrationsDir)
	if err != nil {
		return err
	}
	if len(files) == 0 {
		return fmt.Errorf("`./migrations` directory is empty")
	}
	cArgs := []string{
		"-o", "./migrate/bindata.go",
		"-pkg", "migrate",
		"-ignore=\\\\.go",
		"-ignore=\\\\.DS_Store",
		"./migrations",
	}
	deleteMigrateDir(rootDir)

	fmt.Println(fmt.Sprintf("Running: \"go-bindata %v\"", cArgs))
	c := exec.CommandContext(context.Background(), "go-bindata", cArgs...)
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr

	return c.Run()
}

func init() {
	rootCmd.AddCommand(bindataCmd)
}
