package campaign

import (
	"errors"
	"time"

	"github.com/rs/xid"
)

type Campaign struct {
	ID        string
	Name      string
	CreatedAt time.Time
	Content   string
	Contacts  []Contact
	New       *Campaign
}

type Contact struct {
	Email string
}

func NewCampaign(name, content string, emails []string) (*Campaign, error) {

	if name == "" {
		return nil, errors.New("name is required")
	}

	if content == "" {
		return nil, errors.New("content is required")
	}

	if len(emails) == 0 {
		return nil, errors.New("emails is required")
	}

	contacts := make([]Contact, len(emails))
	for index, email := range emails {
		contacts[index].Email = email
	}

	return &Campaign{
		ID:        xid.New().String(),
		CreatedAt: time.Now(),
		Name:      name,
		Content:   content,
		Contacts:  contacts,
	}, nil
}
