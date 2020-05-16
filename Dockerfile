FROM marceloagmelo/golang-1.13 AS builder

LABEL maintainer="Marcelo Melo marceloagmelo@gmail.com"

USER root

ENV APP_HOME /go/src/github.com/marceloagmelo/go-restore-openshift

ADD . $APP_HOME

WORKDIR $APP_HOME

RUN go mod init && \
    go install && \
    #go get ./... && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o go-restore-openshift && \
    rm -Rf /tmp/* && rm -Rf /var/tmp/*

###
# IMAGEM FINAL
###
FROM centos:7

USER root

ENV GID 23550
ENV UID 23550
ENV USER golang

ENV APP_HOME /go/bin

RUN mkdir -p $APP_HOME && \
    groupadd --gid $GID $USER && useradd --uid $UID -m -g $USER $USER && \
    chown -fR $USER:0 $APP_HOME

COPY Dockerfile $APP_HOME/Dockerfile
WORKDIR $APP_HOME

COPY --from=builder $APP_HOME/go-restore-openshift $APP_HOME/go-restore-openshift
COPY docker-container-start.sh $APP_HOME
COPY views $APP_HOME/views
COPY static $APP_HOME/static
COPY recursosValidos.json $APP_HOME/recursos.json

ENV PATH $APP_HOME:$PATH

EXPOSE 8080

USER ${USER}

ENTRYPOINT [ "./docker-container-start.sh" ]
CMD [ "go-restore-openshift" ]
