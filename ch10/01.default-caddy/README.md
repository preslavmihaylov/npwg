# Instructions
First install caddy from their downloads page for your platform.

Then, run it with default settings:
```
caddy start
```

Afterwards, change its real-time configuration:
```
curl localhost:2019/load -X POST -H "Content-Type: application/json" -d '
{
  "apps": {
    "http": {
      "servers": {
        "hello": {
          "listen": [
            "localhost:2020"
          ],
          "routes": [
            {
              "handle": [
                {
                  "handler": "static_response",
                  "body": "Hello, world!"
                }
              ]
            }
          ]
        }
      }
    }
  }
}'
```

Request a portion of the configuration via caddy's api:
```
curl localhost:2019/config/apps/http/servers/hello/listen
```

Finally, make a call to the handler configured above:
```
curl localhost:2020
```

Add another handler for `:2030`:
```
curl localhost:2019/config/apps/http/servers/test -X POST -H "Content-Type: application/json" -d '
{
  "listen": [
    "localhost:2030"
  ],
  "routes": [
    {
      "handle": [
        {
          "handler": "static_response",
          "body": "Welcome to my temporary test server."
        }
      ]
    }
  ]
}'
```

Make a call to the handler configured above:
```
curl localhost:2030
```

Delete the previous handler:
```
curl localhost:2019/config/apps/http/servers/test -X DELETE
```

Make sure it's down:
```
curl localhost:2019
```
