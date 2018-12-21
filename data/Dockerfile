FROM mariadb:10.3.11

ENV MYSQL_DATABASE parabible_test
ADD ["sql/parabible_data.sql", "/docker-entrypoint-initdb.d/parabible_data.sql"]


