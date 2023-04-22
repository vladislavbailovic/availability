protobuf:
	protoc -I=proto --go_out=. proto/*.proto
cover:
	go test ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out -o coverage.html

job-image:
	docker build . -t availability:job --target job
docker-controller-image:
		docker build . -t availability:docker-controller --target docker-controller

run-docker-controller: docker-controller-image job-image
	docker run --rm \
		-v /var/run/docker.sock:/var/run/docker.sock \
		--name docker-controller \
		availability:docker-controller
