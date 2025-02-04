# Note Taking Application

This project is a simple note-taking application developed using the Go programming language. The application allows users to create, view, update, and delete notes. It is built using a RESTful API architecture.

## Features

- Create notes
- View notes
- Update notes
- Delete notes

## Requirements

- Go 1.16 or newer

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/your_username/not-app.git
   cd not-app

2. Install the necessary modules:

   ```bash
   go mod tidy

3. Running the Application

   ```bash
   go run main.go

## API Usage

1. Create a Note (POST)
   To create a new note, use the following request:

   ```bash
   curl -X POST http://localhost:8080/notes -d '{"title": "Note 1", "content": "This is a note."}' -H "Content-Type: application/json"

2. View All Notes (GET)
   To view all notes, use the following request:

   ```bash
   curl -X GET http://localhost:8080/notes

3. View a Specific Note (GET)
   To view a specific note, use the note's ID in the following request:

   ```bash
   curl -X GET http://localhost:8080/notes/1

   

   
