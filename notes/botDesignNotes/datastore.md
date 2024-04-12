# links
## links maybe helpful for sheets access
* https://robocorp.com/docs/development-guide/google-sheets/interacting-with-google-sheets
* https://stackoverflow.com/questions/27067825/how-to-access-google-spreadsheets-with-a-service-account-credentials
* bring your own private key to GCP SAs https://xebia.com/blog/how-to-create-your-own-google-service-account-key-file/
* functional how to, but relies on insecure key handling and wraps the whole thing in an HTTP request/response package that makes it hard to understand. https://thriveread.com/golang-google-sheets-and-spreadsheet-api/

# decisions
data storage must be free

feels like using BigQuery is just too much more cumbersome. have to be able to sync back and forth (so mgrs can crunch & review numbers)

would prefer a single data store (making bigquery sync even more important, and problematic)

# considering GCP Datastore
free tier might be large enough

still a problem with import/export for managers review, more so than a sheets based DB which they can just query.

HOWEVER, if extended to additional clubs there's less concern over access

could look into dumping CSV files for import?

# notes
found a sheets DAO module, the API module provided by google is ass. DAO module had to be forked (looks like someone's old abandonded PoC project) and adjusted to allow more auth flexibility. it's now one of my repos.

# SA keys
This provides Go code that supposedly generates a key for the SA (auth provided by ADC both locally and in GCE).

https://cloud.google.com/iam/docs/keys-create-delete#iam-service-account-keys-create-go

SA key is generated dynamically by the bot on startup. deletion is deferred, and does not occur in some shutdown scenarios, EG terminating an interactive debug session. automatically deletes all SA keys at startup. if we start sharding this will be problematic

Handling keys dynamically via CLI is no longer necessary but is kept here for convenient reference:
```
ACCOUNT=walrus-sheet-access@prancingwalrus.iam.gserviceaccount.com
gcloud iam service-accounts keys list --iam-account=$ACCOUNT
gcloud iam service-accounts keys create /tmp/test.json --iam-account=$ACCOUNT
gcloud iam service-accounts keys delete --quiet --iam-account=$ACCOUNT KEY_ID
```
KEY_ID can be found in the JSON file: private_key_id