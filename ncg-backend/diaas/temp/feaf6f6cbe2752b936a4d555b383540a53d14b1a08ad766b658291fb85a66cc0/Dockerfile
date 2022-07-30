
FROM gcc:4.9

WORKDIR /app
COPY . .
RUN gcc -o /app/main /app/*.c
RUN chown /app/main

CMD ["/app/main"]
