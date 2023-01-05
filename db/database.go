package db

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"

	boltDB "github.com/boltdb/bolt"
)

const (
	tableName = "StationTable"
)

var database *boltDB.DB
var bucc *boltDB.Bucket

func DatabaseOpen() *boltDB.DB {
	databasePointer, err := boltDB.Open("StationDB", 0600, nil)
	// nil을 줌으로써 디폴트옵션으로 StationDB란 이름의 디비를 오픈함
	// 만약 처음이라면 파일을 생성한다
	// 쉽게말해서 가방을 하나 사서 이름을 붙여준거다 / 가방 = 데이터베이스, 가방에 넣는 물건 = 데이터
	// 볼트디비의 포인터 구조체를 리턴받는다 / 에러리턴은 왜그런건지
	database = databasePointer
	HandleErr(err)
	err = database.Update(func(tx *boltDB.Tx) error {
		HandleErr(err)
		bucc, err := tx.CreateBucketIfNotExists([]byte(tableName))        // StationTable이란 이름의 테이블이 없으면 생성한다
		fmt.Println("DatabaseOpen() :: returned *boltDB.Bucket : ", bucc) // 언더바_로 리턴값을 버리고 있길래 변수선언해서 받음
		return err
	})
	return database
}

func DBClose() {
	DatabaseOpen().Close()
	//DatabaseOpen() 함수를 통해 Close() 함수를 호출하는 이유는
	//DatabaseOpen() 함수는 볼트디비의 type DB struct 를 반환하고
	//Close()함수는 다음처럼 DB struct를 베이스로 두는 "리시버펑션"이기 때문
	//func (db *DB) Close() error {
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

//키를 가지고 데이터를 찾아오는 함수
func SearchStationTable() []byte {
	var data []byte                                 //볼트디비는 인아웃이 바이트배열이므로 디비조회해서 결과를 저장할 변수를 선언함 / 리턴도 할겸
	DatabaseOpen().View(func(tx *boltDB.Tx) error { //select value
		bucket := tx.Bucket([]byte(tableName)) //from tableName
		data = bucket.Get([]byte("GangNam"))   // where key = GangNam
		return nil
	})
	return data
}

func HandleErr(err error) {
	if err != nil {
		log.Panic() //어떤 에러가 발생했는지 로그로 알려주고 프로그램을 종료시키는 함수
	}
}

func FromBytes(i interface{}, data []byte) {
	//들어온 바이트를 읽어서~
	encoder := gob.NewDecoder(bytes.NewReader(data))
	HandleErr(encoder.Decode(i))
	//포인터로 복원해주는 거다? i로 온게 포인터라 그런가 그래서 리턴이 없나
}
