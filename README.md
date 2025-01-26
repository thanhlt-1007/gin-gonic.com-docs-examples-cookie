# gin-gonic.com-docs-examples-cookie

- Set and get a cookie

- Reference: https://gin-gonic.com/docs/examples/cookie/

## gvm

```sh
gvm use go1.23.5
```

## go get .

```sh
go get .
```

## go run .

```sh
go run .
```

## cURL

- GET /cookie

```sh
curl --location 'localhost:8080/cookie'
```

- GET /cookie with cookie

```sh
curl --location 'localhost:8080/cookie' \
--header 'Cookie: gin_cookie=TEST'
```
