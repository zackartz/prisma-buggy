package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/zackartz/prisma-buggy/db"
)

func main() {
	client := db.NewClient()
	if err := client.Prisma.Connect(); err != nil {
		return
	}

	defer func() {
		if err := client.Prisma.Disconnect(); err != nil {
			panic(err)
		}
	}()

	ctx := context.Background()

	u, err := client.User.CreateOne(
		db.User.Username.Set("zack"),
		db.User.Password.Set("a-password"),
		db.User.Email.Set("zackmyers@lavabit.com"),
	).Exec(ctx)

	if err != nil {
		log.Printf("An error occured: %v", err)
		return
	}
	log.Printf("Created user with ID: %s", u.ID)

	updatedUser, err := client.User.FindUnique(
		db.User.ID.Equals(u.ID),
	).Update(
		db.User.Username.Set("mock_username"),
	).Exec(ctx)

	if err != nil {
		log.Panicf("An error occured: %v", err)
	}

	result, _ := json.MarshalIndent(updatedUser, "", "  ")
	fmt.Printf("created post: %s\n", result)
}
