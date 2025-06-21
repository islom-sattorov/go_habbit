# Go Habbit - GORM + SQLite Starter Project

A comprehensive Go starter project demonstrating GORM ORM with SQLite database, featuring proper project structure, repository pattern, and extensive examples.

## 🚀 Features

- **GORM ORM** with SQLite database
- **Repository Pattern** for clean data access
- **Proper project structure** with separation of concerns
- **Comprehensive examples** covering all CRUD operations
- **Model relationships** (One-to-Many, Many-to-Many)
- **GORM hooks** and custom methods
- **JSON serialization** support
- **Database migrations** with auto-migration

## 📁 Project Structure

```
go_habbit/
├── main.go                 # Main application with basic examples
├── go.mod                  # Go module file
├── README.md              # This file
├── models/                # Data models
│   ├── user.go           # User model with relationships
│   ├── post.go           # Post model with relationships
│   └── category.go       # Category model
├── database/             # Database connection
│   └── connection.go     # Database setup and connection
├── repositories/         # Repository pattern implementation
│   └── user_repository.go # User repository with CRUD operations
└── examples/             # Additional examples
    └── repository_example.go # Repository pattern examples
```

## 🛠️ Installation

1. **Clone or create the project:**
   ```bash
   mkdir go_habbit
   cd go_habbit
   ```

2. **Initialize the Go module:**
   ```bash
   go mod init go_habbit
   ```

3. **Install dependencies:**
   ```bash
   go mod tidy
   ```

## 🚀 Quick Start

### Run the main example:
```bash
go run main.go
```

### Run the repository pattern example:
```bash
go run examples/repository_example.go
```

## 📊 Database Models

### User Model
- Basic user information (name, email, age)
- One-to-Many relationship with posts
- GORM hooks for validation
- Custom methods for business logic

### Post Model
- Blog post with title and content
- Many-to-One relationship with users
- Many-to-Many relationship with categories
- Custom methods for content management

### Category Model
- Post categories with name and description
- Many-to-Many relationship with posts
- Utility methods for category management

## 🔧 Examples Included

### Basic GORM Operations (`main.go`)
1. **Database Connection** - Setting up GORM with SQLite
2. **Auto Migration** - Automatic schema creation
3. **Create Operations** - Single and batch record creation
4. **Query Operations** - Finding, filtering, and counting records
5. **Update Operations** - Modifying existing records
6. **Delete Operations** - Removing records
7. **Relationships** - One-to-Many and Many-to-Many associations
8. **Raw SQL** - Custom SQL queries
9. **Complex Queries** - Advanced filtering and aggregation

### Repository Pattern (`examples/repository_example.go`)
1. **Repository Setup** - Creating and using repositories
2. **CRUD Operations** - Create, Read, Update, Delete
3. **Custom Queries** - Age range filtering, counting
4. **Relationship Loading** - Loading associated data
5. **Error Handling** - Proper error management

## 🗄️ Database Schema

The application automatically creates the following tables:

- **users** - User information
- **posts** - Blog posts
- **categories** - Post categories
- **post_categories** - Many-to-Many relationship table

## 🔗 Relationships

- **User → Posts**: One-to-Many (a user can have multiple posts)
- **Post → User**: Many-to-One (each post belongs to one user)
- **Post ↔ Categories**: Many-to-Many (posts can have multiple categories)

## 🎯 Key Features Demonstrated

### GORM Features
- **Auto Migration**: Automatic table creation
- **Hooks**: BeforeCreate, BeforeUpdate, etc.
- **Associations**: Preloading related data
- **Transactions**: Database transaction support
- **Raw SQL**: Custom SQL queries
- **JSON Tags**: Serialization support

### Go Best Practices
- **Repository Pattern**: Clean separation of data access
- **Error Handling**: Proper error management
- **Project Structure**: Organized code layout
- **Documentation**: Comprehensive comments and examples

## 🧪 Testing the Examples

1. **Run the main example:**
   ```bash
   go run main.go
   ```
   This will create a SQLite database file (`habbit.db`) and demonstrate all basic GORM operations.

2. **Run the repository example:**
   ```bash
   go run examples/repository_example.go
   ```
   This demonstrates the repository pattern with proper separation of concerns.

3. **Inspect the database:**
   ```bash
   sqlite3 habbit.db
   .tables
   SELECT * FROM users;
   SELECT * FROM posts;
   SELECT * FROM categories;
   ```

## 📝 Customization

### Adding New Models
1. Create a new file in the `models/` directory
2. Define your struct with GORM tags
3. Add relationships as needed
4. Update the migration in your main function

### Adding New Repositories
1. Create a new file in the `repositories/` directory
2. Implement CRUD operations
3. Add custom query methods as needed

### Database Configuration
Modify `database/connection.go` to:
- Change database file location
- Configure logging levels
- Add connection pooling
- Switch to other databases (PostgreSQL, MySQL, etc.)

## 🔍 SQLite Database

The project uses SQLite for simplicity, but GORM supports many databases:
- **PostgreSQL**: `gorm.io/driver/postgres`
- **MySQL**: `gorm.io/driver/mysql`
- **SQL Server**: `gorm.io/driver/sqlserver`

To switch databases, update the import and connection string in `database/connection.go`.

## 📚 Additional Resources

- [GORM Documentation](https://gorm.io/docs/)
- [Go Modules](https://golang.org/ref/mod)
- [SQLite Documentation](https://www.sqlite.org/docs.html)

## 🤝 Contributing

Feel free to extend this starter project with:
- Additional models and relationships
- More complex query examples
- Web API endpoints
- Testing examples
- Docker configuration

## 📄 License

This project is open source and available under the MIT License. # go_habbit
