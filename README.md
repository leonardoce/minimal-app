# Minimal web app

This is a minimal web application, useful during demos & trainings to
understand how environment variables work and how to expose a web application
into the Internet.

If you find it useful, enjoy!

## Environment variables

* `ENVIRONMENT` contain the environment name that is simply shown in the root
  page of the application
* `COLOR` is used as the color of the text
* `DATABASE_URL` is the connection string to a PostgreSQL database
* `QUERY` is the SQL query used
* `LISTEN_ADDRESS` is the endpoint where we should listen for HTTP requests

## Endpoints

* `/` show a simple page with the desired color and environment
* `/.well-known/check` check if we can connect to PostgreSQL
* `/tx` execute the query in `QUERY`
* `/metrics` Prometheus metrics endpoint
