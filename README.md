# Email Service

This code is implementation of [Coding Challenge](https://github.com/ArentInc/coding-challenge-tools/blob/master/coding_challenge.md).

## Package structure

Basically, the package structure follows the [Standard Go Project Layout](https://github.com/golang-standards/project-layout).

```
.
├── api                     # API Specification
├── cmd                     # Entrypoint of launching app
├── internal                # Code for internal use
│   ├── app                 # Code for app specific
│   │   ├── container       # Create DI container
│   │   ├── http            # Code for API specific
│   │   │   ├── hander      # Request handler
│   │   │   ├── request     # Data object for reuqest
│   │   │   └── response    # Data object for response
│   │   └── service         # Business logic
│   └── pkg                 # Common codes
├── web                     # Frontend codes
├── docker-compose.yaml
├── Dockerfile
├── LICENSE
└── README.md

```

## How to run app

### for local

#### Prepare services

* Amazon SES
  * [Amazon Simple Email Service を設定する](https://docs.aws.amazon.com/ja_jp/ses/latest/dg/setting-up.html)
* SendGrid
  * [SMTPメールの構築](https://sendgrid.kke.co.jp/docs/API_Reference/SMTP_API/building_an_smtp_email.html)

#### Prepare `.env` file

Copy `.env.example` to `.env` and edit it.
Should input these values

* `AWS_ACCESS_KEY_ID`
* `AWS_SECRET_KEY`
* `AWS_SES_ENDPOINT`
* `AWS_REGION`
* `SENDGRID_API_KEY`

#### Execute 

```bash
COMPOSE_DOCKER_CLI_BUILD=1 DOCKER_BUILCDKIT=1 docker-compose up --build
```

## How to execute test

Execute all tests.

```bash
go test ./...
```