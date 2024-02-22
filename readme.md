# Backend for spikeguard

## Dev setup guide
### Requirements:
- go
- MySQL

### Setup
create a .env file in the root of the project containing these values:
```dotenv
MYSQL_USER="your user"
MYSQL_PASSWORD="your password"
MYSQL_DATABASE="your database"
MYSQL_HOST=localhost
MYSQL_PORT="your database port"
PORT="your port of choice"
```

Within your mysql console run the following command:
```mysql
create database "your database";
create table markers
(
    id        bigint auto_increment
        primary key,
    latitude  decimal not null,
    longitude decimal not null
);
```

to run in development mode
```shell
go run .
```
