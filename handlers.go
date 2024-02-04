package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/Altyn119/BayernTeam/api"
	"github.com/gorilla/mux"
)

func GetPlayers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(api.Squad)
}

func GetName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	para := mux.Vars(r)
	PlayerName := para["last_name"]

	PlayerName = strings.ToLower(PlayerName)

	for _, player := range api.Squad {
		if strings.ToLower(player.LastName) == PlayerName {
			json.NewEncoder(w).Encode(player)
			return
		}
	}
	http.Error(w, "Player not found", http.StatusNotFound)
}

func GetID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	para := mux.Vars(r)
	PlayerID := para["id"]

	for _, player := range api.Squad {
		if strconv.Itoa(player.Id) == PlayerID {
			json.NewEncoder(w).Encode(player)
			return
		}
	}
	http.Error(w, "Player not found", http.StatusNotFound)
}

func GetNumber(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	para := mux.Vars(r)
	PlayerNumber := para["number"]

	for _, player := range api.Squad {
		if strconv.Itoa(player.Number) == PlayerNumber {
			json.NewEncoder(w).Encode(player)
			return
		}
	}
	http.Error(w, "Player not found", http.StatusNotFound)
}

func GetPosition(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	pos := params["position"]
	pos = strings.ToLower(pos)
	var matchingPlayers []api.Player

	for _, player := range api.Squad {
		if strings.ToLower(player.Position) == pos {
			matchingPlayers = append(matchingPlayers, player)
		}
	}

	if len(matchingPlayers) > 0 {
		json.NewEncoder(w).Encode(matchingPlayers)
	} else {
		http.Error(w, "Players with the specified position not found", http.StatusNotFound)
	}
}

func GetNation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	nation := params["nation"]

	nation = strings.ToLower(nation)
	var matchingPlayers []api.Player

	for _, player := range api.Squad {
		if strings.ToLower(player.Nation) == nation {
			matchingPlayers = append(matchingPlayers, player)
		}
	}

	if len(matchingPlayers) > 0 {
		json.NewEncoder(w).Encode(matchingPlayers)
	} else {
		http.Error(w, "Players from the specified nation not found", http.StatusNotFound)
	}
}
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")

	healthStatus := "Health!\n\n"

	Description := "It is a simple web API that provides some information about Bayern players. Also it checks app's health./n"

	authorInfo := "\n Altyyyn \n"

	healthCheckResponse := healthStatus + Description + authorInfo

	w.Write([]byte(healthCheckResponse))
}
