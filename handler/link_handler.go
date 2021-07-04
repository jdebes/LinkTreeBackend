package handler

import (
	"net/http"

	"github.com/jdebes/LinkTreeBackend/api"
	"github.com/jdebes/LinkTreeBackend/db"
)

func GetLink(w http.ResponseWriter, r *http.Request) {
	db, err := db.Store(r.Context())
	if err != nil {
		api.Error(w, err, http.StatusInternalServerError)
	}

	err = api.MarshalResponse(w, db.Links())
	if err != nil {
		api.Error(w, err, http.StatusInternalServerError)
	}
}

func PostLink(w http.ResponseWriter, r *http.Request) {

}
