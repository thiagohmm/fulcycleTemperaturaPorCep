FROM golang:1.23 as build
WORKDIR  /app
COPY . .
RUN CGO_ENABLE=0 GOOS=linux GOARCH=amd64 go build -o cloudrun

FROM scratch
WORKDIR /app
COPY --from=build /app/cloudrun .
ENTRYPOINT [ "./cloudrun" ]
