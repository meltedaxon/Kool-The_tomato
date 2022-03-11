package main

import(
	"fmt"
	"log"
	"net/http"
	"github.com/meltedaxon/Koot-The_tomato/routes"
)

func handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "main page")
}

func main(){
	routes.Route();

	log.Fatal(http.ListenAndServe(":8835", nil))
}
