# ttg

**T**ravel **T**erra**g**runt directory tree as a first class passenger! :sunglasses:

## How to use?

* `cd` to the directory of your [locally cloned] Terragrunt/Terraform Git repo;
* run `ttg` command there ([optional] arguments are "path filter" matches);
* use arrow keys to navigate the list and `/` key to search for specific projects;

## How to build?

* `make dep` - install dependencies;
* `make build` - build the `ttg` binary in `cmd/ttg` path;
* `make install` - [optional] install built `ttg` binary under the `${PREFIX}/bin` location;
