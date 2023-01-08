package utils

import (
	"bytes"
	"encoding/gob"
	"log"
	"reflect"
	"unsafe"
)

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

func BytesToString(b []byte) string {
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sh := reflect.StringHeader{bh.Data, bh.Len}
	return *(*string)(unsafe.Pointer(&sh))
}

func Tobytes(i interface{}) []byte { // 어떤 값이든 받을 수 있도록 매개변수 interface로 선언
	var aBuffer bytes.Buffer            // 바이트배열로 만들기위한 틀(aBuffer)
	encoder := gob.NewEncoder(&aBuffer) // 틀을 가지고 재료를 굽는 오븐(encoder) 구입 (틀이 들어가 있는 오븐임)
	HandleErr(encoder.Encode(i))        // 오븐(encoder)안에 틀(aBuffer)이 있는상황에서 재료(i)를 투입하여 오븐(encoder)의 굽기버튼 클릭(Encode())
	return aBuffer.Bytes()              // 틀을 꺼내서 식탁으로 틀(aBuffer)에있는 빵(Bytes())을 리턴 // Bytes()는 bytes 패키지의 Buffer struct의 리시버 펑션이다
}
