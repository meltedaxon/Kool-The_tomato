package router

import(
	"net/http"
	"github.com/meltedaxon/Koot-The_tomato/handler/handler"
)

func Route(){
	//main
	http.HandleFunc("/", handler.MainHandler())

	http.HandleFunc("/posts", handler.PostHandler())
}
