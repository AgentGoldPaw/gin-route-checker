package gin_route_checker

type RouteOptions struct {
	Method string
	StatusCode int
}

type Routes map[string]*RouteOptions