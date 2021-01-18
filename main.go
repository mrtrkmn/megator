package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/mrtrkmn/megator/client"
)

const (
	LIST_OF_MEGA_URLS     = "./resource/mega-link-list.csv"
	OLD_LIST_OF_MEGA_URLS = "./resource/old-mega-link-list.csv"
)

func main() {
	cli := client.New()

	var oldFile, err = os.OpenFile(OLD_LIST_OF_MEGA_URLS, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		log.Fatal(err)
	}
	defer oldFile.Close()
	csvfile, err := os.Open(LIST_OF_MEGA_URLS)
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	r := csv.NewReader(csvfile)
	writer := csv.NewWriter(oldFile)
	defer writer.Flush()
	for {
		// Read each record from csv
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		if err := cli.MegaCLI.DownloadWithDirName("data", record[1]); err != nil {
			panic(err)
		}

		if err := cli.Tar.CompressWithPIGZ(fmt.Sprintf("%s.tar.gz", record[0]), "data"); err != nil {
			panic(err)
		}
		writer.Write([]string{record[0], record[1]})
	}

	if err := os.Remove("./resource/mega-link-list.csv"); err != nil {
		panic(err)
	}

	createFile("./resource/mega-link-list.csv")

}

func createFile(path string) {
	// detect if file exists
	var _, err = os.Stat(path)

	// create file if not exists
	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if err != nil {
			return
		}
		defer file.Close()
	}
}
