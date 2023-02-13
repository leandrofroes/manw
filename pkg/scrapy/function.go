package scrapy

import (
  "log"
  "regexp"
  "strings"
  "strconv"

  "github.com/leandrofroes/manw/pkg/utils"
  "github.com/leandrofroes/manw/pkg/cache"
  
  "github.com/gocolly/colly"
)

func ParseMSDNFunction(search, url string) *utils.API{
  api := utils.API{}

  collector := colly.NewCollector(
    colly.UserAgent("Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:47.0) Gecko/20100101 Firefox/47.0"),
  )

  collector.OnHTML("meta", func(e *colly.HTMLElement){
    if e.Attr("property") == "og:title"{
      funcTitle := strings.Split(strings.ToLower(e.Attr("content")), " ")[0]

      if !strings.Contains(funcTitle, search){
        utils.Warning("Unable to find this Windows function.")
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

  collector.OnHTML("meta", func(e *colly.HTMLElement){
    if e.Attr("name") == "req.dll"{
      if e.Attr("content") != ""{
        api.DLL = e.Attr("content")
        return
      }
    }
    if e.Attr("name") == "APILocation"{
      if strings.Contains(e.Attr("content"), ".dll"){
        api.DLL = e.Attr("content")
        return
      }
    }
  })

  collector.OnHTML("pre", func(e *colly.HTMLElement){
    if e.Index == 0 {
      api.CodeA = e.Text
      api.Argc = strconv.Itoa(len(strings.Split(e.Text, "\n")) - 3)

      return
    }
    if e.Index == 1{
      api.CodeB = e.Text
      return
    }
  })

  collector.OnHTML("p", func(e *colly.HTMLElement){
    re, err := regexp.Compile("^(If the function succeeds|The return value|Returns|This function does|If the function fails|If no error occurs)[^.]+.*[.]")
    utils.CheckError(err)
    match := re.FindString(e.Text)
    
    if match != ""{
      api.Return += match
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
  search = strings.ToLower(search)

  if cachePath != ""{
    if !cache.CheckCache(search, cachePath){
      searchAux := "+function+msdn"
    
      url := GoogleMSDNSearch(search, searchAux)

      if url == ""{
        utils.Warning("Unable to find this Windows function.")
      }
    
      api := ParseMSDNFunction(search, url)
      
      cache.RunFunctionCache(search, cachePath, api)
    } 
  } else {
    searchAux := "+function+msdn"

    url := GoogleMSDNSearch(search, searchAux)
  
    if url == ""{
      utils.Warning("Unable to find this Windows function.")
    }
  
    api := ParseMSDNFunction(search, url)

    utils.PrintMSDNFunc(api)
  }
}