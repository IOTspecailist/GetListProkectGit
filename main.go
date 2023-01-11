package main

/*
20230110
1. 리팩토링진행중
2. http에 데이터 보기하면 에러
*/
import (
	db "GetListProject/db" //database.go의 package를 db로 해야(맨윗줄) 패키지로 인식해서 임포트 가능
	"fmt"
	"log"
	"net/http"
)

var t *db.Team
var t1 db.Team

// type t3 db.Team  이게되네

const (
	Team1 = "TeamUFC"
	Team2 = "TeamRoadFC"
)

func main() {

	defer db.DBClose()

	fmt.Println("==============================Start==============================")
	//////////////////////////////////////////////
	//1. 디비, 테이블, 팀 생성
	//makeTeam()
	//////////////////////////////////////////////

	//////////////////////////////////////////////
	//2. 로컬호스트에 데이터 출력
	handler := http.NewServeMux()
	handler.HandleFunc("/", ViewAtHome)
	fmt.Printf("Listening on http://localhost:%d\n", 4000)            // %s\n에서 :%d 로 변경 (prvSrc_16)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", 4000), handler)) //서버 열기 // port에서 fmt.Sprintf(":%d", port)변경  (prvSrc_16)
	//////////////////////////////////////////////
	fmt.Println("==============================END==============================")

}

func ViewAtHome(rw http.ResponseWriter, r *http.Request) {
	viewTeam()
	//var tarr []t3
	fmt.Fprint(rw, t)
}

func makeTeam() {

	Team1Ranker := []string{"enganu", "colvi", "volkanofseki"}
	Team2Ranker := []string{"MHMan", "HISu", "YYJun"}
	t.CreateTeam(Team1, Team1Ranker)
	t.CreateTeam(Team2, Team2Ranker)
}

func viewTeam() {

	t1.Restore(db.SearchMMACompanyTable(Team1))

}
