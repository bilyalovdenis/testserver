package apiserver

import (
	"html/template"
	"io"
	"net/http"
	"strconv"

	"github.com/bilyalovdenis/testserver/internal/app/store"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type APIServer struct{
	config *Config
	logger *logrus.Logger
	router *mux.Router
	store *store.Store
}

 

func New(config *Config) * APIServer{
	return &APIServer{
		config: config,
		logger: logrus.New(), 
		router: mux.NewRouter(),
	}
}

func (s *APIServer) Start() error{
	if err := s.configureLogger(); err != nil{
		return err
	}

	s.configureRouter()

	if err := s.configureStore(); err != nil {
		return err
	}

	s.logger.Info("starting api server")
	return http.ListenAndServe(s.config.BindAddr, s.router)
}
func (s* APIServer) configureLogger() error{
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil{
		return err
	}
	s.logger.SetLevel(level)
	return nil
}
func (s* APIServer) configureRouter(){
	s.router.HandleFunc("/hello", s.handleHello())
	s.router.HandleFunc("/test", s.handleProduct())
	
} 
func(s *APIServer) handleHello() http.HandlerFunc{

	return func(w http.ResponseWriter, r *http.Request){
		io.WriteString(w, "Hello")
	}
}
func(s *APIServer) handleProduct() http.HandlerFunc{
	return func (w http.ResponseWriter, r *http.Request){
		id, err_id := strconv.Atoi(r.URL.Query().Get("id"))
		if err_id != nil{
			s.logger.Info("неудалось преобразовать")
			return
		}

		p, err := s.store.Product().FindById(id)
		if err == nil{
			s.logger.Info(p.Category)
		} else{
			s.logger.Info("unfind")	
		}

		templ, err_t := template.ParseFiles("./front/templates/product.html")
		if err_t != nil{
			s.logger.Info(err_t)
		}
		templ.Execute(w, p)
	}
}
func(s *APIServer) configureStore() error{
	st := store.New(s.config.Store)
	if err := st.Open(); err != nil{
		return err
	}
	s.store = st

	return nil
}