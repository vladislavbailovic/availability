protobuf:
	protoc -I=proto --go_out=. proto/*.proto
cover:
	go test ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out -o coverage.html

job-image:
	docker build . -f docker/availiability-service \
		-t availability:job --target job
docker-controller-image:
		docker build . -f docker/availiability-service \
			-t availability:docker-controller --target docker-controller

data:
	-docker container stop avbl-data
	docker build . -t availability:mysql -f docker/mysql
	docker run --rm -d -p 63306:3306 --name avbl-data \
		--env MYSQL_ROOT_PASSWORD=root \
		availability:mysql

run-docker-controller: job-image docker-controller-image
	docker run --rm \
		-v /var/run/docker.sock:/var/run/docker.sock \
		--link avbl-data \
		--name avbl-controller \
		availability:docker-controller

reports:
	docker build . -f docker/availiability-service \
		-t availability:reports --target reports
	docker run --rm \
		-v ./tmp:/tmp \
		--link avbl-data \
		--name avbl-reports \
		availability:reports
