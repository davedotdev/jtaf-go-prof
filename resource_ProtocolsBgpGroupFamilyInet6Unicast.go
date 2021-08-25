
package main

import (
    "encoding/xml"
    "fmt"
    "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)


// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex - interface is keyword in golang
type xmlProtocolsBgpGroupFamilyInet6Unicast struct {
	XMLName xml.Name `xml:"configuration"`
	Groups  struct {
		XMLName	xml.Name	`xml:"groups"`
		Name	string	`xml:"name"`
		V_group  struct {
			XMLName xml.Name `xml:"group"`
			V_name  string  `xml:"name"`
			V_unicast  struct {
				XMLName xml.Name `xml:"unicast"`
				V_apply__groups  string  `xml:"apply-groups"`
				V_apply__groups__except  string  `xml:"apply-groups-except"`
				V_apply__macro	struct {
					XMLName xml.Name `xml:"apply-macro"`
					V_name__1  string  `xml:"name"`
					V_data	struct {
						XMLName xml.Name `xml:"data"`
						V_name__2  string  `xml:"name"`
						V_value  string  `xml:"value"`
					} `xml:"data"`
				} `xml:"apply-macro"`
				V_prefix__limit	struct {
					XMLName xml.Name `xml:"prefix-limit"`
					V_apply__groups__1  string  `xml:"apply-groups"`
					V_apply__groups__except__1  string  `xml:"apply-groups-except"`
					V_apply__macro__1	struct {
						XMLName xml.Name `xml:"apply-macro"`
						V_name__3  string  `xml:"name"`
						V_data__1	struct {
							XMLName xml.Name `xml:"data"`
							V_name__4  string  `xml:"name"`
							V_value__1  string  `xml:"value"`
						} `xml:"data"`
					} `xml:"apply-macro"`
					V_maximum  string  `xml:"maximum"`
					V_teardown	struct {
						XMLName xml.Name `xml:"teardown"`
						V_apply__groups__2  string  `xml:"apply-groups"`
						V_apply__groups__except__2  string  `xml:"apply-groups-except"`
						V_apply__macro__2	struct {
							XMLName xml.Name `xml:"apply-macro"`
							V_name__5  string  `xml:"name"`
							V_data__2	struct {
								XMLName xml.Name `xml:"data"`
								V_name__6  string  `xml:"name"`
								V_value__2  string  `xml:"value"`
							} `xml:"data"`
						} `xml:"apply-macro"`
						V_limit__threshold  string  `xml:"limit-threshold"`
						V_idle__timeout	struct {
							XMLName xml.Name `xml:"idle-timeout"`
						} `xml:"idle-timeout"`
					} `xml:"teardown"`
				} `xml:"prefix-limit"`
				V_accepted__prefix__limit	struct {
					XMLName xml.Name `xml:"accepted-prefix-limit"`
					V_apply__groups__3  string  `xml:"apply-groups"`
					V_apply__groups__except__3  string  `xml:"apply-groups-except"`
					V_apply__macro__3	struct {
						XMLName xml.Name `xml:"apply-macro"`
						V_name__7  string  `xml:"name"`
						V_data__3	struct {
							XMLName xml.Name `xml:"data"`
							V_name__8  string  `xml:"name"`
							V_value__3  string  `xml:"value"`
						} `xml:"data"`
					} `xml:"apply-macro"`
					V_maximum__1  string  `xml:"maximum"`
					V_teardown__1	struct {
						XMLName xml.Name `xml:"teardown"`
						V_apply__groups__4  string  `xml:"apply-groups"`
						V_apply__groups__except__4  string  `xml:"apply-groups-except"`
						V_apply__macro__4	struct {
							XMLName xml.Name `xml:"apply-macro"`
							V_name__9  string  `xml:"name"`
							V_data__4	struct {
								XMLName xml.Name `xml:"data"`
								V_name__10  string  `xml:"name"`
								V_value__4  string  `xml:"value"`
							} `xml:"data"`
						} `xml:"apply-macro"`
						V_limit__threshold__1  string  `xml:"limit-threshold"`
						V_idle__timeout__1	struct {
							XMLName xml.Name `xml:"idle-timeout"`
						} `xml:"idle-timeout"`
					} `xml:"teardown"`
				} `xml:"accepted-prefix-limit"`
				V_rib__group	struct {
					XMLName xml.Name `xml:"rib-group"`
					V_ribgroup__name  string  `xml:"ribgroup-name"`
				} `xml:"rib-group"`
				V_add__path	struct {
					XMLName xml.Name `xml:"add-path"`
					V_apply__groups__5  string  `xml:"apply-groups"`
					V_apply__groups__except__5  string  `xml:"apply-groups-except"`
					V_apply__macro__5	struct {
						XMLName xml.Name `xml:"apply-macro"`
						V_name__11  string  `xml:"name"`
						V_data__5	struct {
							XMLName xml.Name `xml:"data"`
							V_name__12  string  `xml:"name"`
							V_value__5  string  `xml:"value"`
						} `xml:"data"`
					} `xml:"apply-macro"`
					V_receive  string  `xml:"receive"`
					V_send	struct {
						XMLName xml.Name `xml:"send"`
						V_apply__groups__6  string  `xml:"apply-groups"`
						V_apply__groups__except__6  string  `xml:"apply-groups-except"`
						V_apply__macro__6	struct {
							XMLName xml.Name `xml:"apply-macro"`
							V_name__13  string  `xml:"name"`
							V_data__6	struct {
								XMLName xml.Name `xml:"data"`
								V_name__14  string  `xml:"name"`
								V_value__6  string  `xml:"value"`
							} `xml:"data"`
						} `xml:"apply-macro"`
						V_path__selection__mode	struct {
							XMLName xml.Name `xml:"path-selection-mode"`
							V_apply__groups__7  string  `xml:"apply-groups"`
							V_apply__groups__except__7  string  `xml:"apply-groups-except"`
							V_apply__macro__7	struct {
								XMLName xml.Name `xml:"apply-macro"`
								V_name__15  string  `xml:"name"`
								V_data__7	struct {
									XMLName xml.Name `xml:"data"`
									V_name__16  string  `xml:"name"`
									V_value__7  string  `xml:"value"`
								} `xml:"data"`
							} `xml:"apply-macro"`
						} `xml:"path-selection-mode"`
						V_prefix__policy  string  `xml:"prefix-policy"`
						V_path__count  string  `xml:"path-count"`
						V_include__backup__path  string  `xml:"include-backup-path"`
						V_multipath  string  `xml:"multipath"`
					} `xml:"send"`
				} `xml:"add-path"`
				V_aigp	struct {
					XMLName xml.Name `xml:"aigp"`
					V_apply__groups__8  string  `xml:"apply-groups"`
					V_apply__groups__except__8  string  `xml:"apply-groups-except"`
					V_apply__macro__8	struct {
						XMLName xml.Name `xml:"apply-macro"`
						V_name__17  string  `xml:"name"`
						V_data__8	struct {
							XMLName xml.Name `xml:"data"`
							V_name__18  string  `xml:"name"`
							V_value__8  string  `xml:"value"`
						} `xml:"data"`
					} `xml:"apply-macro"`
					V_disable  string  `xml:"disable"`
				} `xml:"aigp"`
				V_damping  string  `xml:"damping"`
				V_local__ipv4__address  string  `xml:"local-ipv4-address"`
				V_loops	struct {
					XMLName xml.Name `xml:"loops"`
					V_apply__groups__9  string  `xml:"apply-groups"`
					V_apply__groups__except__9  string  `xml:"apply-groups-except"`
					V_apply__macro__9	struct {
						XMLName xml.Name `xml:"apply-macro"`
						V_name__19  string  `xml:"name"`
						V_data__9	struct {
							XMLName xml.Name `xml:"data"`
							V_name__20  string  `xml:"name"`
							V_value__9  string  `xml:"value"`
						} `xml:"data"`
					} `xml:"apply-macro"`
					V_loops__1  string  `xml:"loops"`
				} `xml:"loops"`
				V_delay__route__advertisements	struct {
					XMLName xml.Name `xml:"delay-route-advertisements"`
					V_apply__groups__10  string  `xml:"apply-groups"`
					V_apply__groups__except__10  string  `xml:"apply-groups-except"`
					V_apply__macro__10	struct {
						XMLName xml.Name `xml:"apply-macro"`
						V_name__21  string  `xml:"name"`
						V_data__10	struct {
							XMLName xml.Name `xml:"data"`
							V_name__22  string  `xml:"name"`
							V_value__10  string  `xml:"value"`
						} `xml:"data"`
					} `xml:"apply-macro"`
					V_always__wait__for__krt__drain  string  `xml:"always-wait-for-krt-drain"`
					V_minimum__delay	struct {
						XMLName xml.Name `xml:"minimum-delay"`
						V_apply__groups__11  string  `xml:"apply-groups"`
						V_apply__groups__except__11  string  `xml:"apply-groups-except"`
						V_apply__macro__11	struct {
							XMLName xml.Name `xml:"apply-macro"`
							V_name__23  string  `xml:"name"`
							V_data__11	struct {
								XMLName xml.Name `xml:"data"`
								V_name__24  string  `xml:"name"`
								V_value__11  string  `xml:"value"`
							} `xml:"data"`
						} `xml:"apply-macro"`
						V_routing__uptime  string  `xml:"routing-uptime"`
						V_inbound__convergence  string  `xml:"inbound-convergence"`
					} `xml:"minimum-delay"`
					V_maximum__delay	struct {
						XMLName xml.Name `xml:"maximum-delay"`
						V_apply__groups__12  string  `xml:"apply-groups"`
						V_apply__groups__except__12  string  `xml:"apply-groups-except"`
						V_apply__macro__12	struct {
							XMLName xml.Name `xml:"apply-macro"`
							V_name__25  string  `xml:"name"`
							V_data__12	struct {
								XMLName xml.Name `xml:"data"`
								V_name__26  string  `xml:"name"`
								V_value__12  string  `xml:"value"`
							} `xml:"data"`
						} `xml:"apply-macro"`
						V_route__age  string  `xml:"route-age"`
						V_routing__uptime__1  string  `xml:"routing-uptime"`
					} `xml:"maximum-delay"`
				} `xml:"delay-route-advertisements"`
				V_nexthop__resolution	struct {
					XMLName xml.Name `xml:"nexthop-resolution"`
					V_apply__groups__13  string  `xml:"apply-groups"`
					V_apply__groups__except__13  string  `xml:"apply-groups-except"`
					V_apply__macro__13	struct {
						XMLName xml.Name `xml:"apply-macro"`
						V_name__27  string  `xml:"name"`
						V_data__13	struct {
							XMLName xml.Name `xml:"data"`
							V_name__28  string  `xml:"name"`
							V_value__13  string  `xml:"value"`
						} `xml:"data"`
					} `xml:"apply-macro"`
					V_no__resolution  string  `xml:"no-resolution"`
					V_preserve__nexthop__hierarchy  string  `xml:"preserve-nexthop-hierarchy"`
				} `xml:"nexthop-resolution"`
				V_defer__initial__multipath__build	struct {
					XMLName xml.Name `xml:"defer-initial-multipath-build"`
					V_apply__groups__14  string  `xml:"apply-groups"`
					V_apply__groups__except__14  string  `xml:"apply-groups-except"`
					V_apply__macro__14	struct {
						XMLName xml.Name `xml:"apply-macro"`
						V_name__29  string  `xml:"name"`
						V_data__14	struct {
							XMLName xml.Name `xml:"data"`
							V_name__30  string  `xml:"name"`
							V_value__14  string  `xml:"value"`
						} `xml:"data"`
					} `xml:"apply-macro"`
					V_maximum__delay__1  string  `xml:"maximum-delay"`
				} `xml:"defer-initial-multipath-build"`
				V_graceful__restart	struct {
					XMLName xml.Name `xml:"graceful-restart"`
					V_apply__groups__15  string  `xml:"apply-groups"`
					V_apply__groups__except__15  string  `xml:"apply-groups-except"`
					V_apply__macro__15	struct {
						XMLName xml.Name `xml:"apply-macro"`
						V_name__31  string  `xml:"name"`
						V_data__15	struct {
							XMLName xml.Name `xml:"data"`
							V_name__32  string  `xml:"name"`
							V_value__15  string  `xml:"value"`
						} `xml:"data"`
					} `xml:"apply-macro"`
					V_long__lived	struct {
						XMLName xml.Name `xml:"long-lived"`
						V_apply__groups__16  string  `xml:"apply-groups"`
						V_apply__groups__except__16  string  `xml:"apply-groups-except"`
						V_apply__macro__16	struct {
							XMLName xml.Name `xml:"apply-macro"`
							V_name__33  string  `xml:"name"`
							V_data__16	struct {
								XMLName xml.Name `xml:"data"`
								V_name__34  string  `xml:"name"`
								V_value__16  string  `xml:"value"`
							} `xml:"data"`
						} `xml:"apply-macro"`
						V_restarter	struct {
							XMLName xml.Name `xml:"restarter"`
							V_apply__groups__17  string  `xml:"apply-groups"`
							V_apply__groups__except__17  string  `xml:"apply-groups-except"`
							V_apply__macro__17	struct {
								XMLName xml.Name `xml:"apply-macro"`
								V_name__35  string  `xml:"name"`
								V_data__17	struct {
									XMLName xml.Name `xml:"data"`
									V_name__36  string  `xml:"name"`
									V_value__17  string  `xml:"value"`
								} `xml:"data"`
							} `xml:"apply-macro"`
							V_stale__time  string  `xml:"stale-time"`
						} `xml:"restarter"`
					} `xml:"long-lived"`
					V_forwarding__state__bit  string  `xml:"forwarding-state-bit"`
				} `xml:"graceful-restart"`
				V_extended__nexthop  string  `xml:"extended-nexthop"`
				V_extended__nexthop__color  string  `xml:"extended-nexthop-color"`
				V_extended__nexthop__tunnel  string  `xml:"extended-nexthop-tunnel"`
				V_no__install  string  `xml:"no-install"`
				V_route__age__bgp__view  string  `xml:"route-age-bgp-view"`
				V_output__queue__priority	struct {
					XMLName xml.Name `xml:"output-queue-priority"`
				} `xml:"output-queue-priority"`
				V_route__refresh__priority	struct {
					XMLName xml.Name `xml:"route-refresh-priority"`
				} `xml:"route-refresh-priority"`
				V_withdraw__priority	struct {
					XMLName xml.Name `xml:"withdraw-priority"`
				} `xml:"withdraw-priority"`
				V_advertise__srv6__service  string  `xml:"advertise-srv6-service"`
				V_accept__srv6__service  string  `xml:"accept-srv6-service"`
				V_protection	struct {
					XMLName xml.Name `xml:"protection"`
				} `xml:"protection"`
				V_topology	struct {
					XMLName xml.Name `xml:"topology"`
					V_name__37  string  `xml:"name"`
					V_apply__groups__18  string  `xml:"apply-groups"`
					V_apply__groups__except__18  string  `xml:"apply-groups-except"`
					V_apply__macro__18	struct {
						XMLName xml.Name `xml:"apply-macro"`
						V_name__38  string  `xml:"name"`
						V_data__18	struct {
							XMLName xml.Name `xml:"data"`
							V_name__39  string  `xml:"name"`
							V_value__18  string  `xml:"value"`
						} `xml:"data"`
					} `xml:"apply-macro"`
					V_community  string  `xml:"community"`
				} `xml:"topology"`
			} `xml:"family>inet6>unicast"`
		} `xml:"protocols>bgp>group"`
	} `xml:"groups"`
	ApplyGroup string `xml:"apply-groups"`
}

// v_ is appended before every variable so it doesn't give any conflict
// with any keyword in golang. ex- interface is keyword in golang
func junosProtocolsBgpGroupFamilyInet6UnicastCreate(d *schema.ResourceData, m interface{}) error {

	var err error
    client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)
	V_apply__groups := d.Get("apply__groups").(string)
	V_apply__groups__except := d.Get("apply__groups__except").(string)
	V_name__1 := d.Get("name__1").(string)
	V_name__2 := d.Get("name__2").(string)
	V_value := d.Get("value").(string)
	V_apply__groups__1 := d.Get("apply__groups__1").(string)
	V_apply__groups__except__1 := d.Get("apply__groups__except__1").(string)
	V_name__3 := d.Get("name__3").(string)
	V_name__4 := d.Get("name__4").(string)
	V_value__1 := d.Get("value__1").(string)
	V_maximum := d.Get("maximum").(string)
	V_apply__groups__2 := d.Get("apply__groups__2").(string)
	V_apply__groups__except__2 := d.Get("apply__groups__except__2").(string)
	V_name__5 := d.Get("name__5").(string)
	V_name__6 := d.Get("name__6").(string)
	V_value__2 := d.Get("value__2").(string)
	V_limit__threshold := d.Get("limit__threshold").(string)
	V_apply__groups__3 := d.Get("apply__groups__3").(string)
	V_apply__groups__except__3 := d.Get("apply__groups__except__3").(string)
	V_name__7 := d.Get("name__7").(string)
	V_name__8 := d.Get("name__8").(string)
	V_value__3 := d.Get("value__3").(string)
	V_maximum__1 := d.Get("maximum__1").(string)
	V_apply__groups__4 := d.Get("apply__groups__4").(string)
	V_apply__groups__except__4 := d.Get("apply__groups__except__4").(string)
	V_name__9 := d.Get("name__9").(string)
	V_name__10 := d.Get("name__10").(string)
	V_value__4 := d.Get("value__4").(string)
	V_limit__threshold__1 := d.Get("limit__threshold__1").(string)
	V_ribgroup__name := d.Get("ribgroup__name").(string)
	V_apply__groups__5 := d.Get("apply__groups__5").(string)
	V_apply__groups__except__5 := d.Get("apply__groups__except__5").(string)
	V_name__11 := d.Get("name__11").(string)
	V_name__12 := d.Get("name__12").(string)
	V_value__5 := d.Get("value__5").(string)
	V_receive := d.Get("receive").(string)
	V_apply__groups__6 := d.Get("apply__groups__6").(string)
	V_apply__groups__except__6 := d.Get("apply__groups__except__6").(string)
	V_name__13 := d.Get("name__13").(string)
	V_name__14 := d.Get("name__14").(string)
	V_value__6 := d.Get("value__6").(string)
	V_apply__groups__7 := d.Get("apply__groups__7").(string)
	V_apply__groups__except__7 := d.Get("apply__groups__except__7").(string)
	V_name__15 := d.Get("name__15").(string)
	V_name__16 := d.Get("name__16").(string)
	V_value__7 := d.Get("value__7").(string)
	V_prefix__policy := d.Get("prefix__policy").(string)
	V_path__count := d.Get("path__count").(string)
	V_include__backup__path := d.Get("include__backup__path").(string)
	V_multipath := d.Get("multipath").(string)
	V_apply__groups__8 := d.Get("apply__groups__8").(string)
	V_apply__groups__except__8 := d.Get("apply__groups__except__8").(string)
	V_name__17 := d.Get("name__17").(string)
	V_name__18 := d.Get("name__18").(string)
	V_value__8 := d.Get("value__8").(string)
	V_disable := d.Get("disable").(string)
	V_damping := d.Get("damping").(string)
	V_local__ipv4__address := d.Get("local__ipv4__address").(string)
	V_apply__groups__9 := d.Get("apply__groups__9").(string)
	V_apply__groups__except__9 := d.Get("apply__groups__except__9").(string)
	V_name__19 := d.Get("name__19").(string)
	V_name__20 := d.Get("name__20").(string)
	V_value__9 := d.Get("value__9").(string)
	V_loops__1 := d.Get("loops__1").(string)
	V_apply__groups__10 := d.Get("apply__groups__10").(string)
	V_apply__groups__except__10 := d.Get("apply__groups__except__10").(string)
	V_name__21 := d.Get("name__21").(string)
	V_name__22 := d.Get("name__22").(string)
	V_value__10 := d.Get("value__10").(string)
	V_always__wait__for__krt__drain := d.Get("always__wait__for__krt__drain").(string)
	V_apply__groups__11 := d.Get("apply__groups__11").(string)
	V_apply__groups__except__11 := d.Get("apply__groups__except__11").(string)
	V_name__23 := d.Get("name__23").(string)
	V_name__24 := d.Get("name__24").(string)
	V_value__11 := d.Get("value__11").(string)
	V_routing__uptime := d.Get("routing__uptime").(string)
	V_inbound__convergence := d.Get("inbound__convergence").(string)
	V_apply__groups__12 := d.Get("apply__groups__12").(string)
	V_apply__groups__except__12 := d.Get("apply__groups__except__12").(string)
	V_name__25 := d.Get("name__25").(string)
	V_name__26 := d.Get("name__26").(string)
	V_value__12 := d.Get("value__12").(string)
	V_route__age := d.Get("route__age").(string)
	V_routing__uptime__1 := d.Get("routing__uptime__1").(string)
	V_apply__groups__13 := d.Get("apply__groups__13").(string)
	V_apply__groups__except__13 := d.Get("apply__groups__except__13").(string)
	V_name__27 := d.Get("name__27").(string)
	V_name__28 := d.Get("name__28").(string)
	V_value__13 := d.Get("value__13").(string)
	V_no__resolution := d.Get("no__resolution").(string)
	V_preserve__nexthop__hierarchy := d.Get("preserve__nexthop__hierarchy").(string)
	V_apply__groups__14 := d.Get("apply__groups__14").(string)
	V_apply__groups__except__14 := d.Get("apply__groups__except__14").(string)
	V_name__29 := d.Get("name__29").(string)
	V_name__30 := d.Get("name__30").(string)
	V_value__14 := d.Get("value__14").(string)
	V_maximum__delay__1 := d.Get("maximum__delay__1").(string)
	V_apply__groups__15 := d.Get("apply__groups__15").(string)
	V_apply__groups__except__15 := d.Get("apply__groups__except__15").(string)
	V_name__31 := d.Get("name__31").(string)
	V_name__32 := d.Get("name__32").(string)
	V_value__15 := d.Get("value__15").(string)
	V_apply__groups__16 := d.Get("apply__groups__16").(string)
	V_apply__groups__except__16 := d.Get("apply__groups__except__16").(string)
	V_name__33 := d.Get("name__33").(string)
	V_name__34 := d.Get("name__34").(string)
	V_value__16 := d.Get("value__16").(string)
	V_apply__groups__17 := d.Get("apply__groups__17").(string)
	V_apply__groups__except__17 := d.Get("apply__groups__except__17").(string)
	V_name__35 := d.Get("name__35").(string)
	V_name__36 := d.Get("name__36").(string)
	V_value__17 := d.Get("value__17").(string)
	V_stale__time := d.Get("stale__time").(string)
	V_forwarding__state__bit := d.Get("forwarding__state__bit").(string)
	V_extended__nexthop := d.Get("extended__nexthop").(string)
	V_extended__nexthop__color := d.Get("extended__nexthop__color").(string)
	V_extended__nexthop__tunnel := d.Get("extended__nexthop__tunnel").(string)
	V_no__install := d.Get("no__install").(string)
	V_route__age__bgp__view := d.Get("route__age__bgp__view").(string)
	V_advertise__srv6__service := d.Get("advertise__srv6__service").(string)
	V_accept__srv6__service := d.Get("accept__srv6__service").(string)
	V_name__37 := d.Get("name__37").(string)
	V_apply__groups__18 := d.Get("apply__groups__18").(string)
	V_apply__groups__except__18 := d.Get("apply__groups__except__18").(string)
	V_name__38 := d.Get("name__38").(string)
	V_name__39 := d.Get("name__39").(string)
	V_value__18 := d.Get("value__18").(string)
	V_community := d.Get("community").(string)
	commit := true

	config := xmlProtocolsBgpGroupFamilyInet6Unicast{}
	config.ApplyGroup = id
	config.Groups.Name = id
	config.Groups.V_group.V_name = V_name
	config.Groups.V_group.V_unicast.V_apply__groups = V_apply__groups
	config.Groups.V_group.V_unicast.V_apply__groups__except = V_apply__groups__except
	config.Groups.V_group.V_unicast.V_apply__macro.V_name__1 = V_name__1
	config.Groups.V_group.V_unicast.V_apply__macro.V_data.V_name__2 = V_name__2
	config.Groups.V_group.V_unicast.V_apply__macro.V_data.V_value = V_value
	config.Groups.V_group.V_unicast.V_prefix__limit.V_apply__groups__1 = V_apply__groups__1
	config.Groups.V_group.V_unicast.V_prefix__limit.V_apply__groups__except__1 = V_apply__groups__except__1
	config.Groups.V_group.V_unicast.V_prefix__limit.V_apply__macro__1.V_name__3 = V_name__3
	config.Groups.V_group.V_unicast.V_prefix__limit.V_apply__macro__1.V_data__1.V_name__4 = V_name__4
	config.Groups.V_group.V_unicast.V_prefix__limit.V_apply__macro__1.V_data__1.V_value__1 = V_value__1
	config.Groups.V_group.V_unicast.V_prefix__limit.V_maximum = V_maximum
	config.Groups.V_group.V_unicast.V_prefix__limit.V_teardown.V_apply__groups__2 = V_apply__groups__2
	config.Groups.V_group.V_unicast.V_prefix__limit.V_teardown.V_apply__groups__except__2 = V_apply__groups__except__2
	config.Groups.V_group.V_unicast.V_prefix__limit.V_teardown.V_apply__macro__2.V_name__5 = V_name__5
	config.Groups.V_group.V_unicast.V_prefix__limit.V_teardown.V_apply__macro__2.V_data__2.V_name__6 = V_name__6
	config.Groups.V_group.V_unicast.V_prefix__limit.V_teardown.V_apply__macro__2.V_data__2.V_value__2 = V_value__2
	config.Groups.V_group.V_unicast.V_prefix__limit.V_teardown.V_limit__threshold = V_limit__threshold
	config.Groups.V_group.V_unicast.V_accepted__prefix__limit.V_apply__groups__3 = V_apply__groups__3
	config.Groups.V_group.V_unicast.V_accepted__prefix__limit.V_apply__groups__except__3 = V_apply__groups__except__3
	config.Groups.V_group.V_unicast.V_accepted__prefix__limit.V_apply__macro__3.V_name__7 = V_name__7
	config.Groups.V_group.V_unicast.V_accepted__prefix__limit.V_apply__macro__3.V_data__3.V_name__8 = V_name__8
	config.Groups.V_group.V_unicast.V_accepted__prefix__limit.V_apply__macro__3.V_data__3.V_value__3 = V_value__3
	config.Groups.V_group.V_unicast.V_accepted__prefix__limit.V_maximum__1 = V_maximum__1
	config.Groups.V_group.V_unicast.V_accepted__prefix__limit.V_teardown__1.V_apply__groups__4 = V_apply__groups__4
	config.Groups.V_group.V_unicast.V_accepted__prefix__limit.V_teardown__1.V_apply__groups__except__4 = V_apply__groups__except__4
	config.Groups.V_group.V_unicast.V_accepted__prefix__limit.V_teardown__1.V_apply__macro__4.V_name__9 = V_name__9
	config.Groups.V_group.V_unicast.V_accepted__prefix__limit.V_teardown__1.V_apply__macro__4.V_data__4.V_name__10 = V_name__10
	config.Groups.V_group.V_unicast.V_accepted__prefix__limit.V_teardown__1.V_apply__macro__4.V_data__4.V_value__4 = V_value__4
	config.Groups.V_group.V_unicast.V_accepted__prefix__limit.V_teardown__1.V_limit__threshold__1 = V_limit__threshold__1
	config.Groups.V_group.V_unicast.V_rib__group.V_ribgroup__name = V_ribgroup__name
	config.Groups.V_group.V_unicast.V_add__path.V_apply__groups__5 = V_apply__groups__5
	config.Groups.V_group.V_unicast.V_add__path.V_apply__groups__except__5 = V_apply__groups__except__5
	config.Groups.V_group.V_unicast.V_add__path.V_apply__macro__5.V_name__11 = V_name__11
	config.Groups.V_group.V_unicast.V_add__path.V_apply__macro__5.V_data__5.V_name__12 = V_name__12
	config.Groups.V_group.V_unicast.V_add__path.V_apply__macro__5.V_data__5.V_value__5 = V_value__5
	config.Groups.V_group.V_unicast.V_add__path.V_receive = V_receive
	config.Groups.V_group.V_unicast.V_add__path.V_send.V_apply__groups__6 = V_apply__groups__6
	config.Groups.V_group.V_unicast.V_add__path.V_send.V_apply__groups__except__6 = V_apply__groups__except__6
	config.Groups.V_group.V_unicast.V_add__path.V_send.V_apply__macro__6.V_name__13 = V_name__13
	config.Groups.V_group.V_unicast.V_add__path.V_send.V_apply__macro__6.V_data__6.V_name__14 = V_name__14
	config.Groups.V_group.V_unicast.V_add__path.V_send.V_apply__macro__6.V_data__6.V_value__6 = V_value__6
	config.Groups.V_group.V_unicast.V_add__path.V_send.V_path__selection__mode.V_apply__groups__7 = V_apply__groups__7
	config.Groups.V_group.V_unicast.V_add__path.V_send.V_path__selection__mode.V_apply__groups__except__7 = V_apply__groups__except__7
	config.Groups.V_group.V_unicast.V_add__path.V_send.V_path__selection__mode.V_apply__macro__7.V_name__15 = V_name__15
	config.Groups.V_group.V_unicast.V_add__path.V_send.V_path__selection__mode.V_apply__macro__7.V_data__7.V_name__16 = V_name__16
	config.Groups.V_group.V_unicast.V_add__path.V_send.V_path__selection__mode.V_apply__macro__7.V_data__7.V_value__7 = V_value__7
	config.Groups.V_group.V_unicast.V_add__path.V_send.V_prefix__policy = V_prefix__policy
	config.Groups.V_group.V_unicast.V_add__path.V_send.V_path__count = V_path__count
	config.Groups.V_group.V_unicast.V_add__path.V_send.V_include__backup__path = V_include__backup__path
	config.Groups.V_group.V_unicast.V_add__path.V_send.V_multipath = V_multipath
	config.Groups.V_group.V_unicast.V_aigp.V_apply__groups__8 = V_apply__groups__8
	config.Groups.V_group.V_unicast.V_aigp.V_apply__groups__except__8 = V_apply__groups__except__8
	config.Groups.V_group.V_unicast.V_aigp.V_apply__macro__8.V_name__17 = V_name__17
	config.Groups.V_group.V_unicast.V_aigp.V_apply__macro__8.V_data__8.V_name__18 = V_name__18
	config.Groups.V_group.V_unicast.V_aigp.V_apply__macro__8.V_data__8.V_value__8 = V_value__8
	config.Groups.V_group.V_unicast.V_aigp.V_disable = V_disable
	config.Groups.V_group.V_unicast.V_damping = V_damping
	config.Groups.V_group.V_unicast.V_local__ipv4__address = V_local__ipv4__address
	config.Groups.V_group.V_unicast.V_loops.V_apply__groups__9 = V_apply__groups__9
	config.Groups.V_group.V_unicast.V_loops.V_apply__groups__except__9 = V_apply__groups__except__9
	config.Groups.V_group.V_unicast.V_loops.V_apply__macro__9.V_name__19 = V_name__19
	config.Groups.V_group.V_unicast.V_loops.V_apply__macro__9.V_data__9.V_name__20 = V_name__20
	config.Groups.V_group.V_unicast.V_loops.V_apply__macro__9.V_data__9.V_value__9 = V_value__9
	config.Groups.V_group.V_unicast.V_loops.V_loops__1 = V_loops__1
	config.Groups.V_group.V_unicast.V_delay__route__advertisements.V_apply__groups__10 = V_apply__groups__10
	config.Groups.V_group.V_unicast.V_delay__route__advertisements.V_apply__groups__except__10 = V_apply__groups__except__10
	config.Groups.V_group.V_unicast.V_delay__route__advertisements.V_apply__macro__10.V_name__21 = V_name__21
	config.Groups.V_group.V_unicast.V_delay__route__advertisements.V_apply__macro__10.V_data__10.V_name__22 = V_name__22
	config.Groups.V_group.V_unicast.V_delay__route__advertisements.V_apply__macro__10.V_data__10.V_value__10 = V_value__10
	config.Groups.V_group.V_unicast.V_delay__route__advertisements.V_always__wait__for__krt__drain = V_always__wait__for__krt__drain
	config.Groups.V_group.V_unicast.V_delay__route__advertisements.V_minimum__delay.V_apply__groups__11 = V_apply__groups__11
	config.Groups.V_group.V_unicast.V_delay__route__advertisements.V_minimum__delay.V_apply__groups__except__11 = V_apply__groups__except__11
	config.Groups.V_group.V_unicast.V_delay__route__advertisements.V_minimum__delay.V_apply__macro__11.V_name__23 = V_name__23
	config.Groups.V_group.V_unicast.V_delay__route__advertisements.V_minimum__delay.V_apply__macro__11.V_data__11.V_name__24 = V_name__24
	config.Groups.V_group.V_unicast.V_delay__route__advertisements.V_minimum__delay.V_apply__macro__11.V_data__11.V_value__11 = V_value__11
	config.Groups.V_group.V_unicast.V_delay__route__advertisements.V_minimum__delay.V_routing__uptime = V_routing__uptime
	config.Groups.V_group.V_unicast.V_delay__route__advertisements.V_minimum__delay.V_inbound__convergence = V_inbound__convergence
	config.Groups.V_group.V_unicast.V_delay__route__advertisements.V_maximum__delay.V_apply__groups__12 = V_apply__groups__12
	config.Groups.V_group.V_unicast.V_delay__route__advertisements.V_maximum__delay.V_apply__groups__except__12 = V_apply__groups__except__12
	config.Groups.V_group.V_unicast.V_delay__route__advertisements.V_maximum__delay.V_apply__macro__12.V_name__25 = V_name__25
	config.Groups.V_group.V_unicast.V_delay__route__advertisements.V_maximum__delay.V_apply__macro__12.V_data__12.V_name__26 = V_name__26
	config.Groups.V_group.V_unicast.V_delay__route__advertisements.V_maximum__delay.V_apply__macro__12.V_data__12.V_value__12 = V_value__12
	config.Groups.V_group.V_unicast.V_delay__route__advertisements.V_maximum__delay.V_route__age = V_route__age
	config.Groups.V_group.V_unicast.V_delay__route__advertisements.V_maximum__delay.V_routing__uptime__1 = V_routing__uptime__1
	config.Groups.V_group.V_unicast.V_nexthop__resolution.V_apply__groups__13 = V_apply__groups__13
	config.Groups.V_group.V_unicast.V_nexthop__resolution.V_apply__groups__except__13 = V_apply__groups__except__13
	config.Groups.V_group.V_unicast.V_nexthop__resolution.V_apply__macro__13.V_name__27 = V_name__27
	config.Groups.V_group.V_unicast.V_nexthop__resolution.V_apply__macro__13.V_data__13.V_name__28 = V_name__28
	config.Groups.V_group.V_unicast.V_nexthop__resolution.V_apply__macro__13.V_data__13.V_value__13 = V_value__13
	config.Groups.V_group.V_unicast.V_nexthop__resolution.V_no__resolution = V_no__resolution
	config.Groups.V_group.V_unicast.V_nexthop__resolution.V_preserve__nexthop__hierarchy = V_preserve__nexthop__hierarchy
	config.Groups.V_group.V_unicast.V_defer__initial__multipath__build.V_apply__groups__14 = V_apply__groups__14
	config.Groups.V_group.V_unicast.V_defer__initial__multipath__build.V_apply__groups__except__14 = V_apply__groups__except__14
	config.Groups.V_group.V_unicast.V_defer__initial__multipath__build.V_apply__macro__14.V_name__29 = V_name__29
	config.Groups.V_group.V_unicast.V_defer__initial__multipath__build.V_apply__macro__14.V_data__14.V_name__30 = V_name__30
	config.Groups.V_group.V_unicast.V_defer__initial__multipath__build.V_apply__macro__14.V_data__14.V_value__14 = V_value__14
	config.Groups.V_group.V_unicast.V_defer__initial__multipath__build.V_maximum__delay__1 = V_maximum__delay__1
	config.Groups.V_group.V_unicast.V_graceful__restart.V_apply__groups__15 = V_apply__groups__15
	config.Groups.V_group.V_unicast.V_graceful__restart.V_apply__groups__except__15 = V_apply__groups__except__15
	config.Groups.V_group.V_unicast.V_graceful__restart.V_apply__macro__15.V_name__31 = V_name__31
	config.Groups.V_group.V_unicast.V_graceful__restart.V_apply__macro__15.V_data__15.V_name__32 = V_name__32
	config.Groups.V_group.V_unicast.V_graceful__restart.V_apply__macro__15.V_data__15.V_value__15 = V_value__15
	config.Groups.V_group.V_unicast.V_graceful__restart.V_long__lived.V_apply__groups__16 = V_apply__groups__16
	config.Groups.V_group.V_unicast.V_graceful__restart.V_long__lived.V_apply__groups__except__16 = V_apply__groups__except__16
	config.Groups.V_group.V_unicast.V_graceful__restart.V_long__lived.V_apply__macro__16.V_name__33 = V_name__33
	config.Groups.V_group.V_unicast.V_graceful__restart.V_long__lived.V_apply__macro__16.V_data__16.V_name__34 = V_name__34
	config.Groups.V_group.V_unicast.V_graceful__restart.V_long__lived.V_apply__macro__16.V_data__16.V_value__16 = V_value__16
	config.Groups.V_group.V_unicast.V_graceful__restart.V_long__lived.V_restarter.V_apply__groups__17 = V_apply__groups__17
	config.Groups.V_group.V_unicast.V_graceful__restart.V_long__lived.V_restarter.V_apply__groups__except__17 = V_apply__groups__except__17
	config.Groups.V_group.V_unicast.V_graceful__restart.V_long__lived.V_restarter.V_apply__macro__17.V_name__35 = V_name__35
	config.Groups.V_group.V_unicast.V_graceful__restart.V_long__lived.V_restarter.V_apply__macro__17.V_data__17.V_name__36 = V_name__36
	config.Groups.V_group.V_unicast.V_graceful__restart.V_long__lived.V_restarter.V_apply__macro__17.V_data__17.V_value__17 = V_value__17
	config.Groups.V_group.V_unicast.V_graceful__restart.V_long__lived.V_restarter.V_stale__time = V_stale__time
	config.Groups.V_group.V_unicast.V_graceful__restart.V_forwarding__state__bit = V_forwarding__state__bit
	config.Groups.V_group.V_unicast.V_extended__nexthop = V_extended__nexthop
	config.Groups.V_group.V_unicast.V_extended__nexthop__color = V_extended__nexthop__color
	config.Groups.V_group.V_unicast.V_extended__nexthop__tunnel = V_extended__nexthop__tunnel
	config.Groups.V_group.V_unicast.V_no__install = V_no__install
	config.Groups.V_group.V_unicast.V_route__age__bgp__view = V_route__age__bgp__view
	config.Groups.V_group.V_unicast.V_advertise__srv6__service = V_advertise__srv6__service
	config.Groups.V_group.V_unicast.V_accept__srv6__service = V_accept__srv6__service
	config.Groups.V_group.V_unicast.V_topology.V_name__37 = V_name__37
	config.Groups.V_group.V_unicast.V_topology.V_apply__groups__18 = V_apply__groups__18
	config.Groups.V_group.V_unicast.V_topology.V_apply__groups__except__18 = V_apply__groups__except__18
	config.Groups.V_group.V_unicast.V_topology.V_apply__macro__18.V_name__38 = V_name__38
	config.Groups.V_group.V_unicast.V_topology.V_apply__macro__18.V_data__18.V_name__39 = V_name__39
	config.Groups.V_group.V_unicast.V_topology.V_apply__macro__18.V_data__18.V_value__18 = V_value__18
	config.Groups.V_group.V_unicast.V_topology.V_community = V_community

    err = client.SendTransaction("", config, commit)
    check(err)
    
    d.SetId(fmt.Sprintf("%s_%s", client.Host, id))
    
	return junosProtocolsBgpGroupFamilyInet6UnicastRead(d,m)
}

