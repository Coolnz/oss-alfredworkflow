BIN = oss

build:
	rm -rf ${BIN}
	go build -o ${BIN}