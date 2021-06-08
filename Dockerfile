FROM golang:alpine3.13 AS build

WORKDIR /app
COPY go.mod go.sum /app/
COPY cmd /app/cmd/
RUN go mod download
RUN go build -o bin/minimal-app cmd/minimal-app/main.go


FROM alpine:3.13

WORKDIR /app
COPY --from=build /app/bin/minimal-app /app/
CMD /app/minimal-app