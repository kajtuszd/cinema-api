# Cinema Api
Simple Gin-gonic API for managing cinema data.

## Configuration and running

Copy and set environment variables:
```shell
$ cp .env.example .env
```

Configure the initialization file for the admin dashboard:
```shell
$ cd scripts
$ cp example-admin.sql admin.sql
```

Set your own admin user in `line 8` of the `admin.sql` file. For example:
```sql
\set db_user admin
```

Build and run the Docker containers:
```shell
$ docker-compose up
```
