# go-http
Steps to build the WebService :
1) Checkout or unzip the code.
2) set your GOPATH env var to <on-your-checkout-machine-path-to>/go-http folder)
3) CD to <...path-to>/go-http/src.
4) run the command :>  'go build -o <name-of-the-binary-you-want-to-creare>' ...
   for  eg: 'go build -o websvc' <- this will build a binary named websvc conatining the webservice
5) Start the webservice by running the binary create step 4. This will start the http server listening on 8080 port by default.
   To choose another port set the env HTTP_PORT=XXXX .for e.g. 'export HTTP_PORT=8090' 
6) On another terminal use cmd : 'curl localhost:$HTTP_PORT'  <= the port you set in the env var or 8080.
   This should output a random Chuck Norris joke.
