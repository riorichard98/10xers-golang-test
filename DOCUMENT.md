# Phone Management API Design Document

## Overview

This document provides the design and usage details for a Phone Management API that allows administrative users to manage phone inventory. The API is built with simplicity and flexibility in mind, utilizing dependency injection to facilitate testing and maintenance, and separating domain logic from HTTP handling to ensure clean code organization.

## Architecture

### Components

- **HTTP Handlers**: Serve as the entry point for API requests. They parse request data, handle HTTP-specific actions, and pass control to the domain logic.
- **Domain Logic**: Encapsulated business logic that performs operations directly related to phone data management. It communicates directly with the database to create, retrieve, update, or delete phone records.
- **Database**: Persistent storage for phone records.
- **Dependency Injection**: Used to inject database connections and other dependencies into the domain logic at runtime, which enhances flexibility and aids in error handling.
- **Middleware**: Implements basic authentication to secure the API, ensuring that only users with valid credentials (admin) can make changes.

### Error Handling

Initial error handling is integrated into the dependency injection mechanism to ensure that any issues with service configuration or dependency resolution are caught and handled gracefully.

### Security

Basic authentication is employed to protect all API endpoints. Clients must provide a valid admin username and password encoded in Base64 format within the `Authorization` header.

## API Endpoints

### 1. Create a New Phone

### Endpoint Information

- **Endpoint:** `/phones`
- **HTTP Method:** `POST`
- **Authentication:** Basic Auth

### Request Details

#### Request Body

```json
{
  "name": "string",
  "brand": "string",
  "price": "number",
  "stock": "integer"
}
```

#### success response Body

```json
"success creating new phone"
```


### 2. Search phone by name

### Endpoint Information

- **Endpoint:** `/phones`
- **HTTP Method:** `GET`
- **Authentication:** Basic Auth

### Request Details

#### Request Query Param

```json
{
    "name":"string"
}
```

#### example success response Body

```json
[
    {
        "id": "8122bd86-a6b1-4b53-bbb4-efde4fc26bdf",
        "ID": 0,
        "CreatedAt": "2024-05-01T12:05:46.668292476+07:00",
        "UpdatedAt": "2024-05-01T12:05:46.668292476+07:00",
        "DeletedAt": null,
        "name": "iphone15pro",
        "brand": "apple",
        "price": 100.99,
        "stock": 100
    }
]
```


### 3. Update phone by id

### Endpoint Information

- **Endpoint:** `/phones/:phoneId`
- **HTTP Method:** `PATCH`
- **Authentication:** Basic Auth

### Request Details

#### Request Body

number of properties is optional in the request body

```json
{
  "name": "string",
  "brand": "string",
  "price": "number",
  "stock": "integer"
}
```

#### example success response Body

```json
"success update phone with id: 8122bd86-a6b1-4b53-bbb4-efde4fc26bdf"
```


### 4. Delete phone by id

### Endpoint Information

- **Endpoint:** `/phones/:phoneId`
- **HTTP Method:** `DELETE`
- **Authentication:** Basic Auth

### Request Details

#### example success response Body

```json
"success delete phone with id: 8122bd86-a6b1-4b53-bbb4-efde4fc26bdf"
```