package handler

import (
	"context"
	"encoding/json"
	"learn-golang/syntax/day_5/usecase"
	"learn-golang/syntax/day_5/utils"
	"log"
	"net/http"
	"time"
)

type UserHandler struct {
	service *usecase.UserService
}

func NewUserHandler(service *usecase.UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	log.Printf("Incoming request: %s %s", r.Method, r.URL)

	if r.Method != http.MethodPost {
		utils.JSON(w, http.StatusMethodNotAllowed, utils.ApiResponse{
			Success: false,
			Error:   http.StatusText(http.StatusMethodNotAllowed),
		})
		return
	}

	var input struct {
		Name string `json:"name"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		utils.JSON(w, http.StatusBadRequest, utils.ApiResponse{
			Success: false,
			Error:   "Invalid request body",
		})
		return
	}

	if input.Name == "" {
		utils.JSON(w, http.StatusBadRequest, utils.ApiResponse{
			Success: false,
			Error:   "name is required",
		})
	}

	ctx := r.Context()
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)

	defer cancel()

	done := make(chan error, 1)

	go func() {
		done <- h.service.Register(input.Name)
	}()

	select {
	case <-ctx.Done():
		utils.JSON(w, http.StatusRequestTimeout, utils.ApiResponse{
			Success: false,
			Error:   "request timed out",
		})
		return
	case err := <-done:
		if err != nil {
			utils.JSON(w, http.StatusInternalServerError, utils.ApiResponse{
				Success: false,
				Error:   err.Error(),
			})

			return
		}
		utils.JSON(w, http.StatusCreated, utils.ApiResponse{
			Success: true,
			Data:    "user registered successfully",
		})
	}

}
