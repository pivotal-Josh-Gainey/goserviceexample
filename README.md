# goserviceexample

This is a simple hello world go app compiled with go v1.19 that can be pushed to cloud foundry as an application.

This application has 2 endpoints:<br>
```/ ``` returns "Hello World!" <br>
```/get-random-number``` returns a random number<br>

It is possible to deploy this app on a custom port with an environment variable in the manifest. 
Specify TCPPORT env var with the port desired for app to listen on. In the provided example manifest, this env var is set to 8095.

1 - clone this repo:
```
https://github.com/pivotal-Josh-Gainey/goserviceexample.git
```

2 - `cd` into `goserviceexample` folder, make desired edits to manifest if applicable, and run `cf push`

This pushes the application with a process healthcheck.



