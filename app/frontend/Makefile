compile:
	go build -o frontend.out

clean:
	go clean
	rm -f frontend.out
	docker stop $(docker ps -a --filter "ancestor=backend" -q) || true
	docker rm $(docker ps -a --filter "ancestor=backend" -q) || true
	docker image remove frontend || true

build:
	docker build . -t frontend
