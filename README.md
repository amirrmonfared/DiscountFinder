# WebCrawler

## Overview
The Project we're try to build is a discount finder from every website just with changing the html we want.

## Setup local development

### Install tools

- [Docker desktop](https://www.docker.com/products/docker-desktop)
- [TablePlus](https://tableplus.com/)
- [Golang](https://golang.org/)
- [Homebrew](https://brew.sh/)
- [Migrate](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)

    ```bash
    brew install golang-migrate
    ```

- [Sqlc](https://github.com/kyleconroy/sqlc#installation)

    ```bash
    brew install golang-migrate
    ```

### Setup infrastructure

- Start postgres container:

    ```bash
    make postgres
    ```

- Create database:

    ```bash
    make createdb
    ```

- Run db migration:

    ```bash
    make migrateup
    ```

### How to run

- Run test:

    ```bash
    make test
    ```
