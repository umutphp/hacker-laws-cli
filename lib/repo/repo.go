package repo

import (
	"fmt"
    "math/rand"
    "time"
    "strings"

    "github.com/fatih/color"
)

type Repo struct {
    Categories []Category
}

type Category struct {
    Name string
    Contents []Content
}

type Content struct {
    Title string
    Slug string
    Body string
}

func New() Repo {
	return Repo{
		Categories: []Category{
			{
				Name: "Laws",
				Contents: []Content{},
			},
			{
				Name: "Principles",
				Contents: []Content{},
			},
		},
	}
}

func NewContent(title string, body string) Content {
	return Content{
		Title: title,
		Body: body,
	}
}

func (r *Repo) RandomContent() Content {
    rand.Seed(time.Now().UnixNano())

    cat     := rand.Intn(1000)% len(r.Categories)
    content := rand.Intn(1000)% len(r.Categories[cat].Contents)

    return r.Categories[cat].Contents[content]
}

func (r *Repo) DisplayContentList() {
	for _, cat := range r.Categories {
		fmt.Println(cat.Name)
		for _, content := range cat.Contents {
			fmt.Println("\t - ", content.Title)
		}
	}
}

func (c *Content) Display() {
	fmt.Println("")

	DisplayTitle(c.Title)

	for _, line := range strings.Split(strings.TrimSuffix(c.Body, "\n"), "\n") {
		DisplayLine(line)
	}
}

func DisplayTitle(title string) {
	d := color.New(color.FgGreen, color.Bold)
	d.Print(strings.Trim(title, "\n"))
}

func DisplayLine(line string) {
	if strings.HasPrefix(line, ">") {
		color.Yellow(strings.Trim(line, "\n"))
		return
	}

	// Skip Wikipedi links
	if strings.HasPrefix(line, "[") {
		return
	}

	color.White(strings.Trim(line, "\n"))
}
