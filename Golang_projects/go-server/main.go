package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8") // to set the content type of body
	fmt.Fprint(w, "<h2>Welcome to my website testing!</h2>")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err : %v", err)
		return
	}
	fmt.Fprintf(w, "POST request successful\n")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name: %s\n", name)
	fmt.Fprintf(w, "Address: %s\n", address)

}

func contactHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "<h1>Contact Page</h1>  <p>To get in touch email me at <a href = \"mailto:yakshgandhi1@gmail.com\">yakshgandhi1@gmail.com</a></p>") // contact page

}

func pathHandler(w http.ResponseWriter, r *http.Request) {
	// if r.URL.Path == "/contact" {
	// 	contactHandler(w, r)
	// } else if r.URL.Path == "/form" {
	// 	formHandler(w, r)
	// } else if r.URL.Path == "/" {
	// 	homeHandler(w, r)
	// }
	switch r.URL.Path {
	case "/":
		homeHandler(w, r)
	case "/contact":
		contactHandler(w, r)
	case "/form":
		formHandler(w, r)
	default:
		// TODO: here we will handle the page not found error
		http.Error(w, "404 page not found", http.StatusNotFound)
		// w.WriteHeader(http.StatusNotFound)
		// fmt.Fprint(w, "404 page not found")
	}

}

func main() {
	http.HandleFunc("/", pathHandler)
	fmt.Println("Starting server at port :8080")
	//setting the localhost
	// if err := http.ListenAndServe(":8080", nil); err != nil {
	// 	log.Fatal(err)
	// }
	http.ListenAndServe(":8080", nil)
}
