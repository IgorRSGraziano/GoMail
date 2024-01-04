package campaign

import (
	"GoMail/src/internal/contract"
	internalerrors "GoMail/src/internal/internal-errors"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type repositoryMock struct {
	mock.Mock
}

var (
	repository = new(repositoryMock)
	service    = Service{Repository: repository}
	campaign   = contract.NewCampaign{
		Name:    "Test",
		Content: "Body",
		Emails:  []string{"teste@gmail.com"},
	}
)

func (r *repositoryMock) Save(campaign *Campaign) error {
	args := r.Called(campaign)
	return args.Error(0)
}

func Test_CreateCampaign(t *testing.T) {

	repository.On("Save", mock.MatchedBy(func(c *Campaign) bool {
		return c.Name == campaign.Name && c.Content == campaign.Content && len(c.Contacts) == len(campaign.Emails)
	})).Return(nil)

	service.Create(campaign)

	repository.AssertExpectations(t)
}

func Test_ReturnError(t *testing.T) {
	a := assert.New(t)

	campaign.Name = ""

	_, err := service.Create(campaign)
	errors.Is(err, internalerrors.ErrInternal)
	a.NotNil(err)
}
