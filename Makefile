IMAGE_PREFIX=u03013112
IMAGE_NAME=$(IMAGE_PREFIX)/ss-tester
all:
	docker build -f ./Dockerfile0 -t u03013112/ss-tester-base .
	GOFLAGS=-mod=vendor CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/ss-tester -a -installsuffix cgo -ldflags '-w'
	docker build -t $(IMAGE_NAME) .
push:
	docker push u03013112/ss-tester-base
	docker push $(IMAGE_NAME)
clean:
	docker rmi  $(IMAGE_NAME)