# go-http
Steps to build the WebService :
1) Checkout or unzip the code.
2) set your GOPATH env var to <on-your-checkout-machine-path-to>/gp-http folder3) CD to <...path-to>/go-http/src.
4) run the command - go build -o < name-of-binary-you->
   eg: go build -o websvc <- this will build a binary websvc conatining the webservice
5) Start the webservice by running the binary create step 4. This will start an httpd deamon listening on 8080 port.
6) On another terminal use cmd : curl localhost:8080
   This should output a random Chuck Norris joke.
