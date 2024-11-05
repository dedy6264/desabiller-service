package models

type (
	ReqUserList struct {
		Username        string `json:"username"`
		Password        string `json:"password"`
		Email           string `json:"email"`
		RoleSegmentId   int    `json:"roleSegmentId"`
		RoleSegmentName string `json:"roleSegmentName"`
		RoleSegmentCode string `json:"roleSegmentCode"`

		Limit   int    `json:"limit"`
		Draw    int    `json:"draw"`
		AscDesc string `json:"ascDesc"`
		SortBy  string `json:"sortBy"`
	}
	RespUserList struct {
		Username        string `json:"username"`
		Password        string `json:"password"`
		Email           string `json:"email"`
		RoleSegmentId   int    `json:"roleSegmentId"`
		RoleSegmentName string `json:"roleSegmentName"`
		RoleSegmentCode string `json:"roleSegmentCode"`
		Role            string `json:"role"`
		HierarchyId     int    `json:"hierarchyId"`
		HierarchyType   string `json:"hierarchyType"`
	}
	ReqLogin struct {
		MerchantOutletUsername string `json:"merchantOutletUsername"`
		MerchantOutletPassword string `json:"merchantOutletPassword"`
	}
	RespLogin struct {
		ID                     int    `json:"id"`
		MerchantOutletName     string `json:"merchantOutletName"`
		MerchantOutletUsername string `json:"merchantOutletUsername"`
		MerchantOutletPassword string `json:"merchantOutletPassword"`
		MerchantId             int    `json:"merchantId"`
		MerchantName           string `json:"merchantName"`
		GroupId                int    `json:"groupId"`
		GroupName              string `json:"groupName"`
		ClientId               int    `json:"clientId"`
		ClientName             string `json:"clientName"`
		Token                  string `json:"token"`
	}
)
