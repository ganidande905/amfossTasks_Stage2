package main
import (
	"github.com/gocolly/colly"
	"bufio"
	"fmt"
	"os"
	"net/url"
	"time"
)
func main(){
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a keyword to search in google: ")
	keyword, _ := reader.ReadString('\n')
	scrapeURL := fmt.Sprintf("https://www.google.com/search?q=%s", url.QueryEscape(keyword))
	c := colly.NewCollector(
		colly.AllowedDomains("google.com","www.google.com"),
	)
	resultCount := 0
	const maxResults = 10
	c.OnHTML("h3",func(h *colly.HTMLElement){
		if resultCount < maxResults{
			fmt.Println(h.Text)
			resultCount++
		}
	})
	c.OnRequest(func(r *colly.Request){
		fmt.Printf("Visiting %s\n",r.URL)
	})
	c.OnError(func(r *colly.Response, e error){
		fmt.Printf("ERror: %s\n",e.Error())
	})
	c.Limit(&colly.LimitRule{
        DomainGlob:  "*",
        Delay:       2 * time.Second,
        RandomDelay: 2 * time.Second,
    })
	c.Visit(scrapeURL)
	
}