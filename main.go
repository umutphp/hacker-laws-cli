package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
    "os"
    "time"

	"hacker-laws-cli/lib/repo"
	"hacker-laws-cli/lib/parser"

    "github.com/alecthomas/chroma/quick"
    colorable "github.com/mattn/go-colorable"
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

    responseString := ParseHackerLawsRepo()

    parser.Parse(responseString, &hackerLaws)

    if arguments[0] == "list" {
        hackerLaws.DisplayContentList()
        return
    }
    
    if arguments[0] == "random" {
        randomContent := hackerLaws.RandomContent()
        // randomContent.Display()

        stdout1 := colorable.NewColorableStdout()
        content := "\n" + "## " + randomContent.Title + "\n"
        content  = content + repo.StripMDTags(randomContent.Body)
        content  = content + "-----------------------------------------------------" + "\n"
        content  = content + "         github.com/dwmkerr/hacker-laws by Dave Kerr " + "\n"
        content  = content + "-----------------------------------------------------" + "\n"

        // colorize and print the documentation
        if err := quick.Highlight(stdout1, string(content), "markdown", "terminal256", "native"); err != nil {
            panic(err)
        }

        // turn off the coloring and reprint a newline
        stdout1.Write([]byte("\033[0m\n"))
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

func ParseHackerLawsRepo() string {
    home, err := os.UserHomeDir()

    if err != nil {
        fmt.Println(err)
        return ""
    }

    cacheFile := home + string(os.PathSeparator) + ".hlcache"

    if FileExists(cacheFile) && FileUptoDate(cacheFile) {
        content, err := ioutil.ReadFile(cacheFile)
        
        if err != nil {
            fmt.Println(err)
        } else {
            return string(content)
        }
    }

    response, err := http.Get("https://raw.githubusercontent.com/dwmkerr/hacker-laws/master/README.md")
    if err != nil {
        fmt.Println(err)
    }
    defer response.Body.Close()

    responseData, err := ioutil.ReadAll(response.Body)

    if err != nil {
        fmt.Println(err)
        return ""
    } 

    responseString := string(responseData)

    file, err := os.Create(cacheFile)

    if err != nil {
        fmt.Println(err)
        return ""
    }

    file.WriteString(responseString)

    return responseString
}

func FileExists(filename string) bool {
    info, err := os.Stat(filename)
    if os.IsNotExist(err) {
        return false
    }
    return !info.IsDir()
}

func FileUptoDate(filename string) bool {
    info, _ := os.Stat(filename)
    modifiedtime := info.ModTime()

    return !IsOlderThanOneDay(modifiedtime)
}

func IsOlderThanOneDay(t time.Time) bool {
      return time.Now().Sub(t) > 24*time.Hour
}
