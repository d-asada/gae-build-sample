package main

import (
	"net/http"
        "html/template"
        "github.com/go-martini/martini"
        "log"
)


func init() {
        m := martini.Classic()
        m.Get("/", index)
        http.Handle("/", m)
}    

func index(w http.ResponseWriter, r *http.Request) {
        t := template.Must(template.ParseFiles("views/index/index.html.tpl"))
        str := "Shaken, not stirred."
        if err := t.ExecuteTemplate(w, "index.html.tpl", str); err != nil {
            log.Fatal(err)
        }
}

