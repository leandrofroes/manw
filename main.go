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

  ./manw [OPTION...] [STRING]

OPTIONS:

  -f, --function  string  Search for a Windows API Function.
  -s, --structure string  Search for a Windows API Structure.    
  -k, --kernel    string  Search for a Windows Kernel Structure.
  -t, --type      string  Search for a Windows Data Type.
  -c, --cache     bool    Enable cache feature.
`

  flag.Usage = func(){
    fmt.Fprintf(os.Stderr, usage)
    os.Exit(1)
  }

  var (
    functionSearch  string
    structureSearch string
    dataTypeSearch  string
    kernelSearch    string
    cacheFlag       bool
  )

  flag.StringVarP(&functionSearch, "function", "f", "", "Search for a Windows API Function.")
  flag.StringVarP(&structureSearch, "structure", "s", "", "Search for a Windows API Structure.")
  flag.StringVarP(&dataTypeSearch, "type", "t", "", "Search for a Windows Data Type.")
  flag.StringVarP(&kernelSearch, "kernel", "k", "", "Search for a Windows Kernel Structure.")
  flag.BoolVarP(&cacheFlag, "cache", "c", false, "Enable cache feature.")

  flag.Parse()

  if(len(os.Args) < 2){
    fmt.Fprintf(os.Stderr, usage)
    os.Exit(1)
  }

  var cachePath string

  if(cacheFlag){
    cachePath = config.Load()
  }

  switch{
    case functionSearch != "":
      scrapy.RunFunctionScraper(functionSearch, cachePath)
    case structureSearch != "":
      scrapy.RunStructureScraper(structureSearch, cachePath)
    case dataTypeSearch != "":
      scrapy.RunTypeScraper(dataTypeSearch, cachePath)
    case kernelSearch != "":
      scrapy.RunKernelScraper(kernelSearch, cachePath)
    default:
      scrapy.RunFunctionScraper(os.Args[1], cachePath)
  }
}
