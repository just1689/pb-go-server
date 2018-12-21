#!/bin/sh

docker stop mariadb
docker rm mariadb
docker run --name mariadb -e MYSQL_ROOT_PASSWORD=toor -e MYSQL_DATABASE=parabible_test -d -p 3306:3306 pb/pb:dev
