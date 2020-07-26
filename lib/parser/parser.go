package parser

import (
	"strings"

	"hacker-laws-cli/lib/repo"
)

func Parse(readme string, r *repo.Repo) {
	var cat = 0

	var parseStatus = false
	var content repo.Content
	for _, line := range strings.Split(strings.TrimSuffix(readme, "\n"), "\n") {
		if IsCategory(line) {
			if !IsCategoryIgnored(line) {
				if LineToTitle(line) == "Laws" {
					cat = 0
				} else {
					if cat == 0 {
						r.Categories[cat].Contents = append(r.Categories[cat].Contents, content)
						content = repo.NewContent("", "")
					}

					cat = 1
				}

				parseStatus = true

				continue
			} else {
				parseStatus = false
			}
		}

		if parseStatus == true {
			if IsContent(line) {
				if content.Title != "" {
					r.Categories[cat].Contents = append(r.Categories[cat].Contents, content)
				}

				content = repo.NewContent(LineToTitle(line), "")
				continue
			}

			content.Body = content.Body + line + "\n"
		}
	}

	r.Categories[cat].Contents = append(r.Categories[cat].Contents, content)
}

func IsCategory(line string) bool {
	return strings.HasPrefix(line, "## ") && !strings.HasPrefix(line, "### ")
}

func IsContent(line string) bool {
	return strings.HasPrefix(line, "### ")
}

func IsCategoryIgnored(line string) bool {
	ignoreList := []string{"Reading List", "Contributing", "TODO", "Introduction", "Translations", "Online Resources", "Podcast"}
	str := LineToTitle(line)
	for _, s := range ignoreList {
		if s == str {
			return true
		}
	}

	return false
}

func LineToTitle(line string) string {
	line = strings.Replace(line, "#", "", -1)
	line = strings.Trim(line, " ")

	return line
}
