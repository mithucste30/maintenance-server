# Maintenance microservice / server

Standalone server, serving static file based folder containing a default message that the site is under maintenance.

# Table of contents

- [Overview](#overview)

## Overview

Image runs a minimal golang based server with a very low memory and cpu footprint to serve anything that is present or mounted to /public.

- Dot files are filtered
- All requests serve 503

Example `docker-compose.yml` configuration:

```yml
version: '3.7'

services:
  mjml:
    image: adrianrudnik/maintenance-server
    ports:
      - 8080:80
    # or expose any other port by using
    # environment:
    #   SERVER_PORT: :80
```
