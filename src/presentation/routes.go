package presentation

import (
	"github/metaufiq/mobile-legend-hero-service/src/common"
	"net/http"
)

func getRoot(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]string)
	resp["data"] = "this is root page, to get heroes data you can access /hero"

	common.CreateJSONResponse(w, resp)
}

func getHero(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]string)
	resp["data"] = "not available"

	common.CreateJSONResponse(w, resp)
}

func ServeRoutes() {
	http.HandleFunc("/", getRoot)
	http.HandleFunc("/hero", getHero)
}
