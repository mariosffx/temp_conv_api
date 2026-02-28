# TempConv
A simple temperature convertion tool written with **Go**

Converts between Celsius, Fahrenheit, and Kelvin.

## Homework
- TempConv from scratch with gRPC + ProtocolBuffers
- public GitHub repository



## Architecture

```
temp_conv_api
|-- convertors.go   Includes functions that convert the temperature
|-- go.mod
|-- handlers.go     Handler functions that run on each request
├── main.go         Main Server application
|-- middleware.go   Configures CORS
└── types.go        Includes Types
```

## Getting Started

### Prerequisites

- [Go](https://go.dev/) 1.21+

### Run the Backend

```bash
go run .
```

The server starts on `http://localhost:8080`.

### Build
```bash
go build
```

### Run the build
```bash
./temp_conv_api
```

## API

| Protocol | Method | Endpoint     | Description           |
| -------- | ------ | ------------ |
| HTTP     | POST   | /api/convert | Convert a temperature |

### POST /api/convert

```json
// Request
{ "value": 100, "from": "celsius", "to": "fahrenheit" }

// Response
{ "original_value": 100, "from": "celsius", "to": "fahrenheit", "result": 212 }
```

### Sample Requests
Convert 30 **Degress** from Celsius to Fahrenheit 

```sh
curl --location 'http://localhost:8080/api/convert' \
--header 'Content-Type: application/json' \
--data '{
    "value": 30,
    "from":"celsius",
    "to":"fahrenheit"
}'
