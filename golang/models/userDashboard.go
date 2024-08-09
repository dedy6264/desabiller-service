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
)
