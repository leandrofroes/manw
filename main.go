package main

import (
  "os"
  "fmt"

  "github.com/leandrofroes/manw/pkg/scrapy"
  "github.com/leandrofroes/manw/pkg/config"

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
-a, --arch      string  Specify the architecture you are looking for.
-n, --syscall   string  Search for a Windows Syscall ID. If you don't use -a the default value is "x86".
-c, --no-cache  bool    Disable the caching feature.
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
    archSearch      string
    syscallSearch   string
    cacheFlag       bool
  )

  flag.StringVarP(&functionSearch, "function", "f", "", "Search for a Windows API Function.")
  flag.StringVarP(&structureSearch, "structure", "s", "", "Search for a Windows API Structure.")
  flag.StringVarP(&dataTypeSearch, "type", "t", "", "Search for a Windows Data Type.")
  flag.StringVarP(&kernelSearch, "kernel", "k", "", "Search for a Windows Kernel Structure.")
  flag.StringVarP(&archSearch, "arch", "a", "x86", "Specify the architecture you are looking for.")
  flag.StringVarP(&syscallSearch, "syscall", "n", "", "Search for a Windows Syscall ID.")
  flag.BoolVarP(&cacheFlag, "no-cache", "c", false, "Disable the caching feature.")

  flag.Parse()

  if(len(os.Args) < 2){
    fmt.Fprintf(os.Stderr, usage)
    os.Exit(1)
  }

  var cachePath string

  if(!cacheFlag){
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
    case syscallSearch != "":
      scrapy.RunSyscallScraper(syscallSearch, archSearch, cachePath)
    default:
      scrapy.RunFunctionScraper(os.Args[1], cachePath)
  }
}