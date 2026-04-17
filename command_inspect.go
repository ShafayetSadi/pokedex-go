package main

import (
	"errors"
	"fmt"

	"github.com/shafayetsadi/pokedexcli/internal/pokeapi"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}

	inspectPokemon(args[0], cfg.pokedex)

	return nil
}

func inspectPokemon(name string, pokedex map[string]pokeapi.Pokemon) {
	p, ok := pokedex[name]
	if !ok {
		fmt.Println("you have not caught that pokemon")
		return
	}

	fmt.Printf("Name: %s\n", p.Name)
	fmt.Printf("Height: %d\n", p.Height)
	fmt.Printf("Weight: %d\n", p.Weight)

	fmt.Println("Stats:")
	for _, stat := range p.Stats {
		fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}

	fmt.Println("Types:")
	for _, t := range p.Types {
		fmt.Printf("  - %s\n", t.Type.Name)
	}
}
