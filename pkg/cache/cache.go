package cache

import (
  "os"
  "strings"
  "io/ioutil"

  "manw/pkg/utils"
)

func addFunctionCache(search, cachePath string, api *utils.API) (entry string){
  entry = strings.ToLower(cachePath + search)

  f, err := os.Create(entry)
  utils.CheckError(err)

  f.WriteString(api.Title + "\n\n")
  f.WriteString("Exported by: " + api.DLL + "\n\n")
  f.WriteString(api.Description + "\n\n")
  f.WriteString(api.CodeA + "\n")

  if api.Return != ""{
    f.WriteString("Return value: " + api.Return + "\n\n")
  }

  if api.CodeB != ""{
    f.WriteString("Example code:\n\n" + api.CodeB + "\n\n")
  }

  f.WriteString("Source: " + api.Source + "\n\n")

  return entry
}

func addStructureCache(search, cachePath string, api *utils.API) (entry string){
  entry = strings.ToLower(cachePath + search)

  f, err := os.Create(entry)
  utils.CheckError(err)

  f.WriteString(api.Title + "\n\n")
  f.WriteString(api.Description + "\n\n")
  f.WriteString(api.CodeA + "\n")

  if api.CodeB != ""{
    f.WriteString(api.CodeB + "\n\n")
  }

  if api.CodeC != ""{
    f.WriteString(api.CodeC + "\n")
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
    if f.Name() == strings.ToLower(search){
      flag := true
      entry := cachePath + search
      utils.GenericFilePrint(entry)
      return flag
    }
  }

  return flag
}

func RunFunctionCache(search, cachePath string, api *utils.API){
  entry := addFunctionCache(search, cachePath, api)
  utils.GenericFilePrint(entry)
}

func RunStructureCache(search, cachePath string, api *utils.API){
  entry := addStructureCache(search, cachePath, api)
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
