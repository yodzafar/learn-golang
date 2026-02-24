package day_5

import (
	"context"
	"errors"
	"learn-golang/syntax/day_5/handler"
	"learn-golang/syntax/day_5/infrastructure"
	"learn-golang/syntax/day_5/middleware"
	"learn-golang/syntax/day_5/usecase"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func Init() {
	repo := &infrastructure.UserRepoDB{}
	service := usecase.NewUserService(repo)
	userHandler := handler.NewUserHandler(service)

	mux := http.NewServeMux()
	mux.HandleFunc("/register", userHandler.Register)

	handlerWithMiddleware := middleware.Logger(
		middleware.ErrorHandler(mux))

	src := &http.Server{
		Addr:    ":8080",
		Handler: handlerWithMiddleware,
	}

	go func() {
		log.Println("Server running on port 8080")
		if err := src.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("ListenAndServe: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := src.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server excited gracefully")
}
