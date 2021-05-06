# go-archetype

Template for [Go](https://www.golang.org) CLI projects.

Convention: project directory = project name

### Usage

Set project variables for configuration:
~~~
set PROJECT=YOUR_PROJECT_NAME
set REPOSITORY=github.com/tischda
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
