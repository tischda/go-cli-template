# go-cli-template

Template for small golang CLI projects hosted on github.com.

## Get started (Windows)

Name your project, for example:
~~~
set PROJECT=test01
set OWNER=tischda
set COPYRIGHT=© 2025 Daniel TISCHER
~~~

Create repository and project folder with [Github CLI](https://github.com/cli/cli):
~~~
gh config set git_protocol ssh
gh repo create %PROJECT% --public --clone --template github.com/tischda/go-cli-template
cd %PROJECT%
go generate windows.go
~~~


### Get started (Unix)

~~~
PROJECT=test01
OWNER=tischda
COPYRIGHT="© 2025 Daniel TISCHER"
~~~

Create repository and project folder:
~~~
gh config set git_protocol ssh
gh repo create $PROJECT --public --clone --template github.com/tischda/go-cli-template
cd $PROJECT
go generate unix.go
~~~



## Release project

~~~
make test

git tag -a v1.0.0 -m "First release"
git push origin v1.0.0
make release
~~~
