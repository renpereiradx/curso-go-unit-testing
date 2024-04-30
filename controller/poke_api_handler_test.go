package controller

import (
	"catching-pokemons/models"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gorilla/mux"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestGetPokemonFromPokeApiSuccess(t *testing.T) {
	c := assert.New(t)

	pokemon, err := GetPokemonFromAPI("bulbasaur")
	c.NoError(err)

	body, err := os.ReadFile("samples/poke_api_read.json")
	c.NoError(err)

	var expected models.PokeApiPokemonResponse

	err = json.Unmarshal(body, &expected)
	c.NoError(err)
	c.Equal(expected, pokemon)
}

func TestGetPokemonFromPokeApiSuccessWithMocks(t *testing.T) {
	c := assert.New(t)
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	body, err := os.ReadFile("samples/poke_api_response.json")
	c.NoError(err)

	id := "bulbasaur"

	request := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", id)

	httpmock.RegisterResponder("GET", request, httpmock.NewStringResponder(200, string(body)))

	pokemon, err := GetPokemonFromAPI(id)
	c.NoError(err)

	var expected models.PokeApiPokemonResponse
	err = json.Unmarshal(body, &expected)
	c.NoError(err)
	c.Equal(expected, pokemon)
}

func TestGetPokemonFromPokeApiInternalServerError(t *testing.T) {
	c := assert.New(t)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	id := "balbasaur"

	body, err := os.ReadFile("samples/poke_api_response.json")
	c.NoError(err)

	request := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", id)

	httpmock.RegisterResponder("GET", request, httpmock.NewStringResponder(500, string(body)))

	_, err = GetPokemonFromAPI(id)
	c.NotNil(err)
	c.EqualError(ErrPokeApiFailure, err.Error())
}

func TestGetPokemon(t *testing.T) {
	c := assert.New(t)

	r, err := http.NewRequest("GET", "/pokemon/{id}", nil)
	c.NoError(err)

	w := httptest.NewRecorder()

	vars := map[string]string{
		"id": "bulbasaur",
	}

	r = mux.SetURLVars(r, vars)

	GetPokemon(w, r)

	expectedBodyResponse, err := os.ReadFile("samples/poke_api_response.json")
	c.NoError(err)

	var expectedPokemon models.PokeApiPokemonResponse
	err = json.Unmarshal(expectedBodyResponse, &expectedPokemon)
	c.NoError(err)

	var actualPokemon models.Pokemon

	err = json.Unmarshal(w.Body.Bytes(), &actualPokemon)
	c.NoError(err)

	c.Equal(http.StatusOK, w.Code)
	c.Equal(expectedPokemon.Id, actualPokemon.Id)
}
