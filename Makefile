up:
	docker-compose up -d --build

stop:
	docker-compose stop

down:
	docker-compose down

restart: down up

unit_test:
	test_status=0;\
	docker-compose -f docker-compose.unit.yaml up --build -d;\
	docker-compose -f docker-compose.unit.yaml run unit_tests go test -v -count=1 ./... || test_status=$$?;\
	docker-compose -f docker-compose.unit.yaml down; echo "status="$$test_status;exit $$test_status;

dev:
	docker-compose -f docker-compose.dev.yaml up -d

dev_stop:
	docker-compose -f docker-compose.dev.yaml stop


image:
	docker build -f ./social/Dockerfile.full -t social:1.1 ./social

dump:
	docker exec master mysqldump -u root --password='123456' soc_db > dump_users.sql

restore:
	docker-compose run --rm master mysql -u root -h master --password=123456 soc_db < dump_users.sql

master:
	docker exec -it master mysql -p -u root soc_db

slave:
	docker exec -it slave mysql -p -u root soc_db
