FROM scratch

ADD api-server /

CMD mkdir config

COPY config/api-server.yaml config/

ENTRYPOINT ["/api-server"]