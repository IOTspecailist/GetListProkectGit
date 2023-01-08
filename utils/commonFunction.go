package utils

import (
	"bytes"
	"encoding/gob"
	"log"
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

func Tobytes(i interface{}) []byte {
	var aBuffer bytes.Buffer
	encoder := gob.NewEncoder(&aBuffer)
	HandleErr(encoder.Encode(i))
	return aBuffer.Bytes()
}
