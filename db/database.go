package db

import (
	"GetListProject/utils"
	"fmt"

	boltDB "github.com/boltdb/bolt"
)

const (
	DataBaseName = "MMACompanyDB"
	TableName    = "MMACompanyTable"
)

//팀 정보 기본 틀
type Team struct {
	TeamName string
	Player   []string
}

var database *boltDB.DB

func DBOpen() *boltDB.DB {
	if database == nil { // go run main.go 할때마다 계속 true이긴 함
		databasePointer, err := boltDB.Open(DataBaseName, 0600, nil)
		fmt.Println("★★★★★★DB_opened★★★★★★")
		fmt.Println("★★DB_created : ", DataBaseName)
		// nil을 줌으로써 디폴트옵션으로  MMACompanyDB 이름의 디비를 오픈함
		// 만약 처음이라면 파일을 생성한다
		// 쉽게말해서 가방을 하나 사서 이름을 붙여준거다 / 가방 = 데이터베이스, 가방에 넣는 물건 = 데이터
		// 볼트디비의 포인터 구조체를 리턴받는다 / 에러리턴은 왜그런건지
		database = databasePointer
		utils.HandleErr(err)
		err = database.Update(func(tx *boltDB.Tx) error {
			utils.HandleErr(err)
			_, err := tx.CreateBucketIfNotExists([]byte(TableName)) // StationTable이란 이름의 테이블이 없으면 생성한다
			fmt.Println("★★DB_Table_created:", TableName)
			return err
		})
	}
	return database
}

func DBClose() {
	DBOpen().Close()
	fmt.Println("★★★★★★DB_closed★★★★★★")
	//DatabaseOpen() 함수를 통해 Close() 함수를 호출하는 이유는
	//DatabaseOpen() 함수는 볼트디비의 type DB struct 를 반환하고
	//Close()함수는 다음처럼 DB struct를 베이스로 두는 "리시버펑션"이기 때문
	//func (db *DB) Close() error {
}

//insert문 호출하는 리시버펑션, 리시버펑션은 struct의 값에 변화를 준다
func (t *Team) CreateTeam(teamName string, teamRanker []string) {
	team := &Team{
		TeamName: teamName,
		Player:   teamRanker,
	}
	Insert_Into_MMACompanyTable(team.TeamName, utils.Tobytes(team.Player))
	fmt.Println("★★★Created Team Name : ", team.TeamName)
	fmt.Println("★★★Created Team Player : ", team.Player)
}

//Table에 데이터 추가
func Insert_Into_MMACompanyTable(data string, byteData []byte) {
	err := DBOpen().Update(func(tx *boltDB.Tx) error { //insert
		bucket := tx.Bucket([]byte(TableName))                        // into tableName
		err := bucket.Put([]byte(data) /*=key*/, byteData /*=value*/) // "byteData" values data
		fmt.Println("★★DB_data inserted")
		return err
	})
	utils.HandleErr(err)
}

//Table로부터 데이터 조회
func SearchMMACompanyTable(s string) []byte {
	fmt.Println("★★DB_searching DataBase..............................")
	var data []byte                           //볼트디비는 인아웃이 바이트배열이므로 디비조회해서 결과를 저장할 변수를 선언함 / 리턴도 할겸
	DBOpen().View(func(tx *boltDB.Tx) error { //select value
		fmt.Println("★★1")
		bucket := tx.Bucket([]byte(TableName)) //from tableName
		fmt.Println("★★2")
		data = bucket.Get([]byte(s)) // where key = GangNam
		fmt.Println("★★3")
		return nil
	})
	fmt.Println("★★4")
	return data
}

func (s *Team) Restore(data []byte) {
	fmt.Println("★★5")
	utils.FromBytes(s, data)
	fmt.Println("★★10")
}
