# PhantomPorts back
Backend `graphql` server for PhantomPorts.com written in `go`.

## Environment Variables
Environment Variable | Description
--- | ---
`MODE` | `DEV` for dev mode, `PROD` for production. Defaults to `DEV`
`PORT` | Port to run server
`HOST` | Host to run server
`LOG_LEVEL` | One of *Verbosity* or *Level* (e.g. `LOG_LEVEL=2` is the same as `LOG_LEVEL=wArN`). See [Log Levels](#log-levels)
`DB_URL` | Database URL
`REFRESH_INTERVAL` | Interval to refresh ports from IANA's registry in milliseconds. Defaults to an hour
`FRONT_UUID_EXPR` | Expiration time for frontend uuid in milliseconds. Defaults to an hour
`REF_FRONT_UUID_PASS` | Password to refresh frontend uuid

## Log Levels
Verbosity | Level (Case Insensitive)
--- | ---
0 | `FATAL`
1 | `ERROR`
2 | `WARN`
3 | `INFO`
4 | `DEBUG`
5 | `TRACE`

## Authorization
Authorization is done using [HTTP authentication](https://developer.mozilla.org/en-US/docs/Web/HTTP/Authentication) with the `Basic` scheme.

The **username** field should be left blank. The **password** field should be a valid **UUID** registered in the database.

## Regenerate From Schema
```bash
go run github.com/99designs/gqlgen generate
```
