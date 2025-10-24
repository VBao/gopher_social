package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// Use to inject all the required service, config for the application
type application struct {
	config config
}

type config struct {
	addr string
}

func (app *application) mount() http.Handler {
	// Khởi tạo 1 router mới dùng chung cho toàn app
	r := chi.NewRouter()

	// Thêm 1 hoặc nhiều (append) middleware vào router
	// Sử dụng Use() từ *chi.Mux từ chi.NewRouter()

	r.Use(middleware.RequestID)

	r.Use(middleware.RealIP)

	// Dùng middleware Logger để track log lại các request
	r.Use(middleware.Logger)

	// Dùng middleware Recoverer để lưu lại log và lịch sử, strack trace
	// của các request lỗi
	r.Use(middleware.Recoverer)

	r.Route("/v1", func(r chi.Router) {
		r.Get("/health", app.healthCheckHandler)
	})

	return r
}

func (app *application) run(mux http.Handler) error {

	srv := &http.Server{
		Addr:         app.config.addr,
		Handler:      mux,
		WriteTimeout: time.Second * 30, // Server handle request time out
		ReadTimeout:  time.Second * 10, // Client read request time out
		IdleTimeout:  time.Minute,
	}

	log.Printf("Server run at %s", app.config.addr)

	return srv.ListenAndServe()
}
