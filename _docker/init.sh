#!/usr/bin/env bash

#NOTE: Even if this script fails, a container is NOT terminated; the remaining startup process just proceeds.

set -ex
cd src/

# create init.sql ファイルを実行
mysql -u root --password=abc123 --default-character-set=utf8mb4 go_basic_db < internal/database/init/init.sql

#We don't (can't) use a migration tool as the SQL files are placed in multiple directories.
for f in $(find internal/database/migration -name '*_table.up.sql' | sort); do
    mysql -u root --password=abc123 --default-character-set=utf8mb4 go_basic_db < $f
done