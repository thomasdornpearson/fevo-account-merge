# fevo-account-merge

Merges accounts provided in an input file provided in json format. An example is provided in accounts.json

Example input
```json
[
  {
    "application": "x",
    "emails": [
      "a",
      "b",
      "c"
    ],
    "name": "Person 1"
  },
  {
    "application": "y",
    "emails": [
      "c",
      "d"
    ],
    "name": "Person 1"
  },
  {
    "application": "y",
    "emails": [
      "c",
      "g"
    ],
    "name": "Person 2"
  }
]
```
Example output
```json
[
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
]
```
# Usage

Requires go 1.18

## From command line in script mode

>cd fevo-account-merge/cmd/account <br />
>go get <br />
>go run main.go -help <br />
>go run main.go -accounts-json -input=../../accounts.json -output=../../accountsOut.json -pretty-print

## Compile project

>cd fevo-account-merge/cmd/account <br />
>go get <br />
>go build

## Running in exe mode
- To pretty print json include the **p** option at the end
>./account -accounts-json -input=../../accounts.json -output=../../accountsOut.json -pretty-print <br />
>./account -accounts-json -input=../../accounts.json -output=../../accountsOut.json


## Running tests
- To run tests
>cd fevo-account-merge/test/account_test <br />
>go test