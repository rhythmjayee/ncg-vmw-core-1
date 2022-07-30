
FROM swift:latest

WORKDIR /app
COPY . .
RUN swiftc -o /app/main /app/*.swift

ENTRYPOINT ["/app/main"]
