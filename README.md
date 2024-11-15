# CLI To-Do Application

A lightweight, command-line To-Do manager built in Go. Easily add, list, and complete tasks directly from the terminal with efficient SQL-backed storage and colored output for priority-based organization.

## Features

- **Add Tasks**: Include name, priority (low, mid, high), and optional due date.
- **List Tasks**: Filter by date, priority, and limit the number of tasks.
- **Mark as Complete**: Update task status by ID.
- **Color-Coded Output**: Quick visual prioritization.

## Installation

1. **Clone and Navigate**:
   ```bash
   git clone https://github.com/yourusername/cli-todo.git && cd cli-todo
2. **Build the Executable**:
   ```bash
   go build -o cli-todo
3. **Run the Application**:
   ```bash
   ./cli-todo

## Usage

1. **Add a Task**:
   ```bash
   ./cli-todo add "Finish report" -i high -d 2024-12-01
2. **List Tasks**:
   ```bash
   ./cli-todo list -d 2024-12-01 -i mid -n 3
3. **Mark a Task as Complete**:
   ```bash
   ./cli-todo complete 3
