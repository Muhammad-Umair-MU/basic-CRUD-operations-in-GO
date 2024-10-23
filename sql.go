package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	ID         int64
	Name       string
	Email      string
	Profession string
}

func connectDatabase() (*sql.DB, error) {
	source := "root:1331@tcp(127.0.0.1:3306)/xflow"
	database, err := sql.Open("mysql", source)
	if err != nil {
		return nil, err
	}

	// Check if the database is connected
	if error := database.Ping(); error != nil {
		return nil, error
	}

	return database, nil
}

func getAllUsers(db *sql.DB) ([]User, error) {
	query := "SELECT id, name, email, profession FROM etsi_employees"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	// To prevent memory loss
	defer rows.Close()

	var users []User

	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Profession); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func createUser(db *sql.DB, name string, email string, profession string) (int64, error) {
	query := "INSERT INTO etsi_employees (name, email, profession) VALUES (?, ?, ?)"
	result, err := db.Exec(query, name, email, profession)
	if err != nil {
		return 0, err
	}
	// Return the inserted ID
	return result.LastInsertId()
}

func getUserByID(db *sql.DB, id int) (string, string, string, error) {
	var name, email, profession string
	query := "SELECT name, email, age FROM etsi_employees WHERE id = ?"
	err := db.QueryRow(query, id).Scan(&name, &email, &profession)
	if err != nil {
		return "", "", "", err
	}
	return name, email, profession, nil
}

func updateUser(db *sql.DB, id int, name string, email string, profession string) (int64, error) {
	query := "UPDATE users SET name = ?, email = ?, profession = ? WHERE id = ?"
	result, err := db.Exec(query, name, email, profession, id)
	if err != nil {
		return 0, err
	}
	// Return the number of rows affected
	return result.RowsAffected()
}

func deleteUser(db *sql.DB, id int) (int64, error) {
	query := "DELETE FROM etsi_employees WHERE id = ?"
	result, err := db.Exec(query, id)
	if err != nil {
		return 0, err
	}
	// Return the number of rows affected
	return result.RowsAffected()
}

func main() {
	db, err := connectDatabase()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	for {
		fmt.Println("Choose an option:")
		fmt.Println("1. Print All Users")
		fmt.Println("2. Create User")
		fmt.Println("3. Get User by ID")
		fmt.Println("4. Update User")
		fmt.Println("5. Delete User")
		fmt.Println("6. Exit")

		//Scan automaticall assgins value with pointing
		var choice int
		_, err := fmt.Scan(&choice)
		if err != nil {
			fmt.Println("Invalid input. Please try again.")
			continue
		}

		switch choice {
		case 1:
			users, err := getAllUsers(db)
			if err != nil {
				fmt.Println("Error fetching users:", err)
			} else {
				fmt.Println("List of all users:")
				for _, user := range users {
					fmt.Printf("ID: %d, Name: %s, Email: %s, Profession: %s\n", user.ID, user.Name, user.Email, user.Profession)
				}
			}
		case 2:
			var name, email, profession string
			fmt.Print("Enter name: ")
			fmt.Scan(&name)
			fmt.Print("Enter email: ")
			fmt.Scan(&email)
			fmt.Print("Enter profession: ")
			fmt.Scan(&profession)

			id, err := createUser(db, name, email, profession)
			if err != nil {
				fmt.Println("Error creating user:", err)
			} else {
				fmt.Printf("User created with ID: %d\n", id)
			}

		case 3:
			var id int
			fmt.Print("Enter id to search: ")
			fmt.Scan(&id)

			userName, email, profession, err := getUserByID(db, id)
			if err != nil {
				fmt.Println("Error fetching user:", err)
			} else {
				fmt.Printf("User found: Name: %s, Email: %s, Profession: %s\n", userName, email, profession)
			}

		case 4:
			var id int
			var name, email, profession string
			fmt.Print("Enter user ID to update: ")
			fmt.Scan(&id)
			fmt.Print("Enter new name: ")
			fmt.Scan(&name)
			fmt.Print("Enter new email: ")
			fmt.Scan(&email)
			fmt.Print("Enter new profession: ")
			fmt.Scan(&profession)

			affectedRows, err := updateUser(db, id, name, email, profession)
			if err != nil {
				fmt.Println("Error updating user:", err)
			} else if affectedRows == 0 {
				fmt.Println("No user found with the given ID.")
			} else {
				fmt.Printf("User with ID %d updated successfully.\n", id)
			}

		case 5:
			var id int
			fmt.Print("Enter user ID to delete: ")
			fmt.Scan(&id)

			affectedRows, err := deleteUser(db, id)
			if err != nil {
				fmt.Println("Error deleting user:", err)
			} else if affectedRows == 0 {
				fmt.Println("No user found with the given ID.")
			} else {
				fmt.Printf("User with ID %d deleted successfully.\n", id)
			}

		case 6:
			fmt.Println("Exiting...")
			return

		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}
