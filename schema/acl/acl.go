package acl

import (
	"fmt"
	"html/template"
	"regexp"
	"strings"
	"time"

	"github.com/pinpt/go-common/datetime"
	"github.com/pinpt/go-common/hash"
)

// RouterConfig ...
type RouterConfig struct {
	Routes map[string]*Route `json:"routes"`
}

// Auth is the authorization related details
type Auth struct {
	Features map[string]*Feature `json:"features,omitempty"`
}

// Feature is a feature description
type Feature struct {
	Description string `json:"description"`

	route *Route
	name  string
}

// Wire will wire a feature to its route
func (f *Feature) Wire(r *Route, name string) {
	f.route = r
	f.name = name
}

// Route is a route definition
type Route struct {
	Name        string `json:"name"`
	Path        string `json:"path"`
	Title       string `json:"title,omitempty"`
	Public      bool   `json:"public"`
	Description string `json:"description,omitempty"`
	Auth        *Auth  `json:"auth,omitempty"`
	Hidden      bool   `json:"hidden"`
	Admin       bool   `json:"admin"`

	CustomizedURN string `json:"-"`
}

func cleanRoute(r string) string {
	i := strings.Index(r, "?")
	if i > 0 {
		return r[0:i]
	}
	return r
}

// URN will return the resource URN for the route
func (r *Route) URN() string {
	if r.CustomizedURN != "" {
		// allow it to be overridden
		return r.CustomizedURN
	}
	return fmt.Sprintf("urn:webapp:%s", cleanRoute(r.Path))
}

func quote(s string) string {
	return fmt.Sprintf(`"%s"`, template.HTMLEscapeString(s))
}

// SQLID returns the escaped value for the primary key
func (r *Route) SQLID() template.HTML {
	return template.HTML(fmt.Sprintf(`"%s"`, r.ID()))
}

// ID returns the primary key for the route
func (r *Route) ID() string {
	return hash.Values(r.Path)
}

// DisplayTitle returns the title to display for the route
func (r *Route) DisplayTitle() string {
	if r.Title != "" {
		return r.Title
	}
	return r.Name
}

// DisplayRoute returns true if the route should be displayed
func (r *Route) DisplayRoute() bool {
	return !r.Public && !r.Hidden
}

// SQLValues returns the sql values for the route
func (r *Route) SQLValues() template.HTML {
	values := []string{
		quote(r.ID()),
		quote(hash.Values(r.Name, r.Path, r.Title, r.Public, r.Description, r.Auth)),
		quote(r.URN()),
		quote(r.Description),
		quote(r.Title),
		fmt.Sprintf("%v", r.Public),
		fmt.Sprintf("%v", r.Hidden),
		fmt.Sprintf("%v", r.Admin),
		fmt.Sprintf("%v", datetime.TimeToEpoch(time.Now())),
	}
	return template.HTML(strings.Join(values, ","))
}

var (
	isPublicRoute, isValidRoute *regexp.Regexp
	routes                      []*Route
)

// IsPublicRoute returns true if the route specified is marked as a public route
func IsPublicRoute(path string) bool {
	return isPublicRoute.MatchString(path)
}

// IsValidRoute returns true if the route specified is valid (meaning it's defined)
func IsValidRoute(path string) bool {
	return isValidRoute.MatchString(path)
}

// IsValidRouteID returns true if the route id is valid and optionally a second route object
func IsValidRouteID(id string) (bool, *Route) {
	for _, r := range routes {
		if r.ID() == id {
			return true, r
		}
	}
	return false, nil
}

// Routes returns the array of routes
func Routes() []*Route {
	return routes
}
