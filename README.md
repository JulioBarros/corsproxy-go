# corsproxy-go

A cors proxy written in go to run on your own servers.

There are useful web apps such as [cors anywhere](https://cors-anywhere.herokuapp.com) and other projects based on node or python but I wanted something that could be compiled into a standalone executable.

By default it will listen on port 8080 on the local interface. If you are running it on a server so you can talk to another process, as I am, you'll want to specify the non public lan interface. It is not recommended that you listen on a public interface because then anyone will be able to access you services.

```
usage: corsproxy [-h|--help] [-p|--port "<value>"] [-i|--interface "<value>"]

                 Add CORS support via a proxy to servers that don't support it.

Arguments:

  -h  --help       Print help information
  -p  --port       The port to listen on. Default: 8080
  -i  --interface  The interface to listen on. localhost or 127.0.0.1 or
                   0.0.0.0 is common. Non routable LAN interface (192.168.X.Y
                   or 10.0.X.Y) is useful. Wan/Public is not recommended..
                   Default: 127.0.0.1
```

Once it is running you can access a service by specifying the host, port and path as the path parameter to the proxy server. For example, if the proxy and server you want to access are running locally something like this should work for you.

```
fetch("http://127.0.0.1:3000/127.0.0.1:8332/", data).then(resp => resp.json());

```

If they are both running on a seperate machine. Try something like this where '192.168.0.101' is replaced by the IP address of your server.

```
fetch("http://192.168.0.101:3000/127.0.0.1:8332/", data).then(resp => resp.json());

```

## Getting started

This project requires Go to be installed. On OS X with Homebrew you can just run `brew install go`.

Running it then should be as simple as:

```console
$ go build main.go
$ ./corsproxy-go
```

### Testing

``make test``
