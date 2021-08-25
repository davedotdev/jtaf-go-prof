resource "junos-device_Policy__OptionsPolicy__StatementTermFromInterface" "policy_redist_lo0_from" {
   resource_name = "export_ifaces"
   name = "policy_redist_lo0"
   name__1 = "policy_redist_lo0"
   interface = "lo0.0"
}

resource "junos-device_Policy__OptionsPolicy__StatementTermThenAccept" "policy_redist_lo0_then" {
   resource_name = "export_ifaces"
   name = "policy_redist_lo0"
   name__1 = "policy_redist_lo0"
}

resource "junos-device_ProtocolsBgpGroupPeer__As" "bgp_neighbor_vqfx02_peer_as" {
    resource_name = "bgp_neighbor_vqfx_02"
    name = "vqfx02"
    peer__as = "65002"
}

resource "junos-device_ProtocolsBgpGroupFamilyInet6Unicast" "bgp_neighbor_vqfx02_inet6_unicast" {
   resource_name = "bgp_neighbor_vqfx_02"
   name = "vqfx02"
   depends_on = [junos-device_ProtocolsBgpGroupFamilyInet6Unicast.bgp_neighbor_vqfx02_inet6_unicast]
}

resource "junos-device_ProtocolsBgpGroupNeighborExport" "bgp_neighbor_vqfx02_export" {
    resource_name = "bgp_neighbor_vqfx_02"
    name = "vqfx02"
    name__1 = "dead:babe:beef::2"
    export = "policy_redist_lo0"
}




