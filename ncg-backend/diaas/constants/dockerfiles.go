package constants

const goDockerfile = `
FROM golang:alpine

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

ENTRYPOINT ["app"]
`

const swiftDockerfile = `
FROM swift:latest

WORKDIR /app
COPY . .
RUN swiftc -o /app/main /app/*.swift

ENTRYPOINT ["/app/main"]
`

const jsDockerfile = `
FROM node:8-alpine

WORKDIR /app
COPY . .

ENTRYPOINT ["node", "/app/main.js"]
`

const cDockerfile = `
FROM gcc:4.9

WORKDIR /app
COPY . .
RUN gcc -o /app/main /app/*.c

ENTRYPOINT ["/app/main"]

`

const cppDockerfile = `
# Get the base Ubuntu image from Docker Hub
FROM ubuntu:latest

WORKDIR /

# Update apps on the base image
RUN apt-get -y update && apt-get install -y

# Install the Clang compiler
RUN apt-get -y install clang

COPY . .

# Use Clang to compile the Test.cpp source file
RUN clang++ -o C /*.cpp

# Run the output program from the previous step
ENTRYPOINT ["./C"]
`

var (
	Dockerfiles = map[string][]byte{
		"go":    []byte(goDockerfile),
		"swift": []byte(swiftDockerfile),
		"js":    []byte(jsDockerfile),
		"c":     []byte(cDockerfile),
		"cpp":   []byte(cppDockerfile),
	}
)
