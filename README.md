## GO RESTFUL APP

A simple RESTful API using Go. Intended for AGMC (by [Alterra](https://www.alterra.id/)) final project.

Notable features:

- Includes a multi-stage Dockerfile, which builds Go binaries (minimum app size).
- Has functional tests for the application's business requirements using default Go testing.
- You need to run the database app (MySQL) on your local machine first before up the app docker image
<!-- - Endpoints documentation was published at  -->

### Setup

1. Create `.env` config file, look at [.env.example](./.env.example) for mandatory _key-value_

2. Build an app using `docker`

    ```bash
    docker build -t <your_tag> .
    ```

    or using `docker-compose`

    ```bash
    docker-compose build
    ```

    or just _up_ published docker image using following [docker-compose.yml](docker-compose.client.yml) file

    ```
    version: "3.7"

    services:
    app:
        image: ymanshur/go-restful:${APP_VERSION}
        container_name: go-restful
        ports:
            - ${HOST_PORT:-8000}:${PORT:-8000}
        expose:
            - ${PORT:-8000}
        environment:
            JWT_SECRET: ${JWT_SECRET:?err}
            DB_USER: ${DB_USER:?err}
            DB_PASS: ${DB_PASS:?err}
            DB_PORT: ${DB_PORT:?err}
            DB_HOST: host.docker.internal
            DB_NAME: ${DB_NAME:?err}
    ```

    Note:
    - `host.docker.internal` means the app will access your local machine host

3. You can _up_ existing DB image, such as MySQL (officially exists at GitHub Hub registry) and use it for the database, but it's optional. You can write your service or extend from [docker-compose.db.yml](docker-compose.db.yml).

    ```
    app_db:
        extends:
            file: docker-compose.db.yml
            service: db
        container_name: go-restful_db
    ```

    and if you need a lite version of phpMyAdmin, just add the following service

    ```
    app_db_adminer:
        extends:
            file: docker-compose.db.yml
            service: adminer
        container_name: go-restful_db_adminer
    ```

4. Run the app using docker-compose

    ```bash
    docker-compose up -d
    ```

    <small>Note: `-d` argument means detached mode</small>

### Testing

- Firstly, you should create `.env.test` file (see [.env.test.example](.env.test.example) for mandatory _key-value_).

- If you don't have a database on your machine, you can use containerized database for testing by run following command.

    ```bash
    docker compose -f docker-compose.test.yml --env-file .env.test up -d
    ```

- Do a functional test using the following command

    ```bash
    make test
    ```

    To show in HTML mode, use

    ```bash
    make test mode=html
    ```
