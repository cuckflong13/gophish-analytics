package runner

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/cuckflong/gophish-analytics/pkg/logging"
)

func (r *Runner) parseFile() {
	csvFile, err := os.Open(r.inputFile)
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	f := csv.NewReader(csvFile)

	// Skipping the first two lines of the file

	// lineSkip := 2
	// for i := 0; i < lineSkip; i++ {
	// 	f.Read()
	// }

	for {
		record, err := f.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		if isValid(record[0]) {
			r.addRecord(record[0], record[1], record[2], record[3])
			logging.Log(logging.DEBUG, fmt.Sprintf("%s|%s|%s|%s", record[0], record[1], record[2], record[3]))
		}
	}
}
