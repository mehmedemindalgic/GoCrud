package apiserver

import (
	"crud/pkg/postgres"
	"crud/src/internal/service/crudusersevice"
	"crud/src/internal/storage/postgres/cruduserstroage"
	"crud/src/internal/transport/http/cruduserhandler"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type apiServer struct {
	port string
	db   *gorm.DB
}

type Option func(*apiServer)

// default port 8080
func WithPort(port string) Option {
	return func(s *apiServer) {
		s.port = port
	}
}

// default db postgres
func WithDatabase(db *gorm.DB) Option {
	return func(s *apiServer) {
		s.db = db
	}
}

// db pass
func New(option ...Option) error {
	apisrvr := &apiServer{port: "8080", db: postgres.ConnectPQ()}
	for _, o := range option {
		o(apisrvr)
	}

	storage := cruduserstroage.New(cruduserstroage.ConnectDb(apisrvr.db))
	service := crudusersevice.New(crudusersevice.WriteDb(storage))
	transport := cruduserhandler.New(cruduserhandler.WriteService(service))

	// //////////////////////////
	// storage.Crate(model.Usermodel{
	// 	Name:  "test543",
	// 	Email: "asdf45",
	// 	Pass:  "123",
	// })
	//////////////////////////
	//////////////////////////
	//////////////////////////
	//request
	// {
	// 	"name": "1123agsfdh423",
	// 	"email": "dsaagsfdsff",
	// 	"pass": "adffdgfadhsssgd"
	// }

	r := mux.NewRouter()
	router := &http.Server{
		Addr:    ":" + apisrvr.port,
		Handler: r,
	}

	r.HandleFunc("/list/", transport.List).Methods("GET")
	r.HandleFunc("/create/", transport.Create).Methods("POST")
	r.HandleFunc("/update/", transport.Update).Methods("PUT")
	r.HandleFunc("/get/", transport.Get).Methods("GET")
	r.HandleFunc("/delete/", transport.Delete).Methods("DELETE")

	apierror := make(chan error)
	shutdown := make(chan os.Signal, 1)

	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		apierror <- router.ListenAndServe()
	}()
	fmt.Println("Server Pid : " + fmt.Sprint(os.Getpid()))
	fmt.Println("Server is running on port " + "< " + apisrvr.port + " >")
	///////////////////////////////
	//http: superfluous response.WriteHeader call from crud/src/internal/transport/http/cruduserhandler.(*handlerusercrud).Update (crud.go:98)
	select {
	case err := <-apierror:
		return fmt.Errorf("listen and server err: %w", err)
	case sig := <-shutdown:
		fmt.Println("\nStarting shutdown", "pid", sig.String())
		return nil
	}
}