func junosProtocolsBgpGroupFamilyInet6UnicastRead(d *schema.ResourceData, m interface{}) error {

	var err error
    client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
	config := &xmlProtocolsBgpGroupFamilyInet6Unicast{}

	err = client.MarshalGroup(id, config)
	check(err)
 	d.Set("name", config.Groups.V_group.V_name)
	d.Set("apply__groups", config.Groups.V_group.V_unicast.V_apply__groups)
	d.Set("apply__groups__except", config.Groups.V_group.V_unicast.V_apply__groups__except)
	d.Set("name__1", config.Groups.V_group.V_unicast.V_apply__macro.V_name__1)
	d.Set("name__2", config.Groups.V_group.V_unicast.V_apply__macro.V_data.V_name__2)
	d.Set("value", config.Groups.V_group.V_unicast.V_apply__macro.V_data.V_value)
	d.Set("apply__groups__1", config.Groups.V_group.V_unicast.V_prefix__limit.V_apply__groups__1)
	d.Set("apply__groups__except__1", config.Groups.V_group.V_unicast.V_prefix__limit.V_apply__groups__except__1)
	d.Set("name__3", config.Groups.V_group.V_unicast.V_prefix__limit.V_apply__macro__1.V_name__3)
	d.Set("name__4", config.Groups.V_group.V_unicast.V_prefix__limit.V_apply__macro__1.V_data__1.V_name__4)
	d.Set("value__1", config.Groups.V_group.V_unicast.V_prefix__limit.V_apply__macro__1.V_data__1.V_value__1)
	d.Set("maximum", config.Groups.V_group.V_unicast.V_prefix__limit.V_maximum)
	d.Set("apply__groups__2", config.Groups.V_group.V_unicast.V_prefix__limit.V_teardown.V_apply__groups__2)
	d.Set("apply__groups__except__2", config.Groups.V_group.V_unicast.V_prefix__limit.V_teardown.V_apply__groups__except__2)
	d.Set("name__5", config.Groups.V_group.V_unicast.V_prefix__limit.V_teardown.V_apply__macro__2.V_name__5)
	d.Set("name__6", config.Groups.V_group.V_unicast.V_prefix__limit.V_teardown.V_apply__macro__2.V_data__2.V_name__6)
	d.Set("value__2", config.Groups.V_group.V_unicast.V_prefix__limit.V_teardown.V_apply__macro__2.V_data__2.V_value__2)
	d.Set("limit__threshold", config.Groups.V_group.V_unicast.V_prefix__limit.V_teardown.V_limit__threshold)
	d.Set("apply__groups__3", config.Groups.V_group.V_unicast.V_accepted__prefix__limit.V_apply__groups__3)
	d.Set("apply__groups__except__3", config.Groups.V_group.V_unicast.V_accepted__prefix__limit.V_apply__groups__except__3)
	d.Set("name__7", config.Groups.V_group.V_unicast.V_accepted__prefix__limit.V_apply__macro__3.V_name__7)
	d.Set("name__8", config.Groups.V_group.V_unicast.V_accepted__prefix__limit.V_apply__macro__3.V_data__3.V_name__8)
	d.Set("value__3", config.Groups.V_group.V_unicast.V_accepted__prefix__limit.V_apply__macro__3.V_data__3.V_value__3)
	d.Set("maximum__1", config.Groups.V_group.V_unicast.V_accepted__prefix__limit.V_maximum__1)
	d.Set("apply__groups__4", config.Groups.V_group.V_unicast.V_accepted__prefix__limit.V_teardown__1.V_apply__groups__4)
	d.Set("apply__groups__except__4", config.Groups.V_group.V_unicast.V_accepted__prefix__limit.V_teardown__1.V_apply__groups__except__4)
	d.Set("name__9", config.Groups.V_group.V_unicast.V_accepted__prefix__limit.V_teardown__1.V_apply__macro__4.V_name__9)
	d.Set("name__10", config.Groups.V_group.V_unicast.V_accepted__prefix__limit.V_teardown__1.V_apply__macro__4.V_data__4.V_name__10)
	d.Set("value__4", config.Groups.V_group.V_unicast.V_accepted__prefix__limit.V_teardown__1.V_apply__macro__4.V_data__4.V_value__4)
	d.Set("limit__threshold__1", config.Groups.V_group.V_unicast.V_accepted__prefix__limit.V_teardown__1.V_limit__threshold__1)
	d.Set("ribgroup__name", config.Groups.V_group.V_unicast.V_rib__group.V_ribgroup__name)
	d.Set("apply__groups__5", config.Groups.V_group.V_unicast.V_add__path.V_apply__groups__5)
	d.Set("apply__groups__except__5", config.Groups.V_group.V_unicast.V_add__path.V_apply__groups__except__5)
	d.Set("name__11", config.Groups.V_group.V_unicast.V_add__path.V_apply__macro__5.V_name__11)
	d.Set("name__12", config.Groups.V_group.V_unicast.V_add__path.V_apply__macro__5.V_data__5.V_name__12)
	d.Set("value__5", config.Groups.V_group.V_unicast.V_add__path.V_apply__macro__5.V_data__5.V_value__5)
	d.Set("receive", config.Groups.V_group.V_unicast.V_add__path.V_receive)
	d.Set("apply__groups__6", config.Groups.V_group.V_unicast.V_add__path.V_send.V_apply__groups__6)
	d.Set("apply__groups__except__6", config.Groups.V_group.V_unicast.V_add__path.V_send.V_apply__groups__except__6)
	d.Set("name__13", config.Groups.V_group.V_unicast.V_add__path.V_send.V_apply__macro__6.V_name__13)
	d.Set("name__14", config.Groups.V_group.V_unicast.V_add__path.V_send.V_apply__macro__6.V_data__6.V_name__14)
	d.Set("value__6", config.Groups.V_group.V_unicast.V_add__path.V_send.V_apply__macro__6.V_data__6.V_value__6)
	d.Set("apply__groups__7", config.Groups.V_group.V_unicast.V_add__path.V_send.V_path__selection__mode.V_apply__groups__7)
	d.Set("apply__groups__except__7", config.Groups.V_group.V_unicast.V_add__path.V_send.V_path__selection__mode.V_apply__groups__except__7)
	d.Set("name__15", config.Groups.V_group.V_unicast.V_add__path.V_send.V_path__selection__mode.V_apply__macro__7.V_name__15)
	d.Set("name__16", config.Groups.V_group.V_unicast.V_add__path.V_send.V_path__selection__mode.V_apply__macro__7.V_data__7.V_name__16)
	d.Set("value__7", config.Groups.V_group.V_unicast.V_add__path.V_send.V_path__selection__mode.V_apply__macro__7.V_data__7.V_value__7)
	d.Set("prefix__policy", config.Groups.V_group.V_unicast.V_add__path.V_send.V_prefix__policy)
	d.Set("path__count", config.Groups.V_group.V_unicast.V_add__path.V_send.V_path__count)
	d.Set("include__backup__path", config.Groups.V_group.V_unicast.V_add__path.V_send.V_include__backup__path)
	d.Set("multipath", config.Groups.V_group.V_unicast.V_add__path.V_send.V_multipath)
	d.Set("apply__groups__8", config.Groups.V_group.V_unicast.V_aigp.V_apply__groups__8)
	d.Set("apply__groups__except__8", config.Groups.V_group.V_unicast.V_aigp.V_apply__groups__except__8)
	d.Set("name__17", config.Groups.V_group.V_unicast.V_aigp.V_apply__macro__8.V_name__17)
	d.Set("name__18", config.Groups.V_group.V_unicast.V_aigp.V_apply__macro__8.V_data__8.V_name__18)
	d.Set("value__8", config.Groups.V_group.V_unicast.V_aigp.V_apply__macro__8.V_data__8.V_value__8)
	d.Set("disable", config.Groups.V_group.V_unicast.V_aigp.V_disable)
	d.Set("damping", config.Groups.V_group.V_unicast.V_damping)
	d.Set("local__ipv4__address", config.Groups.V_group.V_unicast.V_local__ipv4__address)
	d.Set("apply__groups__9", config.Groups.V_group.V_unicast.V_loops.V_apply__groups__9)
	d.Set("apply__groups__except__9", config.Groups.V_group.V_unicast.V_loops.V_apply__groups__except__9)
	d.Set("name__19", config.Groups.V_group.V_unicast.V_loops.V_apply__macro__9.V_name__19)
	d.Set("name__20", config.Groups.V_group.V_unicast.V_loops.V_apply__macro__9.V_data__9.V_name__20)
	d.Set("value__9", config.Groups.V_group.V_unicast.V_loops.V_apply__macro__9.V_data__9.V_value__9)
	d.Set("loops__1", config.Groups.V_group.V_unicast.V_loops.V_loops__1)
	d.Set("apply__groups__10", config.Groups.V_group.V_unicast.V_delay__route__advertisements.V_apply__groups__10)
	d.Set("apply__groups__except__10", config.Groups.V_group.V_unicast.V_delay__route__advertisements.V_apply__groups__except__10)
	d.Set("name__21", config.Groups.V_group.V_unicast.V_delay__route__advertisements.V_apply__macro__10.V_name__21)
	d.Set("name__22", config.Groups.V_group.V_unicast.V_delay__route__advertisements.V_apply__macro__10.V_data__10.V_name__22)
	d.Set("value__10", config.Groups.V_group.V_unicast.V_delay__route__advertisements.V_apply__macro__10.V_data__10.V_value__10)
	d.Set("always__wait__for__krt__drain", config.Groups.V_group.V_unicast.V_delay__route__advertisements.V_always__wait__for__krt__drain)
	d.Set("apply__groups__11", config.Groups.V_group.V_unicast.V_delay__route__advertisements.V_minimum__delay.V_apply__groups__11)
	d.Set("apply__groups__except__11", config.Groups.V_group.V_unicast.V_delay__route__advertisements.V_minimum__delay.V_apply__groups__except__11)
	d.Set("name__23", config.Groups.V_group.V_unicast.V_delay__route__advertisements.V_minimum__delay.V_apply__macro__11.V_name__23)
	d.Set("name__24", config.Groups.V_group.V_unicast.V_delay__route__advertisements.V_minimum__delay.V_apply__macro__11.V_data__11.V_name__24)
	d.Set("value__11", config.Groups.V_group.V_unicast.V_delay__route__advertisements.V_minimum__delay.V_apply__macro__11.V_data__11.V_value__11)
	d.Set("routing__uptime", config.Groups.V_group.V_unicast.V_delay__route__advertisements.V_minimum__delay.V_routing__uptime)
	d.Set("inbound__convergence", config.Groups.V_group.V_unicast.V_delay__route__advertisements.V_minimum__delay.V_inbound__convergence)
	d.Set("apply__groups__12", config.Groups.V_group.V_unicast.V_delay__route__advertisements.V_maximum__delay.V_apply__groups__12)
	d.Set("apply__groups__except__12", config.Groups.V_group.V_unicast.V_delay__route__advertisements.V_maximum__delay.V_apply__groups__except__12)
	d.Set("name__25", config.Groups.V_group.V_unicast.V_delay__route__advertisements.V_maximum__delay.V_apply__macro__12.V_name__25)
	d.Set("name__26", config.Groups.V_group.V_unicast.V_delay__route__advertisements.V_maximum__delay.V_apply__macro__12.V_data__12.V_name__26)
	d.Set("value__12", config.Groups.V_group.V_unicast.V_delay__route__advertisements.V_maximum__delay.V_apply__macro__12.V_data__12.V_value__12)
	d.Set("route__age", config.Groups.V_group.V_unicast.V_delay__route__advertisements.V_maximum__delay.V_route__age)
	d.Set("routing__uptime__1", config.Groups.V_group.V_unicast.V_delay__route__advertisements.V_maximum__delay.V_routing__uptime__1)
	d.Set("apply__groups__13", config.Groups.V_group.V_unicast.V_nexthop__resolution.V_apply__groups__13)
	d.Set("apply__groups__except__13", config.Groups.V_group.V_unicast.V_nexthop__resolution.V_apply__groups__except__13)
	d.Set("name__27", config.Groups.V_group.V_unicast.V_nexthop__resolution.V_apply__macro__13.V_name__27)
	d.Set("name__28", config.Groups.V_group.V_unicast.V_nexthop__resolution.V_apply__macro__13.V_data__13.V_name__28)
	d.Set("value__13", config.Groups.V_group.V_unicast.V_nexthop__resolution.V_apply__macro__13.V_data__13.V_value__13)
	d.Set("no__resolution", config.Groups.V_group.V_unicast.V_nexthop__resolution.V_no__resolution)
	d.Set("preserve__nexthop__hierarchy", config.Groups.V_group.V_unicast.V_nexthop__resolution.V_preserve__nexthop__hierarchy)
	d.Set("apply__groups__14", config.Groups.V_group.V_unicast.V_defer__initial__multipath__build.V_apply__groups__14)
	d.Set("apply__groups__except__14", config.Groups.V_group.V_unicast.V_defer__initial__multipath__build.V_apply__groups__except__14)
	d.Set("name__29", config.Groups.V_group.V_unicast.V_defer__initial__multipath__build.V_apply__macro__14.V_name__29)
	d.Set("name__30", config.Groups.V_group.V_unicast.V_defer__initial__multipath__build.V_apply__macro__14.V_data__14.V_name__30)
	d.Set("value__14", config.Groups.V_group.V_unicast.V_defer__initial__multipath__build.V_apply__macro__14.V_data__14.V_value__14)
	d.Set("maximum__delay__1", config.Groups.V_group.V_unicast.V_defer__initial__multipath__build.V_maximum__delay__1)
	d.Set("apply__groups__15", config.Groups.V_group.V_unicast.V_graceful__restart.V_apply__groups__15)
	d.Set("apply__groups__except__15", config.Groups.V_group.V_unicast.V_graceful__restart.V_apply__groups__except__15)
	d.Set("name__31", config.Groups.V_group.V_unicast.V_graceful__restart.V_apply__macro__15.V_name__31)
	d.Set("name__32", config.Groups.V_group.V_unicast.V_graceful__restart.V_apply__macro__15.V_data__15.V_name__32)
	d.Set("value__15", config.Groups.V_group.V_unicast.V_graceful__restart.V_apply__macro__15.V_data__15.V_value__15)
	d.Set("apply__groups__16", config.Groups.V_group.V_unicast.V_graceful__restart.V_long__lived.V_apply__groups__16)
	d.Set("apply__groups__except__16", config.Groups.V_group.V_unicast.V_graceful__restart.V_long__lived.V_apply__groups__except__16)
	d.Set("name__33", config.Groups.V_group.V_unicast.V_graceful__restart.V_long__lived.V_apply__macro__16.V_name__33)
	d.Set("name__34", config.Groups.V_group.V_unicast.V_graceful__restart.V_long__lived.V_apply__macro__16.V_data__16.V_name__34)
	d.Set("value__16", config.Groups.V_group.V_unicast.V_graceful__restart.V_long__lived.V_apply__macro__16.V_data__16.V_value__16)
	d.Set("apply__groups__17", config.Groups.V_group.V_unicast.V_graceful__restart.V_long__lived.V_restarter.V_apply__groups__17)
	d.Set("apply__groups__except__17", config.Groups.V_group.V_unicast.V_graceful__restart.V_long__lived.V_restarter.V_apply__groups__except__17)
	d.Set("name__35", config.Groups.V_group.V_unicast.V_graceful__restart.V_long__lived.V_restarter.V_apply__macro__17.V_name__35)
	d.Set("name__36", config.Groups.V_group.V_unicast.V_graceful__restart.V_long__lived.V_restarter.V_apply__macro__17.V_data__17.V_name__36)
	d.Set("value__17", config.Groups.V_group.V_unicast.V_graceful__restart.V_long__lived.V_restarter.V_apply__macro__17.V_data__17.V_value__17)
	d.Set("stale__time", config.Groups.V_group.V_unicast.V_graceful__restart.V_long__lived.V_restarter.V_stale__time)
	d.Set("forwarding__state__bit", config.Groups.V_group.V_unicast.V_graceful__restart.V_forwarding__state__bit)
	d.Set("extended__nexthop", config.Groups.V_group.V_unicast.V_extended__nexthop)
	d.Set("extended__nexthop__color", config.Groups.V_group.V_unicast.V_extended__nexthop__color)
	d.Set("extended__nexthop__tunnel", config.Groups.V_group.V_unicast.V_extended__nexthop__tunnel)
	d.Set("no__install", config.Groups.V_group.V_unicast.V_no__install)
	d.Set("route__age__bgp__view", config.Groups.V_group.V_unicast.V_route__age__bgp__view)
	d.Set("advertise__srv6__service", config.Groups.V_group.V_unicast.V_advertise__srv6__service)
	d.Set("accept__srv6__service", config.Groups.V_group.V_unicast.V_accept__srv6__service)
	d.Set("name__37", config.Groups.V_group.V_unicast.V_topology.V_name__37)
	d.Set("apply__groups__18", config.Groups.V_group.V_unicast.V_topology.V_apply__groups__18)
	d.Set("apply__groups__except__18", config.Groups.V_group.V_unicast.V_topology.V_apply__groups__except__18)
	d.Set("name__38", config.Groups.V_group.V_unicast.V_topology.V_apply__macro__18.V_name__38)
	d.Set("name__39", config.Groups.V_group.V_unicast.V_topology.V_apply__macro__18.V_data__18.V_name__39)
	d.Set("value__18", config.Groups.V_group.V_unicast.V_topology.V_apply__macro__18.V_data__18.V_value__18)
	d.Set("community", config.Groups.V_group.V_unicast.V_topology.V_community)

	return nil
}

