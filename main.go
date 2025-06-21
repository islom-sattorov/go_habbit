// Package main demonstrates GORM ORM usage with SQLite database.
// This example shows basic CRUD operations, relationships, and advanced queries
// using GORM with SQLite as the database backend.
package main

import (
	"fmt"

	"go_habbit/database"
)

func main() {
	database.Connect()
	// db := database.GetDB()
	database.Migrate()

	fmt.Println("✅ Database connected and migrated successfully!")

	// Example 1: Create a user
	// fmt.Println("\n📝 Example 1: Creating a user")
	// user := models.User{
	// 	FirstName: "Islomjon", LastName: "Sattorov", Email: "slmsttrv1@gmail.com", Age: 26,
	// }

	// result := db.Create(&user)
	// if result.Error != nil {
	// 	log.Printf("Error creating user: %v", result.Error)
	// } else {
	// 	fmt.Printf("Created user with ID: %d\n", user.ID)
	// }

	// Example 2: Create multiple users
	// fmt.Println("\n📝 Example 2: Creating multiple users")
	// users := []models.User{
	// 	{FirstName: "Islomjon", LastName: "Sattorov", Email: "slmsttrv1@gmail.com", Age: 26},
	// 	{FirstName: "Test", LastName: "User", Email: "jane@example.com", Age: 30},
	// }

	// db.Create(&users)
	// fmt.Printf("Created %d users\n", len(users))

	// // Example 3: Query users
	// fmt.Println("\n🔍 Example 3: Querying users")
	// var allUsers []models.User
	// db.Find(&allUsers)
	// fmt.Printf("Found %d users:\n", len(allUsers))
	// for _, u := range allUsers {
	// 	fmt.Printf("  - %s (%s), Age: %d\n", u.FirstName, u.Email, u.Age)
	// }

	// // Example 4: Find user by email
	// fmt.Println("\n🔍 Example 4: Finding user by email")
	// var foundUser models.User
	// db.Where("email = ?", "slmsttrv1@gmail.com").First(&foundUser)
	// fmt.Printf("Found user: %s\n", foundUser.FirstName)

	// // Example 5: Update user
	// fmt.Println("\n✏️ Example 5: Updating user")
	// db.Model(&foundUser).Update("Age", 26)
	// fmt.Printf("Updated %s's age to 26\n", foundUser.FirstName)

	// // Example 6: Create posts for user
	// fmt.Println("\n📝 Example 6: Creating posts")
	// posts := []Post{
	// 	{Title: "My First Post", Content: "This is my first blog post!", UserID: foundUser.ID},
	// 	{Title: "Learning Go", Content: "Go is an amazing language!", UserID: foundUser.ID},
	// }
	// db.Create(&posts)
	// fmt.Printf("Created %d posts for %s\n", len(posts), foundUser.FirstName)

	// // Example 7: Query with associations
	// // fmt.Println("\n🔍 Example 7: Querying user with posts")
	// // var userWithPosts models.User
	// // db.Preload("Posts").Where("id = ?", foundUser.ID).First(&userWithPosts)
	// // fmt.Printf("User: %s has %d posts:\n", userWithPosts.FirstName, len(userWithPosts.Posts))
	// // for _, post := range userWithPosts.Posts {
	// // 	fmt.Printf("  - %s: %s\n", post.Title, post.Content)
	// // }

	// // Example 8: Create categories
	// fmt.Println("\n📝 Example 8: Creating categories")
	// categories := []Category{
	// 	{Name: "Technology", Description: "Tech-related posts"},
	// 	{Name: "Programming", Description: "Programming tutorials and tips"},
	// 	{Name: "Go Language", Description: "Go programming language content"},
	// }
	// db.Create(&categories)
	// fmt.Printf("Created %d categories\n", len(categories))

	// // Example 9: Many-to-Many relationship
	// fmt.Println("\n🔗 Example 9: Many-to-Many relationships")
	// var techCategory Category
	// db.Where("name = ?", "Technology").First(&techCategory)

	// var firstPost Post
	// db.First(&firstPost)

	// // Associate post with category
	// db.Model(&firstPost).Association("Categories").Append(&techCategory)
	// fmt.Printf("Associated post '%s' with category '%s'\n", firstPost.Title, techCategory.Name)

	// // Example 10: Complex queries
	// fmt.Println("\n🔍 Example 10: Complex queries")
	// var usersOver25 []models.User
	// db.Where("age > ?", 25).Find(&usersOver25)
	// fmt.Printf("Users over 25: %d\n", len(usersOver25))

	// var postCount int64
	// db.Model(&Post{}).Count(&postCount)
	// fmt.Printf("Total posts in database: %d\n", postCount)

	// // Example 11: Delete operations
	// fmt.Println("\n🗑️ Example 11: Delete operations")
	// var userToDelete models.User
	// db.Where("email = ?", "bob@example.com").First(&userToDelete)
	// if userToDelete.ID != 0 {
	// 	db.Delete(&userToDelete)
	// 	fmt.Printf("Deleted user: %s\n", userToDelete.FirstName)
	// }

	// // Example 12: Raw SQL
	// fmt.Println("\n🔍 Example 12: Raw SQL queries")
	// var userCount int
	// db.Raw("SELECT COUNT(*) FROM users").Scan(&userCount)
	// fmt.Printf("Total users (raw SQL): %d\n", userCount)

	// fmt.Println("\n🎉 All examples completed successfully!")
	// fmt.Println("Check the 'habbit.db' file to see your SQLite database.")
}
