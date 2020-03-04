# get-ready-GO

This project is a study case of [go] on how to build an example of an implementation of a service in a microservices architecture.
It is based of the [GORSK] project (GO(lang) Restful Starter Kit), which uses the following technologies:
    1. Echo - HTTP 'framework'
    2. Go-Pg - [PostgreSQL] ORM
    3. JWT-Go - [JWT] Authentication
    4. Zerolog - Structured logging
    5. Bcrypt - Password hashing
    6. Yaml - Unmarshalling YAML config file
    7. Validator - Request validation
    8. lib/pq - [PostgreSQL] driver
    9. zxcvbn-go - Password strength checker
    10. DockerTest - Testing database queries
    11. Testify/Assert - Asserting test results
    12. [go-swagger] - Documentation and testing

I updated the [go-swagger] static files, replaced go dependency manager [dep] for the most recent [go.mod].

### Install dependencies:
```bash
go get -u -f ./...
```

### Run:
Follow the instruction given [here](https://github.com/ribice/gorsk#getting-started).
 > Changed the [PostgreSQL] connection to `postgres://postgres:newPassword@localhost:5432/biadpozi`. The name of the database is the same.

### Debug:
Just run the `Launch` configuration in [Visual Studio Code] defined in the `launch.json` file of the `.vscode` folder.

[//]: # (These are reference links)

[dep]: <https://golang.github.io/dep/>
[go]: <https://golang.org/>
[gomock]: <https://github.com/golang/mock>
[GORSK]: <https://github.com/ribice/gorsk>
[go.mod]: <https://blog.golang.org/using-go-modules>
[go-swagger]: <https://github.com/go-swagger/go-swagger>
[JWT]: <https://jwt.io>
[PostgreSQL]: <https://www.postgresql.org/>
[Visual Studio Code]: <https://code.visualstudio.com/>