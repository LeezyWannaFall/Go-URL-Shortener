.PHONY: build_postgres build_memory down_memory down_postgres create_table delete_table check

build_postgres:
	sudo docker compose --profile with_db up -d --build

build_memory:
	sudo docker compose --profile in_memory up -d --build

down_memory:
	sudo docker compose --profile in_memory down

down_postgres:
	sudo docker compose --profile with_db down

create_table:
	sudo docker exec -i go-url-shortener-db-1 psql -U url-shortener -d urldb < migrations/0001_create_table.sql

delete_table:
	sudo docker exec -i go-url-shortener-db-1 psql -U url-shortener -d urldb < migrations/0001_delete_table.sql

check:
	sudo docker ps -a