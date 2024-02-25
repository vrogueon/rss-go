package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/vrogueon/rss-go/internal/database"
)

func (apiCfg *apiConfig)handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `name`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: ", err))
	}

	apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		
	})

	respondWithJSON(w, 200, struct{}{})
}