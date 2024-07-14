#!/bin/bash

docker run --name postgres-test -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -p 6432:5432 -d postgres:latest

# Wait for the database to start
echo "Waiting for the PostreSQL to start..."
sleep 3

# Create the database and the todo table
docker exec -it postgres-test psql -U postgres -d postgres -c "CREATE DATABASE todoapp"
sleep 3
echo "Database created successfully"

docker exec -it postgres-test psql -U postgres -d todoapp -c "
create table if not exists todos
(
    id bigserial not null primary key,
    title varchar(255) not null,
    description varchar(255) not null,
    is_done boolean not null default false,
    created_at timestamp not null default current_timestamp,
    updated_at timestamp not null default current_timestamp,
    due_date timestamp not null
);"

sleep 3
echo "Product Table created successfully"