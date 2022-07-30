
FROM gcc:4.9

WORKDIR /app
COPY . .
RUN gcc -o /app/main /app/*.c

CMD ["/app/main"]
