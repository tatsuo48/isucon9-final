#!/bin/bash

cd /home/isucon/isutrain
export LANGUAGE=go
# デプロイ作業
docker-compose -f webapp/docker-compose.yml -f webapp/docker-compose.${LANGUAGE}.yml stop
docker-compose -f webapp/docker-compose.yml -f webapp/docker-compose.${LANGUAGE}.yml build
docker-compose -f webapp/docker-compose.yml -f webapp/docker-compose.${LANGUAGE}.yml up -d

# MySQLインデックス貼り
sleep 10
mysql -h 0.0.0.0 -P 13306 -ppassword < webapp/sql/100_index.sql
