# Demo Math Micro Service
This micro service written in Go using the Echo micro framework is a simple example of
how to easily create a micro service able to be containerised.

## Installation
* `$ git clone https://github.com/davidvartanian/demo-math`
* `$ docker-compose up -d --build`
* Service available at `http://localhost:777`

## Endpoints
* GET  /swagger/index.html
* POST /sum
* POST /subtract
* POST /multiply
* POST /divide

## Payload
```
{
  "a": number,
  "b": number
}
```

## Testing
`$ go test`