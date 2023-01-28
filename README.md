## Flight Path Tracker

To create a simple microservice API that can help us understand and track how a particular person's flight path may be queried. The API should accept a request that includes a list of flights, which are defined by a source and destination airport code. These flights may not be listed in order and will need to be sorted to find the total flight paths starting and ending airports.


Run Go app

```shell script
go run ./cmd/server/main.go
```

Or 

```shell script
go build -o flighpathtracker ./cmd/server/main.go
./flighpathtracker
```

| API Requests                    |
|---------------------------------|
| http://localhost:8080/calculate |

## API Specs

### `POST /calculate`
Endpoint to determine the flight path of a person. 

The request payload should have the following request body fields:

```json
{
    "list_of_flights": [
      ["SFO", "EWR"]
    ]
} 
```
 Or
 
 ```json
{
  "list_of_flights": [
    [
      "ATL",
      "EWR"
    ],
    [
      "SFO",
      "ATL"
    ]
  ]
}
 ```
Or
```json
{
  "list_of_flights": [
    [
      "IND",
      "EWR"
    ],
    [
      "SFO",
      "ATL"
    ],
    [
      "GSO",
      "IND"
    ],
    [
      "ATL",
      "GSO"
    ]
  ]
}
```

The response body should be expected by the following fields:

```json
{
  "total_flight_paths": [
    "SFO",
    "EWR"
  ]
}
```

