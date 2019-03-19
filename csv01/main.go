package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
	"text/template"
)

type stat struct {
	Date     string
	Open     float64
	High     float64
	Low      float64
	Close    float64
	Volume   int64
	AdjClose float64
}

var tpl *template.Template

func main() {

	f, err := os.Open("table.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	rr := csv.NewReader(f)

	var stats []stat
	for {
		rcd, err := rr.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		var ns stat
		ns.Date = rcd[0]
		ns.Open, _ = strconv.ParseFloat(rcd[1], 64)
		ns.High, _ = strconv.ParseFloat(rcd[2], 64)
		ns.Low, _ = strconv.ParseFloat(rcd[3], 64)
		ns.Close, _ = strconv.ParseFloat(rcd[4], 64)
		ns.Volume, _ = strconv.ParseInt(rcd[5], 10, 64)
		ns.AdjClose, _ = strconv.ParseFloat(rcd[6], 64)
		stats = append(stats, ns)
	}

	tpl = template.Must(template.ParseFiles("tpl.gohtml"))

	err = tpl.Execute(os.Stdout, stats)
	if err != nil {
		log.Fatalln(err)
	}

}
