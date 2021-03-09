# CMPT 354 Project - Simon Fraser University - Spring 2021 - Group 1

## Instructions

### Initialization

To initialize the DB schema, navigate the the `server/sql` folder and run the following:

```
psql -U <username> -f init_db.sql
```

You can then run the following to verify that the database initialized correctly:

```
psql -U <username> -d dnd_db -c \d
```
