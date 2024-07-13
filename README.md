# discord-rss-go
```
make setup

export DB_USER=<DB_USER>
export DB_PASSWORD=<DB_PASSWORD>
export DB_NAME=<DB_NAME>
export BOT_TOKEN=<BOT_TOKEN>

docker-compose up db -d

DB_HOST=localhost DB_PORT=5432 make migrate-up

docker-compose up app -d
```