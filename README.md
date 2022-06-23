# REST API with Go

A simple REST API written in Go that performs CRUD operations for book management. The purpose of this program is to demonstrate a clean folder structure using best practices. Gorilla/mux is used to route incoming HTTP requests to the correct method handlers.

## Todo

- [x] Original implementation entirely in main.go
- [x] Optimize directory structure
- [ ] Add documentation
- [ ] Use real database

---

## Project structure

The project layout uses the common pattern of splitting the code into the **/cmd** and **/pkg** layout patterns.  

### `/cmd`

The `/cmd` contains the main applications for the project. If you need to have more than one application binary, the name of subdirectories should match the name of the executable application (e.g., /cmd/myapp). It's best practice not to put a lot of code in the application directory.  

### `/pkg`
The library code that can be imported and used by external projects should live in `/pkg` directory. This layout pattern allows the package to be "go gettable". which means you can use the **go get** command to fetch and install the project, its applications, and libraries (e.g., `go get github.com/andrei-k/go-rest-api/pkg`). Use caution with the code placed here because external projects will expect these libraries to work.  

### `/internal`
Your code that is not meant to be reused by external projects should live in the `/internal` directory. Go ensures that these private packages aren't not importable.  

Some notable examples of using this layout pattern include the official [Go Tools](https://github.com/golang/tools), [Kubernetes](https://github.com/kubernetes/kubernetes), and [Docker](https://github.com/docker/compose).  

Good reference for the [Standard Go Project Layout](https://github.com/golang-standards/project-layout)

---

The original code that was contained entirely in the main.go file can be found in `/internal/original_main.go`.