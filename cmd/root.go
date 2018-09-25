package cmd

// +build ignore

import (
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/pinpt/go-common/datetime"
	"github.com/pinpt/go-common/fileutil"
	"github.com/pinpt/go-common/hash"
	"github.com/pinpt/go-schema/schema/acl"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "pinpoint-schema",
	Run: func(cmd *cobra.Command, args []string) {
		dir, _ := cmd.Flags().GetString("webroot")
		if dir == "" {
			panic("environment variable PP_WEBROOT, or flag --webroot, is not defined")
		}
		if !fileutil.FileExists(dir) {
			panic("cannot find " + dir)
		}

		routes := []*acl.Route{}

		if err := generateSQL(routes); err != nil {
			panic(err)
		}
		if err := generateRoutesGo(routes, dir); err != nil {
			panic(err)
		}
		if err := runBindata(); err != nil {
			panic(err)
		}

	},
}

func generateRoutesGo(routes []*acl.Route, dir string) error {

	routefn := filepath.Join(dir, "src", "routes.json")
	if !fileutil.FileExists(routefn) {
		routefn = path.Join(dir, "routes.json")
		if !fileutil.FileExists(routefn) {
			return fmt.Errorf("cannot find %v", routefn)
		}
	}

	cwd, _ := os.Getwd()
	fn := filepath.Join(cwd, "schema", "acl", "routes.go")
	f, err := os.Create(fn)
	if err != nil {
		return err
	}
	defer f.Close()

	fmt.Fprintln(os.Stderr, ">> Loaded route file", fileutil.ResolveFileName(routefn))
	routef, err := os.Open(routefn)
	if err != nil {
		panic(fmt.Errorf("error opening %s. %v", routefn, err))
	}
	defer routef.Close()

	proutes := acl.RouterConfig{}
	if err := json.NewDecoder(routef).Decode(&proutes); err != nil {
		routef.Close()
		routef, _ := os.Open(routefn)
		buf, _ := ioutil.ReadAll(routef)
		routef.Close()
		fmt.Println(">>> ERROR parsing routes:\n", string(buf))
		panic(fmt.Errorf("error decoding %s. %v", routefn, err))
	}

	proutes.Routes["Cost - Salary Information"] = &acl.Route{
		Name:          "Cost - Salary Information",
		Title:         "Salaries",
		Description:   "Disabling salary data will remove salary and cost data from all views, including 2x2 cost oriented matrixes",
		Path:          "",
		CustomizedURN: "urn:feature:cost/all",
		Public:        false,
		Hidden:        false,
		Admin:         false,
	}

	pubRoutes := make([]*acl.Route, 0)
	for k, r := range proutes.Routes {
		r.Name = k
		if r.Public {
			pubRoutes = append(pubRoutes, r)
		}
		if r.Auth != nil && r.Auth.Features != nil {
			for n, f := range r.Auth.Features {
				f.Wire(r, n)
			}
		}
		routes = append(routes, r)
	}

	err = PackageTemplate.Execute(f, struct {
		Timestamp    time.Time
		Dir          string
		Routes       []*acl.Route
		PublicRoutes []*acl.Route
	}{
		Timestamp:    time.Now(),
		Dir:          dir,
		Routes:       routes,
		PublicRoutes: pubRoutes,
	})

	return err
}

func generateSQL(routes []*acl.Route) error {
	cwd, _ := os.Getwd()
	migration := filepath.Join(cwd, "schema", "migrations")
	os.MkdirAll(migration, 0755)
	migration = filepath.Join(migration, "20171108110099_rbac.sql")
	mf, err := os.Create(migration)
	if err != nil {
		return err
	}
	fmt.Fprintln(os.Stderr, ">> Creating migration file", fileutil.ResolveFileName(migration))
	defer mf.Close()
	err = SQLTemplate.Execute(mf, struct {
		Routes          []*acl.Route
		ResourceTable   string
		ResourceColumns string
		IDColumn        string
		RoleTable       string
		RoleColumns     string
		Roles           []role
	}{
		Routes:          routes,
		ResourceTable:   "`acl_resource`",
		ResourceColumns: "`id`,`checksum`,`urn`,`description`,`title`,`public`,`hidden`,`admin`,`created_at`",
		IDColumn:        "`id`",
		RoleTable:       "`acl_role`",
		RoleColumns:     "`id`,`name`,`description`,`created_at`",
		Roles: []role{
			role{hash.Values("admin"), "admin", "Administrative role"},
			role{hash.Values("executive"), "executive", "Executive role"},
			role{hash.Values("manager"), "manager", "Manager role"},
		},
	})
	return err
}

func runBindata() error {

	cArgs := []string{
		"-o", "./schema/migrate/bindata.go",
		"-pkg", "migrate",
		"-ignore=\\\\.go",
		"-ignore=\\\\.DS_Store",
		"./schema/migrations",
	}
	fmt.Println(fmt.Sprintf("Running: go-bindata %v", cArgs))
	c := exec.CommandContext(context.Background(), "go-bindata", cArgs...)
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr

	return c.Run()
}

type role struct {
	id          string
	name        string
	description string
}

func escape(s string) string {
	return fmt.Sprintf(`"%s"`, s)
}

// SQLValues returns sql values
func (r role) SQLValues() template.HTML {
	s := []string{
		escape(r.id),
		escape(r.name),
		escape(r.description),
		fmt.Sprintf("%v", datetime.TimeToEpoch(time.Now())),
	}
	return template.HTML(strings.Join(s, ","))
}

// SQLID returns sql id
func (r role) SQLID() template.HTML {
	return template.HTML(escape(r.id))
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
	rootCmd.Flags().String("webroot", os.Getenv("PP_WEBROOT"), "path to webapp repo")
}
