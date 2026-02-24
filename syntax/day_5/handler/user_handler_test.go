package handler

import (
	"bytes"
	"learn-golang/syntax/day_5/domain"
	"learn-golang/syntax/day_5/usecase"
	"net/http"
	"net/http/httptest"
	"testing"
)

type MockUserRepo struct {
	SavedUsers []domain.User
}

func (m *MockUserRepo) Save(user domain.User) error {
	m.SavedUsers = append(m.SavedUsers, user)
	return nil
}

func (m *MockUserRepo) FindByID(id int) (domain.User, error) {
	for _, u := range m.SavedUsers {
		if u.ID == id {
			return u, nil
		}
	}

	return domain.User{}, nil
}

func TestRegisterHandler(t *testing.T) {
	mockUserRepo := &MockUserRepo{}
	service := usecase.NewUserService(mockUserRepo)
	handler := NewUserHandler(service)

	body := bytes.NewBufferString(`{"name":"Ali"}`)
	req := httptest.NewRequest(http.MethodPost, "/register", body)
	w := httptest.NewRecorder()

	handler.Register(w, req)
	resp := w.Result()

	if resp.StatusCode != http.StatusCreated {
		t.Fatalf("expected status 201, got %d", resp.StatusCode)
	}

	if len(mockUserRepo.SavedUsers) != 1 || mockUserRepo.SavedUsers[0].Name != "Ali" {
		t.Fatal("User was not saved correctly in mock repo")
	}
}
