# Task Tracker CLI

A command-line task management tool written in Go. Tasks are stored locally in a JSON file with no external dependencies.

---

## Features

- Add, update, and delete tasks
- Mark tasks as in-progress or done
- List all tasks or filter by status
- Persistent storage via a local JSON file (auto-created on first run)

---

## Installation

Clone the repository and build the binary:

```bash
git clone <repository-url>
cd task-tracker
go build -o task-cli .
```

---

## Usage

### Add a task

```bash
task-cli add "Buy groceries"
# Output: Task added successfully (ID: 1)
```

### Update a task

```bash
task-cli update 1 "Buy groceries and cook dinner"
```

### Delete a task

```bash
task-cli delete 1
```

### Mark a task as in-progress

```bash
task-cli mark-in-progress 1
```

### Mark a task as done

```bash
task-cli mark-done 1
```

### List tasks

```bash
# All tasks
task-cli list

# Only done tasks
task-cli list done

# Only pending tasks
task-cli list todo

# Only in-progress tasks
task-cli list in-progress
```

---

## Task Properties

Each task contains the following fields:

| Field         | Description                              |
|---------------|------------------------------------------|
| `id`          | Unique identifier                        |
| `description` | Short description of the task            |
| `status`      | Current status: `todo`, `in-progress`, or `done` |
| `createdAt`   | Timestamp when the task was created      |
| `updatedAt`   | Timestamp of the last update             |

---

## Storage

Tasks are saved in a `tasks.json` file in the current working directory. The file is created automatically if it does not exist.

---

## Requirements

- Go 1.18 or later
- No external libraries required


https://roadmap.sh/projects/task-tracker
