package main

import "os"

func calbackExit() error {
	os.Exit(0)
	return nil
}