func junosProtocolsBgpGroupFamilyInet6UnicastUpdate(d *schema.ResourceData, m interface{}) error {

	var err error
    client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
     	V_name := d.Get("name").(string)
	V_apply__groups := d.Get("apply__groups").(string)
	V_apply__groups__except := d.Get("apply__groups__except").(string)
	V_name__1 := d.Get("name__1").(string)
	V_name__2 := d.Get("name__2").(string)
	V_value := d.Get("value").(string)
	V_apply__groups__1 := d.Get("apply__groups__1").(string)
	V_apply__groups__except__1 := d.Get("apply__groups__except__1").(string)
	V_name__3 := d.Get("name__3").(string)
	V_name__4 := d.Get("name__4").(string)
	V_value__1 := d.Get("value__1").(string)
	V_maximum := d.Get("maximum").(string)
	V_apply__groups__2 := d.Get("apply__groups__2").(string)
	V_apply__groups__except__2 := d.Get("apply__groups__except__2").(string)
	V_name__5 := d.Get("name__5").(string)
	V_name__6 := d.Get("name__6").(string)
	V_value__2 := d.Get("value__2").(string)
	V_limit__threshold := d.Get("limit__threshold").(string)
	V_apply__groups__3 := d.Get("apply__groups__3").(string)
	V_apply__groups__except__3 := d.Get("apply__groups__except__3").(string)
	V_name__7 := d.Get("name__7").(string)
	V_name__8 := d.Get("name__8").(string)
	V_value__3 := d.Get("value__3").(string)
	V_maximum__1 := d.Get("maximum__1").(string)
	V_apply__groups__4 := d.Get("apply__groups__4").(string)
	V_apply__groups__except__4 := d.Get("apply__groups__except__4").(string)
	V_name__9 := d.Get("name__9").(string)
	V_name__10 := d.Get("name__10").(string)
	V_value__4 := d.Get("value__4").(string)
	V_limit__threshold__1 := d.Get("limit__threshold__1").(string)
	V_ribgroup__name := d.Get("ribgroup__name").(string)
	V_apply__groups__5 := d.Get("apply__groups__5").(string)
	V_apply__groups__except__5 := d.Get("apply__groups__except__5").(string)
	V_name__11 := d.Get("name__11").(string)
	V_name__12 := d.Get("name__12").(string)
	V_value__5 := d.Get("value__5").(string)
	V_receive := d.Get("receive").(string)
	V_apply__groups__6 := d.Get("apply__groups__6").(string)
	V_apply__groups__except__6 := d.Get("apply__groups__except__6").(string)
	V_name__13 := d.Get("name__13").(string)
	V_name__14 := d.Get("name__14").(string)
	V_value__6 := d.Get("value__6").(string)
	V_apply__groups__7 := d.Get("apply__groups__7").(string)
	V_apply__groups__except__7 := d.Get("apply__groups__except__7").(string)
	V_name__15 := d.Get("name__15").(string)
	V_name__16 := d.Get("name__16").(string)
	V_value__7 := d.Get("value__7").(string)
	V_prefix__policy := d.Get("prefix__policy").(string)
	V_path__count := d.Get("path__count").(string)
	V_include__backup__path := d.Get("include__backup__path").(string)
	V_multipath := d.Get("multipath").(string)
	V_apply__groups__8 := d.Get("apply__groups__8").(string)
	V_apply__groups__except__8 := d.Get("apply__groups__except__8").(string)
	V_name__17 := d.Get("name__17").(string)
	V_name__18 := d.Get("name__18").(string)
	V_value__8 := d.Get("value__8").(string)
	V_disable := d.Get("disable").(string)
	V_damping := d.Get("damping").(string)
	V_local__ipv4__address := d.Get("local__ipv4__address").(string)
	V_apply__groups__9 := d.Get("apply__groups__9").(string)
	V_apply__groups__except__9 := d.Get("apply__groups__except__9").(string)
	V_name__19 := d.Get("name__19").(string)
	V_name__20 := d.Get("name__20").(string)
	V_value__9 := d.Get("value__9").(string)
	V_loops__1 := d.Get("loops__1").(string)
	V_apply__groups__10 := d.Get("apply__groups__10").(string)
	V_apply__groups__except__10 := d.Get("apply__groups__except__10").(string)
	V_name__21 := d.Get("name__21").(string)
	V_name__22 := d.Get("name__22").(string)
	V_value__10 := d.Get("value__10").(string)
	V_always__wait__for__krt__drain := d.Get("always__wait__for__krt__drain").(string)
	V_apply__groups__11 := d.Get("apply__groups__11").(string)
	V_apply__groups__except__11 := d.Get("apply__groups__except__11").(string)
	V_name__23 := d.Get("name__23").(string)
	V_name__24 := d.Get("name__24").(string)
	V_value__11 := d.Get("value__11").(string)
	V_routing__uptime := d.Get("routing__uptime").(string)
	V_inbound__convergence := d.Get("inbound__convergence").(string)
	V_apply__groups__12 := d.Get("apply__groups__12").(string)
	V_apply__groups__except__12 := d.Get("apply__groups__except__12").(string)
	V_name__25 := d.Get("name__25").(string)
	V_name__26 := d.Get("name__26").(string)
	V_value__12 := d.Get("value__12").(string)
	V_route__age := d.Get("route__age").(string)
	V_routing__uptime__1 := d.Get("routing__uptime__1").(string)
	V_apply__groups__13 := d.Get("apply__groups__13").(string)
	V_apply__groups__except__13 := d.Get("apply__groups__except__13").(string)
	V_name__27 := d.Get("name__27").(string)
	V_name__28 := d.Get("name__28").(string)
	V_value__13 := d.Get("value__13").(string)
	V_no__resolution := d.Get("no__resolution").(string)
	V_preserve__nexthop__hierarchy := d.Get("preserve__nexthop__hierarchy").(string)
	V_apply__groups__14 := d.Get("apply__groups__14").(string)
	V_apply__groups__except__14 := d.Get("apply__groups__except__14").(string)
	V_name__29 := d.Get("name__29").(string)
	V_name__30 := d.Get("name__30").(string)
	V_value__14 := d.Get("value__14").(string)
	V_maximum__delay__1 := d.Get("maximum__delay__1").(string)
	V_apply__groups__15 := d.Get("apply__groups__15").(string)
	V_apply__groups__except__15 := d.Get("apply__groups__except__15").(string)
	V_name__31 := d.Get("name__31").(string)
	V_name__32 := d.Get("name__32").(string)
	V_value__15 := d.Get("value__15").(string)
	V_apply__groups__16 := d.Get("apply__groups__16").(string)
	V_apply__groups__except__16 := d.Get("apply__groups__except__16").(string)
	V_name__33 := d.Get("name__33").(string)
	V_name__34 := d.Get("name__34").(string)
	V_value__16 := d.Get("value__16").(string)
	V_apply__groups__17 := d.Get("apply__groups__17").(string)
	V_apply__groups__except__17 := d.Get("apply__groups__except__17").(string)
	V_name__35 := d.Get("name__35").(string)
	V_name__36 := d.Get("name__36").(string)
	V_value__17 := d.Get("value__17").(string)
	V_stale__time := d.Get("stale__time").(string)
	V_forwarding__state__bit := d.Get("forwarding__state__bit").(string)
	V_extended__nexthop := d.Get("extended__nexthop").(string)
	V_extended__nexthop__color := d.Get("extended__nexthop__color").(string)
	V_extended__nexthop__tunnel := d.Get("extended__nexthop__tunnel").(string)
	V_no__install := d.Get("no__install").(string)
	V_route__age__bgp__view := d.Get("route__age__bgp__view").(string)
	V_advertise__srv6__service := d.Get("advertise__srv6__service").(string)
	V_accept__srv6__service := d.Get("accept__srv6__service").(string)
	V_name__37 := d.Get("name__37").(string)
	V_apply__groups__18 := d.Get("apply__groups__18").(string)
	V_apply__groups__except__18 := d.Get("apply__groups__except__18").(string)
	V_name__38 := d.Get("name__38").(string)
	V_name__39 := d.Get("name__39").(string)
	V_value__18 := d.Get("value__18").(string)
	V_community := d.Get("community").(string)
	commit := true

	config := xmlProtocolsBgpGroupFamilyInet6Unicast{}
	config.ApplyGroup = id
	config.Groups.Name = id
	config.Groups.V_group.V_name = V_name
	config.Groups.V_group.V_unicast.V_apply__groups = V_apply__groups
	config.Groups.V_group.V_unicast.V_apply__groups__except = V_apply__groups__except
	config.Groups.V_group.V_unicast.V_apply__macro.V_name__1 = V_name__1
	config.Groups.V_group.V_unicast.V_apply__macro.V_data.V_name__2 = V_name__2
	config.Groups.V_group.V_unicast.V_apply__macro.V_data.V_value = V_value
	config.Groups.V_group.V_unicast.V_prefix__limit.V_apply__groups__1 = V_apply__groups__1
	config.Groups.V_group.V_unicast.V_prefix__limit.V_apply__groups__except__1 = V_apply__groups__except__1
	config.Groups.V_group.V_unicast.V_prefix__limit.V_apply__macro__1.V_name__3 = V_name__3
	config.Groups.V_group.V_unicast.V_prefix__limit.V_apply__macro__1.V_data__1.V_name__4 = V_name__4
	config.Groups.V_group.V_unicast.V_prefix__limit.V_apply__macro__1.V_data__1.V_value__1 = V_value__1
	config.Groups.V_group.V_unicast.V_prefix__limit.V_maximum = V_maximum
	config.Groups.V_group.V_unicast.V_prefix__limit.V_teardown.V_apply__groups__2 = V_apply__groups__2
	config.Groups.V_group.V_unicast.V_prefix__limit.V_teardown.V_apply__groups__except__2 = V_apply__groups__except__2
	config.Groups.V_group.V_unicast.V_prefix__limit.V_teardown.V_apply__macro__2.V_name__5 = V_name__5
	config.Groups.V_group.V_unicast.V_prefix__limit.V_teardown.V_apply__macro__2.V_data__2.V_name__6 = V_name__6
	config.Groups.V_group.V_unicast.V_prefix__limit.V_teardown.V_apply__macro__2.V_data__2.V_value__2 = V_value__2
	config.Groups.V_group.V_unicast.V_prefix__limit.V_teardown.V_limit__threshold = V_limit__threshold
	config.Groups.V_group.V_unicast.V_accepted__prefix__limit.V_apply__groups__3 = V_apply__groups__3
	config.Groups.V_group.V_unicast.V_accepted__prefix__limit.V_apply__groups__except__3 = V_apply__groups__except__3
	config.Groups.V_group.V_unicast.V_accepted__prefix__limit.V_apply__macro__3.V_name__7 = V_name__7
	config.Groups.V_group.V_unicast.V_accepted__prefix__limit.V_apply__macro__3.V_data__3.V_name__8 = V_name__8
	config.Groups.V_group.V_unicast.V_accepted__prefix__limit.V_apply__macro__3.V_data__3.V_value__3 = V_value__3
	config.Groups.V_group.V_unicast.V_accepted__prefix__limit.V_maximum__1 = V_maximum__1
	config.Groups.V_group.V_unicast.V_accepted__prefix__limit.V_teardown__1.V_apply__groups__4 = V_apply__groups__4
	config.Groups.V_group.V_unicast.V_accepted__prefix__limit.V_teardown__1.V_apply__groups__except__4 = V_apply__groups__except__4
	config.Groups.V_group.V_unicast.V_accepted__prefix__limit.V_teardown__1.V_apply__macro__4.V_name__9 = V_name__9
	config.Groups.V_group.V_unicast.V_accepted__prefix__limit.V_teardown__1.V_apply__macro__4.V_data__4.V_name__10 = V_name__10
	config.Groups.V_group.V_unicast.V_accepted__prefix__limit.V_teardown__1.V_apply__macro__4.V_data__4.V_value__4 = V_value__4
	config.Groups.V_group.V_unicast.V_accepted__prefix__limit.V_teardown__1.V_limit__threshold__1 = V_limit__threshold__1
	config.Groups.V_group.V_unicast.V_rib__group.V_ribgroup__name = V_ribgroup__name
	config.Groups.V_group.V_unicast.V_add__path.V_apply__groups__5 = V_apply__groups__5
	config.Groups.V_group.V_unicast.V_add__path.V_apply__groups__except__5 = V_apply__groups__except__5
	config.Groups.V_group.V_unicast.V_add__path.V_apply__macro__5.V_name__11 = V_name__11
	config.Groups.V_group.V_unicast.V_add__path.V_apply__macro__5.V_data__5.V_name__12 = V_name__12
	config.Groups.V_group.V_unicast.V_add__path.V_apply__macro__5.V_data__5.V_value__5 = V_value__5
	config.Groups.V_group.V_unicast.V_add__path.V_receive = V_receive
	config.Groups.V_group.V_unicast.V_add__path.V_send.V_apply__groups__6 = V_apply__groups__6
	config.Groups.V_group.V_unicast.V_add__path.V_send.V_apply__groups__except__6 = V_apply__groups__except__6
	config.Groups.V_group.V_unicast.V_add__path.V_send.V_apply__macro__6.V_name__13 = V_name__13
	config.Groups.V_group.V_unicast.V_add__path.V_send.V_apply__macro__6.V_data__6.V_name__14 = V_name__14
	config.Groups.V_group.V_unicast.V_add__path.V_send.V_apply__macro__6.V_data__6.V_value__6 = V_value__6
	config.Groups.V_group.V_unicast.V_add__path.V_send.V_path__selection__mode.V_apply__groups__7 = V_apply__groups__7
	config.Groups.V_group.V_unicast.V_add__path.V_send.V_path__selection__mode.V_apply__groups__except__7 = V_apply__groups__except__7
	config.Groups.V_group.V_unicast.V_add__path.V_send.V_path__selection__mode.V_apply__macro__7.V_name__15 = V_name__15
	config.Groups.V_group.V_unicast.V_add__path.V_send.V_path__selection__mode.V_apply__macro__7.V_data__7.V_name__16 = V_name__16
	config.Groups.V_group.V_unicast.V_add__path.V_send.V_path__selection__mode.V_apply__macro__7.V_data__7.V_value__7 = V_value__7
	config.Groups.V_group.V_unicast.V_add__path.V_send.V_prefix__policy = V_prefix__policy
	config.Groups.V_group.V_unicast.V_add__path.V_send.V_path__count = V_path__count
	config.Groups.V_group.V_unicast.V_add__path.V_send.V_include__backup__path = V_include__backup__path
	config.Groups.V_group.V_unicast.V_add__path.V_send.V_multipath = V_multipath
	config.Groups.V_group.V_unicast.V_aigp.V_apply__groups__8 = V_apply__groups__8
	config.Groups.V_group.V_unicast.V_aigp.V_apply__groups__except__8 = V_apply__groups__except__8
	config.Groups.V_group.V_unicast.V_aigp.V_apply__macro__8.V_name__17 = V_name__17
	config.Groups.V_group.V_unicast.V_aigp.V_apply__macro__8.V_data__8.V_name__18 = V_name__18
	config.Groups.V_group.V_unicast.V_aigp.V_apply__macro__8.V_data__8.V_value__8 = V_value__8
	config.Groups.V_group.V_unicast.V_aigp.V_disable = V_disable
	config.Groups.V_group.V_unicast.V_damping = V_damping
	config.Groups.V_group.V_unicast.V_local__ipv4__address = V_local__ipv4__address
	config.Groups.V_group.V_unicast.V_loops.V_apply__groups__9 = V_apply__groups__9
	config.Groups.V_group.V_unicast.V_loops.V_apply__groups__except__9 = V_apply__groups__except__9
	config.Groups.V_group.V_unicast.V_loops.V_apply__macro__9.V_name__19 = V_name__19
	config.Groups.V_group.V_unicast.V_loops.V_apply__macro__9.V_data__9.V_name__20 = V_name__20
	config.Groups.V_group.V_unicast.V_loops.V_apply__macro__9.V_data__9.V_value__9 = V_value__9
	config.Groups.V_group.V_unicast.V_loops.V_loops__1 = V_loops__1
	config.Groups.V_group.V_unicast.V_delay__route__advertisements.V_apply__groups__10 = V_apply__groups__10
	config.Groups.V_group.V_unicast.V_delay__route__advertisements.V_apply__groups__except__10 = V_apply__groups__except__10
	config.Groups.V_group.V_unicast.V_delay__route__advertisements.V_apply__macro__10.V_name__21 = V_name__21
	config.Groups.V_group.V_unicast.V_delay__route__advertisements.V_apply__macro__10.V_data__10.V_name__22 = V_name__22
	config.Groups.V_group.V_unicast.V_delay__route__advertisements.V_apply__macro__10.V_data__10.V_value__10 = V_value__10
	config.Groups.V_group.V_unicast.V_delay__route__advertisements.V_always__wait__for__krt__drain = V_always__wait__for__krt__drain
	config.Groups.V_group.V_unicast.V_delay__route__advertisements.V_minimum__delay.V_apply__groups__11 = V_apply__groups__11
	config.Groups.V_group.V_unicast.V_delay__route__advertisements.V_minimum__delay.V_apply__groups__except__11 = V_apply__groups__except__11
	config.Groups.V_group.V_unicast.V_delay__route__advertisements.V_minimum__delay.V_apply__macro__11.V_name__23 = V_name__23
	config.Groups.V_group.V_unicast.V_delay__route__advertisements.V_minimum__delay.V_apply__macro__11.V_data__11.V_name__24 = V_name__24
	config.Groups.V_group.V_unicast.V_delay__route__advertisements.V_minimum__delay.V_apply__macro__11.V_data__11.V_value__11 = V_value__11
	config.Groups.V_group.V_unicast.V_delay__route__advertisements.V_minimum__delay.V_routing__uptime = V_routing__uptime
	config.Groups.V_group.V_unicast.V_delay__route__advertisements.V_minimum__delay.V_inbound__convergence = V_inbound__convergence
	config.Groups.V_group.V_unicast.V_delay__route__advertisements.V_maximum__delay.V_apply__groups__12 = V_apply__groups__12
	config.Groups.V_group.V_unicast.V_delay__route__advertisements.V_maximum__delay.V_apply__groups__except__12 = V_apply__groups__except__12
	config.Groups.V_group.V_unicast.V_delay__route__advertisements.V_maximum__delay.V_apply__macro__12.V_name__25 = V_name__25
	config.Groups.V_group.V_unicast.V_delay__route__advertisements.V_maximum__delay.V_apply__macro__12.V_data__12.V_name__26 = V_name__26
	config.Groups.V_group.V_unicast.V_delay__route__advertisements.V_maximum__delay.V_apply__macro__12.V_data__12.V_value__12 = V_value__12
	config.Groups.V_group.V_unicast.V_delay__route__advertisements.V_maximum__delay.V_route__age = V_route__age
	config.Groups.V_group.V_unicast.V_delay__route__advertisements.V_maximum__delay.V_routing__uptime__1 = V_routing__uptime__1
	config.Groups.V_group.V_unicast.V_nexthop__resolution.V_apply__groups__13 = V_apply__groups__13
	config.Groups.V_group.V_unicast.V_nexthop__resolution.V_apply__groups__except__13 = V_apply__groups__except__13
	config.Groups.V_group.V_unicast.V_nexthop__resolution.V_apply__macro__13.V_name__27 = V_name__27
	config.Groups.V_group.V_unicast.V_nexthop__resolution.V_apply__macro__13.V_data__13.V_name__28 = V_name__28
	config.Groups.V_group.V_unicast.V_nexthop__resolution.V_apply__macro__13.V_data__13.V_value__13 = V_value__13
	config.Groups.V_group.V_unicast.V_nexthop__resolution.V_no__resolution = V_no__resolution
	config.Groups.V_group.V_unicast.V_nexthop__resolution.V_preserve__nexthop__hierarchy = V_preserve__nexthop__hierarchy
	config.Groups.V_group.V_unicast.V_defer__initial__multipath__build.V_apply__groups__14 = V_apply__groups__14
	config.Groups.V_group.V_unicast.V_defer__initial__multipath__build.V_apply__groups__except__14 = V_apply__groups__except__14
	config.Groups.V_group.V_unicast.V_defer__initial__multipath__build.V_apply__macro__14.V_name__29 = V_name__29
	config.Groups.V_group.V_unicast.V_defer__initial__multipath__build.V_apply__macro__14.V_data__14.V_name__30 = V_name__30
	config.Groups.V_group.V_unicast.V_defer__initial__multipath__build.V_apply__macro__14.V_data__14.V_value__14 = V_value__14
	config.Groups.V_group.V_unicast.V_defer__initial__multipath__build.V_maximum__delay__1 = V_maximum__delay__1
	config.Groups.V_group.V_unicast.V_graceful__restart.V_apply__groups__15 = V_apply__groups__15
	config.Groups.V_group.V_unicast.V_graceful__restart.V_apply__groups__except__15 = V_apply__groups__except__15
	config.Groups.V_group.V_unicast.V_graceful__restart.V_apply__macro__15.V_name__31 = V_name__31
	config.Groups.V_group.V_unicast.V_graceful__restart.V_apply__macro__15.V_data__15.V_name__32 = V_name__32
	config.Groups.V_group.V_unicast.V_graceful__restart.V_apply__macro__15.V_data__15.V_value__15 = V_value__15
	config.Groups.V_group.V_unicast.V_graceful__restart.V_long__lived.V_apply__groups__16 = V_apply__groups__16
	config.Groups.V_group.V_unicast.V_graceful__restart.V_long__lived.V_apply__groups__except__16 = V_apply__groups__except__16
	config.Groups.V_group.V_unicast.V_graceful__restart.V_long__lived.V_apply__macro__16.V_name__33 = V_name__33
	config.Groups.V_group.V_unicast.V_graceful__restart.V_long__lived.V_apply__macro__16.V_data__16.V_name__34 = V_name__34
	config.Groups.V_group.V_unicast.V_graceful__restart.V_long__lived.V_apply__macro__16.V_data__16.V_value__16 = V_value__16
	config.Groups.V_group.V_unicast.V_graceful__restart.V_long__lived.V_restarter.V_apply__groups__17 = V_apply__groups__17
	config.Groups.V_group.V_unicast.V_graceful__restart.V_long__lived.V_restarter.V_apply__groups__except__17 = V_apply__groups__except__17
	config.Groups.V_group.V_unicast.V_graceful__restart.V_long__lived.V_restarter.V_apply__macro__17.V_name__35 = V_name__35
	config.Groups.V_group.V_unicast.V_graceful__restart.V_long__lived.V_restarter.V_apply__macro__17.V_data__17.V_name__36 = V_name__36
	config.Groups.V_group.V_unicast.V_graceful__restart.V_long__lived.V_restarter.V_apply__macro__17.V_data__17.V_value__17 = V_value__17
	config.Groups.V_group.V_unicast.V_graceful__restart.V_long__lived.V_restarter.V_stale__time = V_stale__time
	config.Groups.V_group.V_unicast.V_graceful__restart.V_forwarding__state__bit = V_forwarding__state__bit
	config.Groups.V_group.V_unicast.V_extended__nexthop = V_extended__nexthop
	config.Groups.V_group.V_unicast.V_extended__nexthop__color = V_extended__nexthop__color
	config.Groups.V_group.V_unicast.V_extended__nexthop__tunnel = V_extended__nexthop__tunnel
	config.Groups.V_group.V_unicast.V_no__install = V_no__install
	config.Groups.V_group.V_unicast.V_route__age__bgp__view = V_route__age__bgp__view
	config.Groups.V_group.V_unicast.V_advertise__srv6__service = V_advertise__srv6__service
	config.Groups.V_group.V_unicast.V_accept__srv6__service = V_accept__srv6__service
	config.Groups.V_group.V_unicast.V_topology.V_name__37 = V_name__37
	config.Groups.V_group.V_unicast.V_topology.V_apply__groups__18 = V_apply__groups__18
	config.Groups.V_group.V_unicast.V_topology.V_apply__groups__except__18 = V_apply__groups__except__18
	config.Groups.V_group.V_unicast.V_topology.V_apply__macro__18.V_name__38 = V_name__38
	config.Groups.V_group.V_unicast.V_topology.V_apply__macro__18.V_data__18.V_name__39 = V_name__39
	config.Groups.V_group.V_unicast.V_topology.V_apply__macro__18.V_data__18.V_value__18 = V_value__18
	config.Groups.V_group.V_unicast.V_topology.V_community = V_community

    err = client.SendTransaction(id, config, commit)
    check(err)
    
	return junosProtocolsBgpGroupFamilyInet6UnicastRead(d,m)
}

