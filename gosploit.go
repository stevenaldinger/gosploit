package main
//go build -buildmode=plugin -o modules/test/eng/chi.so modules/test/eng/.go

import (
    "github.com/stevenaldinger/gosploit/engine"
)

func main() {

	for {
		engine.RunShell()
	}

}
