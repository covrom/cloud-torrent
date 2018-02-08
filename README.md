<img src="https://user-images.githubusercontent.com/633843/32198822-e59a0fc4-be1d-11e7-9b92-03ce17ba05ba.png" alt="screenshot"/>

**Cloud torrent** is a a self-hosted remote torrent client, written in Go (golang). You start torrents remotely, which are downloaded as sets of files on the local disk of the server, which are then retrievable or streamable via HTTP.

**UPDATED VERSION**
Directories to watching is placed into settings json as a "WatchDirs" array of path's.

### Features

* Single binary
* Cross platform
* Embedded torrent search
* Real-time updates
* Directories watching for .torrent files with fsnotify
* Mobile-friendly
* Fast [content server](http://golang.org/pkg/net/http/#ServeContent)

See [Future Features here](#future-features)

### Install

**Binaries**

[![Releases](https://img.shields.io/github/release/covrom/cloud-torrent.svg)](https://github.com/covrom/cloud-torrent/releases) [![Releases](https://img.shields.io/github/downloads/covrom/cloud-torrent/total.svg)](https://github.com/covrom/cloud-torrent/releases)

See [the latest release](https://github.com/covrom/cloud-torrent/releases/latest) or download and install it now with

**Docker**

``` sh
$ docker run -d -p 3000:3000 -v /path/to/my/downloads:/downloads covrom/cloud-torrent
```

**Source**

*[Go](https://golang.org/dl/) is required to install from source*

``` sh
$ go get -v github.com/covrom/cloud-torrent
```

**VPS**

[Digital Ocean](https://m.do.co/c/011fa87fde07)

  1. Sign in
  2. "Create Droplet"
  3. "One-Click Apps"
  4. "Docker X.X.X on X.X"
  5. Choose server size ("$5/month" is enough)
  6. Choose server location
  7. **OPTIONAL** Add your SSH key
  8. "Create"
  9. You will be emailed the server details (`IP Address: ..., Username: root, Password: ...`)
  10. SSH into the server using these details (Windows: [Putty](https://the.earth.li/~sgtatham/putty/latest/x86/putty.exe), Mac: Terminal)
  11. Follow the prompts to set a new password
  12. Run `cloud-torrent` with:

    docker run --name ct -d -p 63000:63000 \
      --restart always \
      -v /root/downloads:/downloads \
      covrom/cloud-torrent --port 63000

  13. Visit `http://<IP Address from email>:63000/`
  14. **OPTIONAL** In addition to `--port` you can specify the options below

[Vultr](http://www.vultr.com/?ref=6947403-3B)

* Sign in
* Follow the DO tutorial above, very similar steps ("Applications" instead of "One-Click Apps")
* Offers different server locations

[AWS](https://aws.amazon.com)

**Heroku**

Heroku is no longer supported

### Usage

```
$ cloud-torrent --help

  Usage: cloud-torrent [options]

  Options:
  --title, -t        Title of this instance (default Cloud Torrent, env TITLE)
  --port, -p         Listening port (default 3000, env PORT)
  --host, -h         Listening interface (default all)
  --auth, -a         Optional basic auth in form 'user:password' (env AUTH)
  --config-path, -c  Configuration file path (default cloud-torrent.json)
  --key-path, -k     TLS Key file path
  --cert-path, -r    TLS Certicate file path
  --log, -l          Enable request logging
  --open, -o         Open now with your default browser
  --help
  --version, -v

  Version:
    0.X.Y

  Read more:
    https://github.com/covrom/cloud-torrent

```

### Notes

This project is the rewrite of the original [Node version](https://github.com/jpillora/node-torrent-cloud).

![overview](https://docs.google.com/drawings/d/1ekyeGiehwQRyi6YfFA4_tQaaEpUaS8qihwJ-s3FT_VU/pub?w=606&h=305)

Credits to @anacrolix for https://github.com/anacrolix/torrent

Copyright (c) 2017 Jaime Pillora

Copyright (c) 2018 Roman TSovanyan