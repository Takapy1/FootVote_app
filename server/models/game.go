package models

import (

)

type Game struct {
	ID         uint    `json:"id"`
	MyTeam     string  `json:"my_team"`
	EnemyTeam  string  `json:"enemy_team"`
	Date       string  `json:"date"`
}

func GetGames() []*Game {
	Games := []*Game{
		&Game{ID: 1, MyTeam: "日本", EnemyTeam: "サウジアラビア", Date: "2021/10/02"},
		&Game{ID: 2, MyTeam: "オーストラリア", EnemyTeam: "オマーン", Date: "2021/11/06"},
		&Game{ID: 3, MyTeam: "オーストラリア", EnemyTeam: "日本", Date: "2021/11/9"},
	}
	return Games
}



// 後でアソシエーションでチームを管理する

/* 
	後に必要になりそうなカラム
		スタジアム

*/
