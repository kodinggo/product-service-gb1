# Kodinggo - Product Service

## Setup

- First, you need to set the env value according to file `.env.sample`
- To do migration, you need to install `sql-migrate`.
  - Set dbConfig according to file `dbConfig.yml.sample`
  - then run migration

## Run Service

- To run http service, execute `go run main.go httpsrv`
- or `make run` (need to install `https://github.com/cortesi/modd`)
- To run grpc service, execute `make grpc`
