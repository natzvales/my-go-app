package server

import "github.com/natz/go-lib-app/internal/container"

var moduleFactories []ModuleFactory

// RegisterModule registers a module factory to be used when initializing the server
func RegisterModule(factory ModuleFactory) {
	moduleFactories = append(moduleFactories, factory)
}

func LoadModules(c *container.Container) []Module {
	var modules []Module
	for _, factory := range moduleFactories {
		modules = append(modules, factory(c))
	}
	return modules
}
