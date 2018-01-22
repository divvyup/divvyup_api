FROM golang:1.9.2-alpine3.7
# Need to install git
RUN apk update
RUN apk add git sqlite gcc musl-dev

RUN mkdir -p /go/src/github.com/domtheporcupine/divvyup_api
WORKDIR /go/src/github.com/domtheporcupine/divvyup_api

# For now lets use volumes

COPY api/ ./api
COPY config/ ./config
COPY models/ ./models
COPY db/ ./db
COPY app.go .
COPY schema.dev.sql .

RUN go get golang.org/x/crypto/bcrypt
RUN go get github.com/fatih/color
RUN go get github.com/gorilla/mux
RUN go get github.com/mattn/go-sqlite3
RUN go get github.com/dgrijalva/jwt-go

ENV DIVVYUP_API_MODE=demo
ENV DIVVYUP_HOST=http://demo.divvyup.doms.land

CMD [ "go", "run", "app.go" ]