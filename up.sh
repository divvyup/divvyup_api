#! /bin/bash

# Update our git repo
git reset --hard HEAD
git pull

# Then find and kill the container currently running
docker stop $(docker ps | grep "divvyapi" | awk '{print $1;}')
# Rebuild it
docker build -t divvyapi .
# Finally bring it back up
docker run -d -p 3088:3088 divvyapi