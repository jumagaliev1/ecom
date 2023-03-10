FROM golang

WORKDIR /ecom

COPY . .

RUN go mod download
RUN go build -o ecom ./cmd/api/

CMD ["./ecom"]