package campaign

import "GoMail/src/internal/contract"

type Service struct {
	Repository Repository
}

func (s *Service) Create(newCampaign contract.NewCampaign) error {
	return nil
}
