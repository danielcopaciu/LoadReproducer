FROM alpine:3.3

ADD *.go /load-reproducer/

RUN apk add --update bash \
  && apk --update add git bzr \
  && apk --update add go \
  && export GOPATH=/gopath \
  && REPO_PATH="github.com/danielcopaciu/LoadReproducer" \
  && mkdir -p $GOPATH/src/${REPO_PATH} \
  && mv /load-reproducer/* $GOPATH/src/${REPO_PATH} \
  && cd $GOPATH/src/${REPO_PATH} \
  && go get -t ./... \
  && go build \
  && mv LoadReproducer /load-reproducer-app \
  && rm -rf /load-reproducer \
  && apk del go git bzr \
  && rm -rf $GOPATH /var/cache/apk/*

CMD [ "/load-reproducer-app" ]