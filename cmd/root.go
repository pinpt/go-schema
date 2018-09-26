package cmd

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var allCmd = &cobra.Command{
	Use:   "all",
	Short: "Runs clean, genschema, genroutes, and bindata",
	Run: func(cmd *cobra.Command, args []string) {
		cleanAll()
		dir, _ := cmd.Flags().GetString("webroot")
		if err := runGenSchemaCmd(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if err := runGenroutesCmd(dir); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if err := runBindataCmd(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

var rootCmd = &cobra.Command{
	Use:  "pinpoint-schema",
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	rootCmd.SetOutput(color.Output)
	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("err=%s\n", err)
		os.Exit(1)
	}
}
func init() {
	rootCmd.AddCommand(allCmd)
	allCmd.Flags().String("webroot", os.Getenv("PP_WEBROOT"), "path to webapp repo")
}
