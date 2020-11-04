package main

import (
  "os"
  "fmt"
  "log"
  "regexp"
  "strings"
  "github.com/gocolly/colly"
)

type API struct {
  Title       string
  Description string
  Code        string
  Return      string
  ExampleCode string
  Source      string
}

func check(err error){
  if(err != nil){
    log.Fatal(err)
  }
}

func google_search(search_query string) string{
  base_url := "https://www.google.com/search?q="
  url := base_url + search_query + "+msdn"

  var result string

  collector := colly.NewCollector(
    colly.UserAgent("Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:47.0) Gecko/20100101 Firefox/47.0"),
  )

  collector.OnHTML("html", func(e *colly.HTMLElement){
    sellector := e.DOM.Find("div.g")
    for node := range sellector.Nodes{
      item := sellector.Eq(node)
      link_tag := item.Find("a")
      link, _ := link_tag.Attr("href")
      link = strings.Trim(link, " ")
  
      re, err := regexp.Compile("https://docs.microsoft.com/en-us/windows+")
      check(err)
  
      if link != "" && link != "#" && re.MatchString(link) {
        result = link
        return
      }
    }
  })

  collector.OnError(func(r *colly.Response, err error) {
    log.Fatal(err)
	})

  collector.Visit(url)

  return result
}

func parse_msdn(msdn_url string) *API{
  api := API{}

  collector := colly.NewCollector(
    colly.AllowedDomains("docs.microsoft.com"),
    colly.UserAgent("Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:47.0) Gecko/20100101 Firefox/47.0"),
  )

  collector.OnHTML("meta", func(e *colly.HTMLElement){
    if e.Attr("property") == "og:title"{
      api.Title = e.Attr("content")
      return
    }
    if e.Attr("property") == "og:description"{
      api.Description = e.Attr("content")
      return
    }
    if e.Attr("property") == "og:url"{
      api.Source = e.Attr("content")
      return
    }
  })

  collector.OnHTML("pre", func(e *colly.HTMLElement){
    if e.Index == 0 {
      api.Code = e.Text
      return
    }
    if e.Index == 1{
      api.ExampleCode = e.Text
      return
    }
  })

  collector.OnHTML("p", func(e *colly.HTMLElement){
    re, err := regexp.Compile(".*(no error occurs|succeeds|fails|failure|returns|return value|returned).*(no error occurs|succeeds|fails|failure|returns|return value|returned)[^.]+")
    check(err)
    match_str := re.FindString(e.Text)

    if match_str != ""{
      api.Return += match_str + ". "
      api.Return = strings.ReplaceAll(api.Return, "\n", " ",)
    }
  })

  collector.OnError(func(r *colly.Response, err error) {
    log.Fatal(err)
	})

  collector.Visit(msdn_url)

  return &api
}

func main(){
  usage := "Usage: ./manw <function_name>"

  if len(os.Args) != 2{
    log.Fatal(usage)
  }

  search_query := os.Args[1]

  msdn_url := google_search(search_query)

  if msdn_url == ""{
    log.Fatal("[!] ERROR: Unable to find this Windows API function!")
  }

  api := parse_msdn(msdn_url)

  fmt.Printf("%s\n\n", api.Title)
  fmt.Printf("%s\n\n", api.Description)
  fmt.Printf("%s\n\n", api.Code)

  if api.Return != ""{
    fmt.Printf("Return value: %s\n\n", api.Return)
  }

  if api.ExampleCode != ""{
    fmt.Printf("Example code:\n\n%s\n\n", api.ExampleCode)
  }

  fmt.Printf("Source: %s\n\n", api.Source)
}