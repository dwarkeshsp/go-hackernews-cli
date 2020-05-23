package main

import (
	"fmt"

	"github.com/mmcdole/gofeed"
)

func main() {
	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL("https://hnrss.org/frontpage")
	for _, item := range feed.Items {
		fmt.Printf("%s\n%s\n\n", item.Title, item.Link)
	}
}
