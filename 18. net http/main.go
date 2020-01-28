package main

import(
	"fmt"
	"log"
	"net/http"
	//"io/ioutil"
	"encoding/json"
)

type Article struct {
	Title string `json:"title"`
	Desc string `json:"desc"`
}

type Articles []Article

var articles = Articles{
	Article{Title: "Judul pertama", Desc: "deskripsi pertama"},
	Article{Title: "Judul kedua", Desc: "deskripsi kedua"},
}

func main() {
	http.HandleFunc("/", getHome)
	http.HandleFunc("/articles", getArticles)
	http.HandleFunc("/post-article", withLogging(postArticle))
	http.ListenAndServe(":3000", nil)
}

func getHome(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Test home response"))
}

func getArticles(res http.ResponseWriter, req *http.Request) {
	json.NewEncoder(res).Encode(articles)
}

func postArticle(res http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		// body, err := ioutil.ReadAll(req.Body)
		// if err != nil {
		// 	http.Error(res, "Can't read body", http.StatusInternalServerError)
		// }

		// res.Write([]byte(string(body)))

		var newArticle Article
		err := json.NewDecoder(req.Body).Decode(&newArticle)

		if err != nil {
			fmt.Printf("Ada error", err)
		}
		
		articles = append(articles, newArticle)

		json.NewEncoder(res).Encode(articles)
	} else {
		http.Error(res, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func withLogging(next http.HandlerFunc) http.HandlerFunc {
	return func (res http.ResponseWriter, req *http.Request) {
		//fungsi dari middleware nya
		log.Printf("Logging koneksi dari", req.RemoteAddr)
		next.ServeHTTP(res, req)
	}
}