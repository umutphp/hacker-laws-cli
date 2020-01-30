package repo

import (
	"fmt"
    "math/rand"
    "time"
    "strings"
    "regexp"

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

	c.DisplayTitle(c.Title)

	for _, line := range strings.Split(strings.TrimSuffix(c.Body, "\n"), "\n") {
		c.DisplayLine(line)
	}

	c.DisplayFooter()
}

func (c *Content) DisplayTitle(title string) {
	d := color.New(color.FgGreen, color.Bold)
	d.Println("-----------------------------------------------------")
	d.Println(strings.Trim(title, "\n"))
	d.Println("-----------------------------------------------------")
}

func (c *Content) DisplayLine(line string) {
	if strings.HasPrefix(line, ">") {
		color.Yellow(strings.Trim(line, "\n"))
		return
	}

	//color.White(strings.Trim(line, "\n"))
	color.White(strings.Trim(StripMDTags(line), "\n"))
}

func (c *Content) DisplayFooter() {
	d := color.New(color.FgGreen, color.Bold)
	d.Println("-----------------------------------------------------")
	d.Println("         github.com/dwmkerr/hacker-laws by Dave Kerr ")
	d.Println("-----------------------------------------------------")
}

func StripMDTags(line string) string {
	re   := regexp.MustCompile(`\[(.*?)\]\((.*?)\)`)
	line  = re.ReplaceAllString(line, `$1`)

	return line
}

