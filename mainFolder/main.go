package main

import (
	"GetListProject/db"
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
)

func main() {
	fmt.Println("asdf")

	db.SaveIntoStationTable("asdf", Tobytes("asds"))
}

func Tobytes(i interface{}) []byte {
	var aBuffer bytes.Buffer
	encoder := gob.NewEncoder(&aBuffer)
	HandleErr(encoder.Encode(i))
	return aBuffer.Bytes()
}

func HandleErr(err error) {
	if err != nil {
		log.Panic()
	}
}
