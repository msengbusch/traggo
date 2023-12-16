FROM gcr.io/distroless/static
WORKDIR /opt/traggo
ADD build/traggo-server /opt/traggo
EXPOSE 3030
ENTRYPOINT ["./traggo-server"]