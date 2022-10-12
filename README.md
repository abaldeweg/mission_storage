# baldeweg/mission_storage

An app to administer a log file of missions.

## Repositories

- mission_storage <https://github.com/abaldeweg/mission_storage> - Backend
- mission_ui <https://github.com/abaldeweg/mission_ui> - UI

## Requirements

- Golang
- Docker

## Getting Started

First, you need to install [Go](https://go.dev/).

Download the project archive from the [git repository](https://github.com/abaldeweg/desk_storage).

Inside the project directory, you can build the app with the `go build` command.

Run the command `mission_storage`. Depending on the OS you need to add a file extension.

The app will create files where you can edit the missions.

## Env vars

Create a `.env` file to define some settings.

```env
// .env
ENV=prod
GCP_BUCKET_NAME=name
GOOGLE_APPLICATION_CREDENTIALS=service-account-file.json
CORS_ALLOW_ORIGIN=http://localhost:8081
```

- ENV - Set to `prod`, `dev` or `test`
- GCP_BUCKET_NAME - If `gcp-bucket` was chosen as storage, then define the bucket name.
- GOOGLE_APPLICATION_CREDENTIALS - Key file, for auth and buckets
- CORS_ALLOW_ORIGIN - Allowed origins
