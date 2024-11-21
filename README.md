# MTG Backend Project

## Tech stack
- golang
- docker
- postgre db

## Initialize PostgreSQL DB
I am using brew to download postgreSQL

```
brew install postgresql
```

when finish download I run this command based on suggestion

```
brew services start postgresql@14
```

add user to initialize DB

```
/opt/homebrew/opt/postgresql@14/bin/createuser -s <username>
```