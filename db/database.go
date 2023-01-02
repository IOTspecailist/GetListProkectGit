package db

import (
	"bytes"
	"encoding/gob"
	"log"

	boltDB "github.com/boltdb/bolt"
)

const (
	tableName = "StationTable"
)

var database *boltDB.DB

func DatabaseOpen() *boltDB.DB {
	databasePointer, err := boltDB.Open("StationDB", 0600, nil)
	database = databasePointer
	HandleErr(err)
	err = database.Update(func(tx *boltDB.Tx) error {
		HandleErr(err)
		_, err := tx.CreateBucketIfNotExists([]byte(tableName))
		return err
	})
	return database
}

func DBClose() {
	DatabaseOpen().Close()
}

type Station struct {
	StationName string
	Dosi        string
}

func CreateStation() *Station {
	station := &Station{
		StationName: "GangNam",
		Dosi:        "Seoul",
	}
	station.AddStation()
	return station

}

func (s *Station) Restore(data []byte) {
	FromBytes(s, data)
}
func (s *Station) AddStation() {
	SaveIntoStationTable(s.StationName, Tobytes(s))
}

func Tobytes(i interface{}) []byte {
	var aBuffer bytes.Buffer
	encoder := gob.NewEncoder(&aBuffer)
	HandleErr(encoder.Encode(i))
	return aBuffer.Bytes()
}

func SaveIntoStationTable(data string, byteData []byte) {
	err := DatabaseOpen().Update(func(tx *boltDB.Tx) error { //insert
		bucket := tx.Bucket([]byte(tableName))                        // into tableName
		err := bucket.Put([]byte(data) /*=key*/, byteData /*=value*/) // "byteData" values data
		return err
	})
	HandleErr(err)
}

func SearchStationTable() []byte {
	var data []byte
	DatabaseOpen().View(func(tx *boltDB.Tx) error {
		bucket := tx.Bucket([]byte(tableName)) //from tableName
		data = bucket.Get([]byte("GangNam"))   // where key = GangNam
		return nil
	})
	return data
}

func HandleErr(err error) {
	if err != nil {
		log.Panic()
	}
}

func FromBytes(i interface{}, data []byte) {
	//들어온 바이트를 읽어서~
	encoder := gob.NewDecoder(bytes.NewReader(data))
	HandleErr(encoder.Decode(i))
	//포인터로 복원해주는 거다? i로 온게 포인터라 그런가 그래서 리턴이 없나
}
