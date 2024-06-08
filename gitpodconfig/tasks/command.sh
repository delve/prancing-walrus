chmod ugo+w /var/run/docker.sock
gcloud components update --quiet
project=$(gcloud projects list | grep TestWalrus | awk '{ print $1 }')
gcloud config set project $project
