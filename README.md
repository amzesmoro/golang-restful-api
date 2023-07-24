## Simple Golang API
Created with Golang without Framework. This repo contains tests, api specifications.
This repo created with Repository Pattern.

### Example Migration Commands
**Create Migration**
```
migrate create -ext sql -dir db/migrations create_table_categories
```
**Up Migration**
* up all migration
```
migrate -database "postgres://postgres:@localhost:5432/golang_restful_api_migration?sslmode=disable" -path db/migrations up
```
* specific up migration
```
migrate -database "postgres://postgres:@localhost:5432/golang_restful_api_migration?sslmode=disable" -path db/migrations up <number>
```
**Down Migration**
* down all migration (imo, this command is dangerous if app has been complex)
```
migrate -database "postgres://postgres:@localhost:5432/golang_restful_api_migration?sslmode=disable" -path db/migrations down
```
* specific down migration
```
migrate -database "postgres://postgres:@localhost:5432/golang_restful_api_migration?sslmode=disable" -path db/migrations down <number>
```
**Version**
* force change version to valid version (dirty=false)
```
migrate -database "postgres://postgres:@localhost:5432/golang_restful_api_migration?sslmode=disable" -path db/migrations force <last_valid_version_migration>
```
* to know current version
```
migrate -database "postgres://postgres:@localhost:5432/golang_restful_api_migration?sslmode=disable" -path db/migrations version
```
