# How to release a new version? :package:

:bulb: Export `GITHUB_TOKEN` environment variable. You may use git-ignored `secrets.mk` for this purpose:
```
echo "export GITHUB_TOKEN := ghp_<your_very_secret_token_here>" > secrets.mk
```

Run `make full-release` recipe, which is equal to run following recipes, one by one:

* `make github-token` - check if `GITHUB_TOKEN` variable is set;
* `make clean` - cleanup project tree from previously built artifacts;
* `make dep` - ensure all dependencies are installed;
* `make release` - create release artifacts;
* `make next-version-tag` - tag your HEAD with the incremented version;
* `make push-tags` - push your new tag to Git;
* `make github` - create a GitHub release using your artifacts;
