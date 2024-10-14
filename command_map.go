package main

import (
	"errors"
	"fmt"
)

func mapf(cfg *config) error {
	response, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationUrl)
	if err != nil {
		return fmt.Errorf("error in mapf respone: %w", err)
	}

	cfg.nextLocationUrl = &response.Next
	cfg.prevLocationUrl = &response.Previous

	for _, loc := range response.Results {
		fmString := fmt.Sprintf("Name:	%s\nURL:%s\n", loc.Name, loc.URL)
		fmt.Println(fmString)
	}
	return nil
}

func mapb(cfg *config) error {
	if cfg.prevLocationUrl == nil {
		return errors.New("you are on the first page\n")
	}
	response, err := cfg.pokeapiClient.ListLocations(cfg.prevLocationUrl)
	if err != nil {
		return fmt.Errorf("error in mapf respone: %w", err)
	}

	cfg.nextLocationUrl = &response.Next
	cfg.prevLocationUrl = &response.Previous

	for _, loc := range response.Results {
		fmString := fmt.Sprintf("Name:	%s\nURL:%s\n", loc.Name, loc.URL)
		fmt.Println(fmString)
	}
	return nil
}