#!/bin/bash

cd "$(dirname $0)"

mysql -u root -ptaynguyen < ./dbs/reset-dbs.sql