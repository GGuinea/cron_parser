package main

import (
	"cron_expression_parser/parser"
	"fmt"
	"os"
)

func main() {
	cliArgs := os.Args[1:]

	if len(cliArgs) != 1 {
		fmt.Println("Please provide one argument")
		os.Exit(1)
	}

	parser := parser.NewParser()
	err := parser.Parse(cliArgs[0])
	if err != nil {
		panic(err)
	}
	parser.PrintCurrentCronExpression()
}
