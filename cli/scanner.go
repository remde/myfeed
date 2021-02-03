package cli

import (
	"fmt"

	"github.com/remde/myfeed/website"
)

func getArticleURL(websiteTable website.WebsiteTable) string {
	var index int
	_, err := fmt.Scanf("%d", &index)
	if err != nil {
		fmt.Println("error")
	}
	return websiteTable[index][0].URL
}
