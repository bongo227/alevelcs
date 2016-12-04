package main

import (
	"bytes"
	"net/http"
	"text/template"
)

type Root struct {
	Pages []Page
}

type StringTree struct {
	Value    string
	Children []StringTree
}

type Page struct {
	Id           string
	Title        string
	Content      string
	ContentValue string
	SubSections  StringTree
}

var root = Root{
	[]Page{
		Page{
			Id:      "finiteStateMachines",
			Title:   "Finite State Machines",
			Content: "finiteStateMachines.html",
		},
		Page{
			Id:      "externalHardware",
			Title:   "External Hardware",
			Content: "externalHardware.html",
		},
	},
}

var templates = []string{
	"root.html",
	"styles.html",
	"sidebarStyle.html",
	"sidebar.html",
	"pages.html",
	"finiteStateMachines.html",
	"externalHardware.html",
}

func handler(w http.ResponseWriter, r *http.Request) {
	for i, p := range root.Pages {
		t, err := template.ParseFiles(p.Content)
		if err != nil {
			panic(err.Error())
		}

		buf := new(bytes.Buffer)
		t.Execute(buf, nil)
		root.Pages[i].ContentValue = buf.String()
	}

	t, _ := template.ParseFiles(templates...)
	t.Execute(w, root)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
