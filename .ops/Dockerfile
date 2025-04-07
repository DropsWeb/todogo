FROM golang:1.23.0

RUN curl -fsSL \
    https://raw.githubusercontent.com/pressly/goose/master/install.sh |\
    sh

COPY . /builds

WORKDIR /builds/cmd/app

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build  -o /goapi

CMD ["/goapi"]
