package main

import (
	"errors"
	"fmt"

	"github.com/IvanYaremko/repl-poke/pokeapi"
)

func mapf(cfg *config) error {
	response, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationUrl)
	if err != nil {
		return fmt.Errorf("error in mapf respone: %w", err)
	}

	cfg.nextLocationUrl = &response.Next
	cfg.prevLocationUrl = &response.Previous

	for _, loc := range response.Results {
		formatMapMessage(loc)
	}
	return nil
}

func mapb(cfg *config) error {
	if cfg.prevLocationUrl == nil {
		return errors.New("you are on the first page")
	}
	response, err := cfg.pokeapiClient.ListLocations(cfg.prevLocationUrl)
	if err != nil {
		return fmt.Errorf("error in mapf respone: %w", err)
	}

	cfg.nextLocationUrl = &response.Next
	cfg.prevLocationUrl = &response.Previous

	for _, loc := range response.Results {
		formatMapMessage(loc)
	}
	return nil
}

func formatMapMessage(location pokeapi.RespLocationsAreaResults) string {
	format := fmt.Sprintf("Name: %s", location.Name)
	fmt.Println(format)
	return format
}
