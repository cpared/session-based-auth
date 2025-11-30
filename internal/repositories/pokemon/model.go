package pokemon

// Type represents a Pokémon type.
type Type struct {

	// ID is the identifier for this resource.
	ID int `json:"id"`

	// Name is the name for this resource.
	Name string `json:"name"`

	// DamageRelations describes how effective this type is toward others and vice versa.
	DamageRelations TypeRelations `json:"damage_relations"`

	// PastDamageRelations lists how effective this type was toward others in previous generations.
	PastDamageRelations []TypeRelationsPast `json:"past_damage_relations"`

	// GameIndices is a list of game indices relevant to this type by generation.
	GameIndices []GenerationGameIndex `json:"game_indices"`

	// Generation is the generation where this type was first introduced.
	Generation NamedAPIResource `json:"generation"`

	// MoveDamageClass is the damage class inflicted by this type.
	MoveDamageClass NamedAPIResource `json:"move_damage_class"`

	// Names lists localized names for this resource in different languages.
	Names []Name `json:"names"`

	// Pokemon lists the details of Pokémon that have this type.
	Pokemon []TypePokemon `json:"pokemon"`

	// Moves is a list of moves that have this type.
	Moves []NamedAPIResource `json:"moves"`
}

// TypePokemon represents a Pokémon that has this type.
type TypePokemon struct {

	// Slot is the order the Pokémon’s types are listed in.
	Slot int `json:"slot"`

	// Pokemon is the Pokémon that has the referenced type.
	Pokemon NamedAPIResource `json:"pokemon"`
}

// TypeRelations describes attack and defense effectiveness for this type.
type TypeRelations struct {

	// NoDamageTo lists types this type has no effect on.
	NoDamageTo []NamedAPIResource `json:"no_damage_to"`

	// HalfDamageTo lists types this type is not very effective against.
	HalfDamageTo []NamedAPIResource `json:"half_damage_to"`

	// DoubleDamageTo lists types this type is very effective against.
	DoubleDamageTo []NamedAPIResource `json:"double_damage_to"`

	// NoDamageFrom lists types that have no effect on this type.
	NoDamageFrom []NamedAPIResource `json:"no_damage_from"`

	// HalfDamageFrom lists types that are not very effective against this type.
	HalfDamageFrom []NamedAPIResource `json:"half_damage_from"`

	// DoubleDamageFrom lists types that are very effective against this type.
	DoubleDamageFrom []NamedAPIResource `json:"double_damage_from"`
}

// TypeRelationsPast describes damage relations from a previous generation.
type TypeRelationsPast struct {

	// Generation is the last generation in which these damage relations applied.
	Generation NamedAPIResource `json:"generation"`

	// DamageRelations are the relations that applied up to and including the listed generation.
	DamageRelations TypeRelations `json:"damage_relations"`
}

// GenerationGameIndex represents a game index for a type in a specific generation.
type GenerationGameIndex struct {

	// GameIndex is the internal game index.
	GameIndex int `json:"game_index"`

	// Generation is the generation associated with this index.
	Generation NamedAPIResource `json:"generation"`
}

// Name represents a localized name of a resource.
type Name struct {

	// Name is the localized name.
	Name string `json:"name"`

	// Language is the language of the localized name.
	Language NamedAPIResource `json:"language"`
}

// NamedAPIResource references another API resource by name and URL.
type NamedAPIResource struct {

	// Name is the name of the referenced resource.
	Name string `json:"name"`

	// URL is the URL of the referenced resource.
	URL string `json:"url"`
}
