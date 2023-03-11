#!/bin/bash

command = $(wc -l < $1)
echo $command
if [$command == "start"]
then
    echo "Starting development environment"
    # result = $(docker ps -f name="localstack" --format "{{.ID}}")
    # if [[ -n "$result" ]];
    # then
    #     echo "container is running"
    # else
    #     echo "container is not running"
    #     docker-compose -f docker-compose.localstack.yml up -d
    #     sleep 10
    # fi

    docker-compose -f docker-compose.dev-env.yml up -d
elif [$command == "start"]
    echo "Stoping development environment"
    docker-compose -f docker-compose.dev-env.yml down -v --rmi local --remove-orphans

    docker-compose -f docker-compose.localstack.yml down -v --rmi local --remove-orphans
else
    echo "Invalid arguments"
    echo "Usage: dev-env start|stop"
fi