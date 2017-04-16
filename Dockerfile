FROM alpine:3.1
MAINTAINER David James <icyflame198@gmail.com>
ADD golossary /usr/bin/golossary
ADD test.txt /usr/bin/test.txt
ENTRYPOINT ["golossary"]
