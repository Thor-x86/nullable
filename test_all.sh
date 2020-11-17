#!/bin/bash -e

rm -f coverage.txt

dialects=("postgres" "postgres_simple" "mysql" "mysql-legacy" "mariadb" "sqlite") # "mssql")

for dialect in "${dialects[@]}" ; do
  if [ "$GORM_DIALECT" = "" ] || [ "$GORM_DIALECT" = "${dialect}" ]
  then
    GORM_DIALECT=${dialect} go test -coverprofile=profile.out -covermode=atomic $d
    if [ -f profile.out ]; then
        cat profile.out >> coverage.txt
        rm profile.out
    fi
  fi
done
