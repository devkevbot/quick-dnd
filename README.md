# Quick DnD
A planner for tracking Dungeons & Dragons characters, campaigns, spells, items, and more.

## Sample Screenshots

**Note: these screenshots do not cover all available features.**

### Login

![001-login-screen](https://user-images.githubusercontent.com/31908183/133004782-31a77ca6-7b8b-41f2-b680-25c2ef80c6bb.png)

### Dashboard

![002-dashboard](https://user-images.githubusercontent.com/31908183/133004787-3c5ed14e-a6e4-444c-bfe5-0c7fb93f9e29.png)


### Campaign Milestones

![003-campaign-milestones](https://user-images.githubusercontent.com/31908183/133004791-bbd24aed-0b39-472b-8e9f-f75bef09e616.png)

### Character Screen

![004-character-screenpng](https://user-images.githubusercontent.com/31908183/133004796-111803b3-1129-4d4a-932f-8ab1513ed434.png)

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
