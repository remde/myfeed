package cli

import (
	"fmt"

	"github.com/remde/myfeed/website"
)

func getArticleURL(websiteTable website.WebsiteTable) (string, string) {
	var index int
	printChooseArticle()
	_, err := fmt.Scanf("%d", &index)
	if err != nil {
		fmt.Println("error")
	}
	return websiteTable[index][0].URL, websiteTable[index][0].CommentsURL
}

func chooseArticleOrComment(articleURL, articleCommentsURL string) string {
	var option int
	printChooseArticleOrComment()
	_, err := fmt.Scanf("%d", &option)
	if err != nil {
		fmt.Println("error")
	}
	if option == 0 {
		return articleURL
	} else {
		return articleCommentsURL
	}
}
