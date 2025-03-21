package campaign

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCreateNewCampaign(t *testing.T) {
	t.Run("should create a new campaign successfully", func(t *testing.T) {
		assert := assert.New(t)
		name := "Black Friday"
		content := "Last day of discounts of up to 80%"
		emails := []string{"valid@example.com", "another@example.com"}
		now := time.Now().Add(-time.Minute)

		c, _ := NewCampaign(name, content, emails)

		assert.Equal(name, c.Name, "Expected campaign name to be '%s', but got '%s'", name, c.Name)
		assert.Equal(content, c.Content, "Expected campaign content to be '%s', but got '%s'", content, c.Content)
		assert.Equal(len(c.Contacts), len(emails), "Expected %d contacts, but got %d", len(emails), len(c.Contacts))
		assert.Equal(emails[0], c.Contacts[0].Email, "Expected first contact email to be '%s', but got '%s'", emails[0], c.Contacts[0].Email)
		assert.Equal(emails[1], c.Contacts[1].Email, "Expected second contact email to be '%s', but got '%s'", emails[1], c.Contacts[1].Email)
		assert.NotEmpty(c.ID, "Expected campaign ID to be generated, but got an empty value")
		assert.Greater(c.CreatedAt, now, "Expected campaign creation time to be after '%v', but got '%v'", now, c.CreatedAt)
	})
}

func TestCreateNewCampaignWithEmptyNameError(t *testing.T) {
	t.Run("should return error if name is empty", func(t *testing.T) {
		assert := assert.New(t)
		name := ""
		content := "Last day of discounts of up to 80%"
		emails := []string{"valid@example.com", "another@example.com"}

		_, err := NewCampaign(name, content, emails)

		assert.Equal("name is required", err.Error())
	})
}

func TestCreateNewCampaignWithEmptyContentError(t *testing.T) {
	t.Run("should return error if content is empty", func(t *testing.T) {
		assert := assert.New(t)
		name := "Black Friday"
		content := ""
		emails := []string{"valid@example.com", "another@example.com"}

		_, err := NewCampaign(name, content, emails)

		assert.Equal("content is required", err.Error())
	})
}

func TestCreateNewCampaignWithEmptyContactError(t *testing.T) {
	assert := assert.New(t)
	name := "Black Friday"
	content := "Last day of discounts of up to 80%"
	emails := []string{"valid@example.com", ""}

	_, err := NewCampaign(name, content, emails)

	assert.Equal("email contact cannot be empty", err.Error())
}

func TestCreateNewCampaignWithZeroContactsError(t *testing.T) {
	t.Run("should return error if no emails provided", func(t *testing.T) {
		assert := assert.New(t)
		name := "Black Friday"
		content := "Last day of discounts of up to 80%"
		emails := []string{}

		_, err := NewCampaign(name, content, emails)

		assert.Equal("at least one email is required", err.Error())
	})
}

func TestCreateNewCampaignWithValidEmailError(t *testing.T) {
	t.Run("should return error if email is invalid", func(t *testing.T) {
		assert := assert.New(t)
		name := "Black Friday"
		content := "Last day of discounts of up to 80%"
		emails := []string{"invalid-email@"}

		_, err := NewCampaign(name, content, emails)

		assert.Equal("invalid email: invalid-email@", err.Error())
	})
}
