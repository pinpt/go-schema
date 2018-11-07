package cmd

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/pinpt/go-common/fileutil"
	"github.com/spf13/cobra"
)

var schemaCmd = &cobra.Command{
	Use:   "schema",
	Short: "Generate database schema files from protobuf model definitions",
	Run: func(cmd *cobra.Command, args []string) {
		if err := runSchemaCmd(); err != nil {
			panic(err)
		}
		if err := runBindataCmd(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func runSchemaCmd() error {
	rootDir, _ := os.Getwd()
	protoDir := filepath.Join(rootDir, "proto")
	tempDir := filepath.Join(rootDir, "tmp")
	tempProto := filepath.Join(tempDir, "schema.proto")

	if err := generateTempProtoFile(rootDir, protoDir, tempDir, tempProto); err != nil {
		return err
	}
	if err := runProtoc(protoDir, tempDir, tempProto, rootDir); err != nil {
		return err
	}
	deleteAllGoFiles(rootDir)
	if err := moveAllGoFiles(rootDir); err != nil {
		return err
	}
	if err := moveGooseDb(rootDir, tempDir); err != nil {
		return err
	}
	// cleanup
	deleteTempDir(rootDir)
	// Why is this created? No idea
	os.RemoveAll(filepath.Join(rootDir, "schema", "golang"))
	return nil
}

func generateTempProtoFile(rootDir string, protoDir string, tempDir string, tempProto string) error {
	tmpProtocFile := []string{}
	tmpProtocFile = append(tmpProtocFile, "syntax = \"proto3\";", "")
	tmpProtocFile = append(tmpProtocFile, "package schema;", "")
	tmpProtocFile = append(tmpProtocFile, "import \"proto/annotations.proto\";")
	tmpProtocFile = append(tmpProtocFile, "import \"proto/types.proto\";", "")
	tmpProtocFile = append(tmpProtocFile, "option (proto.file).lowercaseEnums=true;")

	files, err := ioutil.ReadDir(protoDir)
	if err != nil {
		return err
	}

	for _, f := range files {
		n := f.Name()
		if strings.HasSuffix(n, ".proto") {
			content, err := ioutil.ReadFile(filepath.Join(protoDir, n))
			if err != nil {
				return err
			}
			tmpProtocFile = append(tmpProtocFile, string(content))
		}
	}
	deleteTempDir(rootDir)
	os.MkdirAll(tempDir, 0755)
	return ioutil.WriteFile(tempProto, []byte(strings.Join(tmpProtocFile, "\n")), 0644)
}

func runProtoc(protoDir string, tempDir string, tempProto string, rootDir string) error {
	vendorDir := filepath.Join(rootDir, "vendor")
	cArgs := []string{}
	cArgs = append(cArgs, fmt.Sprintf("--proto_path=%v", tempDir))
	cArgs = append(cArgs, fmt.Sprintf("-I=%v", protoDir))
	cArgs = append(cArgs, fmt.Sprintf("-I=%v", filepath.Join(vendorDir, "github.com", "jhaynie", "protoc-gen-gator")))
	cArgs = append(cArgs, fmt.Sprintf("-I=%v", filepath.Join(vendorDir, "github.com", "jhaynie", "protoc-gen-gator", "proto")))
	cArgs = append(cArgs, fmt.Sprintf("--gator_out=goose,golang:%v", tempDir))
	cArgs = append(cArgs, tempProto)

	fmt.Println(fmt.Sprintf("Running: \"protoc %v\"", cArgs))
	c := exec.CommandContext(context.Background(), "protoc", cArgs...)
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr

	return c.Run()
}

func moveAllGoFiles(rootDir string) error {

	old := filepath.Join(rootDir, "tmp", "schema", "golang")
	new := filepath.Join(rootDir, "schema")

	files, err := ioutil.ReadDir(old)
	if err != nil {
		return err
	}

	for _, f := range files {
		n := f.Name()
		if strings.HasSuffix(n, ".go") {
			oldPath := filepath.Join(old, n)
			newPath := filepath.Join(new, n)
			fmt.Println(fmt.Sprintf("moving: %v ----> %v", strings.Replace(oldPath, rootDir, "", -1)[1:], strings.Replace(newPath, rootDir, "", -1)[1:]))
			os.Rename(oldPath, newPath)
		}
	}
	return nil
}

func moveGooseDb(rootDir string, tempDir string) error {
	oldFile := filepath.Join(tempDir, "schema", "goose_db.sql")
	newFile := filepath.Join(rootDir, "migrations")
	if !fileutil.FileExists(oldFile) {
		return fmt.Errorf("file `goose_db.sql` does not exist")
	}
	os.MkdirAll(newFile, 0755)
	newFile = filepath.Join(newFile, "20170523183416_init.sql")
	return os.Rename(oldFile, newFile)
}

func init() {
	rootCmd.AddCommand(schemaCmd)
}
