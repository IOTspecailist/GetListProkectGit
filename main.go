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
	"fmt"
	"log"
	"net/http"
)

const (
	Team1 = "TeamUFC"
	Team2 = "TeamRoadFC"
)

var team db.Team
var port int = 4000

func main() {
	fmt.Println("==============================Start==============================")

	defer db.DBClose()

	//searchTeam()
	handler := http.NewServeMux()
	handler.HandleFunc("/", home)
	fmt.Printf("Listening on http://localhost:%d\n", port)            // %s\n에서 :%d 로 변경 (prvSrc_16)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), handler)) //서버 열기 // port에서 fmt.Sprintf(":%d", port)변경  (prvSrc_16)

	fmt.Println("==============================END==============================")

}

func home(rw http.ResponseWriter, r *http.Request) {
	searchTeam()
	fmt.Fprint(rw, team)
}

func searchTeam() {
	var TeamPlayer1 [3]string
	var TeamPlayer2 [3]string

	TeamPlayer1[0] = "enganu"
	TeamPlayer1[1] = "colvi"
	TeamPlayer1[2] = "volkanofseki"
	//db.CreateTeam(Team1, TeamPlayer1)

	TeamPlayer2[0] = "MHMan"
	TeamPlayer2[1] = "HISu"
	TeamPlayer2[2] = "YYJun"
	//db.CreateTeam(Team2, TeamPlayer2)

	team.Restore(db.SearchMMACompanyTable(Team2))

	fmt.Println("+++++++++++++++++++++++Team+++++++++++++++++++++++++++++++")
	fmt.Printf("[%s]", team)
	team.Restore(db.SearchMMACompanyTable(Team1))

	fmt.Printf("[%s]", team)

}

/*
team struct의 상위 struct가 필요하다
fightteam
	ufc
		1
		2
		3

	roadfc
		1
		2
		3

*/
