# Go URL Shortener 
## With Fiber, PostgreSQL, and ORM  

This repository contains a URL shortener server implemented using **Go**, the **Fiber** web framework, **PostgreSQL** as the database, and a **GORM** (Object-Relational Mapping) library for database operations.

## Prerequisites
Before running this server, ensure that you have the following prerequisites installed on your system:

- Go (1.16 or higher)
- PostgreSQL
- Fiber (go get -u github.com/gofiber/fiber/v2)
- GORM (go get -u gorm.io/gorm)
- GORM PostgreSQL driver (go get -u gorm.io/driver/postgres)

## Installation

1. Clone the repository"
   ```bash
   git clone https://github.com/subrotokumar/go-url-shortener.git
   cd your-repo
   ```
2. Install the required dependencies:
   ```bash
   go mode tidy
   ```

3. Set up the PostgreSQL database. Replace the `dsn` with database connection detail in in `server/server.go` file.
   ```go
   dsn := "host=localhost user=postgres password=password dbname=golink port=5432 sslmode=disable"
   ```
4. Run the server:
   ```bash
   go run main.go
   ```
   The server should now be running on http://localhost:3000.

## ENDPOINTS

Sure! Here's a description of the endpoints you provided:

- **GET /link**: This endpoint is used to retrieve all redirects. It expects no parameters. When accessed, it will return a list of all redirects stored in the database.

- **GET /link/:id**: This endpoint is used to retrieve a specific redirect based on its ID. It expects a parameter id in the URL path, representing the ID of the redirect to fetch. When accessed, it will return the details of the specified redirect.

- **POST /link**: This endpoint is used to create a new redirect. It expects a JSON payload in the request body containing the details of the redirect. The payload should include the redirect URL and link value for the new redirect. Upon successful creation, it will return the details of the newly created redirect.
  
- **DELETE /link/:id**: This endpoint is used to delete an existing redirect. It expects id of the redirect, that required to be delete as url param. Upon successful deletion, it will return a success message indicating that the redirect has been deleted.


## Redirect to the Original URL
To redirect to the original URL associated with a shortened URL, simply visit the shortened URL in a web browser or make a GET request to it:

ENDPOINT: **GET /r/:redirect** 

```bash
curl http://localhost:3000/r/:redirect
```

## Database Schema
The server uses the following database schema: 
```sql
CREATE TABLE links (
  id BIGSERIAL PRIMARY KEY,
  redirect TEXT NOT NULL,
  link TEXT NOT NULL UNIQUE,
  clicked BIGINT,
  random BOOLEAN
);
```



In this schema, the table name is links, and it has the following columns:

- `id` (BIGSERIAL): A unique identifier for each link.
- `redirect` (TEXT): The URL to redirect to when the shortened link is - accessed.
- `link` (TEXT): The shortened link itself, which should be unique.
- `clicked` (BIGINT): The number of times the shortened link has been clicked.
- `random` (BOOLEAN): A flag indicating whether the link is randomly generated or not.
  
Note that the **PRIMARY KEY** constraint is applied to the id column, and the UNIQUE constraint is applied to the **link** column to ensure uniqueness.

You can execute this SQL statement using a PostgreSQL client or incorporate it into your database migration scripts or tools.

## License
This project is licensed under the MIT License. See the LICENSE file for more information.




