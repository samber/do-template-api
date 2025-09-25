
# API boilerplate showcasing github.com/samber/do

![Go Version](https://img.shields.io/badge/Go-%3E%3D%201.23-%23007d9c)
![Build Status](https://github.com/samber/do-template-api/actions/workflows/test.yml/badge.svg)
[![Go report](https://goreportcard.com/badge/github.com/samber/do-template-api)](https://goreportcard.com/report/github.com/samber/do)
[![License](https://img.shields.io/github/license/samber/do-template-api)](./LICENSE)

**⚙️ A comprehensive API template demonstrating the `github.com/samber/do` dependency injection library.**

A comprehensive API template project demonstrating the full power of the `github.com/samber/do` dependency injection library. This project implements a complete REST API with PostgreSQL integration, showcasing type-safe dependency injection, modular architecture, and real-world web application concerns.

Perfect as a starting point for new Go web projects or as a learning resource for understanding dependency injection patterns in API applications.

**See also:**

- [do-template-worker](https://github.com/samber/do-template-worker)
- [do-template-cli](https://github.com/samber/do-template-cli)

## 🚀 Install

Clone the repo and install dependencies:

```bash
git clone --depth 1 --branch main https://github.com/samber/do-template-api.git your-project-name
cd your-project-name

docker compose up -d
make deps
make deps-tools
```

## 💡 Features

- **Type-safe dependency injection** - Service registration and resolution using `samber/do`
- **Modular architecture** - Clean separation of concerns with dependency tree visualization
- **REST API framework** - Built with Gin for robust HTTP web services
- **Configuration management** - Environment-based configuration with dependency injection
- **PostgreSQL integration** - Complete database setup with connection pooling and migrations
- **Repository pattern** - Data access layer with injected dependencies
- **Service layer** - Business logic with proper dependency management
- **Application lifecycle** - Health checks and graceful shutdown handling
- **Comprehensive error handling** - Structured logging and error management
- **Production-ready** - Ready to fork and customize for your next API project
- **Extensive documentation** - Inline comments explaining every `do` library feature

## 🚀 Contributing

```sh
# install deps
make deps
make deps-tools

# compile
make build

# build with hot-reload
make watch-run

# test with hot-reload
make watch-test
```

## 🤠 `do` documentation

- [GoDoc: https://godoc.org/github.com/samber/do/v2](https://godoc.org/github.com/samber/do/v2)
- [Documentation](https://do.samber.dev/docs/getting-started)

## 💫 Show your support

Give a ⭐️ if this project helped you!

[![GitHub Sponsors](https://img.shields.io/github/sponsors/samber?style=for-the-badge)](https://github.com/sponsors/samber)

## 📝 License

Copyright © 2025 [Samuel Berthe](https://github.com/samber).

This project is [MIT](./LICENSE) licensed.
