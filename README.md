# Knock 🚪

 `knock` is a simple network reachability tester. I made knock because
 I found myself using `telnet` to test if I could successfully connect
 to remote databases, webservers, and the like fairly often and realized
 that I don't actually need all of the features and historical baggage
 of that software: I just want to know if I can get from point A to B.
 `knock` is designed around this use-case.

## Get Knock

Build from source or take a pre-built binary from the `./dist` folder:

```bash
$ ./build.tcl
```

## Usage

You can call `knock` just like `telnet`:

```bash
$ knock localhost 8080
```

Or with a colon:

```bash
$ knock 192.168.100.1:22
```

Or with option flags:

```bash
$ knock -time 10 -dest example.com -port 80 -net tcp
```

Normally, any problem connecting will be reported to the user
as explicitly as possible. Unlike `telnet`, `knock` will never
prompt the user for interactive input. As soon as a connection
succeeds, `knock` will hang up and report back that success.

## Options

| Option | Description                                             |
| ------ | ------------------------------------------------------- |
| -dest  | Network destination address or name                     |
| -net   | Network type to use (default `tcp`)                     |
| -port  | Port number to knock on destination (default `80`)      |
| -time  | Time to wait in seconds before giving up (default `15`) |

## Networks
Supported network types are `tcp`, `tcp4 `, `tcp6 `, `udp`, `udp4 `, `udp6 `,
`ip`, `ip4 `, `ip6 `, `unix`, `unixgram`, and `unixpacket`. The network types
with `4` in their name are IPv4-only, and those with `6` are IPv6-only.

## Miscellaneous

Knock is open source-software distributed under the 0BSD license.
To report bugs or view source code, see https://www.github.com/nat-418/knock.

