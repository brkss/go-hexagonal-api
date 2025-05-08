#!/bin/bash 

# Define container names with suffix
POSTGRES_NAME="postgres-dex"
REDIS_NAME="redis-dex"

# Check if postgres container exists and is running
if [ ! "$(docker ps -q -f name=$POSTGRES_NAME)" ]; then
    if [ "$(docker ps -aq -f status=exited -f name=$POSTGRES_NAME)" ]; then
        # Cleanup
        docker rm $POSTGRES_NAME
    fi
    # Run postgres container
    docker run --name $POSTGRES_NAME -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -p 5432:5432 -d postgres:15-alpine
    echo "Created and started postgres container"
else
    echo "Postgres container is already running"
fi

# Check if redis container exists and is running
if [ ! "$(docker ps -q -f name=$REDIS_NAME)" ]; then
    if [ "$(docker ps -aq -f status=exited -f name=$REDIS_NAME)" ]; then
        # Cleanup
        docker rm $REDIS_NAME
    fi
    # Run redis container
    docker run --name $REDIS_NAME -p 6379:6379 -d redis:7-alpine
    echo "Created and started redis container"
else
    echo "Redis container is already running"
fi
