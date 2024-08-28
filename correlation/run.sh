#!/bin/sh

sed -i "s/POSTGRESQL_USER/${POSTGRESQL_USER}/g" /app/config.yml
sed -i "s/POSTGRESQL_PASSWORD/${POSTGRESQL_PASSWORD}/g" /app/config.yml
sed -i "s/POSTGRESQL_HOST/${POSTGRESQL_HOST}/g" /app/config.yml
sed -i "s/POSTGRESQL_PORT/${POSTGRESQL_PORT}/g" /app/config.yml
sed -i "s/POSTGRESQL_DATABASE/${POSTGRESQL_DATABASE}/g" /app/config.yml
sed -i "s/ELASTICSEARCH_HOST/${ELASTICSEARCH_HOST}/g" /app/config.yml
sed -i "s/ELASTICSEARCH_PORT/${ELASTICSEARCH_PORT}/g" /app/config.yml
sed -i "s/ERROR_LEVEL/${ERROR_LEVEL}/g" /app/config.yml
sed -i "s/USE_CACHE/${USE_CACHE}/g" /app/config.yml
sed -i "s/USE_SYSTEM_RULES/${USE_SYSTEM_RULES}/g" /app/config.yml
sed -i "s/APPEND_COMPLETED_ALERTS/${APPEND_COMPLETED_ALERTS}/g" /app/config.yml

cd /app && ./correlation