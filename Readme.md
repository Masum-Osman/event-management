# Service Name
Event Management Service

## Introduction

Welcome to the Event Management Service README! This service is designed to help you efficiently manage various events and workshops, along with reservations for those events. Whether you're organizing conferences, seminars, workshops, or any type of event, this service provides the tools you need to streamline the process.

## Table of Contents

- [Installation](#installation)
- [Getting Started](#getting-started)
- [Usage](#usage)
- [Contributing](#contributing)
- [License](#license)



## Requirements

- Go (version 1.20.3)
- MySQL (version 8.0.35)


## Installation

```bash
git clone https://github.com/Masum-Osman/event-management
```

<!-- If you are using `Docker`:

```
cd event_management
docker-compose up -d
```

> [!WARNING]
> Critical content demanding immediate user attention due to potential risks. -->

## Getting Started

```
cd event_management
go build
./event_management
```

`MySQL` dump of dummy data has been added to `/sql_dumps` dir. You can import the data using the following command:
```
cd sql_dumps
mysql -u username -p database_name < Database.sql
```

## API EndPoints

> 1. `.sql` file has been added on `/sql_dumps/Database.sql`
> 2. Postman collection has been added in root dir named `Event Management API Collection.postman_collection.json`

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

## Area of Improvement
- TDD
- Full Usage of ORM
- Had a issue with Docker MySQL. So the project went without dockerization. Though the required docker file has added.