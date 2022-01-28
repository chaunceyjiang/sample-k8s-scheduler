
IMG_NAME ?=chaunceyjiang/sample-k8s-scheduler
IMG_VERSION ?=v0.0.1


binary:
	go build -o sample-k8s-scheduler main.go

build: binary
	docker build -t ${IMG_NAME}:${IMG_VERSION} .
deploy: build
	kubectl apply -f deploy/sample-scheduler.yaml

release: build
	docker push ${IMG_NAME}:${IMG_VERSION}