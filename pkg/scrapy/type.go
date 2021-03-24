package scrapy

import(
  "log"
  "regexp"
  "strings"

  "github.com/leandrofroes/manw/pkg/utils"
  "github.com/leandrofroes/manw/pkg/cache"
  
  "github.com/gocolly/colly"
)

func parseMSDNDataType(search, url string) string{
  var dataTypeInfo string
  collector := colly.NewCollector(
    colly.AllowedDomains("docs.microsoft.com"),
    colly.UserAgent("Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:47.0) Gecko/20100101 Firefox/47.0"),
  )

  collector.OnHTML("tr", func(e *colly.HTMLElement){
    str := strings.ToUpper(search) + "\n"
    re, err := regexp.Compile(str)
    utils.CheckError(err)
    match := re.FindString(e.Text)
    index := strings.Index(e.Text, str)

    if match != "" && index == 1{
      strSlice := strings.Split(e.Text, "\n")
      dataTypeInfo += "\nData Type: "
      for i, str := range strSlice{
        if i > 0 && i < len(strSlice) - 1{
          dataTypeInfo += str + "\n\n"
        }
      }
    }
  })

  collector.OnError(func(r *colly.Response, err error) {
    log.Fatal(err)
  })

  collector.Visit(url)

  return dataTypeInfo
}

func RunTypeScraper(search, cachePath string){
  search = strings.ToLower(search)

  if cachePath != ""{
    if !cache.CheckCache(search, cachePath){
      searchAux := "+windows+data+type+msdn"

      url := GoogleMSDNSearch(search, searchAux)
    
      if url == ""{
        utils.Warning("Unable to find this Windows data type.")
      }
    
      dataTypeInfo := parseMSDNDataType(search, url)

      if dataTypeInfo == ""{
        utils.Warning("Unable to find this Windows data type.")
      }

      cache.RunTypeCache(search, dataTypeInfo, cachePath)
    }
  } else {
    searchAux := "+windows+data+type+msdn"

    url := GoogleMSDNSearch(search, searchAux)
  
    if url == ""{
      utils.Warning("Unable to find this Windows data type.")
    }
  
    dataTypeInfo := parseMSDNDataType(search, url)
    
    if dataTypeInfo == ""{
      utils.Warning("Unable to find this Windows data type.")
    }

    utils.GenericPrint(dataTypeInfo)
  }
}