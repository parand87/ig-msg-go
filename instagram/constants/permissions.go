package constants

type permissionConstants struct {
	PagesShowList           string
	InstagramBasic          string
	PublicProfile           string
	ReadPageMailboxes       string
	BusinessManagement      string
	PagesMessaging          string
	InstagramManageMessages string
	PagesReadEngagement     string
	PagesManageMetadata     string
}

// Permissions this variable is used in external packages
var Permissions = permissionConstants{ //nolint
	PagesShowList:           "pages_show_list",
	InstagramBasic:          "instagram_basic",
	PublicProfile:           "public_profile",
	ReadPageMailboxes:       "read_page_mailboxes",
	BusinessManagement:      "business_management",
	PagesMessaging:          "pages_messaging",
	InstagramManageMessages: "instagram_manage_messages",
	PagesReadEngagement:     "pages_read_engagement",
	PagesManageMetadata:     "pages_manage_metadata",
}
