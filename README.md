# Pustaka-api

An pustaka api is application for a for book information

## Setting up

- Import database MySQL

```
database_pustaka.sql
```

## Endpoints

### Root URL - `http://localhost:8888`

- This is the root url where the API is hosted. Prepend this root before all requests

### GET `/books`

Lists all books in the with name and price details

#### Response

```json
{
  "data": [
    {
      "id": 1,
      "title": "Go Programming",
      "description": "Go is a programming language that makes it easy to build simple, reliable, and efficient software.",
      "price": 100000,
      "rating": 5
    },
    {
      "id": 2,
      "title": "Cosmos v2",
      "description": "Cosmos is a book about the universe v2",
      "price": 100000,
      "rating": 5
    }
  ],
  "msg": "success",
  "status": 200
}
```

### POST `/books`

Adds a book

#### Request

```json
{
  "title": "Brief history of universe",
  "price": 200000,
  "description": "Book about universe",
  "rating": 5
}
```

### Response

```json
{
  "data": {
    "id": 9,
    "title": "Brief history of universe",
    "description": "Book about universe",
    "price": 200000,
    "rating": 5
  },
  "status": 200
}
```

### PUT `/books/:id`

Updates the details of an book

#### Request

Provide atleast one of the fields

```json
{
  "title": "Brief history of universe Ver.II",
  "price": 210000,
  "description": "Book about universe",
  "rating": 4.5
}
```

### Response

``` json
{
    "data": {
        "id": 9,
        "title": "Brief history of universe Ver.II",
        "description": "Book about universe",
        "price": 210000,
        "rating": 4.5
    },
    "status": 200
}
```

### DELETE `/books/:id`

Deletes an book

### Response

``` json
{
    "data": {
        "id": 9,
        "title": "Brief history of universe Ver.II",
        "description": "Book about universe",
        "price": 210000,
        "rating": 4.5
    },
    "status": 200
}
```

### GET `/warehouse`

Lists all warehouses

#### Response

```json
{
  "warehouses": [
    {
      "id": 1,
      "name": "RDU"
    },
    {
      "id": 2,
      "name": "JFK"
    }
  ],
  "count": 2
}
```