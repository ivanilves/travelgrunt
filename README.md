[![codeclimate](https://api.codeclimate.com/v1/badges/857b9074dbf627e4f594/maintainability)](https://codeclimate.com/github/ivanilves/travelgrunt/maintainability)
[![codecov](https://codecov.io/github/ivanilves/travelgrunt/branch/main/graph/badge.svg?token=SW21884ADR)](https://codecov.io/github/ivanilves/travelgrunt)


# travelgrunt

Travel **[Terragrunt](https://terragrunt.gruntwork.io/)**, **[Terraform](https://www.terraform.io/)** and ... ANY [Git] repository tree with no stress, as a first class passenger! :airplane:

`travelgrunt` alleviates the pain of hitting endless `cd`/`<TAB>` combos while navigating inside the repo.

:bulb: Video with a brief introduction and a demo is available [here](https://www.youtube.com/watch?v=iqKDJZhfjj4) :movie_camera:

## How to use?

* `cd` to the directory of your [locally cloned] Git repository;
* run **tg** [alias](#shell-aliases) there :rocket: ([optional] arguments are "path filter" matches);
* use arrow keys to navigate the list and `/` key to search for specific items;

## Configuration
:bulb: If no configuration file found `travelgrunt` will assume repository having only Terragrunt projects inside.

Create `.travelgrunt.yml` file in the root path of your repository. Example config for a random monorepo:

```
rules:
  - prefix: vendor/
    negate: true
  - prefix: terragrunt/
    mode: terragrunt
  - prefix: code/
    name: '.*\.(go|js|css|html)$'
  - prefix: config/
    name: '*.yaml'
```

:arrow_up: Config is essentially a list of sequentially applied path matching rules. Each rule can have these fields:

* `prefix` - literal prefix to be matched against relative directory path;
* `name` - a free form regular expression or a simple glob (`name: '*.go'`) match applied to the file name;
* `mode` - any matching behavior backed by a [custom logic] function from the [`mode`](https://github.com/ivanilves/travelgrunt/tree/main/pkg/config/mode) package;
* `negate` - boolean directive that reverses the meaning of the match, excluding the paths matched;

:bulb: Even while developing `travelgrunt` itself we use it to navigate [package directories](https://github.com/ivanilves/travelgrunt/blob/main/.travelgrunt.yml) of the application :tophat:

## Shell aliases

It is **absolutely required** to use `bash` (or `zsh`) aliases. Start from something like this:
```
alias tg='_tg(){ travelgrunt -out-file ~/.tg-path ${@} && cd "$(cat ~/.tg-path)" }; _tg'
alias tt='_tt(){ travelgrunt -top -out-file ~/.tg-path && cd "$(cat ~/.tg-path)" }; _tt'
```

These lines are usually added to `~/.bashrc` or `~/.zshrc` file, depending on your system and shell of choice.

:bulb: **tt** is a "convenience alias" that brings you to the top level path of your repository.

### Why aliases?
Core feature of this program is the ability to change working directory while staying **inside the current shell**.
This **can not** be done by the program itself, because of `POSIX` security limitations. Without instrumenting
the shell with aliases `travelgrunt` will not work!

### `CTRL+C` / `CTRL+D` behaviour
When key combinations `CTRL+C` or `CTRL+D` get pressed during the execution, following occures:
* `CTRL+C` - program terminates with exit code `1`, under the starting directory path;
* `CTRL+D` - program terminates with exit code `0`, under the directory path currently selected;

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
