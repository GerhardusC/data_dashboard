package handlers

import (
	"log"
	"net/http"
	"data-dashboard/cliArgs"
	"data-dashboard/middleware"
)


func InitHandlers () {
	mux := http.NewServeMux()

	mux.HandleFunc(
		"GET /measurements",
		middleware.LimitRate(allMeasurementsHandler, 0.2, 1),
	)

	mux.HandleFunc("GET /measurements/since", sinceMeasurementsHandler)
	mux.HandleFunc("GET /measurements/between", betweenMeasurementsHandler)

	fs := http.FileServer(http.Dir(cliargs.ServeDir))
	mux.Handle("GET /", fs)

	muxMiddlewareApplied := middleware.NewLogger(mux.ServeHTTP)

	server := http.Server{
		Handler: muxMiddlewareApplied,
	}
	log.Println("Serving on port 80")
	err := server.ListenAndServe()
	log.Println("Failed to serve on port 80")

	if err != nil {
		server.Addr = ":8000"
		log.Println("Falling back to 8000")
		server.ListenAndServe()
	}
}

