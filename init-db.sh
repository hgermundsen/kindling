#!/bin/bash

# Spin up the postgres container & let the postgres warm up
docker-compose up -d
sleep 5

# Create "message" table w/ psql inside container
docker exec kindling_postgres psql -c 'CREATE TABLE message (id serial PRIMARY KEY, title VARCHAR(60) NOT NULL, content TEXT NOT NULL, upvotes INTEGER, downvotes INTEGER, flags INTEGER, creation_time TIMESTAMP);' -U postgres

# Get out of postgres container & stop it
docker-compose stop
