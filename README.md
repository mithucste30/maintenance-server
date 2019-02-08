# Maintenance microservice / server

Standalone server, serving static file based folder containing a default message that the site is under maintenance.

# Table of contents

- [Overview](#overview)
- [Docker](#docker)
- [Serving your own content](#serving-your-own-content)

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
  mjml:
    image: adrianrudnik/maintenance-server
    ports:
      - 8080:80
    # environment:
    #   SERVER_PORT: :80
    # volumes:
    #   - ./my-own-content:/app/public
```

## Serving your own content

Just override the /app/public folder with any folder full of statics, serving a possible index.html as entrypoint. Look at the docker chapter for the commented part about it
