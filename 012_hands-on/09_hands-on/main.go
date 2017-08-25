package main

import (
	"bufio"
	"encoding/csv"
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type tplData struct {
	Heading []string
	Info    [][]string
}

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	dataFile, err := os.Open("table.csv")

	if err != nil {
		log.Fatalln(err)
	}

	r := csv.NewReader(bufio.NewReader(dataFile))
	rawData, err := r.ReadAll()

	if err != nil {
		log.Fatalln(err)
	}

	data := tplData{}
	data.Heading = rawData[0]

	for i, value := range rawData {
		if i > 0 {
			data.Info = append(data.Info, value)
		}
	}

	err = tpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalln(err)
	}
}
