package main

/*
1. hard로 데이터 insert
2. 함수 호출로 콘솔에 전체 출력

*/
import (
	db "GetListProject/db" //database.go의 package를 db로 해야(맨윗줄) 패키지로 인식해서 임포트 가능
	"bytes"
	"encoding/gob"
	"fmt"
	"reflect"
	"unsafe"
)

func main() {

	defer db.DBClose()
	//fmt.Println(db.CreateStation())
	//db.CreateStation().Restore(db.SearchStationTable())
	//fmt.Println(db.CreateStation())
	var asdf = BytesToString(db.SearchStationTable())
	fmt.Println(asdf)
	//db.SaveIntoStationTable("asdf", Tobytes("asds"))
}
func BytesToString(b []byte) string {
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sh := reflect.StringHeader{bh.Data, bh.Len}
	return *(*string)(unsafe.Pointer(&sh))
}
func Tobytes(i interface{}) []byte {

	var aBuffer bytes.Buffer
	encoder := gob.NewEncoder(&aBuffer)
	db.HandleErr(encoder.Encode(i))
	return aBuffer.Bytes()
}
