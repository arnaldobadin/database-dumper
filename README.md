# Simple Database Dumper

Simple program written in GO that dumps database rows into pre-configured formats

## Usage:

Just build and execute its binary.
You can change the executions configuration path passing "-excs pathtoconfig.json" after the binary command. But it is optional and its default is set to relative "executions.json".

```bash
cd src/dumper
go build

dumper -excs pathtoconfig.json
```

## Executions Config

Executions config is just an array of options.

- name: "execution name"
- query: "query to be executed"
- connection.storage: "type of storage, e.g.: mysql/mssql"
- connection.host: "storage host, e.g.: localhost"
- connection.user: "storage user"
- connection.password: "storage. password"
- connection.database: "storage database"
- output.path: "path to store result"
- output.name: "result file name"
- output.extension: "result file format"

Only switch-type fields are:

- connection.storage: "mysql"/"mssql"
- output.extension: "csv"/"pipe"

```json
[
    {
        "name" : "nothing",
        "query" : "SELECT * FROM mytable WHERE active IS NOT NULL;",
        "connection" : {
            "storage" : "mysql",
            "host" : "localhost",
            "user" : "myuser",
            "password" : "mypassword",
            "database" : "mydatabase"
        },
        "output" : {
            "path" : "C:\\Users\\myuser\\Desktop",
            "name" : "result",
            "extension" : "csv"
        }
    }
]
```

