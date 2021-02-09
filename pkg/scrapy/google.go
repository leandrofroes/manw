package scrapy

import (
  "log"
  "regexp"
  "strings"

  "manw/pkg/utils"
  "github.com/gocolly/colly"
)

func GoogleMSDNSearch(search, searchAux string) string{
  baseUrl := "https://www.google.com/search?q="
  url := baseUrl + search + searchAux

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

func GoogleKernelSearch(search, searchAux string) string{
  baseUrl := "https://www.google.com/search?q="

  if(!strings.HasPrefix(search, "_")){
    search = "_" + search
  }

  url := baseUrl + search + searchAux

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

      re, err := regexp.Compile("https://www.vergiliusproject.com/kernels+")
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
