
IMG_NAME ?=chaunceyjiang/sample-k8s-scheduler
IMG_VERSION ?=v0.0.12
IMG ?=${IMG_NAME}:${IMG_VERSION}

binary:
	go build -o sample-k8s-scheduler main.go

build: binary
	sed -i 's@image: .*@image: '"${IMG}"'@' deploy/sample-scheduler.yaml
	docker build -t ${IMG} .
deploy: build
	kubectl apply -f deploy/sample-scheduler.yaml
push: build
	docker push ${IMG}
release: build push deploy
