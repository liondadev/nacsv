package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.Open("./in.json")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	b, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}

	// Parse the json of the StudentFile
	var sf StudentFile
	if err := json.Unmarshal(b, &sf); err != nil {
		panic(err)
	}

	c := sf.Records.ToCSVSlice()

	out, err := os.OpenFile("out.csv", os.O_CREATE|os.O_RDWR, 0755)
	if err != nil {
		panic(err)
	}
	defer out.Close()

	w := csv.NewWriter(out)
	for _, rec := range c {
		if err := w.Write(rec); err != nil {
			panic(err)
		}
	}

	w.Flush()

	if err := w.Error(); err != nil {
		panic(err)
	}

	fmt.Println("Done!")
}
