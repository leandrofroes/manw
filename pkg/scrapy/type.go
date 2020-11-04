package scrapy

import(
  "log"
  "regexp"
  "strings"

  "manw/pkg/utils"
  "manw/pkg/cache"
  "github.com/gocolly/colly"
)

func googleTypeSearch(s string) string{
  baseUrl := "https://www.google.com/search?q="
  url := baseUrl + s + "+windows+data+type+msdn"

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

func parseMSDNDataType(s, url string) string{
  var dataTypeInfo string
  collector := colly.NewCollector(
    colly.AllowedDomains("docs.microsoft.com"),
    colly.UserAgent("Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:47.0) Gecko/20100101 Firefox/47.0"),
  )

  collector.OnHTML("tr", func(e *colly.HTMLElement){
    str := strings.ToUpper(s) + "\n"
    re, err := regexp.Compile(str)
    utils.CheckError(err)
    match := re.FindString(e.Text)
    index := strings.Index(e.Text, str)

    if match != "" && index == 1{
      strSlice := strings.Split(e.Text, "\n")
      dataTypeInfo += "\nData Type: "
      for i, str := range strSlice{
        if(i > 0 && i < len(strSlice) -1){
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

func RunTypeScraper(search, cachePath string, cacheFlag bool){
  if(cacheFlag){
    if(!cache.CheckCache(search, cachePath)){
      url := googleTypeSearch(search)

      if url == ""{
        utils.Warning("Unable to find the provided Windows resource.")
      }

      dataTypeInfo := parseMSDNDataType(search, url)

      cache.RunKernelCache(search, dataTypeInfo, cachePath)
    }
  }else{
    url := googleTypeSearch(search)

    if url == ""{
      utils.Warning("Unable to find the provided Windows resource.")
    }

    dataTypeInfo := parseMSDNDataType(search, url)

    utils.GenericPrint(dataTypeInfo)
  }
}
