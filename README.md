# travelgrunt

Travel **[Terragrunt](https://terragrunt.gruntwork.io/)** directory tree as a first class passenger! :airplane:

## How to use?

* `cd` to the directory of your [locally cloned] Terragrunt/Terraform Git repo;
* run **tg** alias there :rocket: ([optional] arguments are "path filter" matches);
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

## How to build?

* `make dep` - install dependencies;
* `make build` - build the `travelgrunt` binary in `cmd/travelgrunt` path;
* `make install` - [optional] install built `travelgrunt` binary under the `${PREFIX}/bin` location;

## How to release a new version?

:bulb: Set `GITHUB_TOKEN` environmental variable.

Run `make full-release` recipe, which is equal to run following, one by one:

* `make clean` - cleanup project tree from previously built artifacts;
* `make dep` - ensure all dependencies are installed;
* `make release` - create release artifacts;
* `make next-version-tag` - tag your HEAD with the incremented version;
* `make push-tags` - push your new tag to Git;
* `make github` - create a GitHub release using your artifacts;
