# prancing-walrus

# Perms nded
Scope: Bot
* Read messages/view channels
* Send messages
* Use External emoji
* Use external stickers
* Add reactions
* Use slash comands


# Hosting
tbd

# Dev-ing
## prereq
set CONFIG and BOT_TOKEN env vars for this repo in gitpod user settings. CONFIG should be '../devconfig.json', BOT_TOKEN should be the bot token from Discord. Get yer own. The devconfig has a test app id and specific server id so as to not interfere with the 'production' bot.

does app id need to be per dev also? worry about it when it's not just me. at some point consider loading both configs with merge/overwrite logic to reduce duplication