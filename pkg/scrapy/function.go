package scrapy

import (
  "log"
  "regexp"
  "strings"

  "manw/pkg/utils"
  "manw/pkg/cache"
  "github.com/gocolly/colly"
)

func googleFunctionSearch(s string) string{
  baseUrl := "https://www.google.com/search?q="
  url := baseUrl + s + "+function+msdn"

  var result string

  collector := colly.NewCollector(
    colly.UserAgent("Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:47.0) Gecko/20100101 Firefox/47.0"),
  )

  collector.OnHTML("html", func(e *colly.HTMLElement){
    sellector := e.DOM.Find("div.g")
    for node := range sellector.Nodes{
      item := sellector.Eq(node)
      linkTag := item.Find("a")
      link, _ := linkTag.Attr("href")
      link = strings.Trim(link, " ")

      re, err := regexp.Compile("https://docs.microsoft.com/en-us/windows+")
      utils.CheckError(err)

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

func ParseMSDNFunction(url string) *utils.API{
  api := utils.API{}

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

  collector.OnHTML("meta", func(e *colly.HTMLElement){
    if e.Attr("name") == "req.dll"{
      api.DLL = e.Attr("content")
      return
    }
  })

  collector.OnHTML("pre", func(e *colly.HTMLElement){
    if e.Index == 0 {
      api.CodeA = e.Text
      return
    }
    if e.Index == 1{
      api.CodeB = e.Text
      return
    }
  })

  collector.OnHTML("p", func(e *colly.HTMLElement){
    re, err := regexp.Compile(".*(no error occurs|succeeds|fails|failure|returns|return value|returned).*(no error occurs|succeeds|fails|failure|returns|return value|returned)[^.]+")
    utils.CheckError(err)
    match := re.FindString(e.Text)

    if match != ""{
      api.Return += match + ". "
      api.Return = strings.ReplaceAll(api.Return, "\n", " ",)
    }
  })

  collector.OnError(func(r *colly.Response, err error) {
    log.Fatal(err)
  })

  collector.Visit(url)

  return &api
}

func RunFunctionScraper(search, cachePath string){
  if(!cache.CheckCache(search, cachePath)){
    url := googleFunctionSearch(search)

    if url == ""{
      utils.Warning("Unable to find this Windows function.")
    }

    api := ParseMSDNFunction(url)

    cache.RunFunctionCache(search, cachePath, api)
  }
}
