# kratos_template
Kratos Microserivce Platform Template

## TODO

* [TODO](./TODO.md)

## Kratos Framework

* [Kratos](https://go-kratos.dev/docs/)

## Directory structure

```
c1
├─ api
  └─ {service_name}
     └─ {service_type}
        └─ {version}
├─ app
  └─ {service_name}
    └─ {service_type}
      ├─ bin
      ├─ cmd
      ├─ configs
      └─ internal
        ├─ biz
        ├─ conf
        ├─ data
        ├─ server
        └─ service
    └─ migration
├─ bin
├─ db
├─ deploy
├─ dev
├─ infra
├─ pkg
├─ scripts   
└─ third_party

```
### Project Folder

* api: protobuf api folder
* app: app in this product
* bin: binary tool(including db migration tool)
* deploy: docker file and deployment configuration file
* dev:
* infra: docker compose file for infrastructure setup
* pkg: 
* scripts:
* third_party: third party library

### Application Folder

* bin: application complied binary
* cmd: golang entry point main.go
* configs: configuration files
* internal: golang source code
  * biz: business layer logic
  * conf: configuration logic
  * data: data access layer logic
  * server: server layer logic
  * service: service layer logic

## Infrastructure

Use docker-compose to start the database and relative supporting middleware

```bash
docker-compose -f infra/docker-compose.yaml --project-directory ./infra/ up -d
```


### Database

Connect to database as root(rootpwd)

```bash
mysql -uroot -h{host IP} -P 3306 -p
```

Create rgp database user

```sql
CREATE USER 'rgp'@'%' IDENTIFIED BY 'rgppwd';
ALTER USER 'rgp'@'%' IDENTIFIED WITH mysql_native_password BY 'rgppwd';
GRANT ALL PRIVILEGES ON `rgp_%`.* TO `rgp`@`%`;
```

Create rgp fms database

```sql
CREATE DATABASE IF NOT EXISTS rgp_fms_development
DEFAULT CHARACTER SET utf8mb4
COLLATE utf8mb4_0900_ai_ci
DEFAULT ENCRYPTION='N';
```