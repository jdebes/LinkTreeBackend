package handler

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jdebes/LinkTreeBackend/api"
	"github.com/jdebes/LinkTreeBackend/db/repository"
	"github.com/jdebes/LinkTreeBackend/handler/model"
	log "github.com/sirupsen/logrus"
)

func GetLink(w http.ResponseWriter, r *http.Request) {
	// TODO implement ASC/DESC options instead of true false
	// TODO make this query param optional
	shouldOrder, err := strconv.ParseBool(mux.Vars(r)["orderByDate"])
	if err != nil {
		api.Error(w, err, http.StatusBadRequest)
		return
	}

	store, err := repository.Store(r.Context())
	if err != nil {
		api.Error(w, err, http.StatusInternalServerError)
		return
	}

	// TODO remove this hardcoding of the userId and retrieve it via an authentication toke
	links, err := repository.QueryLinks(store, 0, shouldOrder)
	if err != nil {
		api.Error(w, err, http.StatusBadRequest)
	}

	err = api.MarshalResponse(w, links)
	if err != nil {
		api.Error(w, err, http.StatusInternalServerError)
	}
}

func PostLink(w http.ResponseWriter, r *http.Request) {
	store, err := repository.Store(r.Context())
	if err != nil {
		api.Error(w, err, http.StatusInternalServerError)
		return
	}

	var linkRequest model.Link
	err = api.UnmarshalRequest(w, r, &linkRequest)
	if err != nil {
		log.WithError(err).Error("Unable to unmarshal log request")
		return
	}

	createdLink, err := repository.InsertLink(store, linkRequest, 0)
	if err != nil {
		api.Error(w, err, http.StatusInternalServerError)
		return
	}

	err = api.MarshalResponse(w, createdLink)
	if err != nil {
		api.Error(w, err, http.StatusInternalServerError)
	}
}
