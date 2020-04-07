# Samsung TV REST API

---
    
**NOTE:** While the logic is correct, this code does not currently work for my own TV. I am trying to figure out why, but my Python version of the same code seems to work fine. Still investigating. 

---

A very simple REST API to modern Samsung TVs based on a Websocket connection. I can't promise that this works on all TVs, I built this based on a bit of research and testing what method works well with my own TV.  

## Server

### REST API
The API of the server is defined by the [`api/server.yml`](api/server.yml) Swagger specification. 

#### Examples
* `/status`: to the get connection status
* `/power/<on|off>`: turns the TV on or off
* `/key/POWER`: to turn the TV on or off; this is translated to a `KEY_POWER` key event

A list of KEYs that I tested with my TV: 

* POWER
* MENU
* HOME
* SOURCE

### Configuration
The [`configs/config-template.yml`](configs/config-template.yml) offers a template for the service configuration. 

### Build 

Make sure you that
* you have `dep` installed. Visit https://github.com/golang/dep 
* your `GOPATH` and `GOROOT` environments are set properly.

#### Makefile
There is a [`Makefile`](Makefile) provided that offers a number of targets for preparing, building and running the service. To build and run the service against the [`configs/test.yml`](configs/test.yml) configuration, simply call the `run` target:
```
make clean dep run
```

#### Systemd
I currently have a very basic systemd unit file defined under [`init/api-samsungtv.service`](init/api-samsungtv.service). This can be later improved. 

Before using the service definition, make sure that you go through the file and update the `WorkingDirectory` and `ExecStart` to match your installation.  

## License
The code is published under an [MIT license](LICENSE.md). 

## Contributions
Please report issues or feature requests using Github issues. Code contributions can be done using pull requests. 

## Relevant Repositories
* [github.com/Ape/samsungctl](https://github.com/Ape/samsungctl)
* [github.com/mhvis/samsung-tv-control](https://github.com/mhvis/samsung-tv-control)
* [gist.github.com/freman](https://gist.github.com/freman)

