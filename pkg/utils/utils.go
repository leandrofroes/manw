package utils

import (
  "log"
  "fmt"
  "io/ioutil"
)

type API struct {
  Title       string
  Description string
  Code        string
  Return      string
  ExampleCode string
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
