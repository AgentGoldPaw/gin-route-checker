package gin_route_checker

type RouteOptions struct {
	Methods    []string
	StatusCode int
	Headers    map[string]string
}

type Routes map[string]*RouteOptions
