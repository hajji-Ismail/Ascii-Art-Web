FROM golang:1.22.3
WORKDIR  /Ascii-Art
COPY go.mod /Ascii-Art/
RUN go mod download
COPY  . .
RUN go build -o /Ascii-Art/Ascii-Art /Ascii-Art/main.go
CMD [ "go","run","." ]

