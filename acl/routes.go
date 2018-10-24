// Code generated by go generate; DO NOT EDIT.
// This file was generated on 2018-10-24 12:18:03.833281 -0500 CDT m=&#43;9.427355619
// using data from /Users/developer/pinpt/apps/webapp-react
package acl

import (
	"regexp"
	"sort"
)

func init() {
	isPublicRoute = regexp.MustCompile("(^\\/error$|^\\/welcome$)")
	isValidRoute = regexp.MustCompile("(^\\/data\\/sprints$|^\\/data\\/commits$|^\\/data\\/issues$|^\\/data\\/projects$|^\\/teams\\/performance$|^\\/signal\\/changes-per-commit\\/person\\/(\\w+)?$|^\\/data\\/locations$|^\\/data\\/repositories$|^\\/admin\\/roles$|^$|^\\/people\\/performance$|^\\/signal\\/cost\\/team\\/(\\w+)?$|^\\/signal\\/rework-rate\\/team\\/(\\w+)?$|^\\/signal\\/cycle-time\\/team\\/(\\w+)?$|^\\/signal\\/sprint-volatility\\/team\\/(\\w+)?$|^\\/data\\/people$|^\\/$|^\\/issue\\/(\\w+)?$|^\\/issues\\/performance$|^\\/signal\\/initiative-issues\\/team\\/(\\w+)?$|^\\/locations\\/performance$|^\\/signal\\/commits\\/person\\/(\\w+)?$|^\\/signal\\/cycle-time\\/person\\/(\\w+)?$|^\\/signal\\/traceability\\/person\\/(\\w+)?$|^\\/admin\\/cost-center$|^\\/cost-center\\/(\\w+)?$|^\\/repository\\/(\\w+)?$|^\\/sprint\\/(\\w+)?$|^\\/issues\\/workflow\\/(\\w+)?\\/(\\w+)?\\/(\\w+)?$|^\\/signal\\/defects-rate\\/team\\/(\\w+)?$|^\\/signal\\/issues-completed\\/team\\/(\\w+)?$|^\\/signal\\/on-time-delivery\\/team\\/(\\w+)?$|^\\/signal\\/throughput\\/team\\/(\\w+)?$|^\\/file\\/(\\w+)?$|^\\/location\\/(\\w+)?$|^\\/person\\/(\\w+)?$|^\\/signal\\/scheduled-rate\\/team\\/(\\w+)?$|^\\/admin$|^\\/admin\\/mapping$|^\\/error$|^\\/project\\/(\\w+)?$|^\\/signal\\/backlog-change\\/team\\/(\\w+)?$|^\\/signal\\/code-ownership\\/person\\/(\\w+)?$|^\\/data\\/teams$|^\\/welcome$|^\\/team\\/(\\w+)?$|^\\/signal\\/issues-worked\\/person\\/(\\w+)?$|^\\/signal\\/rework-rate\\/person\\/(\\w+)?$|^\\/signal\\/delivered-vs-committed\\/team\\/(\\w+)?$|^\\/signal\\/sprint-health\\/team\\/(\\w+)?$|^\\/commit\\/(\\w+)?$|^\\/language\\/(\\w+)?$|^\\/signal\\/defects-density\\/team\\/(\\w+)?$|^\\/signal\\/innovation-rate\\/team\\/(\\w+)?$|^\\/admin\\/people$)")
	routes = []*Route{
		&Route{
			Name:           "DataSprints",
			Path:	        "/data/sprints",
			Title:          "Data - Sprints",
			Public:         false,
			Hidden:         false,
			Admin:          false,
			Description:    "",
			
		},
		&Route{
			Name:           "DataCommits",
			Path:	        "/data/commits",
			Title:          "Data - Commits",
			Public:         false,
			Hidden:         false,
			Admin:          false,
			Description:    "",
			
		},
		&Route{
			Name:           "DataIssues",
			Path:	        "/data/issues",
			Title:          "Data - Issues",
			Public:         false,
			Hidden:         false,
			Admin:          false,
			Description:    "",
			
		},
		&Route{
			Name:           "DataProjects",
			Path:	        "/data/projects",
			Title:          "Data - Projects",
			Public:         false,
			Hidden:         false,
			Admin:          false,
			Description:    "",
			
		},
		&Route{
			Name:           "TeamPerformance",
			Path:	        "/teams/performance",
			Title:          "Teams - Performance Summary",
			Public:         false,
			Hidden:         false,
			Admin:          false,
			Description:    "Compare teams across a range of signals",
			
		},
		&Route{
			Name:           "PersonChangesPerCommit",
			Path:	        "/signal/changes-per-commit/person/:id",
			Title:          "People - Changes Per Commit",
			Public:         false,
			Hidden:         false,
			Admin:          false,
			Description:    "",
			
		},
		&Route{
			Name:           "DataLocations",
			Path:	        "/data/locations",
			Title:          "Data - Locations",
			Public:         false,
			Hidden:         false,
			Admin:          false,
			Description:    "",
			
		},
		&Route{
			Name:           "DataRepositories",
			Path:	        "/data/repositories",
			Title:          "Data - Repositories",
			Public:         false,
			Hidden:         false,
			Admin:          false,
			Description:    "",
			
		},
		&Route{
			Name:           "AdminRoles",
			Path:	        "/admin/roles",
			Title:          "Roles",
			Public:         false,
			Hidden:         false,
			Admin:          true,
			Description:    "Create and manage roles",
			
		},
		&Route{
			Name:           "Cost - Salary Information",
			Path:	        "",
			Title:          "Salaries",
			Public:         false,
			Hidden:         false,
			Admin:          false,
			Description:    "Disabling salary data will remove salary and cost data from all views, including 2x2 cost oriented matrixes",
			
		},
		&Route{
			Name:           "PersonPerformance",
			Path:	        "/people/performance",
			Title:          "People - Performance Summary",
			Public:         false,
			Hidden:         false,
			Admin:          false,
			Description:    "Compare employees across a range of signals",
			
		},
		&Route{
			Name:           "TeamCost",
			Path:	        "/signal/cost/team/:id",
			Title:          "Teams - Cost",
			Public:         false,
			Hidden:         false,
			Admin:          false,
			Description:    "",
			
		},
		&Route{
			Name:           "TeamReworkPercent",
			Path:	        "/signal/rework-rate/team/:id",
			Title:          "Teams - Rework Rate",
			Public:         false,
			Hidden:         false,
			Admin:          false,
			Description:    "",
			
		},
		&Route{
			Name:           "TeamCycleTime",
			Path:	        "/signal/cycle-time/team/:id",
			Title:          "Teams - Cycle Time",
			Public:         false,
			Hidden:         false,
			Admin:          false,
			Description:    "",
			
		},
		&Route{
			Name:           "TeamSprintVolatility",
			Path:	        "/signal/sprint-volatility/team/:id",
			Title:          "Teams - Sprint Volatility",
			Public:         false,
			Hidden:         false,
			Admin:          false,
			Description:    "",
			
		},
		&Route{
			Name:           "DataPeople",
			Path:	        "/data/people",
			Title:          "Data - People",
			Public:         false,
			Hidden:         false,
			Admin:          false,
			Description:    "",
			
		},
		&Route{
			Name:           "Main",
			Path:	        "/",
			Title:          "",
			Public:         false,
			Hidden:         true,
			Admin:          false,
			Description:    "",
			
		},
		&Route{
			Name:           "Issue",
			Path:	        "/issue/:id",
			Title:          "Issue Detail",
			Public:         false,
			Hidden:         true,
			Admin:          false,
			Description:    "",
			
		},
		&Route{
			Name:           "IssuePerformance",
			Path:	        "/issues/performance",
			Title:          "Work - Performance Summary",
			Public:         false,
			Hidden:         false,
			Admin:          false,
			Description:    "Summary of forecasts and delivery status for larger open issues",
			
		},
		&Route{
			Name:           "TeamStrategicPercent",
			Path:	        "/signal/initiative-issues/team/:id",
			Title:          "Teams - Strategic Issues",
			Public:         false,
			Hidden:         false,
			Admin:          false,
			Description:    "",
			
		},
		&Route{
			Name:           "LocationPerformance",
			Path:	        "/locations/performance",
			Title:          "Location - Performance Summary",
			Public:         false,
			Hidden:         false,
			Admin:          false,
			Description:    "Compare locations across a range of signals",
			
		},
		&Route{
			Name:           "PersonCommits",
			Path:	        "/signal/commits/person/:id",
			Title:          "People - Commits",
			Public:         false,
			Hidden:         false,
			Admin:          false,
			Description:    "",
			
		},
		&Route{
			Name:           "PersonCycleTime",
			Path:	        "/signal/cycle-time/person/:id",
			Title:          "People - Cycle Time",
			Public:         false,
			Hidden:         false,
			Admin:          false,
			Description:    "",
			
		},
		&Route{
			Name:           "PersonTraceability",
			Path:	        "/signal/traceability/person/:id",
			Title:          "People - Traceability",
			Public:         false,
			Hidden:         false,
			Admin:          false,
			Description:    "",
			
		},
		&Route{
			Name:           "AdminCostCenters",
			Path:	        "/admin/cost-center",
			Title:          "Cost Centers",
			Public:         false,
			Hidden:         false,
			Admin:          true,
			Description:    "Create and manage cost centers",
			
		},
		&Route{
			Name:           "CostCenter",
			Path:	        "/cost-center/:id",
			Title:          "Cost Center Detail",
			Public:         false,
			Hidden:         true,
			Admin:          false,
			Description:    "",
			
		},
		&Route{
			Name:           "Repository",
			Path:	        "/repository/:id",
			Title:          "Repository Detail",
			Public:         false,
			Hidden:         true,
			Admin:          false,
			Description:    "",
			
		},
		&Route{
			Name:           "Sprint",
			Path:	        "/sprint/:id",
			Title:          "Sprint Detail",
			Public:         false,
			Hidden:         true,
			Admin:          false,
			Description:    "",
			
		},
		&Route{
			Name:           "IssueWorkflow",
			Path:	        "/issues/workflow/:team/:issueType/:interval",
			Title:          "Work - Issue Workflow",
			Public:         false,
			Hidden:         false,
			Admin:          false,
			Description:    "Analysis of typical issue workflow paths",
			
		},
		&Route{
			Name:           "TeamDefectsPercent",
			Path:	        "/signal/defects-rate/team/:id",
			Title:          "Teams - Defect Rate",
			Public:         false,
			Hidden:         false,
			Admin:          false,
			Description:    "",
			
		},
		&Route{
			Name:           "TeamIssuesCompleted",
			Path:	        "/signal/issues-completed/team/:id",
			Title:          "Teams - Completed Issues",
			Public:         false,
			Hidden:         false,
			Admin:          false,
			Description:    "",
			
		},
		&Route{
			Name:           "TeamOnTimeDelivery",
			Path:	        "/signal/on-time-delivery/team/:id",
			Title:          "Teams - On-time Delivery",
			Public:         false,
			Hidden:         false,
			Admin:          false,
			Description:    "",
			
		},
		&Route{
			Name:           "TeamThroughput",
			Path:	        "/signal/throughput/team/:id",
			Title:          "Teams - Throughput",
			Public:         false,
			Hidden:         false,
			Admin:          false,
			Description:    "",
			
		},
		&Route{
			Name:           "File",
			Path:	        "/file/:id",
			Title:          "File Detail",
			Public:         false,
			Hidden:         true,
			Admin:          false,
			Description:    "",
			
		},
		&Route{
			Name:           "Location",
			Path:	        "/location/:id",
			Title:          "Location Detail",
			Public:         false,
			Hidden:         true,
			Admin:          false,
			Description:    "",
			
		},
		&Route{
			Name:           "Person",
			Path:	        "/person/:id",
			Title:          "Person Detail",
			Public:         false,
			Hidden:         true,
			Admin:          false,
			Description:    "",
			
		},
		&Route{
			Name:           "TeamScheduled",
			Path:	        "/signal/scheduled-rate/team/:id",
			Title:          "Teams - Planned Issues",
			Public:         false,
			Hidden:         false,
			Admin:          false,
			Description:    "",
			
		},
		&Route{
			Name:           "Admin",
			Path:	        "/admin",
			Title:          "",
			Public:         false,
			Hidden:         true,
			Admin:          true,
			Description:    "",
			
		},
		&Route{
			Name:           "AdminMapping",
			Path:	        "/admin/mapping",
			Title:          "Team Mapping",
			Public:         false,
			Hidden:         false,
			Admin:          true,
			Description:    "Associate projects and code repositories with teams",
			
		},
		&Route{
			Name:           "Error",
			Path:	        "/error",
			Title:          "Error",
			Public:         true,
			Hidden:         true,
			Admin:          false,
			Description:    "",
			
		},
		&Route{
			Name:           "Project",
			Path:	        "/project/:id",
			Title:          "Project Detail",
			Public:         false,
			Hidden:         true,
			Admin:          false,
			Description:    "",
			
		},
		&Route{
			Name:           "TeamBacklogChangePercent",
			Path:	        "/signal/backlog-change/team/:id",
			Title:          "Teams - Backlog Change",
			Public:         false,
			Hidden:         false,
			Admin:          false,
			Description:    "",
			
		},
		&Route{
			Name:           "PersonCodeOwnership",
			Path:	        "/signal/code-ownership/person/:id",
			Title:          "People - Code Ownership",
			Public:         false,
			Hidden:         false,
			Admin:          false,
			Description:    "",
			
		},
		&Route{
			Name:           "DataTeams",
			Path:	        "/data/teams",
			Title:          "Data - Teams",
			Public:         false,
			Hidden:         false,
			Admin:          false,
			Description:    "",
			
		},
		&Route{
			Name:           "Welcome",
			Path:	        "/welcome",
			Title:          "Welcome",
			Public:         true,
			Hidden:         true,
			Admin:          false,
			Description:    "",
			
		},
		&Route{
			Name:           "Team",
			Path:	        "/team/:id",
			Title:          "Team Detail",
			Public:         false,
			Hidden:         true,
			Admin:          false,
			Description:    "",
			
		},
		&Route{
			Name:           "PersonIssuesWorked",
			Path:	        "/signal/issues-worked/person/:id",
			Title:          "People - Issues Worked",
			Public:         false,
			Hidden:         false,
			Admin:          false,
			Description:    "",
			
		},
		&Route{
			Name:           "PersonReworkPercent",
			Path:	        "/signal/rework-rate/person/:id",
			Title:          "People - Rework Rate",
			Public:         false,
			Hidden:         false,
			Admin:          false,
			Description:    "",
			
		},
		&Route{
			Name:           "TeamDeliveredVsCommitted",
			Path:	        "/signal/delivered-vs-committed/team/:id",
			Title:          "Teams - Delivered vs. Planned",
			Public:         false,
			Hidden:         false,
			Admin:          false,
			Description:    "",
			
		},
		&Route{
			Name:           "TeamSprintHealth",
			Path:	        "/signal/sprint-health/team/:id",
			Title:          "Teams - Sprint Health",
			Public:         false,
			Hidden:         false,
			Admin:          false,
			Description:    "",
			
		},
		&Route{
			Name:           "Commit",
			Path:	        "/commit/:id",
			Title:          "Commit Detail",
			Public:         false,
			Hidden:         true,
			Admin:          false,
			Description:    "",
			
		},
		&Route{
			Name:           "Language",
			Path:	        "/language/:id",
			Title:          "Language Detail",
			Public:         false,
			Hidden:         true,
			Admin:          false,
			Description:    "",
			
		},
		&Route{
			Name:           "TeamDefectsDensity",
			Path:	        "/signal/defects-density/team/:id",
			Title:          "Teams - Defect Density",
			Public:         false,
			Hidden:         false,
			Admin:          false,
			Description:    "",
			
		},
		&Route{
			Name:           "TeamNewFeatures",
			Path:	        "/signal/innovation-rate/team/:id",
			Title:          "Teams - Innovation Rate",
			Public:         false,
			Hidden:         false,
			Admin:          false,
			Description:    "",
			
		},
		&Route{
			Name:           "AdminPeople",
			Path:	        "/admin/people",
			Title:          "People",
			Public:         false,
			Hidden:         false,
			Admin:          true,
			Description:    "Manage your users",
			
		},
	}
	// sort the routes by path to be consistent
	sort.Slice(routes, func(i, j int) bool {
		return routes[i].Name < routes[j].Name
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