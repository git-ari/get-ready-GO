# get-ready-GO

This project is a study case of [go] on how to build an example of an implementation of a service in a microservices architecture.
It also includes an example of a dockerfile for the service and the script that should be run in the container.
This project uses [go-swagger].

### Install dependencies:
```bash
go get github.com/git-ari/get-ready-GO 
```

### Run:
```bash
go build && ./get-ready-GO
```

### Debug:
Just run the `Launch` configuration in [Visual Studio Code] defined in the `launch.json` file of the `.vscode` folder.

[//]: # (These are reference links)

[go]: <https://golang.org/>
[gomock]: <https://github.com/golang/mock>
[go-swagger]: <https://github.com/go-swagger/go-swagger>
[Visual Studio Code]: <https://code.visualstudio.com/>