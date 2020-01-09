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

	// 在原tpl的基础上，可以再用同样的方法parse 其他template.
	tpl, err = tpl.ParseFiles("templates/tpl2.gohtml", "templates/tpl3.gohtml")
	if err != nil {
		log.Fatalln(err)
	}
	// 也可以直接parse指定文件夹所有template.
	tpl, err = tpl.ParseGlob("templates/*")
	if err != nil {
		log.Fatalln(err)
	}

	// 执行的时候根据template各自的名称来执行对应tpl.
	err = tpl.ExecuteTemplate(os.Stdout, "tpl2.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}
	err = tpl.ExecuteTemplate(os.Stdout, "tpl3.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	// 如果未指定，并在很多template的情况下，会默认执行第一个.
	// 如果有数据，可以将数据传入template中的{{.}}里.
	// 在template里可以assign value 给var，{{$var_a := .}}.
	err = tpl.Execute(os.Stdout, "James Bond")
	if err != nil {
		log.Fatalln(err)
	}
}
