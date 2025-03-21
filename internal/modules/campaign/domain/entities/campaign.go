package campaign

import (
	"errors"
	"net/mail"
	"time"
)

type Contact struct {
	Email string `json:"email"`
}

type Campaign struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Content   string    `json:"content"`
	Contacts  []Contact `json:"contacts"`
	CreatedAt time.Time `json:"created_at"`
}

func isValidEmail(email string) bool {
	emailAddress, err := mail.ParseAddress(email)
	return err == nil && emailAddress.Name == email
}

func NewCampaign(name, content string, emails []string) (*Campaign, error) {
	if name == "" {
		return nil, errors.New("name is required")
	}

	if content == "" {
		return nil, errors.New("content is required")
	}

	if len(emails) == 0 {
		return nil, errors.New("at least one email is required")
	}

	contacts := make([]Contact, len(emails))
	for i, email := range emails {
		if email == "" {
			return nil, errors.New("email contact cannot be empty")
		}

		if !isValidEmail(email) {
			return nil, errors.New("invalid email: " + email)
		}

		contacts[i].Email = email
	}

	return &Campaign{
		ID:        "1",
		Name:      name,
		Content:   content,
		Contacts:  contacts,
		CreatedAt: time.Now(),
	}, nil
}
