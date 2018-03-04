package createDockerFile

import "io/ioutil"

func New() {
	back := []byte(`FROM alpine:3.4
COPY bin /bin/binary_back
CMD chmod +x /binary_back
WORKDIR /bin
ENTRYPOINT /bin/sh -c ./binary_back
`)

	http := []byte(`FROM alpine:3.4
COPY bin /bin/binary_back
CMD chmod +x /binary_back
WORKDIR /bin
ENTRYPOINT /bin/sh -c ./binary_back
`)

	ioutil.WriteFile("./back/Dockerfile", back, 0644)
	ioutil.WriteFile("./http/Dockerfile", http, 0644)
}
