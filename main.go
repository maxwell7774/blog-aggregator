package main

import (
	"log"
	"os"

	"github.com/maxwell7774/blog-aggregator/internal/config"
)

type state struct {
	cfg *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil { log.Fatalf("error reading config: %v", err)
	}

    programState := &state{
        cfg: &cfg,
    }

    cmds := commands{
        registeredCommands: make(map[string]func(*state, command) error),
    }
    cmds.register("login", handlerLogin)
    
    if len(os.Args) < 2 {
        log.Fatal("Usage: cli <command> [args...]")
        return
    }
    
    var cmd command
    cmd.Name = os.Args[1]
    cmd.Args = os.Args[2:]

    err = cmds.run(programState, cmd)
    if err != nil {
        log.Fatal(err)
    }
}
