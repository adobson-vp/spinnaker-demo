FROM golang
WORKDIR /opt/go-app
ADD main.go /opt/go-app
ADD app /opt/go-app/app
ENTRYPOINT  ["go", "run", "main.go"]
