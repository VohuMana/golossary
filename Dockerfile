FROM alpine:3.1
MAINTAINER David James <icyflame198@gmail.com>
ADD golossary /usr/bin/golossary
ADD static/ /usr/bin/static/
WORKDIR /usr/bin/
ENTRYPOINT ["golossary"]
