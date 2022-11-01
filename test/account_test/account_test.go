package account_test

import (
	account "github.com/thomasdornpearson/fevo-account-merge/cmd/account/account_merge"
	"github.com/thomasdornpearson/fevo-account-merge/cmd/account/utils"
	"testing"
	"encoding/json"
	"fmt"
	"reflect"
	"os"
)

type AccountOut struct {
	Applications []string `json:"application"`
	Name         string   `json:"name"`
	Emails       []string `json:"emails"`
}

func TestSucessResult(t *testing.T) {
	var result []AccountOut
	resultStr := account.AccountMerge("accounts-json", "./accountsTest.json", "./accountsOutTest.json", false)
	json.Unmarshal([]byte(resultStr), &result)
	fmt.Println(result)
	testResultStr := `[
		{
		 "application": [
		  "x",
		  "y"
		 ],
		 "name": "Person 1",
		 "emails": [
		  "a",
		  "b",
		  "c",
		  "d",
		  "g"
		 ]
		}
	   ]`

	var testResult []AccountOut
	json.Unmarshal([]byte(testResultStr), &testResult)
	fmt.Println(testResult)
	if !reflect.DeepEqual(testResult, result) {
		t.Error("result should be a success, got", result)
	}
	e := os.Remove("./accountsOutTest.json")
    if e != nil {
        fmt.Println(e)
    }
}

func TestHelp(t *testing.T) {
	result := account.AccountMerge("help", "", "", false)
	if "help" != result {
		t.Error("result should be a help, got", result)
	}
}

func TestJson(t *testing.T) {
	result := utils.JsonPrettyPrint(`{"test":"HelloWorld"}`)
	if len(result) != 25 {
		t.Error("result should a length of 25, got", result)
	}
}