package main

import "os"

func calbackExit(cfg *config, args ...string) error {
	os.Exit(0)
	return nil
}
