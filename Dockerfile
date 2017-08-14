# Version: 0.1.0

# https://hub.docker.com/_/golang/
# sudo docker build -t my-golang-app .

FROM golang:1.8.3

WORKDIR /go/src/app

COPY . .

CMD ["go-wrapper", "download github.com/gorilla/mux"]

CMD ["go-wrapper", "build main.go app.go model.go"]
