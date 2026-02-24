package usecase

import (
	"learn-golang/syntax/day_5/domain"
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

func TestRegister(t *testing.T) {
	mockRepo := &MockUserRepo{}
	service := NewUserService(mockRepo)

	err := service.Register("Ali")

	if err != nil {
		t.Fatal("expected no error, got: ", err)
	}

	if len(mockRepo.SavedUsers) != 1 {
		t.Fatal("expected 1 user, got: ", mockRepo.SavedUsers)
	}

	if mockRepo.SavedUsers[0].Name != "Ali" {
		t.Fatal("expected Ali, got: ", mockRepo.SavedUsers[0].Name)
	}
}
