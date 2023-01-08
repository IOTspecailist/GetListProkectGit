package db

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"

	boltDB "github.com/boltdb/bolt"
)

const (
	DataBaseName = "MMACompanyDB"
	TableName    = "MMACompanyTable"
)

//팀 정보 기본 틀
type Team struct {
	TeamName string
	Player   [3]string
}

var database *boltDB.DB

func DatabaseOpen() *boltDB.DB {
	if database == nil {
		databasePointer, err := boltDB.Open(DataBaseName, 0600, nil)
		fmt.Println("★★DB_opened:")
		// nil을 줌으로써 디폴트옵션으로  MMACompanyDB 이름의 디비를 오픈함
		// 만약 처음이라면 파일을 생성한다
		// 쉽게말해서 가방을 하나 사서 이름을 붙여준거다 / 가방 = 데이터베이스, 가방에 넣는 물건 = 데이터
		// 볼트디비의 포인터 구조체를 리턴받는다 / 에러리턴은 왜그런건지
		database = databasePointer
		HandleErr(err)
		err = database.Update(func(tx *boltDB.Tx) error {
			HandleErr(err)
			_, err := tx.CreateBucketIfNotExists([]byte(TableName)) // StationTable이란 이름의 테이블이 없으면 생성한다
			fmt.Println("★★DB_Table_created:", TableName)
			return err
		})
		fmt.Println("★★DB_created:", DataBaseName)
	}
	return database
}

func DBClose() {
	DatabaseOpen().Close()
	fmt.Println("★★DB_closed")
	//DatabaseOpen() 함수를 통해 Close() 함수를 호출하는 이유는
	//DatabaseOpen() 함수는 볼트디비의 type DB struct 를 반환하고
	//Close()함수는 다음처럼 DB struct를 베이스로 두는 "리시버펑션"이기 때문
	//func (db *DB) Close() error {
}

//팀 만들기(팀이름 하나에 플레이어 여러명)
func CreateTeam(teamName string, player [3]string) *Team {
	team := &Team{
		TeamName: teamName,
		Player:   player,
	}
	fmt.Println("생성된 팀 :", team)
	team.Teaminfo_Into_BoltDB()
	return team
}

//insert문 호출하는 리시버펑션, 리시버펑션은 struct의 값에 변화를 준다
func (s *Team) Teaminfo_Into_BoltDB() {
	fmt.Println("★★DB_input_data:", s)
	Insert_Into_MMACompanyTable(s.TeamName, Tobytes(s))
}

func Tobytes(i interface{}) []byte {
	var aBuffer bytes.Buffer
	encoder := gob.NewEncoder(&aBuffer)
	HandleErr(encoder.Encode(i))
	return aBuffer.Bytes()
}

//Table에 데이터 추가
func Insert_Into_MMACompanyTable(data string, byteData []byte) {
	err := DatabaseOpen().Update(func(tx *boltDB.Tx) error { //insert
		bucket := tx.Bucket([]byte(TableName))                        // into tableName
		err := bucket.Put([]byte(data) /*=key*/, byteData /*=value*/) // "byteData" values data
		fmt.Println("★★DB_data inserted")
		return err
	})
	HandleErr(err)
}

//Table로부터 데이터 조회
func SearchMMACompanyTable(s string) []byte {
	var data []byte                                 //볼트디비는 인아웃이 바이트배열이므로 디비조회해서 결과를 저장할 변수를 선언함 / 리턴도 할겸
	DatabaseOpen().View(func(tx *boltDB.Tx) error { //select value
		bucket := tx.Bucket([]byte(TableName)) //from tableName
		data = bucket.Get([]byte(s))           // where key = GangNam
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

func (s *Team) Restore(data []byte) {
	FromBytes(s, data)
}
