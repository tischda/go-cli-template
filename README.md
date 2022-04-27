# go-cli-template

Template for small [Go](https://www.golang.org) CLI projects.

## Get started

Name your project:
~~~
set OWNER=owner
set PROJECT=my-project
~~~

Create repository and project folder with [Github CLI](https://github.com/cli/cli):
~~~
gh repo create %PROJECT% --public --template github.com/tischda/go-cli-template
git clone git@github.com:%OWNER%/%PROJECT%.git
cd %PROJECT%
go mod init github.com/%OWNER%/%PROJECT%
go generate template.go
~~~

Start coding.

## Add modules

~~~
go mod tidy
go mod vendor
~~~

## Release project

~~~
make test

git tag -a v1.0.0 -m "First release"
git push origin v1.0.0
make release
~~~
