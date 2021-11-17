FROM golang:1.17

WORKDIR /infra-control-server
COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN go build -o main .
EXPOSE 8080

COPY dev-runscript.sh ./
CMD [ "sh", "./dev-runscript.sh" ]
