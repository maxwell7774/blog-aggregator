package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/maxwell7774/blog-aggregator/internal/database"
)

func handlerFeedFollow(s *state, cmd command, user *database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <url>", cmd.Name)
	}
	URL := cmd.Args[0]

	feed, err := s.db.GetFeed(context.Background(), URL)
	if err != nil {
		return fmt.Errorf("couldn't retrieve feed: %w", err)
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

func handlerFeedFollowing(s *state, cmd command, user *database.User) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("usage: %s", cmd.Name)
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

func handlerFeedUnfollow(s *state, cmd command, user *database.User) error {
    if len(cmd.Args) != 1 {
        return fmt.Errorf("usage: %s <url>", cmd.Name)
    }
    URL := cmd.Args[0]

    feed, err := s.db.GetFeed(context.Background(), URL)
    if err != nil {
        return fmt.Errorf("couldn't retrieve feed: %w", err)
    }

    params := database.DeleteFollowFeedParams{
        UserID: user.ID,
        FeedID: feed.ID,
    }
    err = s.db.DeleteFollowFeed(context.Background(), params)
    if err != nil {
        return fmt.Errorf("couldn't unfollow feed: %w", err)
    }

    fmt.Printf("You have unfollowed %s", feed.Name)

    return nil
}
