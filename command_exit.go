package main

import "os"

func calbackExit(cfg *config) error {
	os.Exit(0)
	return nil
}
