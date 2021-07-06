package main

import(
	"fmt"
	"os"
	"cql/repl"
)

const logo = `
  ■■■■■    ■■■■    ■     
 ■■   ■   ■■   ■■  ■     
■■       ■■     ■  ■     
■        ■      ■■ ■     
■        ■      ■■ ■     
■        ■      ■■ ■     
■■       ■■  ■  ■  ■     
 ■■   ■   ■■  ■■■  ■     
  ■■■■     ■■■■■■  ■■■■■■
`

func main(){
	fmt.Printf("%v\n", logo)
	fmt.Printf("version: %v\n", "0.0.1")
	fmt.Printf("CQL is a language that manipulates csv files like SQL\n")

	repl.Start(os.Stdin, os.Stdout)
}