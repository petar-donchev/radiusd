// Code generated by "stringer -type=AttributeType"; DO NOT EDIT

package vendor

import "fmt"

const _AttributeType_name = "MSCHAPResponseMikrotikXmitLimitMikrotikGroupMikrotikWirelessForwardMikrotikWirelessSkipDot1xMikrotikWirelessEncAlgoMikrotikWirelessEncKeyMikrotikRateLimitMikrotikRealmMikrotikHostIPMSCHAPChallengeMikrotikAdvertiseURLMikrotikAdvertiseIntervalMikrotikRecvLimitGigawordsMikrotikXmitLimitGigawordsMSMPPESendKeyMikrotikTotalLimitMikrotikTotalLimitGigawordsMikrotikAddressListMikrotikWirelessMPKeyMikrotikWirelessCommentMikrotikDelegatedIPv6PoolMikrotik_DHCP_Option_SetMikrotik_DHCP_Option_Param_STR1MSCHAP2ResponseMikrotik_Wireless_VLANIDMikrotik_Wireless_VLANIDtypeMSPrimaryDNSServerMSSecondaryDNSServer"

var _AttributeType_index = [...]uint16{0, 14, 31, 44, 67, 92, 115, 137, 154, 167, 181, 196, 216, 241, 267, 293, 306, 324, 351, 370, 391, 414, 439, 463, 494, 509, 533, 561, 579, 599}

func (i AttributeType) String() string {
	i -= 1
	if i >= AttributeType(len(_AttributeType_index)-1) {
		return fmt.Sprintf("AttributeType(%d)", i+1)
	}
	return _AttributeType_name[_AttributeType_index[i]:_AttributeType_index[i+1]]
}
