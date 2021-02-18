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

  collector.OnHTML("pre", func(e *colly.HTMLElement){
    if(e.Attr("class") == "kernelstruct"){
      kernelInfo = e.Text
    }
  })

  collector.OnError(func(r *colly.Response, err error) {
    log.Fatal(err)
  })

  collector.Visit(url)

  return kernelInfo
}

func RunKernelScraper(search, cachePath string){
  if(cachePath != ""){
    if(!cache.CheckCache(search, cachePath)){
      searchAux := "+kernel+struct+nirsoft"

      url := GoogleKernelSearch(search, searchAux)
    
      if url == ""{
        utils.Warning("Unable to find this Windows Kernel structure.")
      }
    
      kernelInfo := parseKernelInfo(url)
      
      cache.RunKernelCache(search, kernelInfo, cachePath)
    }
  } else {
    searchAux := "+kernel+struct+nirsoft"

    url := GoogleKernelSearch(search, searchAux)
  
    if url == ""{
      utils.Warning("Unable to find this Windows Kernel structure.")
    }
  
    kernelInfo := parseKernelInfo(url)
    
    utils.GenericPrint(kernelInfo)
  }
}
