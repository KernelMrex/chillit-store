# Chillit Store

### About

Service for uploading public places data. Will be a part of chillit web-app. 

### Using

Compile `make build` and run `./chillit-store [-config_path=<path>]` or just `make run`

### Configuration

Add file `config.yaml` to working directory.
 
``` yaml
database:
  user: "user"
  password: "password"
  host: "localhost"
  port: 3306
  db: "database"
``` 
