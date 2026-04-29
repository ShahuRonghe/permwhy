# permwhy

`permwhy` is a Linux CLI that explains why a user can or cannot access a path.

This repository currently implements the first MVP in Go:

- `permwhy trace PATH`

## How it works

### trace - Path Traversal
The tool traverses the directory tree from the root to the target path, checking permissions at each level and explaining why access is granted or denied.

### Permission Checking
For each directory in the path:
- Checks if the user can traverse (execute) the directory
- Checks if the user can access the file/directory
- Explains the result with detailed reasoning

## Prerequisites

- Go 1.22 or later

## Build

```bash
go build -o permwhy ./cmd
```

## Examples

```bash
./permwhy trace /etc/hosts
```
