package evaluator

import (
	"cql/ast"
	"cql/object"
	"fmt"
)

func Eval(node ast.Node) object.Object{
	switch node := node.(type){
	case *ast.Program:
		return evalProgram(node)
	case *ast.Identifier:
		return evalIdentifier(node)
	case *ast.CreateDirStatement:
		CreateDirStatement(node)
	case *ast.CreateCsvStatement:
		CreateCsvStatement(node)
	case *ast.UseStatement:
		UseStatement(node)
	case *ast.FromStatement:
		FromStatement(node)
	case *ast.SelectStatement:
		SelectStatement(node)
	}
	return nil
}

func evalProgram(program *ast.Program)object.Object{
	var result object.Object
	for _, statement := range program.Statements{
		result = Eval(statement)
	}

	return result
}

func CreateDirStatement(node *ast.CreateDirStatement){
	fmt.Printf("%v\n", node.String())
}

func CreateCsvStatement(node *ast.CreateCsvStatement){
	fmt.Printf("%v\n", node.String())
}

func UseStatement(node *ast.UseStatement){
	fmt.Printf("%v\n", node.String())
}

func FromStatement(node *ast.FromStatement){
	fmt.Printf("%v\n", node.String())
}

func SelectStatement(node *ast.SelectStatement){
	fmt.Printf("%v\n", node.String())
}

func evalIdentifier(node *ast.Identifier) object.Object{
	return &object.String{ Value: node.Value }
}