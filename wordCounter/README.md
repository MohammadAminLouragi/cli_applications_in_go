# Word Counter

this is a simple word counter application 

## Run Tests

Run all tests in the project:

```bash
go test -v

```
## Build Project

```bash
go mod tidy
go build -o app wordCounter/...

```

## Run Project
-l => count number of lines
-b = count number of bytes
```bash
echo "My first command line tool with go" | ./app.exe
```