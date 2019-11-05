FROM golang:1.12-alpine

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

WORKDIR /app

# beego
RUN go get github.com/beego/bee
RUN go get -u github.com/astaxie/beego
RUN go get github.com/astaxie/beego/context
RUN go get github.com/astaxie/beego/cache
RUN go get github.com/astaxie/beego/orm
RUN go get github.com/go-sql-driver/mysql

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o go-mvc .

EXPOSE 4040

CMD ["./go-mvc"]
