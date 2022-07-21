package database

import (
	"bookstore/dto"
	"bookstore/util"
)

func initSeeder() (err error) {
	userSeeder()
	bookSeeder()
	return
}

func userSeeder() (err error) {
	// Create user's dummy data
	pass, _ := util.HashPassword("test123")
	user := dto.User{
		Name:     "User",
		Email:    "user@mail.com",
		Password: pass,
	}

	DB.Create(&user) // add user data to database
	return
}

func bookSeeder() (err error) {
	// Create book's dummy data
	books := []dto.Book{
		{
			Name:        "The Things You Can See Only When You Slow Down",
			Author:      "Sunim, Haemin",
			Description: "Is it the world that's busy, or is it my mind? The world moves fast, but that doesn't mean we have to. ",
			Publisher:   "Penguin Books Ltd",
			ISBN:        "9780241340660",
		},
		{
			Name:        "Homo Deus",
			Author:      "Yuval Noah Harari",
			Description: "Shows us where mankind is headed in an absolutely clear-sighted and accessible manner' JARVIS COCKER",
			Publisher:   "Random House Uk",
			ISBN:        "9781784703936",
		},
	}

	for _, book := range books {
		DB.Create(&book) // add books data to database
	}
	return
}
