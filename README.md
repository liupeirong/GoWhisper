
# The Whisper Game

This sample uses Github actions to deploy a GO REST API to Azure Container Instances(ACI).  It uses Azure Cli rather than Github aci-deploy action because it deploys two containers - the application container and an Nginx sidecar that offloads SSL. 

## How to set it up

* Base64 encode SSL certificate, key, and Nginx configuration file as documented [here](https://docs.microsoft.com/en-us/azure/container-instances/container-instances-container-group-ssl). Make sure to run `base64 -w 0` to encode to a string without newline, and set the encoded strings in Github secrets as `BASE64_SSL_CRT`, `BASE64_SSL_KEY`, and `BASE64_NGINX_CONF` respectively.
* [deploy-aci.yml](deploy-aci.yml) contains all the information to deploy a multi-container group to ACI as documented [here](https://docs.microsoft.com/en-us/azure/container-instances/container-instances-multi-container-yaml).
* [Github actions](.github/workflows/go.yml) replaces the placeholder variables in [deploy-aci.yml](deploy-aci.yml) with variables set in Github secrets before deploying to ACI. 

## How to run it

* To whisper a message: make a HTTP POST call with the following settings, and it should return http status 200

```
URL: https://<youraci>.<yourlocation>.azurecontainer.io/whisper
Content-Type: application/json
Body: {"sender":"john", "message":"hello from john"} 
```
If the sender is the same as set in environment variable ENV_MYSELF, it will return a text "that's myself", otherwise, it will forward the message to ENV_FORWARDURL.

* To retrieve all whispered messages since the Azure Container Instance started: make a HTTP GET call as following

```
URL: https://<youraci>.<yourlocation>.azurecontainer.io/gossips
```

## Azure Container Instances or Azure App Service
Azure App Service can also run containers. Overall ACI provides more control and flexibility while App Service provides more convenience. For example:
* You don't have to redeploy to update environment variable in App Service, whereas you do in ACI. 
* You can set restart policies for ACI and only get charged when the instances are running, whereas you pay for a plan in App Service.
* App Service takes care of enabling SSL, whereas you need to take care of SSL yourself in ACI.

> This project is built during a Microsoft hackfest.