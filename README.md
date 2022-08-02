# Gobserver

This repository contains an sample application to start using Gorm and other useful libraries.

The main purpose is learn ;)


## Development

### MacOS

#### Requirements

- Docker
- Brew

#### DDBB

- MySQL
```
brew install mysql-client
echo 'export PATH="${PATH}:/opt/homebrew/opt/mysql-client/bin"' >> ~/.zshrc
source ~/.zshrc
make mysql
```

- PostgreSQL
```
brew install libpq
echo 'export PATH="${PATH}:/opt/homebrew/opt/libpq/bin"' >> ~/.zshrc
source ~/.zshrc
make pgsql
```
