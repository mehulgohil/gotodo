# gotodo
A todo task implemented in Golang

## Requirement
Add a task: Users should be able to add a new task to their to-do list along with a due date.
List tasks:
Users should be able to view a list of all the tasks in their to-do list, including the task description and due date.
Complete a task: Users should be able to mark a task as completed.
Remove a task: Users should be able to remove a task from their to-do list.
Search tasks: Users should be able to search for tasks by keywords or due date.

## Usage
Install will install the cli binary
The cli will save the todo task in the `storage.json` file
that will be created when you add the task for the first time.

```bash
go install github.com/mehulgohil/gotodo
gotodo --help
```

## Cobra Commands
```bash
cobra-cli init
cobra-cli add list
```

## Test Commands for local testing
```bash
 go run . add -n test -d test
```

