package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
    "os"

	"hacker-laws-cli/lib/repo"
	"hacker-laws-cli/lib/parser"
)

func main() {
	hackerLaws := repo.New()
    arguments  := os.Args[1:]

    if len(arguments) == 0 {
        DisplayHelp()
        return
    }

    if arguments[0] == "help" {
        DisplayHelp()
        return
    }

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

    parser.Parse(responseString, &hackerLaws)

    if arguments[0] == "list" {
        hackerLaws.DisplayContentList()
        return
    }
    
    if arguments[0] == "random" {
        randomContent := hackerLaws.RandomContent()
        randomContent.Display()
        return
    }
}

func DisplayHelp() {
    fmt.Println("Options for the command:")
    fmt.Println(PadLeft("help", " ", 12), " ", "To display argument list.")
    fmt.Println(PadLeft("list", " ", 12), " ", "To list the laws and principles.")
    fmt.Println(PadLeft("random", " ", 12), " ", "To display random law or principles.")
}

func PadLeft(str, pad string, lenght int) string {
    for {
        str = pad + str
        if len(str) >= lenght {
            return str[0:lenght]
        }
    }
}
