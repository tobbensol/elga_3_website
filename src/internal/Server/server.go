package Server

import (
	"fmt"
	"github.com/tobbensol/elga_3_website/internal/Handlers"
	"log"
	"net/http"
)

func init() {
	http.HandleFunc("/", Handlers.GetMainPage)
	http.HandleFunc("/postReview/", Handlers.PostReview)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("UI/static"))))
}

func InitServer() {
	fmt.Println("Server hosted at: http://localhost:8000/")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
