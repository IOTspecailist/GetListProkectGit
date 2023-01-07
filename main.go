package main

/*
20230107
역말고 mma로 재설계
1. 데이터베이스1개, 이름은 MMACompanyDB
2. 테이블 1개, 이름은 MMACompanyTable
3. 테이블은 키, 값 구조이며 키는 단체명, 값은 선수명


*/
import (
	db "GetListProject/db" //database.go의 package를 db로 해야(맨윗줄) 패키지로 인식해서 임포트 가능
	"bytes"
	"encoding/gob"
	"reflect"
	"unsafe"
)

const (
	Team1 = "TeamUFC"
	Team2 = "TeamRoadFC"
)

var TeamPlayer []string

func main() {

	defer db.DBClose()
	TeamPlayer[0] = "enganu"
	TeamPlayer[1] = "colvi"
	TeamPlayer[2] = "JungChanSung"
	db.CreateTeam(Team1, TeamPlayer)

	TeamPlayer[0] = "MHMan"
	TeamPlayer[1] = "HISu"
	TeamPlayer[2] = "YYJun"
	db.CreateTeam(Team2, TeamPlayer)

	//db.CreateTeam().Restore(db.SearchStationTable())
	//fmt.Println(db.CreateTeam())
	//var asdf = BytesToString(db.SearchStationTable())
	//fmt.Println(asdf)
	//db.Insert_Into_MMACompanyTable("asdf", Tobytes("asds"))
}

func BytesToString(b []byte) string {
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sh := reflect.StringHeader{bh.Data, bh.Len}
	return *(*string)(unsafe.Pointer(&sh))
}

func Tobytes(i interface{}) []byte { // 어떤 값이든 받을 수 있도록 매개변수 interface로 선언
	var aBuffer bytes.Buffer            // 바이트배열로 만들기위한 틀(aBuffer)
	encoder := gob.NewEncoder(&aBuffer) // 틀을 가지고 재료를 굽는 오븐(encoder) 구입 (틀이 들어가 있는 오븐임)
	db.HandleErr(encoder.Encode(i))     // 오븐(encoder)안에 틀(aBuffer)이 있는상황에서 재료(i)를 투입하여 오븐(encoder)의 굽기버튼 클릭(Encode())
	return aBuffer.Bytes()              // 틀을 꺼내서 식탁으로 틀(aBuffer)에있는 빵(Bytes())을 리턴 // Bytes()는 bytes 패키지의 Buffer struct의 리시버 펑션이다
}
