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
	case *ast.Target:
		return evalTarget(node)
	case *ast.CreateStatement:
		evalCreateStatement(node)
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

func evalCreateStatement(node *ast.CreateStatement){
	target := Eval(node.Target)
	name := Eval(node.Name)

	fmt.Printf("EXECUTE CREATE OPERATION\n")
	fmt.Printf("%v\n", target)
	fmt.Printf("%v\n", name)
}

func evalTarget(node *ast.Target) object.Object{
	return &object.String{ Value: node.Value }
}

func evalIdentifier(node *ast.Identifier) object.Object{
	return &object.String{ Value: node.Value }
}