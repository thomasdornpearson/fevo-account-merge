package main

import (
	"flag"
	"os"
	account "github.com/thomasdornpearson/fevo-account-merge/cmd/account/account_merge"
)

func main() {
	var action string
	accountsJson := flag.Bool("accounts-json", false, "provide path to accounts json file, and an option p for print, always writes to accountOut.json")
	help := flag.Bool("help", false, "Example: go run main.go -accounts-json ./accounts.json p")
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