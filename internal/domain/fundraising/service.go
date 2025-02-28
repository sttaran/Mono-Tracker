package fundraising

type Service struct {
	storage      IFundraisingStorage
	infoProvider IFundraisingInfoProvider
}

func NewService(s IFundraisingStorage, ip IFundraisingInfoProvider) IFundraisingService {
	return &Service{
		storage:      s,
		infoProvider: ip,
	}
}

func (s Service) GetFundraisings(dto *FetchListDTO) (*FetchListResponse, error) {
	list, err := s.storage.GetFundraisings(dto)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (s Service) CreateFundraising(fundraising *Fundraising) (int, error) {
	_, err := s.infoProvider.IsFundraisingValid(string(fundraising.URL))
	if err != nil {
		return 0, err
	}

	id, err := s.storage.CreateFundraising(fundraising)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s Service) DeleteFundraising(id int) error {
	err := s.storage.DeleteFundraising(id)
	if err != nil {
		return err
	}
	return nil
}

func (s Service) SynchronizeFundraising(id int, initial bool) error {
	f, err := s.storage.GetFundraisingById(id)
	if err != nil {
		return err
	}
	values, err := s.infoProvider.SynchronizeFundraising(string(f.URL), initial)
	if err != nil {
		return err
	}

	if initial {
		err = s.storage.UpdateFundraising(id, &Fundraising{
			ID:          id,
			Name:        values.Name,
			Description: values.Description,
			Goal:        values.Goal,
			URL:         f.URL,
		})
		return nil
	}

	err = s.storage.AddFundraisingHistory(id, values.Raised)

	return nil
}
