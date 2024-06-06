# Events

## Description

This is a backend program written in Go for managing events. Users can sign up, log in, and perform various actions for especial endpoints. For example, only users who are owners can update or delete their own events. for create an event or register an event or cancel the registraition login is required.

## Project Structure

```sh
.
├── db
│   └── db.go
├── Dockerfile
├── events.db
├── go.mod
├── go.sum
├── handlers
│   ├── events.go
│   ├── register.go
│   └── users.go
├── LICENSE
├── main.go
├── models
│   ├── event.go
│   ├── registration.go
│   └── user.go
├── README.md
└── routes
    ├── middleware.go
    └── routes.go

4 directories, 16 files
```

## Getting Started

### Prerequisites

- Docker
- or if you wanna run it locally and golang installed in your system just need internet :)

## Installation
To run this project locally, follow these steps:

Clone the repository:
   ```sh
   git clone https://github.com/younesious/go-events-backend.git
   cd go-events-backend
   go run ./main.go
   ```

Or if prefer Docker, follow these steps:
1. Build the Docker image:
```sh
docker build -t go-events-app .
```
2. Run the Docker container:
```sh
docker run -p 8080:8080 -v /path/to/local/db:/root/events.db go-events-app
```
Replace `/path/to/local/db` with the actual path where you want to store the database on your host system.

#### HTTP API Endpoints

-   **Sign Up**: Register as a new user: `POST   /signup`
-   **Log In**: Authenticate with your credentials: `POST   /login`
-   **View Events**: Retrieve a list of events: `GET    /events`
-   **View an Event**: Retrieve a special event: `GET    /events/:id`
-   **Create Event**: Add a new event: `POST   /events`
-   **Update Event**: Modify an existing event (only allowed by event owner): `PUT    /events/:id`
-   **Delete Event**: Remove an event (only allowed by event owner): `DELETE /events/:id`
-   **Register for Event**: Sign up for an event: `POST   /events/:id/register`
-   **Cancel Registration**: Withdraw from an event: `DELETE /events/:id/register`

License
-------

This project is licensed under the MIT License - see the [LICENSE](https://github.com/younesious/events/blob/master/LICENSE) file for details.

### Contributing

Feel free to contribute and I'll be happy to see you :)
