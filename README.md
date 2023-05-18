
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




