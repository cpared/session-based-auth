package pokemon

import (
	"context"
	repositories "session-based-auth/internal/repositories/pokemon"
)

type PokemonRepository interface {
	GetPokemonTypeByID(ctx context.Context, name string) *repositories.Type
}

type Service struct {
	pr PokemonRepository
}

func New(pr PokemonRepository) *Service {
	return &Service{
		pr: pr,
	}
}

func (s *Service) GetPokemonDataByID(ctx context.Context,id string) *repositories.Type {
	return s.pr.GetPokemonTypeByID(ctx, id)
}
