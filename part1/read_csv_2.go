//Handling unintended data
package main

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

type CsvDataRecord struct {
	SpepalLen  float64
	SpeWidth   float64
	Score      float64
	Point      float64
	Name       string
	ParseError error //parse data error
}

func main() {
	//read file
	f, err := os.Open("data2.csv")
	if err != nil {
		log.Println("read data error")
		return
	}

	defer f.Close()

	var data []CsvDataRecord

	//create a csv reader
	reader := csv.NewReader(f)
	var i int
	for {
		i = i + 1

		row, err := reader.Read()
		if err == io.EOF {
			break
		}

		// log.Println("current row: ", i)
		//read data one by one
		var csvRecord CsvDataRecord

		for index, val := range row {
			var num float64

			if index == 3 {
				if num, err = strconv.ParseFloat(val, 64); err != nil {
					str := fmt.Sprintf("%d row,val is not number,unexpected type in field %d\n", i, index)
					log.Print(str)
					csvRecord.ParseError = errors.New(str)
					break
				}
			}

			//last one
			//When the data in column 4 is empty, the line is directly ignored.
			if index == 4 {
				if val == "" {
					str := fmt.Sprintf("%d row,string is empty,unexpected type in field %d\n", i, index)
					log.Print(str)
					csvRecord.ParseError = errors.New(str)
					break
				}

				csvRecord.Name = val
				continue
			}

			//others list
			if num, err = strconv.ParseFloat(val, 64); err != nil {
				str := fmt.Sprintf("%d row,val is not number,unexpected type in field %d\n", i, index)
				log.Print(str)
				csvRecord.ParseError = errors.New(str)

				break
			}

			switch index {
			case 0:
				csvRecord.SpepalLen = num
			case 1:
				csvRecord.SpeWidth = num
			case 2:
				csvRecord.Score = num
			case 3:
				csvRecord.Point = num
			}

		}

		//append data to data
		if csvRecord.ParseError == nil {
			data = append(data, csvRecord)
		}
	}

	log.Println("right data len: ", len(data))
	log.Println("read data: ", data)

	bytes, _ := json.Marshal(data)
	log.Println(string(bytes))

}

/**
2019/01/11 22:35:00 2 row,string is empty,unexpected type in field 4
2019/01/11 22:35:00 4 row,val is not number,unexpected type in field 3
2019/01/11 22:35:00 5 row,val is not number,unexpected type in field 2
2019/01/11 22:35:00 right data len:  3
2019/01/11 22:35:00 read data:  [{4.3 3.2 1.1 0.18 iris-heige1222 <nil>} {4.3 2.5 2.1 0.5 33 <nil>} {4.3 31.2 7.1 1.1 iris-heigeo323 <nil>}]
2019/01/11 22:35:00 [{"SpepalLen":4.3,"SpeWidth":3.2,"Score":1.1,"Point":0.18,"Name":"iris-heige1222","ParseError":null},{"SpepalLen":4.3,"SpeWidth":2.5,"Score":2.1,"Point":0.5,"Name":"33","ParseError":null},{"SpepalLen":4.3,"SpeWidth":31.2,"Score":7.1,"Point":1.1,"Name":"iris-heigeo323","ParseError":null}]
*/
