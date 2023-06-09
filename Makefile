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
		--env AVBL_DBCONN_URI="root:root@tcp(avbl-data:3306)/narfs" \
		--name avbl-controller \
		availability:docker-controller

cnc-api:
	docker build . -f docker/availiability-service \
		-t availability:api-cnc --target api-cnc
	docker run --rm \
		--link avbl-data \
		--env AVBL_DBCONN_URI="root:root@tcp(avbl-data:3306)/narfs" \
		--env AVBL_API_SECRET_CNC="test" \
		--env AVBL_API_PORT_CNC="3666" \
		-p 3666:3666 \
		--name avbl-api-cnc \
		availability:api-cnc
data-api:
	docker build . -f docker/availiability-service \
		-t availability:api-data --target api-data
	docker run --rm \
		--link avbl-data \
		--env AVBL_DBCONN_URI="root:root@tcp(avbl-data:3306)/narfs" \
		--env AVBL_API_PORT_DATA="3667" \
		-p 3667:3667 \
		--name avbl-api-data \
		availability:api-data
api:
	docker build . -f docker/availiability-service \
		-t availability:api-cnc --target api-cnc
	docker build . -f docker/availiability-service \
		-t availability:api-data --target api-data
	docker run --rm -d \
		--link avbl-data \
		--env AVBL_DBCONN_URI="root:root@tcp(avbl-data:3306)/narfs" \
		--env AVBL_API_SECRET_CNC="test" \
		--env AVBL_API_PORT_CNC="3666" \
		-p 3666:3666 \
		--name avbl-api-cnc \
		availability:api-cnc
	docker run --rm -d \
		--link avbl-data \
		--env AVBL_DBCONN_URI="root:root@tcp(avbl-data:3306)/narfs" \
		--env AVBL_API_PORT_DATA="3667" \
		-p 3667:3667 \
		--name avbl-api-data \
		availability:api-data

reports:
	docker build . -f docker/availiability-service \
		-t availability:reports --target reports
	docker run --rm \
		-v ./tmp:/tmp \
		--link avbl-data \
		--env AVBL_DBCONN_URI="root:root@tcp(avbl-data:3306)/narfs" \
		--name avbl-reports \
		availability:reports
