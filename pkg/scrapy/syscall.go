package scrapy

import(
  "encoding/json"
  "net/http"
  "io/ioutil"
  "strings"
  "regexp"

  "github.com/leandrofroes/manw/pkg/utils"
  "github.com/leandrofroes/manw/pkg/cache"
)

func parseSyscallRepo(search, url string) map[string]interface{}{
  r, err := http.Get(url)
  utils.CheckError(err)

	var jsonData map[string]interface{}

  body, err := ioutil.ReadAll(r.Body)
  utils.CheckError(err)

  re, err := regexp.Compile("\"+" + search + "\":")
  utils.CheckError(err)
  match := re.FindString(strings.ToLower(string(body)))

  if(match == ""){
    utils.Warning("Unable to find this Windows Syscall ID.")
  }

  err = json.Unmarshal(body, &jsonData)
  utils.CheckError(err)

	return jsonData
}

func RunSyscallScraper(search, arch, cachePath string){
  var url string

  search = strings.ToLower(search)

  if(arch == "x64" || arch == "amd64" || arch == "x86_64" ){
    url = "https://raw.githubusercontent.com/j00ru/windows-syscalls/master/x64/json/nt-per-system.json"
    arch = "_x64"
  } else if(arch == "x86" || arch == "i386" || arch == "80386"){
    url = "https://raw.githubusercontent.com/j00ru/windows-syscalls/master/x86/json/nt-per-system.json"
    arch = "_x86"
  } else {
    utils.Warning("Missing architecture (-a) value.")
  }

  if(cachePath != ""){
    if(!cache.CheckSyscallCache(search, arch, cachePath)){
      jsonData := parseSyscallRepo(search, url)
      cache.RunSyscallCache(&jsonData, search, arch, cachePath)
    }
  } else {
      jsonData := parseSyscallRepo(search, url)
      utils.PrintSyscallJson(&jsonData, search)
  }
}