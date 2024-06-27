### http-test
simple test app for my ocp lab that will just deploy a go http-server with a persistent volume attached.

### How to
- localhost:8080/headers
- localhost:8080/mem
- localhost:8080/hello

### deploy to ocp
$ bash <(curl -s https://raw.githubusercontent.com/backchristoffer/http-test/master/setup/install.sh)

### teardown/uninstall
$ bash <(curl -s https://raw.githubusercontent.com/backchristoffer/http-test/master/setup/teardown.sh)
