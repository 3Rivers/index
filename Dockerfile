FROM alpine
ADD html /html
ADD index-web /index-web
WORKDIR /
ENTRYPOINT [ "/index-web" ]
