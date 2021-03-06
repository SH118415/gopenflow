package oxm

const (
	NXM_OF_IN_PORT = OFPXMC_NXM_0<<OXM_CLASS_SHIFT | iota<<OXM_FIELD_SHIFT
	NXM_OF_ETH_DST
	NXM_OF_ETH_SRC
	NXM_OF_ETH_TYPE
	NXM_OF_VLAN_TCI
	NXM_OF_IP_TOS
	NXM_OF_IP_PROTO
	NXM_OF_IP_SRC
	NXM_OF_IP_DST
	NXM_OF_TCP_SRC
	NXM_OF_TCP_DST
	NXM_OF_UDP_SRC
	NXM_OF_UDP_DST
	NXM_OF_ICMP_TYPE
	NXM_OF_ICMP_CODE
	NXM_OF_ARP_OP
	NXM_OF_ARP_SPA
	NXM_OF_ARP_TPA
)

const (
	NXM_NX_REG0 = OFPXMC_NXM_1<<OXM_CLASS_SHIFT | iota<<OXM_FIELD_SHIFT
	NXM_NX_REG1
	NXM_NX_REG2
	NXM_NX_REG3
	NXM_NX_REG4
	NXM_NX_REG5
	NXM_NX_REG6
	NXM_NX_REG7
	_
	_
	_
	_
	_
	_
	_
	_
	NXM_NX_TUN_ID
	NXM_NX_ARP_SHA
	NXM_NX_ARP_THA
	NXM_NX_IPV6_SRC
	NXM_NX_IPV6_DST
	NXM_NX_ICMPV6_TYPE
	NXM_NX_ICMPV6_CODE
	NXM_NX_ND_TARGET
	NXM_NX_ND_SLL
	NXM_NX_ND_TLL
	NXM_NX_IP_FRAG
	NXM_NX_IPV6_LABEL
	NXM_NX_IP_ECN
	NXM_NX_IP_TTL
	_
	NXM_NX_TUN_IPV4_SRC
	NXM_NX_TUN_IPV4_DST
	NXM_NX_PKT_MARK
	NXM_NX_TCP_FLAGS
	NXM_NX_DP_HASH
	NXM_NX_RECIRC_ID
	NXM_NX_CONJ_ID
	NXM_NX_TUN_GBP_ID
	NXM_NX_TUN_GBP_FLAGS
)
