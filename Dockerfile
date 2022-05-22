FROM golang:1.18 as build

WORKDIR /app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading
# them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN CGO_ENABLED=0 go build -v -o main .

FROM gcr.io/distroless/static

WORKDIR /

COPY --from=build /app/main main

CMD ["/main"]
