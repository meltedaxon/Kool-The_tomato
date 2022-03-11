package handler

import(
	"fmt"
	"net/http"
)

func MainHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "main page")
}

func PostHandler(w http.ResponseWriter, r *http.Request){
	switch r.Method{
	case "GET":
		if r.Form.postNum != null{
			fmt.Fprintf(w, "post detail page")
		}else{
			fmt.Fprintf(w, "post list page");
		}
	default:
		http.Error(w, "405 Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
}
