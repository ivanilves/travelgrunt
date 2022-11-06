[![codeclimate](https://api.codeclimate.com/v1/badges/857b9074dbf627e4f594/maintainability)](https://codeclimate.com/github/ivanilves/travelgrunt/maintainability)
[![codecov](https://codecov.io/github/ivanilves/travelgrunt/branch/main/graph/badge.svg?token=SW21884ADR)](https://codecov.io/github/ivanilves/travelgrunt)


# travelgrunt

Travel **[Terragrunt](https://terragrunt.gruntwork.io/)** directory tree as a first class passenger! :airplane:

## How to use?

* `cd` to the directory of your [locally cloned] Terragrunt/Terraform Git repo;
* run **tg** [alias](#shell-aliases) there :rocket: ([optional] arguments are "path filter" matches);
* use arrow keys to navigate the list and `/` key to search for specific items;

## Shell aliases

It is **absolutely required** to use `bash` (or `zsh`) aliases. Start from something like this:
```
alias tg='_tg(){ travelgrunt -out-file ~/.tg-path ${@} && cd "$(cat ~/.tg-path)" }; _tg'
alias tt='_tt(){ travelgrunt -top -out-file ~/.tg-path && cd "$(cat ~/.tg-path)" }; _tt'
```

:bulb: **tt** is a "convenience alias" that brings you to the top level path of your repository.

### Why aliases?
Core feature of this program is the ability to change working directory while staying **inside the current shell**.
This **can not** be done by the program itself, because of `POSIX` security limitations. Without instrumenting
the shell with aliases `travelgrunt` will not work!

## How to install? :neckbeard:

### Install with `homebrew`:

```
brew tap ivanilves/tap
brew install ivanilves/tap/travelgrunt
```

### Install latest binary with `cURL` + `sh`:
Install latest release binary into `/usr/local/bin` (default):

```
curl -s -f \
  https://raw.githubusercontent.com/ivanilves/travelgrunt/main/scripts/install-latest.sh | sh
```

Install latest release binary into `/somewhere/else/bin`:
```
curl -s -f \
  https://raw.githubusercontent.com/ivanilves/travelgrunt/main/scripts/install-latest.sh \
  | PREFIX=/somewhere/else sh
```

## How to build? :construction:

* `make dep` - install/ensure dependencies;
* `make build` - build the `travelgrunt` binary in `cmd/travelgrunt` path;
* `make install` - [optional] install built `travelgrunt` binary under the `${PREFIX}/bin` location;

## How to release a new version? :package:

:bulb: Make sure you have push permissions for this repository!

Run `make release` recipe, which will:
* check, if you are on a `main` branch;
* pull the latest `main` branch from remote;
* calculate the next release version (update `MAJOR`.`MINOR` [here](https://github.com/ivanilves/travelgrunt/blob/main/Makefile#L2) if needed);
* tag the branch tip with the version calculated and push tag to remote then;
* [GoReleaser](https://github.com/ivanilves/travelgrunt/blob/main/.goreleaser.yml) will take care of everything else :sunglasses:
