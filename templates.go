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
	// 用Execute的os方法写出tpl.
	err := tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}
	// 创建以index.html为文件名的文件nf.
	nf, err := os.Create("index.html")
	if err != nil {
		log.Fatalln(err)
	}
	// 用Execute的方法将tpl内容写进nf.
	err = tpl.Execute(nf, nil)
	if err != nil {
		log.Fatalln(err)
	}
	defer nf.Close()

}
