# CMPT 354 Project - Simon Fraser University - Spring 2021 - Group 1

## Instructions

### Initialization

To initialize the DB schema, first create the DB if it does not already exist:

```
psql -U <username> -c "CREATE DATABASE <db_name>;"
```

For a production environment with a blank slate, navigate to `server/sql`, and run the following to initialize the schema:
```
psql -U <username> -d <db_name> -f init_db.sql
```

For a testing environment, with some predefined values already placed in, make sure `test_run` is set to `true` in your `config.yml`, then run `go build`, and then run the resulting executable.

You can then run the following to verify that the database initialized correctly:

```
psql -U <username> -d <db_name> -c \d
```
