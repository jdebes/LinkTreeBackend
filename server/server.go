package server

import (
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/jdebes/LinkTreeBackend/db/repository"
	"github.com/jdebes/LinkTreeBackend/handler"
	log "github.com/sirupsen/logrus"
)

type RootHandler struct {
	router    *mux.Router
	mockStore *repository.MockStore
}

func (f *RootHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ctx = repository.WithStore(ctx, f.mockStore)

	f.router.ServeHTTP(w, r.WithContext(ctx))
}

func NewServer() *http.Server {
	mockStore := repository.NewMockStore()

	return &http.Server{
		Addr: ":8080",
		Handler: &RootHandler{
			router:    newRouter(),
			mockStore: mockStore,
		},
		ReadTimeout:       15 * time.Second,
		ReadHeaderTimeout: 15 * time.Second,
		WriteTimeout:      15 * time.Second,
	}
}

func newRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/links", handler.GetLink).
		Queries("orderByDate", "{orderByDate}").
		Methods(http.MethodGet)
	r.HandleFunc("/links", handler.PostLink).
		Methods(http.MethodPost)

	return r
}

func InitLogging() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}
