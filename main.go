package main

/*
20230110
1. Team struct의 상위 struct가 필요한줄 알았는데 아니다
2. ufc팀과 roadfc팀이 각각 다른 TeamName을 가지므로, DB에 insert할때 다른 키를 가진 데이터가된다
{키, [값]}
{ufc, [은가누, 볼카, 콜비]}
{roadfc, [이윤준, 황인수, 명현만]}

전체조회펑션, 선택조회펑션이 나뉘면 된다
*/
import (
	db "GetListProject/db" //database.go의 package를 db로 해야(맨윗줄) 패키지로 인식해서 임포트 가능
	"fmt"
	"net/http"
)

const (
	Team1 = "TeamUFC"
	Team2 = "TeamRoadFC"
)

//var team db.Team
//var port int = 4000

type FightTeam struct {
	organagation []Team
}

//팀 정보 기본 틀
type Team struct {
	TeamName string
	Player   []string
}

func main() {
	var ddddd [3]string
	ddddd[0] = "dsds"
	ddddd[1] = "dsds"
	ddddd[2] = "dsds"
	Team := &[]Team{"UFC", ddddd}
	FightTeam := &FightTeam{Team}

	// FightTeam.organagation[0] = Team

	임시주석처리
	//fmt.Println("==============================Start==============================")

	// defer db.DBClose()

	// //searchTeam()
	// handler := http.NewServeMux()
	// handler.HandleFunc("/", home)
	// fmt.Printf("Listening on http://localhost:%d\n", port)            // %s\n에서 :%d 로 변경 (prvSrc_16)
	// log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), handler)) //서버 열기 // port에서 fmt.Sprintf(":%d", port)변경  (prvSrc_16)

	// fmt.Println("==============================END==============================")

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
