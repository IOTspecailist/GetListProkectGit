package main

/*
20230110
1. 리팩토링진행중
2. http에 데이터 보기하면 에러
*/
import (
	db "GetListProject/db" //database.go의 package를 db로 해야(맨윗줄) 패키지로 인식해서 임포트 가능
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var t [2]*db.Team
var t1 []*db.Team

// type t3 db.Team  이게되네

const (
	Team1 = "TeamUFC"
	Team2 = "TeamRoadFC"
)

var port string

func main() {

	//defer db.DBClose()

	fmt.Println("==============================Start==============================")
	router := mux.NewRouter() // 고릴라먹스 초기화
	router.Use(jsonContentTypeMiddleware)
	router.HandleFunc("/", status).Methods("GET") //get방식으로 루트(/)에 오면 status함수를 실행한다
	port = fmt.Sprintf(":%d", 4000)
	fmt.Printf("Listening on http://localhost%s\n", port) // 7
	log.Fatal(http.ListenAndServe(port, router))          //8
	fmt.Println("===============================END===============================")

	//////////////////////////////////////////////
	//1. 디비, 테이블, 팀 생성
	//makeTeam()
	//////////////////////////////////////////////

	//////////////////////////////////////////////
	//2. 로컬호스트에 데이터 출력
	//handler := http.NewServeMux()
	//handler.HandleFunc("/", ViewAtHome)
	//fmt.Printf("Listening on http://localhost:%d\n", 4000)            // %s\n에서 :%d 로 변경 (prvSrc_16)
	//log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", 4000), handler)) //서버 열기 // port에서 fmt.Sprintf(":%d", port)변경  (prvSrc_16)
	//////////////////////////////////////////////

}

func status(rw http.ResponseWriter, r *http.Request) {
	Team1Ranker := []string{"enganu", "colvi", "volkanofseki"}
	Team2Ranker := []string{"MHMan", "HISu", "YYJun"}
	teammade1 := db.CreateTeam1(Team1, Team1Ranker)
	teammade2 := db.CreateTeam1(Team2, Team2Ranker)
	// t[0] = teammade1
	// t[1] = teammade2
	t1 := []*db.Team{teammade1, teammade2}
	json.NewEncoder(rw).Encode(t1)

}

//http핸들러 함수를 반환하는 미들웨어다. 응답변수에 헤더를 세팅해서 반환한다
func jsonContentTypeMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(rw, r)
	})
}

// func ViewAtHome(rw http.ResponseWriter, r *http.Request) {
// 	viewTeam()
// 	//var tarr []t3
// 	fmt.Fprint(rw, t)
// }

// func makeTeam() {

// 	Team1Ranker := []string{"enganu", "colvi", "volkanofseki"}
// 	Team2Ranker := []string{"MHMan", "HISu", "YYJun"}
// 	t.CreateTeam(Team1, Team1Ranker)
// 	t.CreateTeam(Team2, Team2Ranker)
// }

// func viewTeam() {

// 	t1.Restore(db.SearchMMACompanyTable(Team1))

// }
