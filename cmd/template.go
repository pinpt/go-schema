package cmd

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"

	"github.com/pinpt/go-schema/acl"
)

// SQLTemplate ...
var SQLTemplate *template.Template

// PackageTemplate ...
var PackageTemplate *template.Template

func init() {
	rep := regexp.MustCompile("(:\\w+)")
	makePath := func(in string) string {
		i := strings.Index(in, "?")
		if i > 0 {
			in = in[0:i]
		}
		in = rep.ReplaceAllStringFunc(in, func(s string) string {
			return "(\\\\w+)?"
		})
		in = strings.Replace(in, "/", "\\\\/", -1)
		return "^" + in + "$"
	}
	makePathArray := func(routes []*acl.Route) template.HTML {
		parts := []string{}
		for _, route := range routes {
			parts = append(parts, makePath(route.Path))
		}
		if len(parts) == 0 {
			// if no path, we need to just return some non-sensical value which will cause the regexp to fail
			return template.HTML("0238450923850923835902834lkjsdfdlkajskdfjklasdfa")
		}
		sort.SliceStable(parts, func(i, j int) bool {
			return parts[i] < parts[j]
		})
		return template.HTML(strings.Join(parts, "|"))
	}
	noescape := func(str string) template.HTML {
		return template.HTML(str)
	}
	fm := template.FuncMap{
		"makePath":      makePath,
		"makePathArray": makePathArray,
		"noescape":      noescape,
	}
	filename, _ := filepath.Abs("./cmd/templates/package.tpl")
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(fmt.Sprintf("Can't read file %v", filename))
		fmt.Println(err)
		os.Exit(-1)
	}
	PackageTemplate = template.Must(template.New("").Funcs(fm).Parse(string(file)))

	filename, _ = filepath.Abs("./cmd/templates/sql.tpl")
	file, err = ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(fmt.Sprintf("Can't read file %v", filename))
		fmt.Println(err)
		os.Exit(-1)
	}

	SQLTemplate = template.Must(template.New("").Funcs(fm).Parse(string(file)))
}
