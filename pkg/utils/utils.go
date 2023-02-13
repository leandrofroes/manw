package utils

import (
  "log"
  "fmt"
  "io/ioutil"
  "strings"
)

type API struct {
  Title       string
  Description string
  CodeA       string
  CodeB       string
  CodeC       string
  Return      string
  Source      string
  DLL         string
  Argc        string
}

func CheckError(err error){
  if err != nil{
    log.Fatal(err)
  }
}

func Warning(s string){
  log.Fatal("[!] " + s)
}

func GenericFilePrint(entry string){
  data, err := ioutil.ReadFile(entry)
  CheckError(err)
  fmt.Print(string(data))
}

func GenericPrint(data string){
  fmt.Print(string(data))
}

func PrintMSDNFunc(api *API){
  fmt.Printf(api.Title + "\n\n")

  if api.DLL != ""{
    fmt.Printf("Exported by: " + api.DLL + "\n\n")
  }
  
  fmt.Printf("Number of arguments: " + api.Argc + "\n\n")

  fmt.Printf(api.Description + "\n\n")
  fmt.Printf(api.CodeA + "\n")

  if api.Return != ""{
    fmt.Printf("Return value: " + api.Return + "\n\n")
  }

  if api.CodeB != ""{
    fmt.Printf("Example code:\n\n" + api.CodeB + "\n\n")
  }

  fmt.Printf("Source: " + api.Source + "\n\n")
}

func PrintMSDNStructure(api *API){
  fmt.Printf(api.Title + "\n\n")
  fmt.Printf(api.Description + "\n\n")
  fmt.Printf(api.CodeA + "\n")

  if api.CodeB != ""{
    fmt.Printf(api.CodeB + "\n\n")
  }

  if api.CodeC != ""{
    fmt.Printf(api.CodeC + "\n")
  }

  fmt.Printf("Source: " + api.Source + "\n\n")
}

func PrintSyscallJson(data *map[string]interface{}, search string){
  for k, v := range *data {
    if strings.HasPrefix(k, "Windows"){
      fmt.Printf("%s\n", k)
    } else if !strings.Contains(k, "Nt"){
      fmt.Printf("\t- %s: ", k)
    }
    if strings.ToLower(k) == strings.ToLower(search){
      switch v.(type){
        case float64:
          fmt.Printf("%2.f\n", v)
      }
    }
    switch v.(type) {
      case map[string]interface{}:
        tmp := v.(map[string]interface{})
        PrintSyscallJson(&tmp, search)
    }
  }
}