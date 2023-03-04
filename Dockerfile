FROM golang:1.19
WORKDIR /app
COPY geffws /app/geffws
EXPOSE 8083