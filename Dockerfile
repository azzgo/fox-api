FROM scratch

COPY ./bin/  /

EXPOSE 9999

ENTRYPOINT ["/app"]
