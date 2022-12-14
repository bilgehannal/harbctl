package operations

import (
	"github.com/bilgehannal/harbctl/internal/pkg/args"
)

type OperationBuilder interface {
	CreateOperation(args.Args) Operation
}

type Operation interface {
	Run()
}

func GetOperationBuilderMap() map[string]OperationBuilder {
	return map[string]OperationBuilder{
		args.VerbLogin: LoginOperationBuilder{},
		"get":          GettingOperationBuilder{},
	}
}

func GetOperationBuilder(args args.Args) OperationBuilder {
	operationBuilders := GetOperationBuilderMap()
	if val, ok := operationBuilders[args.Verb.Value]; ok {
		return val
	}
	return EmptyOperationBuilder{}
}

func GetOperation(args args.Args) Operation {
	return GetOperationBuilder(args).CreateOperation(args)
}
