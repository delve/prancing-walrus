# prancing-walrus

# TODO
* log to gcloud. and maybe also a file, but roll it in the code.
* adjust spreadsheet integration to role and gathering columns
* make log messages show correct caller, not `"caller":"log/log.go:20"`
* enable recaching
* enable user updates. include failsafe for values out of bounds
* make all config values read from env and default to configfile if not found

# Perms nded
Scope: Bot
* Read messages/view channels
* Send messages
* Use External emoji
* Use external stickers
* Add reactions
* Use slash comands

# Hosting
GCP GCE. cloud run was too expensive.

* create project
* enable GCE API
* create VM
* * sudo apt install git
* * curl --output /tmp/go1.21.3.linux-amd64.tar.gz https://dl.google.com/go/go1.21.3.linux-amd64.tar.gz
* * sudo rm -rf /usr/local/go && sudo tar -C /usr/local -xzf /tmp/go1.21.3.linux-amd64.tar.gz
* * git clone https://github.com/delve/prancing-walrus.git
* * git checkout <prodbranch>
* set startup script TODO: secrets management yo?
```
export BOT_TOKEN=<token>
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

rest call to gen VM. note the token replace
```
{
  "canIpForward": false,
  "confidentialInstanceConfig": {
    "enableConfidentialCompute": false
  },
  "cpuPlatform": "Intel Broadwell",
  "creationTimestamp": "2023-11-04T07:25:41.305-07:00",
  "deletionProtection": false,
  "description": "",
  "disks": [
    {
      "architecture": "X86_64",
      "guestOsFeatures": [
        {
          "type": "UEFI_COMPATIBLE"
        },
        {
          "type": "VIRTIO_SCSI_MULTIQUEUE"
        },
        {
          "type": "GVNIC"
        },
        {
          "type": "SEV_CAPABLE"
        }
      ],
      "type": "PERSISTENT",
      "licenses": [
        "projects/debian-cloud/global/licenses/debian-11-bullseye"
      ],
      "deviceName": "instance-1",
      "autoDelete": true,
      "source": "projects/prancingwalrus/zones/us-central1-a/disks/instance-1",
      "index": 0,
      "boot": true,
      "kind": "compute#attachedDisk",
      "mode": "READ_WRITE",
      "interface": "SCSI",
      "diskSizeGb": "10"
    }
  ],
  "displayDevice": {
    "enableDisplay": false
  },
  "fingerprint": "yzZ_LFQrweM=",
  "id": "2100367559267969676",
  "keyRevocationActionType": "NONE",
  "kind": "compute#instance",
  "labelFingerprint": "42WmSpB8rSM=",
  "lastStartTimestamp": "2023-11-04T09:25:47.148-07:00",
  "lastStopTimestamp": "2023-11-04T09:24:34.804-07:00",
  "machineType": "projects/prancingwalrus/zones/us-central1-a/machineTypes/e2-micro",
  "metadata": {
    "fingerprint": "vJ_4fb_h8cY=",
    "kind": "compute#metadata",
    "items": [
      {
        "value": "export BOT_TOKEN=<BOT_TOKEN>\nexport GOCACHE='/home/delve202/.cache/go-build'\nexport GOMODCACHE='/home/delve202/go/pkg/mod'\nexport GOPATH='/home/delve202/go'\ncd /home/delve202/prancing-walrus/walrusbot\ngit config --global --add safe.directory /home/delve202/prancing-walrus\necho \"Git Pull\" > /tmp/log.txt\necho \"---------\" >> /tmp/log.txt\ngit pull &>> /tmp/log.txt\necho \"---------\" >> /tmp/log.txt\necho \"Go Env\" >> /tmp/log.txt\necho \"---------\" >> /tmp/log.txt\n/usr/local/go/bin/go env &>> /tmp/log.txt\necho \"---------\" >> /tmp/log.txt\necho \"Go Run\" >> /tmp/log.txt\necho \"---------\" >> /tmp/log.txt\n/usr/local/go/bin/go run . &>> /tmp/log.txt",
        "key": "startup-script"
      },
      {
        "value": "delve202:ecdsa-sha2-nistp256 AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBJ0/BUNaYxqwnRfsK+MQ6YkQL2mslXfP8I82DMVS21EHimv+jtI0F2D43Tfrj5r9CJTUq1a93G4SbsTll7Ea01w= google-ssh {\"userName\":\"delve202@gmail.com\",\"expireOn\":\"2023-11-04T16:38:28+0000\"}\ndelve202:ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDXk1AMD1MLD5bpbW55rd5OUbLeSuNsg/PmYABQ6vruFEb83uXegAH+R+8wh2Mjk/Gi3yykbc7jn8wv2tUdYYefhQH+peR44v4q/jHgbBlKm0N2FGwGXgb5l2CTyzfNgl8spBf0C3Bo08Gzx1dyNcMqekLdsfr8/4vX0R8T02crgcyOSVHnaa+3VEdZFhlX42RD9bOe84okDVpwgS7MAggw7fE53dJqAl7jJiSS3DbE209/ZBM0gPw8a3928k1fEd/NgzUTmGS9m75sXKj3fxbxS3WNIMYl8eyJ5EOlmHbBtmCr0UUmcVRhZDB8Nx4Mm1aIpDBbvgk0WDQ2TeyMlz4P google-ssh {\"userName\":\"delve202@gmail.com\",\"expireOn\":\"2023-11-04T16:38:46+0000\"}",
        "key": "ssh-keys"
      }
    ]
  },
  "name": "instance-1",
  "networkInterfaces": [
    {
      "stackType": "IPV4_ONLY",
      "name": "nic0",
      "network": "projects/prancingwalrus/global/networks/default",
      "accessConfigs": [
        {
          "name": "External NAT",
          "type": "ONE_TO_ONE_NAT",
          "natIP": "35.223.149.173",
          "kind": "compute#accessConfig",
          "networkTier": "PREMIUM"
        }
      ],
      "subnetwork": "projects/prancingwalrus/regions/us-central1/subnetworks/default",
      "networkIP": "10.128.0.2",
      "fingerprint": "8mlsfLpB9xE=",
      "kind": "compute#networkInterface"
    }
  ],
  "reservationAffinity": {
    "consumeReservationType": "ANY_RESERVATION"
  },
  "resourceStatus": {
    "scheduling": {}
  },
  "scheduling": {
    "onHostMaintenance": "MIGRATE",
    "provisioningModel": "STANDARD",
    "automaticRestart": true,
    "preemptible": false
  },
  "selfLink": "projects/prancingwalrus/zones/us-central1-a/instances/instance-1",
  "serviceAccounts": [
    {
      "email": "883671929937-compute@developer.gserviceaccount.com",
      "scopes": [
        "https://www.googleapis.com/auth/devstorage.read_only",
        "https://www.googleapis.com/auth/logging.write",
        "https://www.googleapis.com/auth/monitoring.write",
        "https://www.googleapis.com/auth/servicecontrol",
        "https://www.googleapis.com/auth/service.management.readonly",
        "https://www.googleapis.com/auth/trace.append"
      ]
    }
  ],
  "shieldedInstanceConfig": {
    "enableSecureBoot": false,
    "enableVtpm": true,
    "enableIntegrityMonitoring": true
  },
  "shieldedInstanceIntegrityPolicy": {
    "updateAutoLearnPolicy": true
  },
  "shieldedVmConfig": {
    "enableSecureBoot": false,
    "enableVtpm": true,
    "enableIntegrityMonitoring": true
  },
  "shieldedVmIntegrityPolicy": {
    "updateAutoLearnPolicy": true
  },
  "startRestricted": false,
  "status": "RUNNING",
  "tags": {
    "items": [
      "https-server"
    ],
    "fingerprint": "nNZ0SA7CJyk="
  },
  "zone": "projects/prancingwalrus/zones/us-central1-a"
}
```

# Dev-ing
## prereq
set CONFIG, APIKey, and BOT_TOKEN env vars for this repo in gitpod user settings. CONFIG should be '../devconfig.json', BOT_TOKEN should be the bot token from Discord. Get yer own. The devconfig has a test app id and specific server id so as to not interfere with the 'production' bot.

does app id need to be per dev also? worry about it when it's not just me. at some point consider loading both configs with merge/overwrite logic to reduce duplication