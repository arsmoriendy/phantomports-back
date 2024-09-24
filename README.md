# Opor Back
Backend `graphql` server for opor written in `go`.

## Environment Variables
Environment Variable | Description
--- | ---
`MODE` | `DEV` for dev mode, `PROD` for production. Defaults to `DEV`
`PORT` | Port to run server
`HOST` | Host to run server
`LOG_LEVEL` | One of *Verbosity* or *Level* (e.g. `LOG_LEVEL=2` is the same as `LOG_LEVEL=wArN`). See [Log Levels](#log-levels)

## Log Levels
Verbosity | Level (Case Insensitive)
--- | ---
0 | `FATAL`
1 | `ERROR`
2 | `WARN`
3 | `INFO`
4 | `DEBUG`
5 | `TRACE`

> [!NOTE]
> This repo was originally apart of [the archived monorepo](https://github.com/arsmoriendy/opor)
