package scrapy

import (
  "log"
  "regexp"
  "strings"

  "github.com/leandrofroes/manw/pkg/utils"
  "github.com/leandrofroes/manw/pkg/cache"

  "github.com/gocolly/colly"
)

func ParseMSDNStructure(search, url string) *utils.API{
  api := utils.API{}

  collector := colly.NewCollector(
    colly.UserAgent("Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:47.0) Gecko/20100101 Firefox/47.0"),
  )

  collector.OnHTML("meta", func(e *colly.HTMLElement){
    if e.Attr("property") == "og:title"{
      strucTitle := strings.Split(strings.ToLower(e.Attr("content")), " ")[0]

      if !strings.Contains(strucTitle, search){
        utils.Warning("Unable to find this Windows structure.")
      }
      
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
      api.CodeA = e.Text
      return
    }
    if e.Index == 1 {
      api.CodeB = e.Text
      return
    }
    if e.Index == 2 {
      api.CodeC = e.Text
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

func RunStructureScraper(search, cachePath string){
  search = strings.ToLower(search)

  if cachePath != ""{
    if !cache.CheckCache(search, cachePath){
      searchAux := "+msdn"

      url := GoogleMSDNSearch(search, searchAux)
    
      if url == ""{
        utils.Warning("Unable to find this Windows structure.")
      }
    
      api := ParseMSDNStructure(search, url)

      cache.RunStructureCache(search, cachePath, api)
    }
  } else {
    searchAux := "+api+msdn"

    url := GoogleMSDNSearch(search, searchAux)
  
    if url == ""{
      utils.Warning("Unable to find this Windows structure.")
    }
  
    api := ParseMSDNStructure(search, url)

    utils.PrintMSDNStructure(api)
  }
}