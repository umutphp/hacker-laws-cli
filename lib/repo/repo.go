package repo

import (
	"fmt"
    "math/rand"
    "time"
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
	// Seed the random number generator using the current time (nanoseconds since epoch):
        rand.Seed(time.Now().UnixNano())

        // Hard to predict...but is it possible? I know the day, and hour, probably minute...
        cat := rand.Intn(1000)% len(r.Categories)

        content := rand.Intn(1000)% len(r.Categories[cat].Contents)

        return r.Categories[cat].Contents[content]
}

func (c *Content) Display() {
	fmt.Println(c.Title)
	fmt.Println(c.Body)
}