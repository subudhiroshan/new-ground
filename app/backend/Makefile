compile:
	go build -o backend.out

clean:
	go clean
	rm -f backend.out
	docker stop $(docker ps -a --filter "ancestor=backend" -q) || true
	docker rm $(docker ps -a --filter "ancestor=backend" -q) || true
	docker image remove backend || true

build:
	docker build . -t backend
