package utils

import (
	"bytes"
	"encoding/json"
)

func Find(email string, parents map[string]string) string {
	if parents[email] == email {
		return email
	} else {
		return parents[email]
	}
}

func JsonPrettyPrint(in string) string {
    var out bytes.Buffer
    err := json.Indent(&out, []byte(in), "", "\t")
    if err != nil {
        return in
    }
    return out.String()
}