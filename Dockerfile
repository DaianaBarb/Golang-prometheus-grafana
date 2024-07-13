 FROM golang:1.22

# WORKDIR /app
# COPY /app .
# ARG USERNAME=root
# USER $USERNAME
# RUN go build main.go
# EXPOSE 8080
# ENTRYPOINT ["./main"]

WORKDIR /app
COPY /app ./
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-gs-ping
EXPOSE 8080

# Run
CMD ["/docker-gs-ping"]

