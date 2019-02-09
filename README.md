# Maintenance microservice / server

Standalone server, serving static file based folder containing a default message that the site is under maintenance.

# Table of contents

- [Overview](#overview)
- [Docker](#docker)
- [Serving your own content](#serving-your-own-content)
- [Acting as Traefik fallback server](#acting-as-traefik-fallback-server)

## Overview

Image runs a minimal golang based server with a very low memory and cpu footprint to serve anything that is present or mounted to /public.

Features:

- Dot files are filtered
- All requests serve 503

## Docker

Example `docker-compose.yml` configuration:

```yml
version: '3.7'

services:
  fallback:
    image: adrianrudnik/maintenance-server
    ports:
      - 8080:80
    # environment:
    #   SERVER_PORT: :80
    # volumes:
    #   - ./my-index.html:/app/index.html
```

## Serving your own content

This image is responsible to serve a HTTP status code 503. Serving any kind of visible content is not. You do not want any kind of risk of an HTTP status code 200 because you placed a nice maintenance favicon.ico into this. This is for 503-ing your clients until your main service comes backup up.

Due to a very tricky thing in Golangs http.ServeFile(s) and http.Fileserver I can only provide you with the mechanic to serve a single html file. There is no clean way to make sure the server will not fall back to HTTP status code 200 if a file is "found" or corrupt the response header by trying to force the status code myself and still rely on std libs security.

Embed images as base64 data streams if needed.

You can override the html file by mounting or placing it to `/app/index.html`

## Acting as Traefik fallback server

If you serve containers through Traefik and you need to have a proper maintenance page during container cycling, you can attach this image as fallback with lower priority like this:

```yml
version: '3.7'

networks:
  proxy:
    external: true

services:
  server:
    image: adrianrudnik/maintenance-server
    networks:
      - proxy6
    labels:
      traefik.enable: true
      traefik.docker.network: proxy6
      traefik.frontend.rule: "HostRegexp:{catchall:.*}"
      traefik.frontend.priority: 1
```

Other containers will have higher priority, so this one will only be considered if the other one went AWOL.
