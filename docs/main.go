package main

import (
	"fmt"
	"os"

	"github.com/preferit/cotalk"
)

func main() {
	page := cotalk.Presentation().Page()
	if err := page.SaveAs("index.html"); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
