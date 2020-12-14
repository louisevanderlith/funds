FROM scratch

COPY cmd/cmd .

EXPOSE 8092

ENTRYPOINT [ "./cmd" ]