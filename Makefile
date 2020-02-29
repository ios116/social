up-p:
	docker-compose -f docker-compose.prod.yaml up -d --build

stop-p:
	docker-compose -f docker-compose.prod.yaml stop

down-p:
	docker-compose down

restart: down up

unit:
	test_status=0;\
	docker-compose -f deployments/docker-compose.unit.yaml up -d ;\
	docker-compose -f deployments/docker-compose.unit.yaml run unit_tests go test -v -count=1 ./... || test_status=$$?;\
	docker-compose -f deployments/docker-compose.unit.yaml down; echo "status="$$test_status;exit $$test_status;

up:
	docker-compose up -d

rt: 
	docker-compose restart tarantool

r: 
	docker-compose restart 

rw: 
	docker-compose restart web.ru

lw:
	docker-compose logs -f --tail=10 web.ru

lt:
	docker-compose logs -f --tail=10 tarantool


stop:
	docker-compose stop

image_multistage:
	docker build -f ./social/Dockerfile.multistage -t social:1.1 ./social

dump:
	docker exec master mysqldump -u root --password='123456' --single-transaction --set-gtid-purged=OFF soc_db > dump_users.sql

restore:
	docker-compose run --rm master mysql -u root -h master --password=123456 soc_db < dump_users.sql

master:
	docker exec -it master mysql -p -u root soc_db

slave:
	docker exec -it slave mysql -p -u root soc_db

slave2:
	docker exec -it slave2 mysql -p -u root soc_db
