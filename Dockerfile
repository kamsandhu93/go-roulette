ARG GO_VERSION=1.17.5

FROM golang:${GO_VERSION}-alpine as Builder

ENV GOPATH=$HOME/go
ENV PATH=$PATH:$GOPATH/bin

RUN apk update && \
    apk add make

WORKDIR /root

COPY . /root
RUN make install && make build

FROM golang:${GO_VERSION}-alpine as main
COPY --from=Builder  /root/go-roulette  /root/go-roulette

ENV PORT=8080
ENV HOST=0.0.0.0

CMD  ["/root/go-roulette"]
EXPOSE 8080/TCP

#RUN apk update && \
#    apk add make docker-compose
#
#RUN  mkdir -p ~/.docker && \
#  echo > ~/.docker/config.json '{ "credsStore": "ecr-login" }' && \
#  mkdir -p ~/.aws && \
#  echo -en > ~/.aws/config '[default]\nregion = eu-west-2\noutput = json\n'
#
#ENV PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin
#ENTRYPOINT ["/bin/sh", "-c"]
