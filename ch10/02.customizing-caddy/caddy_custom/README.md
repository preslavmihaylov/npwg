# Steps

Build the custom binary \w added modules/adapters:
```
go build
```

Make sure the modules are setup properly:
```
./caddy_custom list-modules | grep "toml\|restrict_prefix"
```

Start the server with the toml configuration:
```
./caddy_custom start --config caddy.toml --adapter toml
```

Run the upstream backend server:
```
go run backend/*.go
```

Open `http://localhost:2020` in your browser. Your file server will return the `index.html`. You should see a gopher rendered.

Try accessing a restricted file - `http://localhost:2020/files/.secret`. You should get a 404.
