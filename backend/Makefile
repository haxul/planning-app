migrateup:
	docker run -v /home/haxul/Development/workspaces/go/src/github.com/haxul/planning-app/backend/migrations:/migrations --network host migrate/migrate -path /migrations -database "postgresql://haxul:test@localhost:5432/planning_db?sslmode=disable" -verbose up
