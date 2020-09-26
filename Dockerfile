FROM golang:1.15 AS build

WORKDIR /src
COPY . .
RUN ls  
RUN go build -o /out/visadiscordbot api/main.go

FROM centos:latest AS bin
COPY --from=build /out/visadiscordbot /