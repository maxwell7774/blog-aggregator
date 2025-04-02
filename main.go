package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/maxwell7774/blog-aggregator/internal/config"
	"github.com/maxwell7774/blog-aggregator/internal/database"
)

type state struct {
    db *database.Queries
	cfg *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

    db, err := sql.Open("postgres", cfg.DatabaseURL);
    if err != nil {
        log.Fatalf("error opening db connection: %v", err)
    }

    dbQueries := database.New(db)

	programState := &state{
		cfg: &cfg,
        db: dbQueries,
	}

	cmds := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
    cmds.register("users", handlerUsers)
    cmds.register("agg", handlerAgg)
    cmds.register("addfeed", middlewareLoggedIn(handlerAddFeed))
    cmds.register("feeds", handlerListFeeds)
    cmds.register("follow", middlewareLoggedIn(handlerFeedFollow))
    cmds.register("following", middlewareLoggedIn(handlerFeedFollowing))
    cmds.register("unfollow", middlewareLoggedIn(handlerFeedUnfollow))

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
