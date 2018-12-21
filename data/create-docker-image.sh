#!/bin/sh

wget https://storage.googleapis.com/pb-sql/parabible_data.sql -O sql/parabible_data.sql
docker build -t pb/pb:dev .