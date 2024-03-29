# GO MVC
- This a small boilerplate GoLang web application with MVC architecture built with `beego` framework

# About BEEG0
- [Beego](https://beego.me/quickstart) is an open source framework to build and develop applications in the Go way
- It is easy to use, intelligent, modular and provides high performance

## Requirements

If you are wanting to build and develop this, you will need the following items installed.

- Go version 1.12+
- Docker

# Installation steps

Follow the steps to run the application

- clone the repository

  ```
  git clone https://github.com/rsvijay2009/go-mvc.git
  ```

- Default configuration for MySQL

    ```
    Username: root
    Password: root
    ```
This can be changed by modifying the `.env` file

- Run the below commands to create and run the application inside a Docker container

  ```
  docker-compose up --build -d
  docker-compose up
  ```

- Your application will up and running in http://localhost:4040/

# Screenshots

## Login
![Login](screenshots/login.png)

## Register
![Register](screenshots/register.png)

## Profile
![Profile](screenshots/profile.png)

## Notepad
![Notepad](screenshots/notepad.png)

## References

[Beego quick start](https://beego.me/quickstart)

[Beego documentation](https://beego.me/docs/intro/)

[GOORM documentation](https://gorm.io/docs/)