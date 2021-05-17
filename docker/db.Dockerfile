FROM postgres:9.4

COPY ./docker/custom.cnf /etc/psql/conf.d/custom.cnf