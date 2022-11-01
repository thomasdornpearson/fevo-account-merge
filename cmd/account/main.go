package main

import (
	"flag"
	"os"
	account "github.com/thomasdornpearson/fevo-account-merge/cmd/account/account_merge"
)

func main() {
	var action string
	accountsJson := flag.Bool("accounts-json", false, "provide path to accounts json file and output file, and an option -pretty-print for print")
	help := flag.Bool("help", false, "Example: go run main.go -accounts-json -input=../../accounts.json -output=../../accountsOut.json -pretty-print")
	input := flag.String("input", "", "Input file")
	output := flag.String("output", "", "Output file")
	print := flag.Bool("pretty-print", false, "Print the output to the console")
	flag.Parse()
	if (*accountsJson == true) {
		action = "accounts-json"
	} else if (*help == true) {
		action = "help"
	}	
	account.AccountMerge(action, *input, *output, *print)
	os.Exit(1)
}