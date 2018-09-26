// Code generated by go generate; DO NOT EDIT.
// This file was generated on {{ .Timestamp }}
// using data from {{ .Dir }}
package acl

import (
	"regexp"
	"sort"
)

func init() {
	isPublicRoute = regexp.MustCompile("({{makePathArray .PublicRoutes}})")
	isValidRoute = regexp.MustCompile("({{makePathArray .Routes}})")
	routes = []*Route{
		{{- range .Routes }}
		&Route{
			Name:           "{{ .Name }}",
			Path:	        "{{ .Path }}",
			Title:          "{{ .Title }}",
			Public:         {{ .Public }},
			Hidden:         {{ .Hidden }},
			Admin:          {{ .Admin }},
			Description:    "{{ .Description }}",
			{{if .Auth -}}
			Auth: &Auth{
				Features: map[string]*Feature{
					{{ range $key, $feature := .Auth.Features -}}
					"{{ $key }}": &Feature{
						Description: "{{ $feature.Description }}",
					},
					{{- end }}
				},
			},
			{{- end }}
		},
		{{- end }}
	}
	// sort the routes by path to be consistent
	sort.Slice(routes, func(i, j int) bool {
		return routes[i].Name {{ noescape "<" }} routes[j].Name
	})
	// initialize the route data
	for _, route := range routes {
		if route.Auth != nil {
			for name, feature := range route.Auth.Features {
				feature.name = name
				feature.route = route
			}
		}
	}
}