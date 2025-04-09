package	main	

import(
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func	scrapeHandler(w	http.ResponseWriter,	r	*http.Request){
	url	:=	r.URL.Query().Get("url")
	if	url	==	""{
		http.Error(w,"URL not found",http.StatusBadRequest)
	}

	resp,err	:=	http.Get(url)
	if	err	!=	nil{
		http.Error(w,	fmt.Sprintf("Error on request: %v",	err),http.StatusInternalServerError)
		return
	}
	defer	resp.Body.Close()

	doc,err	:=	goquery.NewDocumentFromReader(resp.Body)
	if	err	!=	nil{
		http.Error(w,	fmt.Sprintf("Error at read HTML: %v",	err),http.StatusInternalServerError)
		return
	}

	selection	:=	doc.Find("div.content.clearfix")
	text	:=	selection.Text()

	if	text	==	""{
		text	=	"Not found."
	}

	w.Header().Set("Content-Type",	"application/json")
	json.NewEncoder(w).Encode(text)
}

func	main(){
	http.HandleFunc("/scrape",	scrapeHandler)

	fmt.Println("Server running on http://0.0.0.0>8080/scrape?url=https://...")
	log.Fatal(http.ListenAndServe(":8080",nil))
}
