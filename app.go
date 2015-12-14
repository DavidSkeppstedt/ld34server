package main

import (
	"github.com/DavidSkeppstedt/ld34server/game"
	"github.com/DavidSkeppstedt/ld34server/host"
	"runtime"
)

func init() {
	numcpu := runtime.NumCPU()
	runtime.GOMAXPROCS(numcpu) // Try to use all available CPUs.
}
func main() {
	game.StartGame()
	host.ListenAndServe()
}
