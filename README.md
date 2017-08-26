# Isomorphic Test

How does a single binary act as both a client and a server?

That's the question this test application sets out to answer.

## CLI

### User commands

The user will have a single useful command: `iso generate`, which will return an UUID.

### Internal commands

The server lifecycle can be controlled with a pair of commands: `start` and `stop`.

## Startup

Initially, there is no instance of `iso` running.  The user makes an initial request:

`iso generate`

The client will make a RPC request to a default port.

When the server does not respond, it will spawn a new process, starting the server instance.  The server will open itself to RPC commands via the default port.

The client will then reissue the RPC request to get the generated UUID.

If there is another application at the default port, the client must error out; the server cannot be started.

Alternatley, the server can be started deliberately with the `iso start` command.

## Shutdown

The server should shut down after a period of inactivity.  Alternately, the server can be stopped proactively by issuing an `iso stop` command.