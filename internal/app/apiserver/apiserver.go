package apiserver

import (
	"io"
	"net/http"

	"github.com/evgensr/http-rest-api/internal/app/store"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// APIserver ...
type APIserver struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
	store  *store.Store
}

// New ...
func New(Config *Config) *APIserver {
	return &APIserver{
		config: Config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

// Start ...
func (s *APIserver) Start() error {

	if err := s.configerLogger(); err != nil {
		return err

	}

	s.ConfigerRouter()

	if err := s.configerStore(); err != nil {
		return err
	}

	s.logger.Info("Staring apiserver")
	return http.ListenAndServe(s.config.BindAddr, s.router)
}

func (s *APIserver) configerLogger() error {

	level, err := logrus.ParseLevel(s.config.LogLevel)

	if err != nil {
		return err
	}

	s.logger.SetLevel(level)

	return nil

}

func (s *APIserver) ConfigerRouter() {

	s.router.HandleFunc("/hello", s.HandlerHello())

}

func (s *APIserver) configerStore() error {
	st := store.New(s.config.Store)
	if err := st.Open(); err != nil {
		return err
	}

	s.store = st
	return nil

}

func (s *APIserver) HandlerHello() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "hello")
	}

}
