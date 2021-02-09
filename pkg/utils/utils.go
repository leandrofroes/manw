package utils

import (
  "log"
  "fmt"
  "io/ioutil"
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
}

func CheckError(err error){
  if(err != nil){
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
  fmt.Printf("Exported by: " + api.DLL + "\n\n")
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