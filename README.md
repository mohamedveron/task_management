# task_management

This app is developed using go-chi https://github.com/go-chi/chi router and oapi-codegen https://github.com/deepmap/oapi-codegen to generate client for apis

# Use api.yml file in the swagger ui to see the description of the rest apis

## Setup

Must have golang installed version >= 1.12

make file consists of 4 steps: generate, test, build, run
you can run all of them 

```bash
make all
```

# Start the http server:

```bash
make run
```
### Build docker image

```bash
docker build --tag task-management .
```

If you have issue with code generation in generate step you can copy the api/api.gen.go file in repo and run go run main.go to start the server

### Notes:

1- I didn't use any database just in memory hash maps for both tasks and users

2- You have to add users first before adding tasks
