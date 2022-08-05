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

## Sample Queries

Use the `sample/queries` directories to check the samples and use them with CURLs

- Add Organizations (Could be 1 or more):

```
curl -X POST -H "Content-Type: application/json" -d @add_org.json http://127.0.0.1:8080/v1/org/add
```

- Add 1 MlModel:
```
curl -X POST -H "Content-Type: application/json" -d @add_mlmodel.json http://127.0.0.1:8080/v1/ml/add
```

- Add X MlModels:
```
curl -X POST -H "Content-Type: application/json" -d @add_mlmodels.json http://127.0.0.1:8080/v1/ml/addBulk
```
