FROM golang:1.22.3
WORKDIR /Ascii-Art
COPY  go.mod .
COPY . .
RUN go build -o . main.go
EXPOSE 8080
CMD [ "go","run","." ]