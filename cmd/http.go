package cmd

func ServeHTTP() {}

type Dependency struct{}

func dependencyInject() Dependency {
	return Dependency{}
}
