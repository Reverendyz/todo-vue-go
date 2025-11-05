# TODO PROJECT

- [TODO PROJECT](#todo-project)
  - [Backend service](#backend-service)
  - [Todo Web](#todo-web)

Basic todo project to test things

## Backend service

Built in [golang programming language](https://go.dev/). Exposes an application to `TODO_BACKEND_HOST` and `TODO_BACKEND_PORT`. Defaults to localhost:8080.

It connects to mongodb by using `TODO_BACKEND_DATABASE_URI`.

## Todo Web

Web frontend build with [vuejs.org](https://vuejs.org/). Set `VITE_API_URL` to let the frontend to know where to look.
