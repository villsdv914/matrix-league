package Backend_Challenge

import (
	"context"
	"github.com/gorilla/mux"
	"league_matrix/common"
	"league_matrix/internal/handlers"
	"league_matrix/internal/service"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Execute() {

	server := startRESTServer(LoadHandler)
	waitForShutdown()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := server.Shutdown(ctx)
	if err != nil {
		log.Fatal(context.Background(), err, "Error while stopping REST server")
	}
	log.Println("Successfully shutdown REST server")

}

func LoadHandler(router *mux.Router) {
	matrixService := service.NewMatrixService()
	matrixHandler := handlers.NewMatrixHandler(matrixService)

	router.HandleFunc("/echo", matrixHandler.FileUpload)
	router.HandleFunc("/matrix-string", matrixHandler.MatrixStringHandler).Methods(http.MethodGet)
	router.HandleFunc("/matrix-invert", matrixHandler.MatrixInvertHandler).Methods(http.MethodGet)
	router.HandleFunc("/matrix-flatten", matrixHandler.MatrixFlattenHandler).Methods(http.MethodGet)
	router.HandleFunc("/matrix-sum", matrixHandler.MatrixSumHandler).Methods(http.MethodGet)
	router.HandleFunc("/matrix-multiply", matrixHandler.MatrixMultiplyHandler).Methods(http.MethodGet)
}

func startRESTServer(handlers func(*mux.Router)) *http.Server {

	router := mux.NewRouter()

	router.NotFoundHandler = http.HandlerFunc(NotFoundHandler)
	router.MethodNotAllowedHandler = http.HandlerFunc(MethodNotAllowedHandler)

	if handlers != nil {
		subrouter := router.PathPrefix("").Subrouter()
		handlers(subrouter)
	}

	server := &http.Server{
		Handler:      router,
		Addr:         ":8080",
		ReadTimeout:  time.Duration(30) * time.Second,
		WriteTimeout: time.Duration(30) * time.Second,
	}

	// Start server
	go func() {
		log.Println(context.Background(), "Starting REST server on port:", 8080)
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatal(context.Background(), err, "REST server cannot be started")
		}
	}()

	return server
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	common.WriteResponse(r.Context(), w, common.ApplicationJSON, http.StatusNotFound, []byte(common.ResourceUnknown))
}

func MethodNotAllowedHandler(w http.ResponseWriter, r *http.Request) {
	common.WriteResponse(r.Context(), w, common.ApplicationJSON, http.StatusMethodNotAllowed, []byte(common.ActionUnavailable))
}
func waitForShutdown() {
	interruptChan := make(chan os.Signal, 1)
	signal.Notify(
		interruptChan,
		os.Interrupt,
		syscall.SIGINT|syscall.SIGALRM|syscall.SIGBUS|syscall.SIGABRT,
		syscall.SIGTERM)

	s := <-interruptChan
	log.Println(context.Background(), "Signal received: ", s)
}
