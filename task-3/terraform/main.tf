# main.tf

provider "google" {
  project = "my-test-gcp-project-for-rancher"
  region  = "us-central1"
  zone    = "us-central1-a"
}

resource "google_compute_network" "default" {
  name = "default"
}

resource "google_compute_instance" "vm_instance" {
  name         = "vm-instance"
  machine_type = "e2-medium"
  zone         = "us-central1-a"
  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-9"
    }
  }
  network_interface {
    network = google_compute_network.default.self_link
    access_config {
    }
  }
  metadata = {
    ssh-keys = "ubuntu:${file("~/.ssh/id_rsa.pub")}"
  }
}


# terraform init
# terraform apply
# ssh ubuntu@<VM_IP_ADDRESS> -i ~/.ssh/id_rsa
