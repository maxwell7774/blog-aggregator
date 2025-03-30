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

    feedFollowParams := database.CreateFollowFeedParams{
        ID: uuid.New(),
        CreatedAt: time.Now(),
        UpdatedAt: time.Now(),
        UserID: currentUser.ID,
        FeedID: feed.ID,
    }
    follow, err := s.db.CreateFollowFeed(context.Background(), feedFollowParams)
    if err != nil {
        return fmt.Errorf("couldn't follow new feed: %w", err)
    }
	fmt.Printf("%s is now following %s\n", follow.UserName, follow.FeedName)

	return nil
}

func handlerListFeeds(s *state, cmd command) error {
	if len(cmd.Args) > 0 {
		return fmt.Errorf("usage: %s", cmd.Name)
	}

	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't retrieve feeds: %w", err)
	}

	for _, feed := range feeds {
		user, err := s.db.GetUserById(context.Background(), feed.UserID)
		if err != nil {
			return fmt.Errorf("couldn't get user: %w", err)
		}
		printFeed(feed, user)
		fmt.Println("=====================================")
	}

	return nil
}

func printFeed(f database.Feed, u database.User) {
	fmt.Printf("* ID:            %s\n", f.ID)
	fmt.Printf("* Created:       %v\n", f.CreatedAt)
	fmt.Printf("* Updated:       %v\n", f.UpdatedAt)
	fmt.Printf("* Name:          %s\n", f.Name)
	fmt.Printf("* URL:           %s\n", f.Url)
	fmt.Printf("* User:          %s\n", u.Name)
}
