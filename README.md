# Golang API Course - Final Project (@ Estiam)

This is an API that needs fixing.

You have been studying how to write web servers in golang.

You are tasked to fixing this API.

## 1. What's wrong with the API?

This API written in Golang utilizes the Echo Framework. It has two endpoints

> GET /live

Checks whether the API is up and running.

> GET /users/:id

Returns the full information about a user: including username and password :warning:

I bet you already know what's wrong, right? How can an API return an User's whole info?

This project will require the student to **fill up missing code**  and to **remove insecure endpoints** from the webapp.

### 2. Project requirements (8 points)

We can divide the assignment into two parts.

1. (1 point) The endpoint `GET /users/:id` should be removed. It is not a safe endpoint and you would be laid off for implementing that. It is there in order to show you how to connect to the database and perform SQL queries on it. However, it cannot stay.

2. (7 points) The endpoint `POST /users` that accept a body conforms this `JSON` should be implemented:

```
POST /users
```

```json
{
    "username": "string",
    "password": "string"
}
```

You will need to write a **service method to create a new user**:
* Create new unique UUID to store in ID using the UUID library;
* Hash the password using bcrypt package;
    - Load bcrypt secret on `.env` file. See how app's configuration was done.
    - Save the hash, not the original password.

:warning: This endpoint is worth 7 points, totalizing 8 points for the whole assignment.

### 3. Extra: Login (2 points)

Students that implement an extra `POST /login` endpoint will be awarded 2 points.

```
POST /login
```
Body:
```json
{
    "username": "string",
    "password": "string"
}
```

Response:
```json
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MTIzNDU2Nzg5LCJuYW1lIjoiSm9zZXBoIn0.OpOSSw7e485LOP5PrzScxHb7SR6sAOMRckfFwi4rp7o
```

This endpoint should create a [JWT](https://github.com/golang-jwt/jwt) based on the user's info and return it to the user.

Some decisions
* How long [will the JWT last](https://jwt.io/)?
You can set the expiration time when creating the JWT.
* What should be returned in case the user doest not exist?
* What should be returned in case the password is wrong and generating the JWT wasn't possible?

This is an extra requirement and it's worth 2 extra points.

### 4. Assessment

:warning: **The project should be submitted as a compressed file (zip/tar/rar) on Microsoft Teams. It can be done in groups of up to 5 people. Only one submission per group.**

:warning: **The names of the group members should be inside the code in commentary on the main module.**