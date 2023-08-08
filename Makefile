compile:
	make -C app/backend compile
	make -C app/frontend compile

build:
	make -C app/backend build
	make -C app/frontend build

clean:
	make -C app/backend clean
	make -C app/frontend clean
	docker compose down --rmi local

up:
	docker compose up -d --build

down:
	docker compose down

reset:
	curl localhost:8080/zero

test: up
	python tests/e2e.py > test_output.json
	cat test_output.json | jq > test_output_formatted.json
