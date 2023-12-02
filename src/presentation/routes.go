package presentation

import (
	"fmt"
	"github/metaufiq/mobile-legend-hero-service/src/common"
	"net/http"
)

func getRoot(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]string)
	resp["data"] = "cannot access your data here"

	common.CreateJSONResponse(w, resp)
}

func getHero(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]any)
	db := common.InnitDB()
	rows, err := db.Query("SELECT * FROM hero_counter")

	if err != nil {

		fmt.Println(err.Error())
		return
	}

	var data []map[string]int
	for rows.Next() {
		var id int
		var heroId int
		var counterHeroId int
		err = rows.Scan(&id, &heroId, &counterHeroId)
		if err != nil {
			fmt.Println(err.Error())
			break
		}
		hero := map[string]int{
			"id":            id,
			"heroId":        heroId,
			"counterHeroId": counterHeroId,
		}
		data = append(data, hero)
	}
	resp["data"] = data

	common.CreateJSONResponse(w, resp)
}

func ServeRoutes() {
	http.HandleFunc("/", getRoot)
	http.HandleFunc("/hero", getHero)
}
