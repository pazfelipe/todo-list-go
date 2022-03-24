# TODO LIST - GOLANG

<p align="center">
<img src="https://jollycontrarian.com/images/a/ab/Dramatic_Look_Gopher.gif" width="280"/>
</p>

## A simple project to learn golang

![golang](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![postgre](https://img.shields.io/badge/PostgreSQL-316192?style=for-the-badge&logo=postgresql&logoColor=white)
![mysql](https://img.shields.io/badge/MySQL-00000F?style=for-the-badge&logo=mysql&logoColor=white)
![mongodb](https://img.shields.io/badge/MongoDB-4EA94B?style=for-the-badge&logo=mongodb&logoColor=white)
![sqlite](https://img.shields.io/badge/SQLite-07405E?style=for-the-badge&logo=sqlite&logoColor=white)

---
![maintained](https://img.shields.io/badge/Maintained%3F-yes-green.svg)
![size](https://badge-size.herokuapp.com/pazfelipe/todo-list-go/main/main.go)
![stars](https://img.shields.io/github/stars/pazfelipe/todo-list-go.svg)

### How to run

- Install golang 1.17+
- CD to the project folder
- Run `go run main.go`

### How to test

There are 5 URLs that can be accessed:

| Method     | URL                          | Response                 |
| ---------- | ---------------------------- | ------------------------ |
| __GET__    | `http://localhost:8080/`     | returns the todo list    |
| __POST__   | `http://localhost:8080/`     | creates a new todo       |
| __GET__    | `http://localhost:8080/{id}` | returns an specific todo |
| __PUT__    | `http://localhost:8080/{id}` | updates a todo           |
| __DELETE__ | `http://localhost:8080/{id}` | deletes a todo           |

### Next steps

- [ ] Integration with Postgres
- [ ] Integration with Redis
- [ ] Integration with MongoDB
- [ ] Integration with SQLite
- [ ] Unity Testing

<p align="center">
  </br>
  </br>
  </br>
  <img src="https://i.giphy.com/media/XsGH5TII7EzGUvTN1g/giphy.webp" width="250" />
</p>