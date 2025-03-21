package campaign

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewCampaign_Success(t *testing.T) {
	t.Run("should create a new valid campaign successfully", func(t *testing.T) {
		assert := assert.New(t)
		name := "Black Friday"
		content := "Last day of discounts of up to 80%"
		emails := []string{"valid@example.com", "another@example.com"}
		now := time.Now().Add(-time.Minute)

		campaign, err := NewCampaign(name, content, emails)

		assert.NoError(err)
		assert.Equal(name, campaign.Name, "Expected campaign name to be '%s', but got '%s'", name, campaign.Name)
		assert.Equal(content, campaign.Content, "Expected campaign content to be '%s', but got '%s'", content, campaign.Content)
		assert.Len(campaign.Contacts, 2, "Expected %d contacts, but got %d", len(emails), len(campaign.Contacts))
		assert.Equal(emails[0], campaign.Contacts[0].Email, "Expected first contact email to be '%s', but got '%s'", emails[0], campaign.Contacts[0].Email)
		assert.Equal(emails[1], campaign.Contacts[1].Email, "Expected second contact email to be '%s', but got '%s'", emails[1], campaign.Contacts[1].Email)
		assert.NotEmpty(campaign.ID, "Expected campaign ID to be generated, but got an empty value")
		assert.Greater(campaign.CreatedAt, now, "Expected campaign creation time to be after '%v', but got '%v'", now, campaign.CreatedAt)
	})
}

// TDT - Table Driven Test
func TestNewCampaign_Validation(t *testing.T) {
	tests := []struct {
		name         string
		inputName    string
		inputContent string
		inputEmails  []string
		expectedErr  string
	}{
		{"Empty name", "", "content", []string{"valid@example.com"}, "name is required"},
		{"Empty content", "name", "", []string{"valid@example.com"}, "content is required"},
		{"Empty email", "name", "content", []string{""}, "email contact cannot be empty"},
		{"Empty contact (no emails)", "name", "content", []string{}, "at least one email is required"},
		{"Invalid email", "name", "content", []string{"invalid-email@"}, "invalid email: invalid-email@"},
		{"Empty name with only spaces", "     ", "content", []string{"valid@example.com"}, "name is required"},
		{"Empty content with only spaces", "name", "     ", []string{"valid@example.com"}, "content is required"},
		{"Empty email with only spaces", "name", "content", []string{"     "}, "email contact cannot be empty"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewCampaign(tt.inputName, tt.inputContent, tt.inputEmails)
			assert.EqualError(t, err, tt.expectedErr)
		})
	}
}

/*

Note:

[:] is used to manipulate parts of arrays or slices.

Sintax:
[low:high]

slice[:] Copies the entire slice.
slice[a:] From index a to the end.
slice[:b] From start to index b (excludes).
slice[a:b] From index a to b (excludes b).

*/
