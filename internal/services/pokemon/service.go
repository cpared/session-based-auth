package pokemon

import repositories "session-based-auth/internal/repositories/pokemon"

type PokemonRepository interface {
	GetPokemonTypeByID(name string) *repositories.Type
}

type Service struct {
	pr PokemonRepository
}

func New(pr PokemonRepository) *Service {
	return &Service{
		pr: pr,
	}
}

func (s *Service) GetPokemonDataByID(id string) *repositories.Type {
	return s.pr.GetPokemonTypeByID(id)
}
