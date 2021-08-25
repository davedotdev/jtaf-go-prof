provider "junos-device" {
    host = "vqfx01.simpledemo.net"
    port = 830
    username = "root"
    password = "Passw0rd"
    sshkey = ""
}

module "qfx_1" {
  source = "./qfx_1"

  providers = {
    junos-device = junos-device
  }
}

resource "junos-device_commit" "commit-main" {
  resource_name = "commit"
  depends_on = [module.qfx_1]
}
