package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/maxwell7774/blog-aggregator/internal/database"
)

func handlerBrowse(s *state, cmd command, user *database.User) error {
	if len(cmd.Args) > 1 {
		return fmt.Errorf("usage: %s (optional)<limit>", cmd.Name)
	}

	limit := 2
	if len(cmd.Args) == 1 {
		if specifiedLimit, err := strconv.Atoi(cmd.Args[0]); err == nil {
			limit = specifiedLimit
		} else {
			return fmt.Errorf("invalid limit: %w", err)
		}
	}

	posts, err := s.db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
	})
	if err != nil {
		return fmt.Errorf("couldn't get posts for user: %w", err)
	}

	for _, post := range posts {
		printPost(post)
	}

	return nil
}

func printPost(post database.GetPostsForUserRow) {
	fmt.Printf("%s from %s\n", post.PublishedAt.Time.Format("Mon Jan 2"), post.FeedName)
	fmt.Printf("--- %s ---\n", post.Title)
	fmt.Printf("    %v\n", post.Description.String)
	fmt.Printf("Link: %s\n", post.Url)
}
