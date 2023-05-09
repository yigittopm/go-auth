#FROM golang:1.14 as build
#WORKDIR /src
#COPY go.sum go.mod ./
#RUN go mod download
#COPY . .
#RUN CGO_ENABLED=0 go build -o /bin/app ./cmd/main.go
#
#FROM alpine
#COPY --from=build /bin/app /bin/app
#ENTRYPOINT ["/bin/app"]

FROM golang:1.20 as build
WORKDIR /app
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
COPY go.mod go.sum ./
RUN go mod download
COPY ./ ./
RUN go build main.go

FROM scratch as final
COPY --from=build /app/main ./
ENTRYPOINT ["./main"]


