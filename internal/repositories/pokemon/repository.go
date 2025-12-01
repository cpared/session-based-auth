package pokemon

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
)

// Pokemon documentation API --> https://pokeapi.co/docs/v2#types
const (
	URL = "https://pokeapi.co/api/v2/type/{id}"
)

type Repository struct {
	client *resty.Client
}

func New() *Repository {
	return &Repository{
		client: resty.New(),
	}
}

// TODO: This is a bad practice to retorn the same data access object (DAO) that handle the repository
// This should convert into a data transfer object (DTO) that should be a domain object
func (r *Repository) GetPokemonTypeByID(ctx context.Context,name string) *Type {
	resp, err := r.client.R().Get(strings.ReplaceAll(URL, "{id}", name))
	if err != nil {
		fmt.Println("cannot get pokemon data err: %v", err)
		return nil
	}

	var data Type
	if err := json.Unmarshal(resp.Body(), &data); err != nil {
		fmt.Println("cannot unmarshal json response err: %v", err)
		return nil
	}

	return &data
}
