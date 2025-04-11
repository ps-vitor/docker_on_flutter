package	main	

import(
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/PuerkitoBio/goquery"
	"crypto/tls"
	"time"
	"strings"
)

type	ScrapeResult	struct{
	URL	string	'json:"url"'
	Text	string	'json:"text"'
	Err	string	'json:"err,omitempty"'
}

func	scrapePage(url	string,ch	chan<-ScrapeResult){
	customClient:=&http.Client{
		Timeout:time.Second*10,
		Transport:&http.Transport{
			TLSClientConfig:&tls.Config{InsecureSkipVerify:true},
		},
	}

	resp,err:=customClient.Get(url)
	if	err!=nil{
		ch<-ScrapeResult{URL:url,Text:"",Err:fmt.Sprintf("Request error: %v",err)}
		return	
	}
	defer	resp.Body.Close()

	doc,err:=goquery.NewDocumentFromReader(resp.Body)
	if	err!=nil{
		ch<-ScrapeResult{URL:url,Text="",Err:fmt.Sprintf("Parse error: %v",err)}
		return
	}

	selection:=doc.find("div.content.clearfix")
	text:=selection.Text()
	if	text==""{
		text="Not found."
	}

	ch<-ScrapeResult(URL:url,Text:text}
}

func	scrapeHandler(w	http.ResponseWriter,	r	*http.Request){
	urls:=r.URL.Query()["url"]
	if	len(urls)==0{
		http.Error(w,"No URLs provided",http.StatusBadRequest)
		return
	}
	results:=make([]ScrapeResult,0)
	ch:=make(chan	ScrapeResult)

	for	_,url:=range	urls{
		go	scrapePage(url,ch)
	}

	for	range	urls{
		go	scrapePage(url,ch)
	}

	for	range	urls{
		results=append(results,<-ch)
	}

	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(results)
}


func	main(){
	http.HandleFunc("/scrape",	scrapeHandler)
	fmt.Println("Server running on http://0.0.0.0:8080/scrape?url=")
	log.Fatal(http.ListenAndServe(":8080",nil))
}
