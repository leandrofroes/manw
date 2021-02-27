package scrapy

import (
  "log"
  "regexp"
  "strings"

  "github.com/leandrofroes/manw/pkg/utils"
  "github.com/leandrofroes/manw/pkg/cache"
  
  "github.com/gocolly/colly"
)

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
  if(cachePath != ""){
    if(!cache.CheckCache(search, cachePath)){
      searchAux := "+api+function+msdn"
    
      url := GoogleMSDNSearch(search, searchAux)
    
      if url == ""{
        utils.Warning("Unable to find this Windows function.")
      }
    
      api := ParseMSDNFunction(url)
      
      cache.RunFunctionCache(search, cachePath, api)
    } 
  } else {
    searchAux := "+api+function+msdn"

    url := GoogleMSDNSearch(search, searchAux)
  
    if url == ""{
      utils.Warning("Unable to find this Windows function.")
    }
  
    api := ParseMSDNFunction(url)

    utils.PrintMSDNFunc(api)
  }
}