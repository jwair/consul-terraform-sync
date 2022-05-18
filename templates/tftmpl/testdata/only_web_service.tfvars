# This file is generated by Consul-Terraform-Sync.
#
# The HCL blocks, arguments, variables, and values are derived from the
# operator configuration for Consul-Terraform-Sync. Any manual changes to
# this file may not be preserved and could be overwritten by a subsequent
# update.
#
# Task: test
# Description: user description for task named 'test'

services = {
  "web.worker-01.dc1" = {
    id              = "web"
    name            = "web"
    kind            = ""
    address         = "1.1.1.1"
    port            = 8000
    meta            = {}
    tags            = ["tag_a", "tag_b"]
    namespace       = ""
    status          = "passing"
    node            = "worker-01"
    node_id         = "39e5a7f5-2834-e16d-6925-78167c9f50d8"
    node_address    = "127.0.0.1"
    node_datacenter = "dc1"
    node_tagged_addresses = {
      lan      = "127.0.0.1"
      lan_ipv4 = "127.0.0.1"
      wan      = "127.0.0.1"
      wan_ipv4 = "127.0.0.1"
    }
    node_meta = {
      consul-network-segment = ""
    }
    cts_user_defined_meta = {}
  },
}
