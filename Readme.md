# Your Service Name
Event Management Service

## Introduction

The Event Management Service simplifies event-related tasks, offering a clean API for creating, updating, and querying events. It aims to streamline event management, making it easy for developers to integrate event functionalities into their applications.

## Features

Event creation with details such as title, start time, and end time.
Querying events based on different parameters like date, title, etc.


## Requirements

- Go (version x.x.x)
- MySQL (version x.x.x)


## Installation

```bash
https://github.com/Masum-Osman/event-management
cd your-repo
go build
./your-service
```
## API EndPoints

### Event List API `GET`
```
http://127.0.0.1:8080/v1/events?page=0
```
### Event List API Response
```
{
    "events": [
        {
            "id": 4,
            "title": "Event 4",
            "start_at": "2023-11-23 10:00:00",
            "end_at": "2023-11-03 18:00:00"
        },
        {
            "id": 3,
            "title": "Event 3",
            "start_at": "2023-12-02 08:00:00",
            "end_at": "2023-11-02 16:00:00"
        },
        {
            "id": 2,
            "title": "Event 2",
            "start_at": "2023-12-01 10:00:00",
            "end_at": "2023-11-01 18:00:00"
        },
        {
            "id": 1,
            "title": "Event 1",
            "start_at": "2023-11-30 09:00:00",
            "end_at": "2023-10-31 17:00:00"
        }
    ],
    "pagination": {
        "total": 4,
        "per_page": 10,
        "total_pages": 0,
        "current_page": 1
    }
}
```

### Event Details API `GET`
```
http://127.0.0.1:8080/v1/events/1
```
### Event Details API Response
```
{
    "id": 1,
    "title": "Event 1",
    "start_at": "2023-11-30 09:00:00",
    "end_at": "2023-10-31 17:00:00",
    "total_workshops": 2
}
```

### Workshop List API `GET`
```
http://127.0.0.1:8080/v1/workshops/1

```
### Workshop List API Response
```
{
    "id": 2,
    "title": "Workshop B",
    "start_at": "2023-10-31 14:00:00",
    "end_at": "2023-10-31 16:00:00",
    "workshops": {
        "id": 1,
        "title": "Workshop A",
        "description": "Description of Workshop A",
        "start_at": "2023-10-31 10:00:00",
        "end_at": "2023-10-31 12:00:00",
    }
}
```

### Workshop Details API `GET`
```
http://127.0.0.1:8080/v1/workshops/1

```
### Workshop Details API Response
```
{
    "id": 1,
    "title": "Workshop A",
    "description": "Description of Workshop A",
    "start_at": "2023-10-31 10:00:00",
    "end_at": "2023-10-31 12:00:00",
    "total_reservations": 2
}
```

### Reservation API `POST`
```
http://127.0.0.1:8080/v1/reservations

{
    "name" : "Mac Rio",
    "email" : "mac@gmail.com"
}
```
### Reservation API Response
```
{
    "reservation": {
        "id": 7,
        "name": "Mac Rio",
        "email": "mac@gmail.com"
    },
    "event": {
        "id": 1,
        "title": "Event 1",
        "start_at": "2023-10-31 09:00:00",
        "end_at": "2023-10-31 17:00:00"
    },
    "workshop": {
        "id": 1,
        "title": "Workshop A",
        "description": "Description of Workshop A",
        "start_at": "2023-10-31 10:00:00",
        "end_at": "2023-10-31 12:00:00"
    }
}
```