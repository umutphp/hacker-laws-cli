package main

import (
	"fmt"
	"net/http"
	"io/ioutil"

	"hacker-laws-cli/lib/repo"
	"hacker-laws-cli/lib/parser"
)

func main() {
	hackerLaws := repo.New()

	response, err := http.Get("https://raw.githubusercontent.com/dwmkerr/hacker-laws/master/README.md")
    if err != nil {
        fmt.Println(err)
    }
    defer response.Body.Close()

    responseData, err := ioutil.ReadAll(response.Body)
    if err != nil {
        fmt.Println(err)
    }
 
    responseString := string(responseData)

    fmt.Println(hackerLaws)

    parser.Parse(responseString, &hackerLaws)

    randomContent := hackerLaws.RandomContent()

    randomContent.Display()
}