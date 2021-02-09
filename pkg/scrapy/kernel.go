package scrapy

import(
  "log"

  "manw/pkg/utils"
  "manw/pkg/cache"
  "github.com/gocolly/colly"
)

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
  searchAux := "+windows+kernel+vergilius+project"
    url := GoogleKernelSearch(search, searchAux)

    if url == ""{
      utils.Warning("Unable to find this Windows Kernel structure.")
    }

    kernelInfo := parseKernelInfo(url)

    cache.RunKernelCache(search, kernelInfo, cachePath)
  }
}
