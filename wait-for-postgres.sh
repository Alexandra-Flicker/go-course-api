#!/bin/sh
set -e

host="$1"
shift
cmd="$@"

until PGPASSWORD=$DB_PASSWORD psql -h "$host" -U "$DB_USER" -d "$DB_NAME" -c '\q' >/dev/null 2>&1; do
  echo "⏳ Waiting for PostgreSQL ($DB_PASSWORD, $DB_USER, $DB_NAME) at $host..."
  sleep 1
done

echo "✅ PostgreSQL is ready"
exec $cmd
