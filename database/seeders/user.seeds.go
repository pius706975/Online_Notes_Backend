package seeders

import "github.com/pius706975/backend/database/models"

// both passwords are the same, pius1234

var UserSeeds = models.Users {
	{
		Name: "Anak Gahol",
		Username: "pipiw",
		Email: "piusrestiantoro02@gmail.com",
		Password: "$2a$10$CSYg2gWZQuOYGlO9eNMI6ONE5KsEfPF0wbu7oqBxIDUGYhxnNS9QW",
		Gender: "Male",
		Country: "Irak",
		DateOfBirth: "",
		MobileNumber: "",
		Role: "user",
		Image: "https://res.cloudinary.com/dccomkorf/image/upload/v1684937069/vehiclerental/1684937065423363500.jpg",
		IsVerified: true,
	},

	{
		Name: "Admin Tamvan",
		Username: "admintamvan",
		Email: "piusrestiantoro2@gmail.com",
		Password: "$2a$10$L9vLCvsHQHGayTQtHC1WwOBl7ItV2n/hTFlQzPQ8Get2IgFbFTRDG",
		Gender: "Female",
		Country: "Russia",
		DateOfBirth: "",
		MobileNumber: "",
		Role: "admin",
		Image: "",
		IsVerified: true,
	},
}