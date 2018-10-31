package cmd

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/pinpt/go-common/datetime"
	"github.com/pinpt/go-common/fileutil"
	"github.com/pinpt/go-common/hash"
	"github.com/pinpt/go-schema/acl"
	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate all code gen related files",
	Run: func(cmd *cobra.Command, args []string) {
		dir, _ := cmd.Flags().GetString("webroot")
		if err := runGenerateCmd(dir); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if err := runBindataCmd(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func runGenerateCmd(dir string) error {
	if dir == "" {
		return fmt.Errorf("environment variable PP_WEBROOT, or flag --webroot, is not defined")
	}
	if !fileutil.FileExists(dir) {
		return fmt.Errorf("cannot find " + dir)
	}

	routes, err := generateRoutesGoFile(dir)
	if err != nil {
		return err
	}
	if err := generateSQL(routes); err != nil {
		return err
	}
	return nil
}

// keep the date the same so that subsequent generations will always use the same
// date and the diff will not change if the routes are the same
var timeConst = time.Unix(1541002044312, 0)

func generateRoutesGoFile(dir string) ([]*acl.Route, error) {
	routefn := filepath.Join(dir, "src", "routes.json")
	if !fileutil.FileExists(routefn) {
		routefn = path.Join(dir, "routes.json")
		if !fileutil.FileExists(routefn) {
			return nil, fmt.Errorf("cannot find %v", routefn)
		}
	}

	cwd, _ := os.Getwd()
	fn := filepath.Join(cwd, "acl", "routes.go")
	f, err := os.Create(fn)
	if err != nil {
		return nil, err
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

	routes := make([]*acl.Route, 0)
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

	sort.SliceStable(routes, func(i, j int) bool {
		return routes[i].Name < routes[j].Name
	})

	err = PackageTemplate.Execute(f, struct {
		Timestamp    time.Time
		Dir          string
		Routes       []*acl.Route
		PublicRoutes []*acl.Route
	}{
		Timestamp:    timeConst,
		Dir:          dir,
		Routes:       routes,
		PublicRoutes: pubRoutes,
	})

	return routes, err
}

func generateSQL(routes []*acl.Route) error {
	cwd, _ := os.Getwd()
	migration := filepath.Join(cwd, "migrations")
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

func init() {
	rootCmd.AddCommand(generateCmd)
	generateCmd.Flags().String("webroot", os.Getenv("PP_WEBROOT"), "path to webapp repo")
}
