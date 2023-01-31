## catchall

Run Go app

```shell script
go run ./cmd/main.go
```

Or 

```shell script
go build -o catchall ./cmd/main.go
./catchall
```

| API Requests                                         |
|------------------------------------------------------|
| http://localhost:8080/events/<domain_name>/delivered |
| http://localhost:8080/events/<domain_name>/bounced   |
| http://localhost:8080/domains/<domain_name>          |

