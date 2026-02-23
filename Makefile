.PHONY: start_service down_service create_table delete_table check check_table

start_service:
	sudo docker compose up -d --build

down_service:
	sudo docker compose down

create_table:
	sudo docker exec -i go-url-shortener-db-1 psql -U url-shortener -d urldb < migrations/0001_create_table.sql

delete_table:
	sudo docker exec -i go-url-shortener-db-1 psql -U url-shortener -d urldb < migrations/0001_delete_table.sql

check:
	sudo docker ps -a

check_table:
	sudo docker exec -i go-url-shortener-db-1 psql -U url-shortener -d urldb < migrations/0001_check_table.sql