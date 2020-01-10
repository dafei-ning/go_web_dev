package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type Language struct {
	Name               string
	Efficiency         string
	Developer_friendly bool
}

type OS struct {
	Name       string
	Popularity string
}

type LanOS struct {
	Languages []Language
	OSs       []OS
}

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

	/*
	 * 关于 passing data
	 * Slice, Map, Struct
	 */

	// 如果未指定，并在很多template的情况下，会默认执行第一个.
	// 如果有数据，可以将数据传入template中的{{.}}里.
	// 在template里可以assign value 给var，{{$var_a := .}}.
	err = tpl.Execute(os.Stdout, "James Bond")
	if err != nil {
		log.Fatalln(err)
	}

	// slice -
	// <li>{{.}}</li> ==> {{range .}}, {{range $index, $element := .}}
	tplSlice := template.Must(template.ParseFiles("templates/tplSlice.gohtml"))
	_slice := []string{"hello", "world", "is", "shit"}
	err = tplSlice.Execute(os.Stdout, _slice)
	if err != nil {
		log.Fatalln(err)
	}

	tplSlice2 := template.Must(template.ParseFiles("templates/tplSlice2.gohtml"))
	_slice2 := []string{"hello", "world", "is", "shit"}
	err = tplSlice2.Execute(os.Stdout, _slice2)
	if err != nil {
		log.Fatalln(err)
	}

	// map -
	// <li>{{.}}</li> ==> {{range .}},
	// <li>{{$key}} -  {{$val}}</li> => {{range $key, $val := .}}
	tplMap := template.Must(template.ParseFiles("templates/tplMap.gohtml"))
	_map := map[string]string{
		"C++":  "No1",
		"Go":   "No2",
		"Java": "No3",
		"Dart": "No4",
	}
	err = tplMap.Execute(os.Stdout, _map) // all values.
	if err != nil {
		log.Fatalln(err)
	}

	tplMap2 := template.Must(template.ParseFiles("templates/tplMap2.gohtml"))
	_map2 := map[string]string{
		"C++":  "No1",
		"Go":   "No2",
		"Java": "No3",
		"Dart": "No4",
	}
	err = tplMap2.Execute(os.Stdout, _map2) // all key - value pair.
	if err != nil {
		log.Fatalln(err)
	}

	// struct -
	// <li>{{.Name}} - {{.Efficiency}} - {{.Developer_friendly}}</li>
	//
	// {{$x := .Name}}
	// {{$y := .Efficiency}}
	// {{$z := .Developer_friendly}}
	// <li>{{$x}} -  {{$y}} - {{$z}}</li>
	tplStruct := template.Must(template.ParseFiles("templates/tplStruct.gohtml"))
	_struct1 := Language{
		Name:               "C++",
		Efficiency:         "high",
		Developer_friendly: false,
	}
	err = tplStruct.Execute(os.Stdout, _struct1)
	if err != nil {
		log.Fatalln(err)
	}

	tplStruct2 := template.Must(template.ParseFiles("templates/tplStruct2.gohtml"))
	_struct2 := Language{
		Name:               "Java",
		Efficiency:         "high",
		Developer_friendly: true,
	}
	err = tplStruct2.Execute(os.Stdout, _struct2)
	if err != nil {
		log.Fatalln(err)
	}

	// struct slice -
	// <li>{{.Name}}-{{.Efficiency}}-{{.Developer_friendly}}</li> ==> {{range .}}
	_struct3 := Language{
		Name:               "Go",
		Efficiency:         "high",
		Developer_friendly: true,
	}

	_struct4 := Language{
		Name:               "Python",
		Efficiency:         "low",
		Developer_friendly: true,
	}

	tplStructSlice := template.Must(template.ParseFiles("templates/tplStructSlice.gohtml"))
	_structSlice := []Language{_struct1, _struct2, _struct3, _struct4}
	err = tplStructSlice.Execute(os.Stdout, _structSlice)
	if err != nil {
		log.Fatalln(err)
	}

	// Struct Slice Struct -
	_struct5 := OS{
		Name:       "Windows",
		Popularity: "Common users",
	}

	_struct6 := OS{
		Name:       "Linux",
		Popularity: "Developer users",
	}

	_languages := []Language{_struct1, _struct2, _struct3, _struct4}
	_OSs := []OS{_struct5, _struct6}
	_lanOS := LanOS{
		Languages: _languages,
		OSs:       _OSs,
	}

	tplStructSliceStruct := template.Must(template.ParseFiles("templates/tplStructSliceStruct.gohtml"))
	err = tplStructSliceStruct.Execute(os.Stdout, _lanOS)
	if err != nil {
		log.Fatalln(err)
	}
}
