# go-cli-template

Template for small Windows CLI projects hosted on github.com.

## Get started

Name your project, for example:
~~~
set OWNER=tischda
set PROJECT=test01
set COPYRIGHT=© 2025 Daniel TISCHER
~~~

Create repository and project folder with [Github CLI](https://github.com/cli/cli):
~~~
gh config set git_protocol ssh
gh repo create %PROJECT% --public --clone --template github.com/tischda/go-cli-template
cd %PROJECT%
go generate template.go
~~~

Start coding.

## Release project

~~~
make test

git tag -a v1.0.0 -m "First release"
git push origin v1.0.0
make release
~~~
