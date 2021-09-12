# Quick DnD
A planner for tracking Dungeons & Dragons characters, campaigns, spells, items, and more.

## Initialization

### Database

* **Prerequisite**: PostgreSQL is installed and running on localhost:5432.
* **Prerequisite**: `psql` is installed and added to `PATH`.

Create a fresh database:

```sh
psql -U postgres -h localhost -c "CREATE DATABASE <db_name>;"
```

Create the tables in the database:

```sh
psql -U postgres -h localhost -d <db_name> -f ./server/sql/init_db.sql
```

You can then run the following to verify that the database initialized correctly:

```sh
psql -U postgres -h localhost -d <db_name> -c "\d"
```

### Server code

* **Prerequisite: Go version 1.15 or later is installed and added to `PATH`.**

* Create a new `config.yml` file inside the `server` directory. The required fields can be found in `server/sample_config.yml`

### Frontend code

* **Prerequisite: `node` and `npm` are installed and added to `PATH`.**

* Follow the instructions in `client/README.md`.
