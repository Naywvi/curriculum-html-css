package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"text/template"
)

var templatesDir = os.Getenv("TEMPLATES_DIR")

func runcv(w http.ResponseWriter, r *http.Request) {
	template.Must(template.ParseFiles(filepath.Join(templatesDir, "../templates/cv.html"))).Execute(w, "")
}
func start(w http.ResponseWriter, r *http.Request) {
	template.Must(template.ParseFiles(filepath.Join(templatesDir, "../static/home.html"))).Execute(w, "")
	/*go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("ok")
		http.Redirect(w, r, "/login?step=2", http.StatusFound)

	}()*/
}
func main() {
	http.Handle("/", http.FileServer(http.Dir("../static")))
	http.HandleFunc("/home.html", start)
	http.HandleFunc("/cv", runcv)
	fmt.Printf("Started server successfully on http://localhost:8089/\n")
	http.ListenAndServe(":8089", nil)
}
