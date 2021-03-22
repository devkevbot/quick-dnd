# CMPT 354 Project - Simon Fraser University - Spring 2021 - Group 1

## Initialization

### Git

* **Prerequisite**: `git` is installed and added to `PATH`.

From this directory, run `git config core.autocrlf false`

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

### Go code

* **Prerequisite: Go version 1.15 or later is installed and added to
  `PATH`.**

Create a new `config.yml` file inside the `server` directory:

* Refer to `server/sample_config.yml` for the required fields.

* For generating the `jwt_signing_key` field `node` can
  be used if installed:

 ```node
    > node
    > crypto.randomBytes(64).toString('hex');
 ```

### Website code

* **Prerequisite: `node` and `npm` are installed and added to `PATH`.**

Navigate to the `client/` folder and run `npm install`.

Follow the instructions in `client/README.md`.

## Testing

* Make sure `test_run` is set to `true` in `server/config.yml`

* Run `go run .` from the `server/` directory.
