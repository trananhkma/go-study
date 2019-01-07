package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/crawler", func(response http.ResponseWriter, request *http.Request) {
		//url := "https://thanhnien.vn/thoi-su/hinh-anh-nguoi-dan-ong-dung-xe-de-di-ve-sinh-giua-duong-cao-toc-gay-phan-no-1037558.html"
		url, _ := request.URL.Query()["url"]
		resp, e := http.Get(url[0]);

		if e != nil{
			panic(e);
		}

		defer resp.Body.Close()

		doc, err := goquery.NewDocumentFromReader(resp.Body)

		if err != nil {
			panic(err)
		}


		title := doc.Find(".details__headline").Text()
		fmt.Fprintln(response,"<html><body><h3>Title: ", title, "</h3><p></p>")

		des := doc.Find(".sapo").Text()
		fmt.Fprintln(response,"<h3>Desciption: ", des, "</h3><p></p>")

		body, _ := doc.Find(".cms-body").Html()
		fmt.Fprintln(response,"<h3>Body: </h3><p></p>")
		fmt.Fprintln(response, body)

		time := doc.Find(".details__meta .meta time").Text()
		t := strings.Split(time, " - ")
		time = t[1] + " " + t[0]
		fmt.Fprintln(response,"<h3>Time: ", time, "</h3><p></p>")

		thumb, _ := doc.Find(".storyavatar").Attr("src")
		fmt.Fprintln(response,"<h3>Thumbnail: <a href='", thumb, "'>Link</a></h3><p></p>")

		fmt.Fprintln(response,"</body></html>")
	})

	http.ListenAndServe("localhost:6969",nil);
}
