# Taiga to Mattermost Integration
## Usage
1. Build a binary from source or use one from Releases page and put it on your server.
1. Adjust `config.json` to your needs (select language and add mattermost bot token).
1. Run the binary or use [OpenRC script](etc/init.d/taiga-mm) to start it as a service.
1. In your Taiga instance go to `Settings` -> `INTEGRATIONS` -> `WEBHOOKS` and add a webhook:
- Name: `Mattermost`
- URL: `http://<server-address>/taiga-integration/channel/<mattermost-channel-name>`
- Service secret key: `<random-garbage>`

## Development
Debug run:
``` bash
go run main.go
```
Build
``` bash
go build -ldflags '-linkmode external -extldflags "-static"'
```
Production run:
``` bash
./taiga-mm
```

## Config

| Field             | Explanation                                           | Example
| ----------------- | ----------------------------------------------------- | ---------------------------------------------
| mattermost_token  | Mattermost API token                                  | `123456abc78901de234567fghi8jk9l0`
| usernames         | list of taiga usernames with their mattermost analogs | `{"taiga_username1": "mattermost_username1"}`
| host              | integration server ip                                 | `0.0.0.0`
| port              | integration server port                               | `8080`
| language          | messages language                                     | `en`
| mattermost_server | Mattermost server address                             | `mattermost.example.com`