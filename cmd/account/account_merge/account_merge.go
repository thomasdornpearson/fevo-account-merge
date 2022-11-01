package account_merge

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/fatih/set"
	"github.com/thomasdornpearson/fevo-account-merge/cmd/account/utils"
	"io/ioutil"
	"os"
	"reflect"
	"sort"
)

type Accounts struct {
	Accounts []Account `json:"accounts"`
}

type Account struct {
	Application string   `json:"application"`
	Name        string   `json:"name"`
	Emails      []string `json:"emails"`
}

type AccountOut struct {
	Applications []string `json:"application"`
	Name         string   `json:"name"`
	Emails       []string `json:"emails"`
}

func AccountMerge(action string, readFile string, outFile string, print bool) string {
	switch {
	case action == "accounts-json":
		var names = make(map[string]string)
		var applications = make(map[string]set.Interface)
		var emails = make(map[string]string)
		var common = make(map[string]set.Interface)
		jsonFile, err := os.Open(readFile)
		defer jsonFile.Close()
		if err != nil {
			fmt.Println(err)
		} else {
			byteValue, err := ioutil.ReadAll(jsonFile)
			if err != nil {
				fmt.Println(err)
			} else {
				var accounts []Account
				json.Unmarshal(byteValue, &accounts)
				for _, account := range accounts {
					for _, email := range account.Emails {
						names[email] = account.Name
						if _, ok := applications[email]; ok == false {
							applications[email] = set.New(set.NonThreadSafe)
						}
						applications[email].Add(account.Application)
						emails[email] = email
					}
				}
				for _, account := range accounts {
					parentEmail := account.Emails[0]
					for i := 1; i < len(account.Emails); i++ {
						emails[account.Emails[i]] = parentEmail
					}
				}
				fmt.Println(emails)
				for _, account := range accounts {
					parentEmail := utils.Find(account.Emails[0], emails)
					if _, ok := common[parentEmail]; ok == false {
						common[parentEmail] = set.New(set.NonThreadSafe)
						common[parentEmail].Add(parentEmail)
					}
					for i := 1; i < len(account.Emails); i++ {
						if _, ok := common[parentEmail]; ok == true {
							common[parentEmail].Add(account.Emails[i])
							for _, value := range applications[account.Emails[i]].List() {
								applications[parentEmail].Add(value)
							}
						}
					}
				}
				var result []AccountOut
				for _, parent := range reflect.ValueOf(common).MapKeys() {
					stringParent := parent.Interface().(string)
					emails := set.StringSlice(common[stringParent])
					sort.Strings(emails)
					applications := set.StringSlice(applications[stringParent])
					sort.Strings(applications)
					name := names[stringParent]
					result = append(result, AccountOut{
						Name:         name,
						Emails:       emails,
						Applications: applications,
					})
				}
				e, err := json.Marshal(result)
				if err != nil {
					fmt.Println(err)
				} else {
					stringReq := utils.JsonPrettyPrint(string(e))
					if print == true {
						fmt.Println(stringReq)
					}
					file, _ := json.MarshalIndent(result, "", " ")
					err = ioutil.WriteFile(outFile, file, 0644)
					if err != nil {
						fmt.Println(err)
					}
					return stringReq
				}
			}
		}
		return "failed"
	case action == "help":
		fmt.Println("\n" + `go run main.go -accounts-json -input=../../accounts.json -output=../../accountsOut.json -pretty-print -> To run as the script version` + "\n" +
						   `./account -accounts-json -input=../../accounts.json -output=../../accountsOut.json -pretty-print -> Run and print results in console` + "\n" +
						   `./account -accounts-json -input=../../accounts.json -output=../../accountsOut.json -> Run in quiet mode`)
		return "help"	
	default:
		flag.Usage()
		return "usage"
	}
	return "no_case"
}
