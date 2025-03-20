package campaign

import "time"

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

func NewCampaign(name, content string, emails []string) *Campaign {

	contacts := make([]Contact, len(emails))
	for i, email := range emails {
		contacts[i].Email = email
	}

	return &Campaign{
		ID:        "1", // For now, hard-coded
		Name:      name,
		Content:   content,
		Contacts:  contacts,
		CreatedAt: time.Now(),
	}
}
