package main

import (
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
		_, err := tx.CreateBucketIfNotExists([]byte(tableName))
		return err
	})
	return database
}

func DBClose() {
	DatabaseOpen().Close()
}

func SaveIntoStationTable(data string, byteData []byte) {
	err := DatabaseOpen().Update(func(tx *boltDB.Tx) error { //insert
		bucket := tx.Bucket([]byte(tableName))                      // into tableName
		err := bucket.Put([]byte(data) /*key*/, byteData /*value*/) // "byteData" values data
		return err
	})
	HandleErr(err)
}

func SearchStationTable() []byte {
	var data []byte
	DatabaseOpen().View(func(tx *boltDB.Tx) error {
		bucket := tx.Bucket([]byte(tableName))
		data = bucket.Get([]byte("GangNamStation"))
		return nil
	})
	return data
}

func HandleErr(err error) {
	if err != nil {
		log.Panic()
	}
}
