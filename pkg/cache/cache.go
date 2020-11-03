package cache

import (
	"os"
	"fmt"
	"strings"
	"io/ioutil"

	"manw/pkg/utils"
	"manw/pkg/scrapy"
)

func printCache(entry string){
	data, err := ioutil.ReadFile(entry)
	utils.CheckError(err)
	fmt.Print(string(data))
}

func addCacheEntry(funcName, path string) (entry string){
	entry = strings.ToLower(path + "/" + funcName)

	url := scrapy.GoogleSearch(funcName, "api")
	api := scrapy.ParseMSDNAPI(url)

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
  
func CheckCache(funcName, path string) (flag bool){
	files, err := ioutil.ReadDir(path)
	utils.CheckError(err)

	flag = false

	for _, f := range files {
		if f.Name() == funcName{
		flag := true
		entry := path + "/" + funcName
		printCache(entry)
		return flag
		}
	}

	return flag
}

func RunCacheScraper(search, cachePath string){
	entry := addCacheEntry(search, cachePath)
	printCache(entry)
}