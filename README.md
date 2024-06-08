https://gitpod.io/#https://github.com/delve/prancing-walrus
# prancing-walrus
## TODO
* fix debug launch profile so it runs main.go for whichever module the active file is part of
### Before 1.0
* automate version tagging & publishing (makefile)
* automate deployment (makefile)
* replace as many Panics as possible with proper error handling
* Extract user input sanitizing into the data model somehow
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
* make log messages show correct caller, not `"caller":"log/log.go:20"`
* make all config values read from env and default to configfile if not found
* enable user updates. include failsafe for values out of bounds
* in config the values pulled from env should be tested, if len == 0 panic
* setup a sidechannel method to restart the host VM. consider ways to manage container version through it as well
* several spots that assume only one guild (EG ctx.Session.State.Guilds[0]). these should be fixed

# Discord perms nded
Scopes: Bot + applications.commands
* General Perms
* * Manage Roles
* * Read messages/view channels
* Text Perms
* * Send messages
* * Create public threads
* * Create private threads
* * Send messages in threads
* * Manage messages
* * Manage threads
* * Embed links
* * Read message history
* * Use External emoji
* * Use external stickers
* * Add reactions
* * Use slash comands
* * Use external emoji
* * Use external stickers
* * Add reactions
* * Use slash commands

Invite URL for test app

https://discord.com/oauth2/authorize?client_id=1170099827611291761&permissions=534992219200&scope=bot+applications.commands

Invite URL for prod app

https://discord.com/oauth2/authorize?client_id=1169332084813348865&permissions=534992219200&scope=bot+applications.commands

# Dev-ing
## prereq
set CONFIG, APIKey, and BOT_TOKEN env vars for this repo in gitpod user settings. CONFIG should be '../devconfig.json', BOT_TOKEN should be the bot token from Discord. Get yer own. The devconfig has a test app id and specific server id so as to not interfere with the 'production' bot.

does app id need to be per dev also? worry about it when it's not just me. at some point consider loading both configs with merge/overwrite logic to reduce duplication

PROD BOT_TOKEN and sheets APIKey stored in GCP Secrets Manager

* to enable local service account impersonation through ADC for dev work:
* * mymail=the email address of your gcp user id
* * prj=`gcloud config get-value project`;useremail="user:${mymail}"
* * to enable SA token creation `gcloud projects add-iam-policy-binding ${prj} --member=${useremail} --role="roles/iam.serviceAccountTokenCreator"`
* * note that this can apparently take up to ten minutes to actually apply, which is incredibly frustrating when you're debugging something.
* * also note that apparently 'roles/owner' does NOT include this, for whatever reason.

## DB modules
During workspace setup the gsheets and sheetdb module repos are cloned into /workspace.The workspace's `go.work` file include replacement directives for these modules such that you will always load those from your local copy (which should be checked out at the version tag from walrusbot's `go.mod` file, HOWEVER TODO: need to automate which version tag is checked out). Therefore if there is a bug in one of these modules you can edit it locally to troubleshoot and fix it, tag and push the new version, and finally update the version in the `go.mod` files of the prancing-walrus repo.

## Run
`F5` for interactive debugging in VSCode with the 'Launch Package' profile
`make run` to compile and execute local code
`make image` to create a buildpack image
`make runimage` to run a local image
## Deploy
`make publish` will login to gcp and push the file. NOTE: this will only work if you have permissions in the relevant GCP project, which you probably don't have
# Hosting
GCP GCE. cloud run was too expensive.

TODO: consider Terraform :(

* create project
* enable APIs
* * GCE
* * Cloud logging
* * Artifact Registry
* * Identity and Access Management (IAM) API
* create Artifact Registy docker registry colocated with GCE instance (usc1) with a cleanup policy keep latest 5 images. immutable tags seems to prevent deletion, so don't enable it to avoid ballooning storage & cost
* create service account `walrus-sheet-access`
* * --command to create SA--
* * add permissions
* * * prj=`gcloud config get-value project`;saId="serviceAccount:walrus-sheet-access@prancingwalrus.iam.gserviceaccount.com"
* * * for logging `gcloud projects add-iam-policy-binding ${prj} --member=${saId} --role=roles/logging.logWriter`
* * * for artifact registry `gcloud projects add-iam-policy-binding ${prj} --member=${saId} --role=roles/artifactregistry.reader`
* * * for secrets manager read access `gcloud projects add-iam-policy-binding ${prj} --member=${saId} --role="roles/secretmanager.secretAccessor"`
* * * te generate SA key for sheets access `gcloud projects add-iam-policy-binding ${prj} --member=${saId} --role="roles/iam.serviceAccountKeyAdmin"`

* * * add SA email to gsheet permissions as editor
* setup secretmanager
* * add bot token
* * add API key (this is for sheets, probably don't need it anymore?!?)
* * create key for SA from IAM Service Accounts > ...s > Manage keys > Add Key > Create New Key > JSON, create
* * upload key file to secret manager, delete local file.
* create VM
* * select deploy container 
```
us-central1-docker.pkg.dev/prancingwalrus/prancing-walrus/prancing-walrus:v0.01
```
* * service account: walrus-sheet-access
* * set custom metadata google-logging-enabled	true
