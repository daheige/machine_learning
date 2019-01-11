package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
)

func main() {
	//read file
	f, err := os.Open("data.csv")
	if err != nil {
		log.Println("read data error")
		return
	}

	defer f.Close()

	//create a csv reader
	reader := csv.NewReader(f)
	reader.FieldsPerRecord = 5 //point fields number

	var rawData [][]string

	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Println("read current row error: ", err)
			continue
		}

		//append data to rawData
		rawData = append(rawData, row)
	}

	log.Println("all data: ")
	log.Println(rawData)
	log.Println(rawData[0])
}

/**
 * exec result:
 * 2019/01/11 21:44:19 all data:
2019/01/11 21:44:19 [[4.3 3.2 1.1 0.18 iris-heige1] [4.3 3.1 0.1 0.6 iris-heige2] [4.3 2.2 2.1 0.5 iris-heige3] [4.3 1.2 5.1 0.3 iris-heige0] [4.3 33.2 6.1 0.1 iris-heigei] [4.3 31.2 7.1 1.1 iris-heigeo]]
2019/01/11 21:44:19 [4.3 3.2 1.1 0.18 iris-heige1]
*/
