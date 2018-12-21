FROM mariadb:10.3.11

ENV MYSQL_DATABASE parabible_test

ADD ["sql/parabible_data.sql", "/var/lib/mysql/parabible_data.sql"]

RUN /etc/init.d/mysql start
