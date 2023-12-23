https://gitpod.io/#https://github.com/delve/prancing-walrus
# prancing-walrus
## TODO
### Before 1.0
* automate version tagging & publishing (makefile)
* automate deployment (makefile)
* replace as many Panics as possible with proper error handling
* on /refreshassignments add a discord name check and warn on incorrect names.
* Add backoff-retry logic around accessing the spreadsheet
* Look into backoff-retry in the discord commands, to recover from discord outages
* find a way to publish logs to walrus wranglers
### After 1.0
* extra tabs for adhoc notes. EG: what war is next, and a generic note.
* reorganize with https://github.com/FedorLap2006/disgolf/blob/master/examples/modules/main.go ?
* automatic recaching once when someone is missing an assignment (if older than (3?)hours)
* automatic recaching every 24 hours (if older than (3?)hours)
* backoff retry on recaching
* grab the roleid for `manager` dynamically
* make log messages show correct caller, not `"caller":"log/log.go:20"`
* make all config values read from env and default to configfile if not found
* enable user updates. include failsafe for values out of bounds
* in config the values pulled from env should be tested, if len == 0 panic
* setup a sidechannel method to restart the host VM. consider ways to manage container version through it as well

# Discord perms nded
Scope: Bot
* Read messages/view channels
* Send messages
* Use External emoji
* Use external stickers
* Add reactions
* Use slash comands
# Dev-ing
## prereq
set CONFIG, APIKey, and BOT_TOKEN env vars for this repo in gitpod user settings. CONFIG should be '../devconfig.json', BOT_TOKEN should be the bot token from Discord. Get yer own. The devconfig has a test app id and specific server id so as to not interfere with the 'production' bot.

does app id need to be per dev also? worry about it when it's not just me. at some point consider loading both configs with merge/overwrite logic to reduce duplication

PROD BOT_TOKEN and sheets APIKey stored in GCP Secrets Manager
## Run
`F5` for interactive debugging in VSCode with the 'Launch Package' profile
`make run` to compile and execute local code
`make image` to create a buildpack image
`make runimage` to run a local image
## Deploy
`make publish` will login to gcp and push the file. NOTE: this will only work if you have permissions in the relevant GCP project, which you probably don't have
# Hosting
GCP GCE. cloud run was too expensive.

* create project
* enable APIs
* * GCE
* * Cloud logging
* * Artifact Registry
* create Artifact Registy docker registry colocated with GCE instance (usc1) with a cleanup policy keep latest 5 images. immutable tags seems to prevent deletion, so don't enable it to avoid ballooning storage & cost


deploy container 
us-central1-docker.pkg.dev/prancingwalrus/prancing-walrus/prancing-walrus:v0.01
container env
BOT_TOKEN
<insert BOT_TOKEN here>
APIKey
<insert APIKey here>


* create VM
* * sudo apt install git
* * curl --output /tmp/go1.21.3.linux-amd64.tar.gz https://dl.google.com/go/go1.21.3.linux-amd64.tar.gz
* * sudo rm -rf /usr/local/go && sudo tar -C /usr/local -xzf /tmp/go1.21.3.linux-amd64.tar.gz
* * git clone https://github.com/delve/prancing-walrus.git
* * git checkout <prodbranch>
* set startup script TODO: secrets management yo?
```
export BOT_TOKEN=<token>
export APIKey=<key>
export GOCACHE='/home/delve202/.cache/go-build'
export GOMODCACHE='/home/delve202/go/pkg/mod'
export GOPATH='/home/delve202/go'
cd /home/delve202/prancing-walrus/walrusbot
git config --global --add safe.directory /home/delve202/prancing-walrus
echo "Git Pull" > /tmp/log.txt
echo "---------" >> /tmp/log.txt
git pull &>> /tmp/log.txt
echo "---------" >> /tmp/log.txt
echo "Go Env" >> /tmp/log.txt
echo "---------" >> /tmp/log.txt
/usr/local/go/bin/go env &>> /tmp/log.txt
echo "---------" >> /tmp/log.txt
echo "Go Run" >> /tmp/log.txt
echo "---------" >> /tmp/log.txt
/usr/local/go/bin/go run . &>> /tmp/log.txt
```
* reboot or manually exec startup script
