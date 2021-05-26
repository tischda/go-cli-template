# go-cli-template

Template for [Go](https://www.golang.org) CLI projects.

### Usage

~~~
set PROJECT=my-project
set MODULE=github.com/tischda/%PROJECT%

gh repo create %PROJECT% --confirm --public --template github.com/tischda/go-cli-template
cd %PROJECT%
go mod init %MODULE%
go generate
~~~

Start coding.

### Add modules

~~~
go mod tidy
go mod vendor
~~~

### Release

~~~
make test
make release
~~~
