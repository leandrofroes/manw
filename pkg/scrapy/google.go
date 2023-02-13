package scrapy

import (
  "log"
  "regexp"
  "strings"

  "github.com/leandrofroes/manw/pkg/utils"
  
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
    sellector := e.DOM.Find("a")
    for node := range sellector.Nodes{
      item := sellector.Eq(node)
      link, _ := item.Attr("href")

      re, err := regexp.Compile(".microsoft.com/en-us/+")
      utils.CheckError(err)

      if re.MatchString(link) {
        tmpUrl := strings.Split(link, "=")[5]
        result = strings.Split(tmpUrl, "&")[0]
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

  if !strings.HasPrefix(search, "_"){
    search = "_" + search
  }

  url := baseUrl + strings.ToUpper(search) + searchAux

  var result string

  collector := colly.NewCollector(
    colly.UserAgent("Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:47.0) Gecko/20100101 Firefox/47.0"),
    colly.DetectCharset(),
  )

  collector.OnHTML("html", func(e *colly.HTMLElement){
    sellector := e.DOM.Find("a")
    for node := range sellector.Nodes{
      item := sellector.Eq(node)
      link, _ := item.Attr("href")

      re, err := regexp.Compile("https://www.nirsoft.net/kernel_struct/+")
      utils.CheckError(err)

      if re.MatchString(link) {
        tmpUrl := strings.Split(link, "=")[5]
        result = strings.Split(tmpUrl, "&")[0]
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