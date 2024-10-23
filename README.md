# User Management Console Application

This GitHub repository contains a simple Go-based command-line interface (CLI) application for managing user data. The application connects to a MySQL database and provides options to create, read, update, and delete user records. It uses the `go-sql-driver/mysql` driver to interact with the database.

## Features
- **Connect to MySQL Database**: Establishes a connection to a MySQL database using the credentials and database details.
- **CRUD Operations**:
  - **Create**: Add a new user to the database with a name, email, and profession.
  - **Read**: Retrieve all users or fetch a specific user by their ID.
  - **Update**: Modify an existing user’s name, email, or profession by their ID.
  - **Delete**: Remove a user from the database by their ID.
- **Interactive CLI**: The application provides a simple interactive menu for user management.

## How to Use
1. **Database Setup**: Ensure you have a MySQL database running with a table named `etsi_employees`:
   ```sql
   CREATE TABLE etsi_employees (
       id INT AUTO_INCREMENT PRIMARY KEY,
       name VARCHAR(255),
       email VARCHAR(255),
       profession VARCHAR(255)
   );

# Run the Application:

1. Clone the repository.
2. Modify the database connection string (in the `connectDatabase` function) with your MySQL credentials and database name.
3. Run the Go application:

    ```bash
    go run main.go
    ```

# Menu Options:

Choose from the following options:
- Print all users
- Create a new user
- Retrieve a user by ID
- Update a user
- Delete a user
- Exit the application

# Pros:
- Simple setup steps that require only basic MySQL and Go knowledge.
- Minimal configuration, with a default structure ready to use.

# Cons:
- Manual setup of MySQL and database table required before running the app.
- No Docker or automated setup scripts provided, which could make deployment harder for some users.

---

# Dependencies:

- **Go 1.18+**: Ensure Go is installed on your system.
- **MySQL**: A running MySQL instance for data storage.
- **go-sql-driver/mysql**: The MySQL driver for Go.

## Pros:
- Minimal dependencies ensure the application is lightweight.
- Standard and widely used MySQL driver (`go-sql-driver/mysql`) for Go applications.

## Cons:
- Limited to MySQL, lacking flexibility for users who prefer other databases.
- May require additional packages for advanced features like database migrations or ORMs.

---

# Error Handling:

The application includes error handling for:
- Invalid inputs
- Database connection issues
- Failed SQL queries

## Pros:
- Basic error handling ensures that the application doesn't crash on common issues.
- Provides informative feedback for database connection problems.

## Cons:
- Error handling is simple, without custom error messages or advanced logging.
- Doesn't provide detailed logs or diagnostics for troubleshooting complex issues.

