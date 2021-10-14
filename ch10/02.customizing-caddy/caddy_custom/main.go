package main

import (
	cmd "github.com/caddyserver/caddy/v2/cmd"
	_ "github.com/caddyserver/caddy/v2/modules/standard"

	// Injecting custom modules into Caddy
	_ "caddy_custom/restrictprefix"
	_ "caddy_custom/tomladapter"
)

func main() {
	cmd.Main()
}
