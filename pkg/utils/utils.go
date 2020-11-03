package utils

import (
  "log"
)

func CheckError(err error){
  if(err != nil){
    log.Fatal(err)
  }
}

func Warning(s string){
  log.Fatal("[!] " + s)
}