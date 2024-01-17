# Events Management Rest API in Go
This is an event management Rest api written using Gin web framerwork
It also uses sqlite3 database to store the events data

## How to Run
- Clone and cd the repo
- go run .

## Api
- GET /events : Returns all existing events stored in the database
- GET /events/:id : Returns event with given id or 404 if no event found with given id
- POST /events : Add event and store into database

## Schema of the event
- id: INTEGER, PRIMARY KEY with auto increment ID
- name TEXT NOT NULL
- description TEXT NOT NULL,
- location TEXT NOT NULL,
- dateTime DATETIME NOT NULL,
- userId INTEGER NOT NULL

## WIP Features
- Add JWT authentication for users
- DELETE /events/:id : Delete events
