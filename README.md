## GO RESTFUL APP

The day 8 & 10 AGMC (by [Alterra](https://www.alterra.id/)) submission project.

Notable features:

- Includes a multi-stage Dockerfile, which actually build Go binaries (minimum app size).
- Has functional tests for application's business requirements using default Go testing.
- You need to run database app (MySQL) in your local machine first before up the app docker image
- You can access the endpoints at https://go-restful.herokuapp.com/
<!-- - Endpoints documentation was published at  -->

### Setup

1. Create `.env` config file, look at [.env.example](./.env.example) for mandatory _key-value_

2. Build app using `docker`

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
    - `host.docker.internal` means app will access your local machine host

3. You can _up_ existing db image such as mysql (officialy exist at Github Hub registry) and use it for database, but its optional. You can write your own service or extend from [docker-compose.db.yml](docker-compose.db.yml).

    ```
    app_db:
        extends:
            file: docker-compose.db.yml
            service: db
        container_name: go-restful_db
    ```

    and if you need lite version of phpMyAdmin just add following service

    ```
    app_db_adminer:
        extends:
            file: docker-compose.db.yml
            service: adminer
        container_name: go-restful_db_adminer
    ```

4. Run app using docker-compose

    ```bash
    docker-compose up -d
    ```

    <small>Note: `-d` argument means detached mode</small>

### Testing

- Firstly, you should create `.env.test` file (see [.env.test.example](.env.test.example) for mandatory _key-value_).

- If you have not database in you machine, you can use containerized database for testing by run following command.

    ```bash
    docker compose -f docker-compose.test.yml --env-file .env.test up -d
    ```

- Do functional test using following command

    ```bash
    make test
    ```

    To show in html mode, use

    ```bash
    make test mode=html
    ```
