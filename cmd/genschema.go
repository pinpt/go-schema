package cmd

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var genSchemaCmd = &cobra.Command{
	Use: "genschema",
	Run: func(cmd *cobra.Command, args []string) {

		cwd, _ := os.Getwd()
		protoDir := filepath.Join(cwd, "proto")

		tmpProtocFile := []string{}
		tmpProtocFile = append(tmpProtocFile, "")
		tmpProtocFile = append(tmpProtocFile, "syntax = \"proto3\";")
		tmpProtocFile = append(tmpProtocFile, "")
		tmpProtocFile = append(tmpProtocFile, "package schema;")
		tmpProtocFile = append(tmpProtocFile, "")
		tmpProtocFile = append(tmpProtocFile, "import \"proto/annotations.proto\";")
		tmpProtocFile = append(tmpProtocFile, "import \"proto/types.proto\";")
		tmpProtocFile = append(tmpProtocFile, "")
		tmpProtocFile = append(tmpProtocFile, "option (proto.file).lowercaseEnums=true;")

		files, err := ioutil.ReadDir(protoDir)
		if err != nil {
			log.Fatal(err)
		}

		for _, f := range files {
			n := f.Name()
			if strings.HasSuffix(n, ".proto") {
				fmt.Println(f.Name())
				content, err := ioutil.ReadFile(filepath.Join(protoDir, n))
				if err != nil {
					fmt.Println(err)
					continue
				}
				tmpProtocFile = append(tmpProtocFile, string(content))
			}
		}
		schemaProtocDir := filepath.Join(cwd, "tmp")
		os.RemoveAll(schemaProtocDir)
		os.MkdirAll(schemaProtocDir, 0755)

		schemaProtoc := filepath.Join(schemaProtocDir, "schema.proto")
		os.Remove(schemaProtoc)

		ioutil.WriteFile(schemaProtoc, []byte(strings.Join(tmpProtocFile, "\n")), 0644)

		cArgs := []string{}
		cArgs = append(cArgs, fmt.Sprintf("--proto_path=%v", schemaProtocDir))
		cArgs = append(cArgs, fmt.Sprintf("-I=%v", protoDir))
		// cArgs = append(cArgs, fmt.Sprintf("-I=%v", filepath.Join(cwd, "proto")))
		cArgs = append(cArgs, fmt.Sprintf("-I=%v", filepath.Join(os.Getenv("GOPATH"), "src", "github.com", "jhaynie", "protoc-gen-gator")))
		cArgs = append(cArgs, fmt.Sprintf("-I=%v", filepath.Join(os.Getenv("GOPATH"), "src", "github.com", "jhaynie", "protoc-gen-gator", "proto")))
		cArgs = append(cArgs, fmt.Sprintf("-I=%v", filepath.Join(os.Getenv("GOPATH"), "src")))
		cArgs = append(cArgs, fmt.Sprintf("--gator_out=goose,golang:%v", schemaProtocDir))
		cArgs = append(cArgs, schemaProtoc)

		fmt.Println(fmt.Sprintf("Running: %v", cArgs))
		c := exec.CommandContext(context.Background(), "protoc", cArgs...)
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr

		if err := c.Run(); err != nil {
			panic(err)
		}
		if err := deleteAllGoFiles(); err != nil {
			panic(err)
		}
		if err := moveAllGoFiles(); err != nil {
			panic(err)
		}
	},
}

func deleteAllGoFiles() error {

	cwd, _ := os.Getwd()
	src := filepath.Join(cwd, "schema")
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

func moveAllGoFiles() error {
	cwd, _ := os.Getwd()
	old := filepath.Join(cwd, "tmp", "schema", "golang")
	new := filepath.Join(cwd, "schema")

	files, err := ioutil.ReadDir(old)
	if err != nil {
		return err
	}
	for _, f := range files {
		n := f.Name()
		if strings.HasSuffix(n, ".go") {
			oldPath := filepath.Join(old, n)
			newPath := filepath.Join(new, n)
			fmt.Println(fmt.Sprintf("moving %v, %v", oldPath, newPath))
			os.Rename(oldPath, newPath)
		}
	}
	return nil
}
func init() {
	rootCmd.AddCommand(genSchemaCmd)

}
