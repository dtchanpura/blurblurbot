# Starting up the bot

Configuration is very minimal currently. Only one Environment variable is needed for starting the bot.

## From Shell

* Export the variable and execute the binary from latest release page on Github

```bash
export TG_BOT_TOKEN="ABCD1234"
./path/to/blurblurbot
```

## From `systemd`

* A .service file has been added in etc folder which can be **edited** and then copied to ~/.config/systemd/user
* Then reload the systemd daemon
```
systemctl --user daemon-reload
```
* Then start or stop the service with following commands
```
systemctl --user start blurblurbot
systemctl --user stop blurblurbot
```

## From Docker Image

* Build the image.

```bash
cd path/to/repository
docker build -t "dtchanpura/blurblurbot:0.1.2" .
# .... Wait for it to build.
docker run --name blurblurbot --env TG_BOT_TOKEN="ABCD1234" -p 18080:8080 dtchanpura/blurblurbot:0.1.2
```
