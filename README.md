# GO MVC
- This a small boilerplate GoLang web application with MVC architecture built with `beego` framework

# About BEEG0
- `Beego` is an open source framework to build and develop applications in the Go way
- It is easy to use, intelligent, modular and provides high performance

## Requirements

If you are wanting to build and develop this, you will need the following items installed.

- Go version 1.12+
- Docker

# Up and running

Follow the steps to run the application

- clone the repository

  ```
  git clone https://github.com/rsvijay2009/go-mvc.git
  ```

- Run the below command to build the go binary for the docker container

    ```
    cd go-mvc
    go build webapp
    ```

- Run the below commands to create and run the application inside a Docker container

  ```
  docker-compose up --build -d
  docker-compose up
  ```

- Your application will up and running in http://localhost:4040/

# Useful docker commands

# List all the containers

```
docker container ls -a
```

# Access database container

```
docker container exec -it {DB_CONTAINER_NAME} bash
```

# References

[Beego quick start](https://beego.me/quickstart)

[Beego documentation](https://beego.me/docs/intro/)

[GOORM documentation](https://gorm.io/docs/)