func junosProtocolsBgpGroupFamilyInet6UnicastDelete(d *schema.ResourceData, m interface{}) error {

	var err error
    client := m.(*ProviderConfig)

    id := d.Get("resource_name").(string)
    
    _, err = client.DeleteConfig(id)
    check(err)

    d.SetId("")
    
	return nil
}

func junosProtocolsBgpGroupFamilyInet6Unicast() *schema.Resource {
	return &schema.Resource{
		Create: junosProtocolsBgpGroupFamilyInet6UnicastCreate,
		Read: junosProtocolsBgpGroupFamilyInet6UnicastRead,
		Update: junosProtocolsBgpGroupFamilyInet6UnicastUpdate,
		Delete: junosProtocolsBgpGroupFamilyInet6UnicastDelete,

        Schema: map[string]*schema.Schema{
            "resource_name": &schema.Schema{
                Type:    schema.TypeString,
                Required: true,
            },
			"name": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_group",
			},
			"apply__groups": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_group.V_unicast. Groups from which to inherit configuration data",
			},
			"apply__groups__except": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_group.V_unicast. Don't inherit configuration data from these groups",
			},
			"name__1": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_group.V_unicast.V_apply__macro. Name of the macro to be expanded",
			},
			"name__2": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_group.V_unicast.V_apply__macro.V_data. Keyword part of the keyword-value pair",
			},
			"value": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_group.V_unicast.V_apply__macro.V_data. Value part of the keyword-value pair",
			},
			"apply__groups__1": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_group.V_unicast.V_prefix__limit. Groups from which to inherit configuration data",
			},
			"apply__groups__except__1": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_group.V_unicast.V_prefix__limit. Don't inherit configuration data from these groups",
			},
			"name__3": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_group.V_unicast.V_prefix__limit.V_apply__macro__1. Name of the macro to be expanded",
			},
			"name__4": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_group.V_unicast.V_prefix__limit.V_apply__macro__1.V_data__1. Keyword part of the keyword-value pair",
			},
			"value__1": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_group.V_unicast.V_prefix__limit.V_apply__macro__1.V_data__1. Value part of the keyword-value pair",
			},
			"maximum": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_group.V_unicast.V_prefix__limit. Maximum number of prefixes from a peer",
			},
			"apply__groups__2": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_group.V_unicast.V_prefix__limit.V_teardown. Groups from which to inherit configuration data",
			},
			"apply__groups__except__2": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_group.V_unicast.V_prefix__limit.V_teardown. Don't inherit configuration data from these groups",
			},
			"name__5": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_group.V_unicast.V_prefix__limit.V_teardown.V_apply__macro__2. Name of the macro to be expanded",
			},
			"name__6": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_group.V_unicast.V_prefix__limit.V_teardown.V_apply__macro__2.V_data__2. Keyword part of the keyword-value pair",
			},
			"value__2": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_group.V_unicast.V_prefix__limit.V_teardown.V_apply__macro__2.V_data__2. Value part of the keyword-value pair",
			},
			"limit__threshold": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_group.V_unicast.V_prefix__limit.V_teardown. Percentage of prefix-limit to start warnings",
			},
			"apply__groups__3": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_group.V_unicast.V_accepted__prefix__limit. Groups from which to inherit configuration data",
			},
			"apply__groups__except__3": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_group.V_unicast.V_accepted__prefix__limit. Don't inherit configuration data from these groups",
			},
			"name__7": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_group.V_unicast.V_accepted__prefix__limit.V_apply__macro__3. Name of the macro to be expanded",
			},
			"name__8": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_group.V_unicast.V_accepted__prefix__limit.V_apply__macro__3.V_data__3. Keyword part of the keyword-value pair",
			},
			"value__3": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_group.V_unicast.V_accepted__prefix__limit.V_apply__macro__3.V_data__3. Value part of the keyword-value pair",
			},
			"maximum__1": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_group.V_unicast.V_accepted__prefix__limit. Maximum number of prefixes accepted from a peer",
			},
			"apply__groups__4": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_group.V_unicast.V_accepted__prefix__limit.V_teardown__1. Groups from which to inherit configuration data",
			},
			"apply__groups__except__4": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_group.V_unicast.V_accepted__prefix__limit.V_teardown__1. Don't inherit configuration data from these groups",
			},
			"name__9": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_group.V_unicast.V_accepted__prefix__limit.V_teardown__1.V_apply__macro__4. Name of the macro to be expanded",
			},
			"name__10": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_group.V_unicast.V_accepted__prefix__limit.V_teardown__1.V_apply__macro__4.V_data__4. Keyword part of the keyword-value pair",
			},
			"value__4": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_group.V_unicast.V_accepted__prefix__limit.V_teardown__1.V_apply__macro__4.V_data__4. Value part of the keyword-value pair",
			},
			"limit__threshold__1": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_group.V_unicast.V_accepted__prefix__limit.V_teardown__1. Percentage of prefix-limit to start warnings",
			},
			"ribgroup__name": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_group.V_unicast.V_rib__group. Name of the routing table group",
			},
			"apply__groups__5": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_group.V_unicast.V_add__path. Groups from which to inherit configuration data",
			},
			"apply__groups__except__5": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_group.V_unicast.V_add__path. Don't inherit configuration data from these groups",
			},
			"name__11": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_group.V_unicast.V_add__path.V_apply__macro__5. Name of the macro to be expanded",
			},
			"name__12": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_group.V_unicast.V_add__path.V_apply__macro__5.V_data__5. Keyword part of the keyword-value pair",
			},
			"value__5": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_group.V_unicast.V_add__path.V_apply__macro__5.V_data__5. Value part of the keyword-value pair",
			},
			"receive": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_group.V_unicast.V_add__path. Receive multiple paths from peer",
			},
			"apply__groups__6": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_group.V_unicast.V_add__path.V_send. Groups from which to inherit configuration data",
			},
			"apply__groups__except__6": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_group.V_unicast.V_add__path.V_send. Don't inherit configuration data from these groups",
			},
			"name__13": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_group.V_unicast.V_add__path.V_send.V_apply__macro__6. Name of the macro to be expanded",
			},
			"name__14": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_group.V_unicast.V_add__path.V_send.V_apply__macro__6.V_data__6. Keyword part of the keyword-value pair",
			},
			"value__6": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_group.V_unicast.V_add__path.V_send.V_apply__macro__6.V_data__6. Value part of the keyword-value pair",
			},
			"apply__groups__7": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_group.V_unicast.V_add__path.V_send.V_path__selection__mode. Groups from which to inherit configuration data",
			},
			"apply__groups__except__7": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_group.V_unicast.V_add__path.V_send.V_path__selection__mode. Don't inherit configuration data from these groups",
			},
			"name__15": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_group.V_unicast.V_add__path.V_send.V_path__selection__mode.V_apply__macro__7. Name of the macro to be expanded",
			},
			"name__16": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_group.V_unicast.V_add__path.V_send.V_path__selection__mode.V_apply__macro__7.V_data__7. Keyword part of the keyword-value pair",
			},
			"value__7": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_group.V_unicast.V_add__path.V_send.V_path__selection__mode.V_apply__macro__7.V_data__7. Value part of the keyword-value pair",
			},
			"prefix__policy": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_group.V_unicast.V_add__path.V_send. Perform add-path only for prefixes that match policy",
			},
			"path__count": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_group.V_unicast.V_add__path.V_send. Number of paths to advertise",
			},
			"include__backup__path": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_group.V_unicast.V_add__path.V_send. Number of backup paths to advertise",
			},
			"multipath": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_group.V_unicast.V_add__path.V_send. Include only multipath contributor routes",
			},
			"apply__groups__8": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_group.V_unicast.V_aigp. Groups from which to inherit configuration data",
			},
			"apply__groups__except__8": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_group.V_unicast.V_aigp. Don't inherit configuration data from these groups",
			},
			"name__17": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_group.V_unicast.V_aigp.V_apply__macro__8. Name of the macro to be expanded",
			},
			"name__18": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_group.V_unicast.V_aigp.V_apply__macro__8.V_data__8. Keyword part of the keyword-value pair",
			},
			"value__8": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_group.V_unicast.V_aigp.V_apply__macro__8.V_data__8. Value part of the keyword-value pair",
			},
			"disable": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_group.V_unicast.V_aigp. Disable sending and receiving of AIGP attribute",
			},
			"damping": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_group.V_unicast. Enable route flap damping",
			},
			"local__ipv4__address": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_group.V_unicast. Local IPv4 address",
			},
			"apply__groups__9": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_group.V_unicast.V_loops. Groups from which to inherit configuration data",
			},
			"apply__groups__except__9": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_group.V_unicast.V_loops. Don't inherit configuration data from these groups",
			},
			"name__19": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_group.V_unicast.V_loops.V_apply__macro__9. Name of the macro to be expanded",
			},
			"name__20": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_group.V_unicast.V_loops.V_apply__macro__9.V_data__9. Keyword part of the keyword-value pair",
			},
			"value__9": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_group.V_unicast.V_loops.V_apply__macro__9.V_data__9. Value part of the keyword-value pair",
			},
			"loops__1": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_group.V_unicast.V_loops. AS-Path loop count",
			},
			"apply__groups__10": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_group.V_unicast.V_delay__route__advertisements. Groups from which to inherit configuration data",
			},
			"apply__groups__except__10": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_group.V_unicast.V_delay__route__advertisements. Don't inherit configuration data from these groups",
			},
			"name__21": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_group.V_unicast.V_delay__route__advertisements.V_apply__macro__10. Name of the macro to be expanded",
			},
			"name__22": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_group.V_unicast.V_delay__route__advertisements.V_apply__macro__10.V_data__10. Keyword part of the keyword-value pair",
			},
			"value__10": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_group.V_unicast.V_delay__route__advertisements.V_apply__macro__10.V_data__10. Value part of the keyword-value pair",
			},
			"always__wait__for__krt__drain": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_group.V_unicast.V_delay__route__advertisements. Wait for KRT-queue drain for more-specific prefixes",
			},
			"apply__groups__11": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_group.V_unicast.V_delay__route__advertisements.V_minimum__delay. Groups from which to inherit configuration data",
			},
			"apply__groups__except__11": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_group.V_unicast.V_delay__route__advertisements.V_minimum__delay. Don't inherit configuration data from these groups",
			},
			"name__23": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_group.V_unicast.V_delay__route__advertisements.V_minimum__delay.V_apply__macro__11. Name of the macro to be expanded",
			},
			"name__24": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_group.V_unicast.V_delay__route__advertisements.V_minimum__delay.V_apply__macro__11.V_data__11. Keyword part of the keyword-value pair",
			},
			"value__11": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_group.V_unicast.V_delay__route__advertisements.V_minimum__delay.V_apply__macro__11.V_data__11. Value part of the keyword-value pair",
			},
			"routing__uptime": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_group.V_unicast.V_delay__route__advertisements.V_minimum__delay. Min delay(sec) advertisement after RPD start",
			},
			"inbound__convergence": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_group.V_unicast.V_delay__route__advertisements.V_minimum__delay. Min delay(sec) advertisement after source-peer sent all routes",
			},
			"apply__groups__12": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_group.V_unicast.V_delay__route__advertisements.V_maximum__delay. Groups from which to inherit configuration data",
			},
			"apply__groups__except__12": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_group.V_unicast.V_delay__route__advertisements.V_maximum__delay. Don't inherit configuration data from these groups",
			},
			"name__25": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_group.V_unicast.V_delay__route__advertisements.V_maximum__delay.V_apply__macro__12. Name of the macro to be expanded",
			},
			"name__26": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_group.V_unicast.V_delay__route__advertisements.V_maximum__delay.V_apply__macro__12.V_data__12. Keyword part of the keyword-value pair",
			},
			"value__12": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_group.V_unicast.V_delay__route__advertisements.V_maximum__delay.V_apply__macro__12.V_data__12. Value part of the keyword-value pair",
			},
			"route__age": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_group.V_unicast.V_delay__route__advertisements.V_maximum__delay. Max delay(sec) advertisement route age",
			},
			"routing__uptime__1": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_group.V_unicast.V_delay__route__advertisements.V_maximum__delay. Max delay(sec) advertisement after RPD start",
			},
			"apply__groups__13": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_group.V_unicast.V_nexthop__resolution. Groups from which to inherit configuration data",
			},
			"apply__groups__except__13": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_group.V_unicast.V_nexthop__resolution. Don't inherit configuration data from these groups",
			},
			"name__27": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_group.V_unicast.V_nexthop__resolution.V_apply__macro__13. Name of the macro to be expanded",
			},
			"name__28": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_group.V_unicast.V_nexthop__resolution.V_apply__macro__13.V_data__13. Keyword part of the keyword-value pair",
			},
			"value__13": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_group.V_unicast.V_nexthop__resolution.V_apply__macro__13.V_data__13. Value part of the keyword-value pair",
			},
			"no__resolution": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_group.V_unicast.V_nexthop__resolution. Consider nexthop good without resolution attempt",
			},
			"preserve__nexthop__hierarchy": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_group.V_unicast.V_nexthop__resolution. Attempt preserving resolved nexthop chain in forwarding",
			},
			"apply__groups__14": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_group.V_unicast.V_defer__initial__multipath__build. Groups from which to inherit configuration data",
			},
			"apply__groups__except__14": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_group.V_unicast.V_defer__initial__multipath__build. Don't inherit configuration data from these groups",
			},
			"name__29": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_group.V_unicast.V_defer__initial__multipath__build.V_apply__macro__14. Name of the macro to be expanded",
			},
			"name__30": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_group.V_unicast.V_defer__initial__multipath__build.V_apply__macro__14.V_data__14. Keyword part of the keyword-value pair",
			},
			"value__14": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_group.V_unicast.V_defer__initial__multipath__build.V_apply__macro__14.V_data__14. Value part of the keyword-value pair",
			},
			"maximum__delay__1": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_group.V_unicast.V_defer__initial__multipath__build. Max delay(sec) multipath build after peer is up",
			},
			"apply__groups__15": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_group.V_unicast.V_graceful__restart. Groups from which to inherit configuration data",
			},
			"apply__groups__except__15": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_group.V_unicast.V_graceful__restart. Don't inherit configuration data from these groups",
			},
			"name__31": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_group.V_unicast.V_graceful__restart.V_apply__macro__15. Name of the macro to be expanded",
			},
			"name__32": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_group.V_unicast.V_graceful__restart.V_apply__macro__15.V_data__15. Keyword part of the keyword-value pair",
			},
			"value__15": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_group.V_unicast.V_graceful__restart.V_apply__macro__15.V_data__15. Value part of the keyword-value pair",
			},
			"apply__groups__16": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_group.V_unicast.V_graceful__restart.V_long__lived. Groups from which to inherit configuration data",
			},
			"apply__groups__except__16": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_group.V_unicast.V_graceful__restart.V_long__lived. Don't inherit configuration data from these groups",
			},
			"name__33": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_group.V_unicast.V_graceful__restart.V_long__lived.V_apply__macro__16. Name of the macro to be expanded",
			},
			"name__34": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_group.V_unicast.V_graceful__restart.V_long__lived.V_apply__macro__16.V_data__16. Keyword part of the keyword-value pair",
			},
			"value__16": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_group.V_unicast.V_graceful__restart.V_long__lived.V_apply__macro__16.V_data__16. Value part of the keyword-value pair",
			},
			"apply__groups__17": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_group.V_unicast.V_graceful__restart.V_long__lived.V_restarter. Groups from which to inherit configuration data",
			},
			"apply__groups__except__17": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_group.V_unicast.V_graceful__restart.V_long__lived.V_restarter. Don't inherit configuration data from these groups",
			},
			"name__35": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_group.V_unicast.V_graceful__restart.V_long__lived.V_restarter.V_apply__macro__17. Name of the macro to be expanded",
			},
			"name__36": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_group.V_unicast.V_graceful__restart.V_long__lived.V_restarter.V_apply__macro__17.V_data__17. Keyword part of the keyword-value pair",
			},
			"value__17": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_group.V_unicast.V_graceful__restart.V_long__lived.V_restarter.V_apply__macro__17.V_data__17. Value part of the keyword-value pair",
			},
			"stale__time": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_group.V_unicast.V_graceful__restart.V_long__lived.V_restarter. Stale time in seconds or dhms notation (1..16777215)",
			},
			"forwarding__state__bit": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_group.V_unicast.V_graceful__restart. Control forwarding-state flag negotiation",
			},
			"extended__nexthop": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_group.V_unicast. Extended nexthop encoding",
			},
			"extended__nexthop__color": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_group.V_unicast. Resolve using extended color nexthop",
			},
			"extended__nexthop__tunnel": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_group.V_unicast. Use BGP tunnel attribute",
			},
			"no__install": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_group.V_unicast. Dont install received routes in forwarding",
			},
			"route__age__bgp__view": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_group.V_unicast. Maintain BGP route's age based on Update messages only",
			},
			"advertise__srv6__service": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_group.V_unicast. Advertise SRv6 service",
			},
			"accept__srv6__service": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_group.V_unicast. Accept SRv6 service",
			},
			"name__37": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_group.V_unicast.V_topology. Topology name",
			},
			"apply__groups__18": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_group.V_unicast.V_topology. Groups from which to inherit configuration data",
			},
			"apply__groups__except__18": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_group.V_unicast.V_topology. Don't inherit configuration data from these groups",
			},
			"name__38": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_group.V_unicast.V_topology.V_apply__macro__18. Name of the macro to be expanded",
			},
			"name__39": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_group.V_unicast.V_topology.V_apply__macro__18.V_data__18. Keyword part of the keyword-value pair",
			},
			"value__18": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_group.V_unicast.V_topology.V_apply__macro__18.V_data__18. Value part of the keyword-value pair",
			},
			"community": &schema.Schema{
				Type:    schema.TypeString,
				Optional: true,
				Description:    "xpath is: config.Groups.V_group.V_unicast.V_topology. Community to identify multi topology routes",
			},
		},
	}
}