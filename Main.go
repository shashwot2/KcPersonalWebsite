package main

import (
  //"fmt"
  "html/template"
  "net/http"
  "log"
)
var tpl *template.Template
func init() {
  tpl = template.Must(template.ParseGlob("public/templates/*.gohtml"))
}
func main() {
  http.HandleFunc("/", indexHandler)
  log.Fatal(http.ListenAndServe(":8080", nil))
  
}

func indexHandler(w http.ResponseWriter, req *http.Request) {
  err := tpl.ExecuteTemplate(w, "index.gohtml", nil)
  if err != nil {
    log.Println(err)
    http.Error(w, "Internal server error", http.StatusInternalServerError)
    return
  }
}