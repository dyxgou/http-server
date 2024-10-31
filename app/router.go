package main

import "sync"

type Router interface {
	Has(string) bool
	Add(string)
}

type Routes struct {
	routesMu sync.Mutex
	routes   map[string]struct{}
}

func NewRoutes() *Routes {
	return &Routes{
		routes:   make(map[string]struct{}),
		routesMu: sync.Mutex{},
	}
}

func (r *Routes) Has(path string) bool {
	_, ok := r.routes[path]

	return ok
}

func (r *Routes) Add(path string) {
	r.routesMu.Lock()
	defer r.routesMu.Unlock()
	r.routes[path] = struct{}{}
}
