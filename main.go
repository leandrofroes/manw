package main

import (
  "os"
  "fmt"

  "manw/pkg/scrapy"
  "manw/pkg/config"

  flag "github.com/spf13/pflag"
)

func main(){
  usage := `NAME

  manw - A multiplatform command line search engine for Windows API.

SYNOPSIS: 

  ./manw [OPTION]... [STRING]

OPTIONS:

  -a, --api     string  Search for a Windows API Function/Structure.
  -c, --cache           Enable caching feature.
  -k, --kernel  string  Search for a Windows Kernel Structure.
  -t, --type    string  Search for a Windows Data Type.

  `

  if(len(os.Args) < 2){
    fmt.Fprintf(os.Stderr, usage)
    os.Exit(1)
  }

  flag.Usage = func(){
    fmt.Fprintf(os.Stderr, usage)
    os.Exit(1)
  }

  cachePath := config.Load()

  apiFlag := flag.StringP("api", "a", "", "Search for a Windows API Function/Structure.")
  cache := flag.BoolP("cache", "c", false, "Enable caching feature.")
  dataTypeFlag := flag.StringP("type", "t", "", "Search for a Windows Data Type.")
  kernelFlag := flag.StringP("kernel", "k", "", "Search for a Windows Kernel Structure.")

  flag.Parse()

  apiSearch := *apiFlag
  cacheFlag := *cache
  dataTypeSearch := *dataTypeFlag
  kernelSearch := *kernelFlag

  switch {
  case apiSearch != "":
    scrapy.RunAPIScraper(apiSearch, cachePath, cacheFlag)
  case dataTypeSearch != "":
    scrapy.RunTypeScraper(dataTypeSearch, cachePath, cacheFlag)
  case kernelSearch != "":
    scrapy.RunKernelScraper(kernelSearch, cachePath, cacheFlag)
  default:
    scrapy.RunAPIScraper(os.Args[1], cachePath, cacheFlag)
  }

}
