package main

/*

1. 소스 한줄씩 분석
2. 리팩토링


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
	// //fmt.Println(db.CreateStation())
	// //db.CreateStation().Restore(db.SearchStationTable())
	// //fmt.Println(db.CreateStation())
	var asdf = BytesToString(db.SearchStationTable())
	fmt.Println(asdf)
	// //db.SaveIntoStationTable("asdf", Tobytes("asds"))
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

/*
1. 값이 i로 들어옴
2. i를 []byte(바이트배열)로 만들기 위해 bytes 패키지 사용
3. bytes 패키지를 encoding 하기 위해
3. aBuffer 변수에 bytes 패키지의 Buffer Struct를 선언 (초기화 아직안함)
4. encoder 변수를 encoding/gob 패키지의 NewEncoder()로 선언과 동시에 초기화
	NewEncoder()는 io패키지의 Write interface를 매개변수로 받고, *Encoder struct를 리턴함
	NewEncoder()의 매개변수타입이 interface이므로  &aBuffer를 매개변수로 받는게 가능하고
	encoder 변수의 타입은 *gob.Encoder struct로 정의됨
5. Encode()는 gob.Encoder struct의 리시버 펑션이고 매개변수를 interface 타입으로 받음
아 이렇게 하면 끝이 없네
*/
