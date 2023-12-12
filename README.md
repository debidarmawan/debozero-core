# debozero-core
Backend Serivces for DeboZero Apps

# Migrations
1. To create migration file run the following command ```migrate create -ext sql -dir db/migrations -seq create_users_table```
2. To run migrations run the following command ```migrate -database ${POSTGRESQL_URL} -path db/migrations up```
3. To reverse migration run the following command ```migrate -database ${POSTGRESQL_URL} -path db/migrations down```