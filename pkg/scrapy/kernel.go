package scrapy

import(
  "log"
  "regexp"
  "strings"

  "manw/pkg/utils"
  "manw/pkg/cache"
  "github.com/gocolly/colly"
)

func googleKernelSearch(s string) string{
  baseUrl := "https://www.google.com/search?q="

  if(!strings.HasPrefix(s, "_")){
    s = "_" + s
  }

  url := baseUrl + s + "+windows+kernel+vergilius+project"

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

func parseKernelInfo(url string) string{
  var kernelInfo string
  collector := colly.NewCollector(
    colly.UserAgent("Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:47.0) Gecko/20100101 Firefox/47.0"),
  )

  collector.OnHTML("div", func(e *colly.HTMLElement){
    if(e.Attr("id") == "copyblock"){
      kernelInfo += e.Text
    }
    if(e.Attr("class") == "maincross"){
      kernelInfo += "\n" + e.Text + "\n\n"
    }
  })

  collector.OnError(func(r *colly.Response, err error) {
    log.Fatal(err)
  })

  collector.Visit(url)

  return kernelInfo
}

func RunKernelScraper(search, cachePath string){
  if(!cache.CheckCache(search, cachePath)){
    url := googleKernelSearch(search)

    if url == ""{
      utils.Warning("Unable to find this Windows Kernel structure.")
    }

    kernelInfo := parseKernelInfo(url)

    cache.RunKernelCache(search, kernelInfo, cachePath)
  }
}
