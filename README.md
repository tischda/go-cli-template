# go-archetype

Template for [Go](https://www.golang.org) CLI projects.

Convention: project directory = YOUR_PROJECT_NAME

### Usage

Set environment variables for configuration:
~~~
set PROJECT=YOUR_PROJECT_NAME
set REPOSITORY=github.com/tischda/YOUR_PROJECT_NAME
~~~

Download and configure:
~~~
wget --no-check-certificate https://github.com/tischda/go-archetype/archive/refs/heads/master.zip
unzip -j -d %PROJECT% master.zip
rm master.zip
cd %PROJECT%
make -f Makefile.init
~~~

### Clean-up

~~~
rm Makefile.init
~~~

Start coding.

### Adding modules

~~~
go mod tidy
go mod vendor
~~~
