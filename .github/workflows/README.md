
# The Whisper Game

This sample kicks off Github action to deploy a GO REST API to Azure Container Instance. 

* Whisper a message - make a HTTP POST call with the following settings, and it should return http status 200

```
URL: http://<youraci>.<yourlocation>.azurecontainer.io:5000/whisper
Content-Type: application/json
Body: {"sender":"john", "message":"hello from john"} 
```
If the sender is the same as in environment variable ENV_MYSELF, it will return a text "that's myself", otherwise, it will forward the message to ENV_FORWARDURL.

* Retrieve all whispered gossips since the Azure Container Instance started
```
URL: http://<youraci>.<yourlocation>.azurecontainer.io:5000/gossips
```

## TODO
* Add unit tests
* Offload SSL to a sidecar 
