package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	_tpl, _err := template.ParseFiles("templates/tpl1.gohtml")
	tpl = template.Must(_tpl, _err)
}

func main() {
	err := tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
