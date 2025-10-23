# This code is compatible with Terraform 4.25.0 and versions that are backwards compatible to 4.25.0.
# For information about validating this Terraform code, see https://developer.hashicorp.com/terraform/tutorials/gcp-get-started/google-cloud-platform-build#format-and-validate-the-configuration

resource "google_compute_instance" "walrus-1" {
  boot_disk {
    auto_delete = true
    device_name = "walrus-1"

    initialize_params {
      image = "projects/cos-cloud/global/images/cos-stable-109-17800-66-54"
      size  = 10
      type  = "pd-standard"
    }

    mode = "READ_WRITE"
  }

  can_ip_forward      = false
  deletion_protection = false
  enable_display      = false

  labels = {
    container-vm = "cos-stable-109-17800-66-54"
    goog-ec-src  = "vm_add-tf"
  }

  machine_type = "e2-micro"

  # revoked API tokens inline. this file is out of date anyway and shouldn't be used.
  # tokens are stored in GCP secret manager for runtime access.
  metadata = {
    gce-container-declaration = "spec:\n  containers:\n  - name: walrus-1\n    image: us-central1-docker.pkg.dev/prancingwalrus/prancing-walrus/prancing-walrus:v0.03\n    env:\n    - name: BOT_TOKEN\n      value: MTE2OTMzMjA4NDgxMzM0ODg2NQ.GzLVAi.LtbFRgtG7DPavXcfBVUNEugui5eDfj0c079-L4\n    - name: APIKey\n      value: AIzaSyCYHifpsoQfDgtFDJD590Z_KwProE7xvwM\n    stdin: false\n    tty: false\n  restartPolicy: Always\n# This container declaration format is not public API and may change without notice. Please\n# use gcloud command-line tool or Google Cloud Console to run Containers on Google Compute Engine."
  }

  name = "walrus-1"

  network_interface {
    access_config {
      network_tier = "PREMIUM"
    }

    subnetwork = "projects/prancingwalrus/regions/us-central1/subnetworks/default"
  }

  scheduling {
    automatic_restart   = true
    on_host_maintenance = "MIGRATE"
    preemptible         = false
    provisioning_model  = "STANDARD"
  }

  service_account {
    email  = "walrus-sheet-access@prancingwalrus.iam.gserviceaccount.com"
    scopes = ["https://www.googleapis.com/auth/cloud-platform"]
  }

  shielded_instance_config {
    enable_integrity_monitoring = true
    enable_secure_boot          = false
    enable_vtpm                 = true
  }

  zone = "us-central1-a"
}
