GOCMD:=go
GORUN:=${GOCMD} run
GOBUILD:=${GOCMD} build
GOTEST:=${GOCMD} test

run:
	${GORUN} main.go
test-all:
	${GOTEST} ./internal/...
test-all-cv:
	${GOTEST} ./internal/... -v -cover
clean-cache:
	${GOCMD} clean -testcache ./...
test-no-cache: clean-cache test-all

doc:
	swag init
dev-migrate-up:
	sql-migrate up -env="development"
dev-migrate-down:
	sql-migrate down -env="development"
seed:
	${GORUN} cmd/seeder/main.go
flush:
	docker exec Mysql mysql -uroot -proot -e \
	"SELECT CONCAT('TRUNCATE TABLE ',table_schema,'.',TABLE_NAME, ';') from  information_schema.tables where table_schema in ('Planet') \
	AND table_name not like '%migrations'" | grep "Planet" | xargs -I {} docker exec Mysql mysql -uroot -proot -e {}

flush-seed: flush seed

.PHONY: doc dev-migrate-up dev-migrate-down seed-for-catalog