package main

import (
	"errors"
	"fmt"
	"math/rand"

	"github.com/shafayetsadi/pokedexcli/internal/pokeapi"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}

	name := args[0]
	fmt.Printf("Throwing a Pokeball at %s...\n", name)
	pokemon, err := cfg.pokeapiClient.GetPokemon(name)
	if err != nil {
		return err
	}

	if catchPokemon(pokemon) {
		fmt.Printf("%s was caught!\n", name)
		cfg.pokedex[name] = pokemon
	} else {
		fmt.Printf("%s escaped!\n", name)
	}
	return nil
}

func catchPokemon(pokemon pokeapi.Pokemon) bool {
	chance := rand.Intn(pokemon.BaseExperience + 20)
	if chance < 40 {
		return true
	} else {
		return false
	}
}
