# Starting point for docker image
# Building the go application

FROM golang AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o golang_api

# Creating Docker image from Go app we built previously
FROM debian
WORKDIR /app
COPY --from=builder /app/golang_api .
RUN chmod +x /app/golang_api
EXPOSE 8080
CMD ["./golang_api"]


