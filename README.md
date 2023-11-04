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
BOT_TOKEN=<token>
cd /home/delve202/prancing-walrus/walrusbot
/usr/local/go/bin/go run .
```
* reboot or manually exec startup script

rest call to gen VM
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
  "fingerprint": "Ha8y9KIlJGo=",
  "id": "2100367559267969676",
  "keyRevocationActionType": "NONE",
  "kind": "compute#instance",
  "labelFingerprint": "42WmSpB8rSM=",
  "lastStartTimestamp": "2023-11-04T07:25:48.865-07:00",
  "machineType": "projects/prancingwalrus/zones/us-central1-a/machineTypes/e2-micro",
  "metadata": {
    "fingerprint": "swIPQWFZSUE=",
    "kind": "compute#metadata",
    "items": [
      {
        "value": "delve202:ecdsa-sha2-nistp256 AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBO56TzPbk7tJHLyDuN3aKzD0feIPiEs+3bV/5KKVxI+R2WL7EFVlUKuAJ/vRZdfJQ9sdalNsIZLuS3jP9SgbyBY= google-ssh {\"userName\":\"delve202@gmail.com\",\"expireOn\":\"2023-11-04T14:29:30+0000\"}\ndelve202:ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAEaxtzuwxwNrW3NPMk6tUnLz5xKx5bO8yDaOPu15ts/XTSAr60uXy6VUmBx6qF8V9c+qUwVgz7fFK3904ZOdcxxb+OZfAeKhnQ9uFqFlXHO2Ke7mcl7KB+S1S6uhLnEIp6jikA8h+1/j4i8U2MKqBZUuYXwbfi2PXE1I95Qkc960eXnGlV0YuStYVfa/p5FZmoGq6tm4iJK0ksoXz1WLeP/TtlpdBuLDiDzOJWu84a8CEYPxla6HPPxLXCXPNwDjBvmZ8So31uzdfIXWmSwTk8vWaWISz/mPvBm6Vtu47gQ79Pnx9HAIAetpYHMfa3/LgtL3Hh+ez4v6XiNDjg9Mwp8= google-ssh {\"userName\":\"delve202@gmail.com\",\"expireOn\":\"2023-11-04T14:29:48+0000\"}",
        "key": "ssh-keys"
      },
      {
        "value": "BOT_TOKEN=MTE2OTMzMjA4NDgxMzM0ODg2NQ.GS9fQU.gCfolVuWGpi0WdbTlBeIAhhk4oCcElspgczMYY\ncd /home/delve202/prancing-walrus/walrusbot\n/usr/local/go/bin/go run .",
        "key": "startup-script"
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
  "resourceStatus": {},
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
set CONFIG and BOT_TOKEN env vars for this repo in gitpod user settings. CONFIG should be '../devconfig.json', BOT_TOKEN should be the bot token from Discord. Get yer own. The devconfig has a test app id and specific server id so as to not interfere with the 'production' bot.

does app id need to be per dev also? worry about it when it's not just me. at some point consider loading both configs with merge/overwrite logic to reduce duplication