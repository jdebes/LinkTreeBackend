package server

import (
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/jdebes/LinkTreeBackend/db"
	"github.com/jdebes/LinkTreeBackend/db/repository"
	"github.com/jdebes/LinkTreeBackend/handler"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

type RootHandler struct {
	router    *mux.Router
	db        *sqlx.DB
	mockStore *repository.MockStore
}

func (f *RootHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ctx = db.WithDB(ctx, f.db)
	ctx = repository.WithStore(ctx, f.mockStore)

	f.router.ServeHTTP(w, r.WithContext(ctx))
}

func NewServer() *http.Server {
	sqlDb, err := db.NewDB(db.Config{
		User:     "user",
		Password: "password",
		Host:     "127.0.0.1",
		Port:     3306,
		Database: "linktree",
	})
	if err != nil {
		panic(err)
	}

	mockStore := repository.NewMockStore()

	return &http.Server{
		Addr: ":8080",
		Handler: &RootHandler{
			router:    newRouter(),
			db:        sqlDb,
			mockStore: mockStore,
		},
		ReadTimeout:       15 * time.Second,
		ReadHeaderTimeout: 15 * time.Second,
		WriteTimeout:      15 * time.Second,
	}
}

func newRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/link", handler.GetLink).Methods(http.MethodGet)
	r.HandleFunc("/link", handler.PostLink).Methods(http.MethodPost)

	return r
}

func InitLogging() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}
