package util

import (
	"catching-pokemons/models"
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParsePokemonSuccess(t *testing.T) {
	// test code
	assert := assert.New(t)

	body, err := os.ReadFile("samples/pokeapi_response.json")
	assert.NoError(err)
	var response models.PokeApiPokemonResponse
	err = json.Unmarshal(body, &response)
	assert.NoError(err)
	parsePokemon, err := ParsePokemon(response)
	assert.NoError(err)

	body, err = os.ReadFile("samples/api_response.json")
	assert.NoError(err)

	var expectedPokemon models.Pokemon
	err = json.Unmarshal(body, &expectedPokemon)
	assert.NoError(err)
	assert.Equal(expectedPokemon, parsePokemon)
}

func TestParsePokemonTypeNotFound(t *testing.T) {
	assert := assert.New(t)

	body, err := os.ReadFile("samples/pokeapi_response.json")
	assert.NoError(err)
	var response models.PokeApiPokemonResponse
	err = json.Unmarshal(body, &response)
	assert.NoError(err)

	response.PokemonType = []models.PokemonType{}
	_, err = ParsePokemon(response)
	assert.NotNil(err)
	assert.EqualError(ErrNotFoundPokemonType, err.Error())
}
