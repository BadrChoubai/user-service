FROM golang:latest AS api

WORKDIR /compiler

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o user-service ./cmd/user-service/main.go

FROM scratch AS prod

WORKDIR /production

COPY --from=api /compiler/user-service .

EXPOSE 8080
CMD ["./user-service"]

