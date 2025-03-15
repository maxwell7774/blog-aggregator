package main

import (
	"fmt"
	"log"

	"github.com/maxwell7774/blog-aggregator/internal/config"
)

func main() {
    cfg, err := config.Read()
    if err != nil {
        log.Fatalf("error reading config: %v", err)
    }
    fmt.Printf("Read config: %+v\n", cfg)

    err = cfg.SetUser("adam")
    if err != nil {
        log.Fatalf("error setting user: %v", err)
    }

    cfg, err = config.Read()
    if err != nil {
        log.Fatalf("error reading config: %v", err)
    }
    fmt.Printf("Read config again: %+v\n", cfg)
}
