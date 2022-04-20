FROM postgres:14.2-alpine

COPY ./db.sql /docker-entrypoint-initdb.d/db.sql
CMD ["postgres"]