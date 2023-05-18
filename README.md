# Gin-Gorm API Template

This repository contains a RESTful API implementation using the Gin web framework, Gorm as the Object-Relational Mapping (ORM) library, and PostgreSQL as the underlying database.

The API serves as a foundation for building web applications and services with robust routing, database connectivity, and query management. It provides a structured architecture to handle HTTP requests, perform CRUD (Create, Read, Update, Delete) operations on the database, and return appropriate responses.

## Getting started

#### Clone the repository

```bash
git clone https://github.com/Fabrizio-rar/Gin-Gorm.git
```

#### Set up your work environment

- Ensure you have Go installed in your machine, for more information access https://go.dev/
- Create your .env file using the provided example, add the desired port in which the application will run and the URI for the PostgreSQL database

#### Run the application

Run the run.sh script to start the application


## API Reference

#### Create an user 

```http
  POST /create_user
```
Request body: 
```json
{
    "Name": "Name",
    "Gender": "Gender",
    "Email": "Email",
    "Password": "Password"
}
```

Creates a new user

#### Get all users

```http
  GET /get_all_users
```

Returns all of the users stored in the database.

#### Get user

```http
  GET /get_user/${email}
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `email`   | `string` | **Required**. User email |

Returns a user by its email

#### Delete user

```http
  POST /delete_user
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `email`   | `string` | **Required**. User email |

Deletes a user and all its entries by its email

#### Create an entry

```http
  POST /create_entry
```
Request body: 
```json
{
    "Email": "Email",
    "Title": "Title",
    "Content": "Content"
}
```

Creates an entry associated with a users email

#### Get an entry

```http
  GET /get_entry?email=example@email.com&title=ExampleTitle
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `email`   | `string` | **Required**. User email |
| `title`   | `string` | **Required**. Title of the entry |

Returns the entry with the specific title from a user

#### Get all entries from a user

```http
  GET /get_user_entries/${email}
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `email`   | `string` | **Required**. User email |

Returns all the entries from a user

#### Delete an entry

```http
  POST /delete_entry?email=example@email.com&title=ExampleTitle
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `email`   | `string` | **Required**. User email |
| `title`   | `string` | **Required**. Title of the entry |

Deletes an entry with the specific title from a user

#### Update an entry

```http
  POST /update_entry
```

Request body: 
```json
{
    "Email": "Email",
    "Title": "Title",
    "Content": "Content"
}
```

Updates an entry's content from a specific user and title




