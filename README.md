# Localdealer

* endpoints for the frontend
* serves database entries

REST API with :
* [Echo Framework](https://echo.labstack.com/guide/migration)
* Postgres

## Configs

Config file located in ```config.yaml``` on the root of project. 

##example config.yaml
```yaml
app_name: kiezsupport
app_version: 0.0.1
app_env: dev

dev:
  adapter: postgre
  database: kiezsupport
  username: root
  password: rootpassword
  host: localhost     #defaults to 127.0.0.1
  port: 5432
  idle_conns: 10
  open_conns: 100
  sslmode: disable
``