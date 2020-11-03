package scrapy

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	"manw/pkg/utils"
	"github.com/gocolly/colly"
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

func GoogleSearch(s, searchType string) string{
	baseUrl := "https://www.google.com/search?q="

	var (
		url string
		result string
		regexURL string
	)

	switch{
		case searchType == "api":
			url = baseUrl + s + "+msdn"
			regexURL = "https://docs.microsoft.com/en-us/windows+"
		case searchType == "type":
			url = baseUrl + s + "+windows+data+type+msdn"
			regexURL = "https://docs.microsoft.com/en-us/windows+"
		case searchType == "kernel":
			url = baseUrl + s + "+windows+kernel+vergilius+project"
			regexURL = "https://www.vergiliusproject.com/kernels+"
	}
  
	collector := colly.NewCollector(
	  colly.UserAgent("Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:47.0) Gecko/20100101 Firefox/47.0"),
	)
  
	collector.OnHTML("html", func(e *colly.HTMLElement){
	  sellector := e.DOM.Find("div.g")
	  for node := range sellector.Nodes{
		item := sellector.Eq(node)
		linkTag := item.Find("a")
		link, _ := linkTag.Attr("href")
		link = strings.Trim(link, " ")
	
		re, err := regexp.Compile(regexURL)
		utils.CheckError(err)
	
		if link != "" && link != "#" && re.MatchString(link) {
		  result = link
		  return
		}
	  }
	})
  
	collector.OnError(func(r *colly.Response, err error) {
	  log.Fatal(err)
	})
  
	collector.Visit(url)
  
	return result
}

func ParseMSDNAPI(url string) *API{
	api := API{}
  
	collector := colly.NewCollector(
	  colly.AllowedDomains("docs.microsoft.com"),
	  colly.UserAgent("Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:47.0) Gecko/20100101 Firefox/47.0"),
	)
  
	collector.OnHTML("meta", func(e *colly.HTMLElement){
	  if e.Attr("property") == "og:title"{
		api.Title = e.Attr("content")
		return
	  }
	  if e.Attr("property") == "og:description"{
		api.Description = e.Attr("content")
		return
	  }
	  if e.Attr("property") == "og:url"{
		api.Source = e.Attr("content")
		return
	  }
	})
  
	collector.OnHTML("meta", func(e *colly.HTMLElement){
	  if e.Attr("name") == "req.dll"{
		  api.DLL = e.Attr("content")
		  return
		}
	})
  
	collector.OnHTML("pre", func(e *colly.HTMLElement){
	  if e.Index == 0 {
		api.Code = e.Text
		return
	  }
	  if e.Index == 1{
		api.ExampleCode = e.Text
		return
	  }
	})
  
	collector.OnHTML("p", func(e *colly.HTMLElement){
	  re, err := regexp.Compile(".*(no error occurs|succeeds|fails|failure|returns|return value|returned).*(no error occurs|succeeds|fails|failure|returns|return value|returned)[^.]+")
	  utils.CheckError(err)
	  match := re.FindString(e.Text)
  
	  if match != ""{
		api.Return += match + ". "
		api.Return = strings.ReplaceAll(api.Return, "\n", " ",)
	  }
	})
  
	collector.OnError(func(r *colly.Response, err error) {
	  log.Fatal(err)
	})
  
	collector.Visit(url)
  
	return &api
  }

func parseMSDNDataType(s, url string){
	collector := colly.NewCollector(
		colly.AllowedDomains("docs.microsoft.com"),
		colly.UserAgent("Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:47.0) Gecko/20100101 Firefox/47.0"),
		)

	collector.OnHTML("tr", func(e *colly.HTMLElement){
		str := strings.ToUpper(s) + "\n"
		re, err := regexp.Compile(str)
		utils.CheckError(err)
		match := re.FindString(e.Text)
		index := strings.Index(e.Text, str)

		if match != "" && index == 1{
			strSlice := strings.Split(e.Text, "\n")
			fmt.Printf("\nData Type: ")
			for i, str := range strSlice{
				if(i > 0 && i < len(strSlice) -1){
					fmt.Printf("%s\n\n", str)
				}
			}
		}
	})

	collector.OnError(func(r *colly.Response, err error) {
		log.Fatal(err)
	})

	collector.Visit(url)
}

func parseKernelInfo(url string) {
	collector := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:47.0) Gecko/20100101 Firefox/47.0"),
		)

	collector.OnHTML("div", func(e *colly.HTMLElement){
		if(e.Attr("id") == "copyblock"){
			fmt.Println(e.Text)
		}
		if(e.Attr("class") == "maincross"){
			fmt.Printf("\n%s\n\n", e.Text)
		}
	})

	collector.OnError(func(r *colly.Response, err error) {
		log.Fatal(err)
	})

	collector.Visit(url)
	
  }

func printMSDNAPINoCache(api *API){
	fmt.Printf("%s - %s\n\n", api.Title, api.DLL)
	fmt.Printf("%s\n\n", api.Description)
	fmt.Printf("%s\n\n", api.Code)

	if api.Return != ""{
		fmt.Printf("Return value: %s\n\n", api.Return)
	}

	if api.ExampleCode != ""{
		fmt.Printf("Example code:\n\n%s\n\n", api.ExampleCode)
	}

	fmt.Printf("Source: %s\n\n", api.Source)
}

func RunScraper(s, searchType string){
	url := GoogleSearch(s, searchType)

	if url == ""{
		utils.Warning("Unable to find the provided Windows information.")
	}

	switch{
		case searchType == "api":
			api := ParseMSDNAPI(url)
			printMSDNAPINoCache(api)
		case searchType == "type":
			parseMSDNDataType(s, url)
		case searchType == "kernel":
			parseKernelInfo(url)
	}
}
