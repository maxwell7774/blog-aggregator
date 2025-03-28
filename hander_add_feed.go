package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/maxwell7774/blog-aggregator/internal/database"
)

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.Args) != 2 {
		return fmt.Errorf("usage: %s <name> <url>", cmd.Name)
	}

	currentUser, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("couldn't retrieve current user from database: %w", err)
	}

	createFeedParams := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.Args[0],
		Url:       cmd.Args[1],
		UserID:    currentUser.ID,
	}
	feed, err := s.db.CreateFeed(context.Background(), createFeedParams)
	if err != nil {
		return fmt.Errorf("couldn't create new feed: %w", err)
	}

	fmt.Printf("Created new feed: %v\n", feed)

	return nil
}
