<div align="center" style="margin-bottom:20px">
  <!-- <img src=".assets/banner.png" alt="logger" /> -->
  <div align="center">
    <a href="https://github.com/blugnu/logger/actions/workflows/qa.yml"><img alt="build-status" src="https://github.com/blugnu/logger/actions/workflows/qa.yml/badge.svg?branch=master&style=flat-square"/></a>
    <a href="https://goreportcard.com/report/github.com/blugnu/logger" ><img alt="go report" src="https://goreportcard.com/badge/github.com/blugnu/logger"/></a>
    <a><img alt="go version >= 1.14" src="https://img.shields.io/github/go-mod/go-version/blugnu/logger?style=flat-square"/></a>
    <a href="https://github.com/blugnu/logger/blob/master/LICENSE"><img alt="MIT License" src="https://img.shields.io/github/license/blugnu/logger?color=%234275f5&style=flat-square"/></a>
    <a href="https://coveralls.io/github/blugnu/logger?branch=master"><img alt="coverage" src="https://img.shields.io/coveralls/github/blugnu/logger?style=flat-square"/></a>
    <a href="https://pkg.go.dev/github.com/blugnu/logger"><img alt="docs" src="https://pkg.go.dev/badge/github.com/blugnu/logger"/></a>
  </div>
</div>

<br>

# logger

A package that provides an adaptable logger implementation to be used by other re-usable modules that wish to emit logs using a logger supplied by the consuming project.

Such packages would provide a mechanism for "injecting" a logger exported by this package.  Automatic log enrichment (from `context`) can also be established by those modules implementing an `init()` function to register an enrichment func.

## How It Works

Despite the name, this `logger` package does not implement an actual logger, rather it provides a type that delegates logging calls to an _adapter_.  A consuming project will configure whatever logger it wishes and then pass that logger into a dependent package by supplying a logger from _this_ package along with an appropriate _adapter_.

Adapters are provided in this package for `logrus` and the standard `log` package.

A `NulAdapter` is also provided which produces no log output what-so-ever (logging to NUL).

<br>
<hr>
<br>

## How to Use Logger

### In an Application or Service Project

_IF NEEDED_: Implement an `Adapter` (if one does not already exist for your preferred logger)

1. Configure your logger
2. Wrap your project logger in a `Logger` using the appropriate `Adapter`
4. Pass the `Logger` into any modules/packages that support shared logging
3. _OPTIONAL:_ Register `Enrichment` functions provided by your project and any other modules (they do not need to be logging via a `Logger` in order to provide enrichment)
5. Enjoy your logs

<br>

### In a Module/Package

1. Provide a means for a `Logger` to be configured for use by your module/package
    - **EITHER:** import this module and use the `Logger` type directly 
    - **OR:** employ an interface type (declared in your module/package) that is compatible with those methods of `Logger` that your code uses
2. _OPTIONAL_: Export an `Enrichment` function
3. _OPTIONAL_: Add log enrichment data to `context` where appropriate and use `errorcontext` to return errors with context for enriched logging 
4. Write logs from your code using the configured `Logger`
    - Ensure that logs are not written if _no_ `Logger` is configured (provide a _no-op_ implementation of your defined interface _or_ initialise a default `Logger` with a `NulAdapter` if you don't mind taking a dependency on this `logger` module)
