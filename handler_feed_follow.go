package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/maxwell7774/blog-aggregator/internal/database"
)

func handlerFeedFollow(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <url>", cmd.Name)
	}
	URL := cmd.Args[0]

	feed, err := s.db.GetFeed(context.Background(), URL)
	if err != nil {
		return fmt.Errorf("couldn't retrieve feed: %w", err)
	}

	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("couldn't retrieve current user: %w", err)
	}

	params := database.CreateFollowFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	}

	follow, err := s.db.CreateFollowFeed(context.Background(), params)
	if err != nil {
		return fmt.Errorf("couldn't follow feed: %w", err)
	}

	fmt.Printf("%s is now following %s\n", follow.UserName, follow.FeedName)

	return nil
}

func handlerFeedFollowing(s *state, cmd command) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("usage: %s", cmd.Name)
	}

	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("couldn't retrieve current user: %w", err)
	}

	follows, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("couldn't retrieve following: %w", err)
	}

	for _, follow := range follows {
        fmt.Printf("* %s\n", follow.FeedName)
	}
    return nil
}
