## install migrate cli
brew install golang-migrate
migrate -source file://path/to/migrations -database postgres://localhost:5432/database up
## create migrations
migrate create -ext sql -dir store/migrations -seq create_users_table
## docker
docker run -v \`pwd\`:/migrations --network host migrate/migrate -path=/migrations/ -database postgres://calendar:123456@localhost:5432/calendar?sslmode=disable up
## docker-compose
docker-compose run migrations -path /migrations/ -database postgres://calendar:123456@localhost:5432/calendar?sslmode=disable up
## implementation
impl 'p *PgStorage' usersservice/internal/domain/entities.RoleRepository >> role.go