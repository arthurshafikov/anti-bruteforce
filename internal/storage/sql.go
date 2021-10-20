package storage

type Storage struct{}

func NewStorage(dsn string) *Storage {
	return &Storage{}
}

func (s *Storage) AddToWhiteList(ip string) error {
	return nil
}

func (s *Storage) RemoveFromWhiteList(ip string) error {
	return nil
}

func (s *Storage) AddToBlackList(ip string) error {
	return nil
}

func (s *Storage) RemoveFromBlackList(ip string) error {
	return nil
}
