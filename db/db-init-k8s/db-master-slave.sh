#!/bin/sh

HOST=`hostname -s`
ORD=${HOST##*-}
HOST_TEMPLATE=${HOST%-*}

case $ORD in
  0)
  echo "host    replication     all     all     md5" >> /var/lib/postgresql/data/pg_hba.conf
  echo "archive_mode = on"  >> /etc/postgresql/postgresql.conf
  echo "archive_mode = on"  >> /etc/postgresql/postgresql.conf
  echo "archive_command = '/bin/true'"  >> /etc/postgresql/postgresql.conf
  echo "archive_timeout = 0"  >> /etc/postgresql/postgresql.conf
  echo "max_wal_senders = 8"  >> /etc/postgresql/postgresql.conf
  echo "wal_keep_segments = 32"  >> /etc/postgresql/postgresql.conf
  echo "wal_level = replica"  >> /etc/postgresql/postgresql.conf
  echo "hot_standby = on"  >> /etc/postgresql/postgresql.conf
  ;;
  *)
  pg_ctl -D /var/lib/postgresql/data/ -m fast -w stop
  rm -rf /var/lib/postgresql/data/*
  PGPASSWORD=postgres pg_basebackup -h ${HOST_TEMPLATE}-0.db-service -w -U replicator -p 5432 -D /var/lib/postgresql/data -Fp -Xs -P -R
  pg_ctl -D /var/lib/postgresql/data/ -w start
  ;;
esac

