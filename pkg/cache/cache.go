package cache

import (
  "os"
  "strings"
  "io/ioutil"

  "manw/pkg/utils"
)

func addAPICache(search, cachePath string, api *utils.API) (entry string){
  entry = strings.ToLower(cachePath + search)

  f, err := os.Create(entry)
  utils.CheckError(err)

  f.WriteString(api.Title + " - " + api.DLL + "\n\n")
  f.WriteString(api.Description + "\n\n")
  f.WriteString(api.Code + "\n\n")

  if api.Return != ""{
    f.WriteString("Return value: " + api.Return + "\n\n")
  }

  if api.ExampleCode != ""{
    f.WriteString("Example code:\n\n" + api.ExampleCode + "\n\n")
  }

  f.WriteString("Source: " + api.Source + "\n\n")

  return entry
}

func addGenericCache(search, data, cachePath string) (entry string){
  entry = strings.ToLower(cachePath + search)

  f, err := os.Create(entry)
  utils.CheckError(err)

  f.WriteString(data)

  return entry
}

func CheckCache(search, cachePath string) (flag bool){
  files, err := ioutil.ReadDir(cachePath)
  utils.CheckError(err)

  flag = false

  for _, f := range files {
    if f.Name() == search{
      flag := true
      entry := cachePath + search
      utils.GenericFilePrint(entry)
      return flag
    }
  }

  return flag
}

func RunAPICache(search, cachePath string, api *utils.API){
  entry := addAPICache(search, cachePath, api)
  utils.GenericFilePrint(entry)
}

func RunTypeCache(search, dataTypeInfo, cachePath string){
  entry := addGenericCache(search, dataTypeInfo, cachePath)
  utils.GenericFilePrint(entry)
}

func RunKernelCache(search, kernelInfo, cachePath string){
  entry := addGenericCache(search, kernelInfo, cachePath)
  utils.GenericFilePrint(entry)
}
