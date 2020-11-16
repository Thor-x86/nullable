#!/bin/bash -e

dialects=("postgres" "postgres_simple" "mysql" "mysql-legacy" "mariadb" "sqlite") # "mssql")

for dialect in "${dialects[@]}" ; do
  if [ "$GORM_DIALECT" = "" ] || [ "$GORM_DIALECT" = "${dialect}" ]
  then
    GORM_DIALECT=${dialect} go test  --tags "json1"
  fi
done
