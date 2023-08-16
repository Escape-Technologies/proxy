# Escape proxy

A simple proxy that can be used to allow escape to connect to some APIs.

## Requirements

To use this proxy, you need to have an [Escape](https://escape.tech) account.

Before installing the proxy make shure you have the following:

- `ESCAPE_ORGANIZATION_ID` : Your Escape organization id (can be found in [organization/settings](https://app.escape.tech/organization/settings/)).
- `ESCAPE_API_KEY` : The API key of your Escape account (can be found in [user/profile](https://app.escape.tech/user/profile/)).

## Install

You have multiple options to install the proxy:

- Docker image (covered in this tutorial)
- From source ([go to releases](https://github.com/Escape-Technologies/proxy/releases/latest))

You now need to run the proxy with the following environment variables:

- `ESCAPE_ORGANIZATION_ID`: Your organization id.
- `ESCAPE_API_KEY` : Your API key.
- `PORT` : The port on which the proxy will listen (default: `8080`).

```bash
docker run -it --rm --name escape-proxy \
    -e ESCAPE_ORGANIZATION_ID=xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx \
    -e ESCAPE_API_KEY=xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx \
    -p 8080:8080 escapetech/proxy
```

You can find in the example folder more deployment examples.
Feel free to contribute and add your own.
