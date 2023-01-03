package main

/*
1. []byte 타입으로 데이터 출력하기 --완
	바이트배열 선언 및 사용법
	바이트, 비트, 스트링 개념 구분
	변수 선언방식
	한글을 표현하는 char-set 정리했음 헷갈려서
	구글드라이브에
2. []byte 타입을 받는 함수에 데이터 전달하기 --완
3. []byte 타입 데이터 읽어서 string 출력 하는 원리(함수분석 -이건안함 고언어가 기본제공하는 함수 이용해서 만든거라)
4. 소스 한줄씩 분석
5. 리팩토링


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
