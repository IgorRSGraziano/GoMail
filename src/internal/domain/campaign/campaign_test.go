package campaign

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	name     = "Campaign"
	content  = "<div>Hello!</div>"
	contacts = []string{"email1@e.com", "email2@e.com"}
)

func Test_NewCampaign(t *testing.T) {
	a := assert.New(t)
	now := time.Now().Add(-time.Minute)

	campaign, _ := NewCampaign(name, content, contacts)

	a.NotNil(campaign.ID)
	a.Equal(campaign.Name, name)
	a.Equal(campaign.Content, content)
	a.Equal(len(campaign.Contacts), len(contacts))
	a.Equal(len(campaign.Contacts), len(contacts))
	a.Greater(campaign.CreatedAt, now)
}

func Test_NewCampaignInvalid(t *testing.T) {
	a := assert.New(t)

	_, errNameEmpty := NewCampaign("", content, contacts)
	a.NotNil(errNameEmpty)

	_, errContentEmpty := NewCampaign(name, "", contacts)
	a.NotNil(errContentEmpty)

	_, errContactsEmpty := NewCampaign(name, content, []string{})
	a.NotNil(errContactsEmpty)
}
