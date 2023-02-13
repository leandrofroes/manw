package cache

import (
  "os"
  "strings"
  "io/ioutil"
  "fmt"

  "github.com/leandrofroes/manw/pkg/utils"
)

func addFunctionCache(search, cachePath string, api *utils.API) (entry string){
  entry = strings.ToLower(cachePath + search)

  f, err := os.Create(entry)
  utils.CheckError(err)

  f.WriteString(api.Title + "\n\n")

  if api.DLL != ""{
    f.WriteString("Exported by: " + api.DLL + "\n\n")
  }
  
  f.WriteString("Number of arguments: " + api.Argc + "\n\n")

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

func parseSyscallJson(data *map[string]interface{}, search string, f *os.File){
  for k, v := range *data {
    if strings.HasPrefix(k, "Windows"){
      f.WriteString(k + "\n")
    } else if !strings.Contains(k, "Nt"){
      f.WriteString("\t- " + k + ": ")
    }
    if strings.ToLower(k) == strings.ToLower(search){
      switch v.(type){
        case float64:
          s := fmt.Sprintf("%2.f\n", v)
          f.WriteString(s)
      }
    }
    switch v.(type) {
      case map[string]interface{}:
        tmp := v.(map[string]interface{})
        parseSyscallJson(&tmp, search, f)
    }
  }
}

func addSyscallCache(data *map[string]interface{}, search, arch, cachePath string) string{
  entry := strings.ToLower(cachePath + search + arch)

  f, err := os.Create(entry)
  utils.CheckError(err)

  parseSyscallJson(data, search, f)

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
  search = strings.ToLower(search)

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

func CheckSyscallCache(search, arch, cachePath string) (flag bool){
  files, err := ioutil.ReadDir(cachePath)
  utils.CheckError(err)

  flag = false
  search = strings.ToLower(search + arch)

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

func RunSyscallCache(syscallJson *map[string]interface{}, search, arch, cachePath string){
  entry := addSyscallCache(syscallJson, search, arch, cachePath)
  utils.GenericFilePrint(entry)
}