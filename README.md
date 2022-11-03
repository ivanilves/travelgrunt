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

## How to install?

### Install with homebrew

```
brew tap ivanilves/tap
brew install ivanilves/tap/travelgrunt
```

### Install latest binary with cURL
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

## How to build?

* `make dep` - install/ensure dependencies;
* `make build` - build the `travelgrunt` binary in `cmd/travelgrunt` path;
* `make install` - [optional] install built `travelgrunt` binary under the `${PREFIX}/bin` location;

[How to release a new version?](RELEASE.md) :package:
