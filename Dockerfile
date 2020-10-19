FROM golang:1.15 AS build

WORKDIR /src
COPY . .
RUN ls  
RUN go build -o /out/visadiscordbot api/main.go

FROM centos:7 AS bin
ENV VISA_BOT_CONFIG=/config/config.yaml
COPY --from=build /src/api/config.yaml /config/
COPY --from=build /out/visadiscordbot /
EXPOSE 80
ENTRYPOINT [ "./visadiscordbot" ] 