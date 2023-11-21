run:
	CompileDaemon -command="./golang-auth-api"

migrate:
	go run migrate/migrate.go