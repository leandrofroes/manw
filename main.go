package main

import (
  "os"
  "fmt"

  "manw/pkg/scrapy"
  "manw/pkg/config"
  "manw/pkg/cache"

  flag "github.com/spf13/pflag"
)

func main(){
  usage := `NAME

  manw - A multiplatform command line search engine for Windows API.
  
SYNOPSIS: 

  ./manw [-a] [-c] [-k] [-t]
          
OPTIONS:

  -a, --api string    Search for a Windows API Function/Structure.
  -c, --cache         Enable caching feature.
  -k, --kernel string Search for a Windows Kernel Structure.
  -t, --type string   Search for a Windows Data Type.

`

  flag.Usage = func(){
    fmt.Fprintf(os.Stderr, usage)
    os.Exit(1)
  }

  cachePath := config.Load()

  apiFlag := flag.StringP("api", "a", "", "Search for a Windows API Function/Structure.")
  cacheFlag := flag.BoolP("cache", "c", false, "Enable caching feature.")
  dataTypeFlag := flag.StringP("type", "t", "", "Search for a Windows Data Type.")
  kernelFlag := flag.StringP("kernel", "k", "", "Search for a Windows Kernel Structure.")

  flag.Parse()

  apiSearch := *apiFlag
  cacheEnabled := *cacheFlag
  dataTypeSearch := *dataTypeFlag
  kernelSearch := *kernelFlag

  switch {
    case cacheEnabled && apiSearch != "":
      if(!cache.CheckCache(apiSearch, cachePath)){
        cache.RunCacheScraper(apiSearch, cachePath)
      }
    case apiSearch != "":
      scrapy.RunScraper(apiSearch, "api")
    case dataTypeSearch != "":
      scrapy.RunScraper(dataTypeSearch, "type")
    case kernelSearch != "":
      scrapy.RunScraper(kernelSearch, "kernel")
    default:
      fmt.Fprintf(os.Stderr, usage)
      os.Exit(1)
  }

}
