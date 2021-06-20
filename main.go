package main

import (
	"github/ByungHaLee/nomadcoin/explorer"
	"github/ByungHaLee/nomadcoin/rest"
)

func main() {
	go explorer.Start(3000)
	rest.Start(4000)
}
