package protocol

import "encoding/json"

type Config struct {
	Chats       map[int64]ConfigChat `json:"chats"`
	Server      ConfigServer         `json:"server"`
	User        ConfigUser           `json:"user"`
	Hash        string               `json:"hash"`
	Experiments ConfigExperiments    `json:"experiments"`
}

type ConfigExperiments struct {
	CallsAndroidWtp            string `json:"calls-android-wtp"`
	KeepBackgroundSocket       string `json:"keep-background-socket"`
	MediaTransform             string `json:"media-transform"`
	CallsCheckAuState          bool   `json:"calls-check-au-state"`
	SkipStatusFirstAudioUpload bool   `json:"skip-status-first-audio-upload"`
	CallsAndroidH265S          int    `json:"calls-android-h265-s"`
	CallsSdkIosFastJoin        int    `json:"calls-sdk-ios-fast-join"`
	AppAbTestExp               bool   `json:"app.ab.test.exp"`
	CallsAndroidEarlySetOffer  bool   `json:"calls-android-early-set-offer"`
	CallsSdkH265Prioritized    bool   `json:"calls-sdk-h265-prioritized"`
	CallsSdkIceSize            int    `json:"calls-sdk-ice-size"`
	CallsSdkWtDns              int    `json:"calls-sdk-wt-dns"`
}

type ConfigUser struct {
	ChatPushNotification      string          `json:"CHATS_PUSH_NOTIFICATION"`
	PushDetails               bool            `json:"PUSH_DETAILS"`
	PushSound                 string          `json:"PUSH_SOUND"`
	PhoneNumberPrivacy        string          `json:"PHONE_NUMBER_PRIVACY"`
	InactiveTTL               string          `json:"INACTIVE_TTL"`
	ShowReadMark              bool            `json:"SHOW_READ_MARK"`
	AudioTranscriptionEnabled bool            `json:"AUDIO_TRANSCRIPTION_ENABLED"`
	SearchByPhone             string          `json:"SEARCH_BY_PHONE"`
	IncomingCall              string          `json:"INCOMING_CALL"`
	DoubleTapReactionDisabled bool            `json:"DOUBLE_TAP_REACTION_DISABLED"`
	SafeModeNoPin             bool            `json:"SAFE_MODE_NO_PIN"`
	ChatPushSound             string          `json:"CHATS_PUSH_SOUND"`
	DoubleTapReactionValue    json.RawMessage `json:"DOUBLE_TAP_REACTION_VALUE"`
	FamilyProtection          string          `json:"FAMILY_PROTECTION"`
	Hidden                    bool            `json:"HIDDEN"`
	ChatInvite                string          `json:"CHATS_INVITE"`
	PushNewContacts           bool            `json:"PUSH_NEW_CONTACTS"`
	UnsafeFiles               bool            `json:"UNSAFE_FILES"`
	DontDisturbUntil          int             `json:"DONT_DISTURB_UNTIL"`
	AltKeyboard               bool            `json:"ALT_KEYBOARD"`
	ContentLevelAccess        bool            `json:"CONTENT_LEVEL_ACCESS"`
	StickersSuggest           string          `json:"STICKERS_SUGGEST"`
	SafeMode                  bool            `json:"SAFE_MODE"`
	MCallPushNotification     string          `json:"M_CALL_PUSH_NOTIFICATION"`
}

type ConfigChat struct {
	Led              bool `json:"led"`
	DontDisturbUntil int  `json:"dontDisturbUntil"`
	Vibr             bool `json:"vibr"`
	Sound            bool `json:"sound"`
	FavIndex         int  `json:"favIndex"`
}

type ConfigServer struct {
	SetUnreadTimeout                  int                   `json:"set-unread-timeout"`
	AccountRemovalEnabled             bool                  `json:"account-removal-enabled"`
	AppearanceMultiThemeScreenEnabled bool                  `json:"appearance-multi-theme-screen-enabled"`
	Gce                               bool                  `json:"gce"`
	GroupCallChatSupport              bool                  `json:"group-call-chat-support"`
	CallRate                          CallRate              `json:"call-rate"`
	ImageSize                         int                   `json:"image-size"`
	Gcce                              bool                  `json:"gcce"`
	LiveStreams                       bool                  `json:"live-streams"`
	MaxVideoDurationDownload          int                   `json:"max-video-duration-download"`
	MaxMsgLength                      int                   `json:"max-msg-length"`
	MarkdownMenu                      int                   `json:"markdown-menu"`
	PhonePrivacyConfig                bool                  `json:"phone-privacy-config"`
	QuotesEnabled                     bool                  `json:"quotes-enabled"`
	AsyncTracer                       int                   `json:"async-tracer"`
	ImageWidth                        int                   `json:"image-width"`
	ReactionsSettingsEnabled          bool                  `json:"reactions-settings-enabled"`
	InviteByLinkScreen                bool                  `json:"invite-by-link-screen"`
	CallsEndpoint                     string                `json:"calls-endpoint"`
	SendLocationEnabled               bool                  `json:"send-location-enabled"`
	GcFromP2P                         bool                  `json:"gc-from-p2p"`
	MaxThemeLength                    int                   `json:"max-theme-length"`
	CallDontUseVpnForRtp              bool                  `json:"callDontUseVpnForRtp"`
	Lgce                              bool                  `json:"lgce"`
	PollsInP2PChats                   bool                  `json:"polls-in-p2p-chats"`
	ChannelsSuggestsFolder            bool                  `json:"channels-suggests-folder"`
	ScheduledMessagesEnabled          bool                  `json:"scheduled-messages-enabled"`
	NonContactComplaintsEnabled       bool                  `json:"non-contact-complaints-enabled"`
	MsgGetReactionsPageSize           int                   `json:"msg-get-reactions-page-size"`
	CallsSdkWtEnabled                 bool                  `json:"calls-sdk-wt-enabled"`
	JoinRequests                      bool                  `json:"join-requests"`
	NotContactPlaceholder             bool                  `json:"not-contact-placeholder"`
	JsDownloadDelegate                bool                  `json:"js-download-delegate"`
	DefaultReactionsSettings          ReactionSettings      `json:"default-reactions-settings"`
	February2326Theme                 bool                  `json:"february-23-26-theme"`
	Wud                               bool                  `json:"wud"`
	VideoMsgEnabled                   bool                  `json:"video-msg-enabled"`
	OrgProfile                        bool                  `json:"org-profile"`
	Grse                              bool                  `json:"grse"`
	PostLinkEnabled                   bool                  `json:"post-link-enabled"`
	EditTimeout                       int                   `json:"edit-timeout"`
	ReactionsMax                      int                   `json:"reactions-max"`
	ViewsCountEnabled                 bool                  `json:"views-count-enabled"`
	TypingEnabledFILE                 bool                  `json:"typing-enabled-FILE"`
	LebedevThemeEnabled               bool                  `json:"lebedev-theme-enabled"`
	MaxParticipants                   int                   `json:"max-participants"`
	AudioTranscriptionLocales         []json.RawMessage     `json:"audio-transcription-locales"` //TODO
	ChannelProfileInviteLink          bool                  `json:"channel-profile-invite-link"`
	WelcomeStickerIds                 []int                 `json:"welcome-sticker-ids"`
	BotsChannelAdding                 bool                  `json:"bots-channel-adding"`
	FilePreview                       bool                  `json:"file-preview"`
	NewAdminPermissions               bool                  `json:"new-admin-permissions"`
	InviteLong                        string                `json:"invite-long"`
	MaxFavoriteStickerSets            int                   `json:"max-favorite-sticker-sets"`
	EnableAudioMessagesTranscription  bool                  `json:"enable-audio-messages-transcription"`
	CallsFakebossIncomingCallEnabled  bool                  `json:"calls-fakeboss-incoming-call-enabled"`
	MoscowThemeEnabled                bool                  `json:"moscow-theme-enabled"`
	CalcVideoWave                     bool                  `json:"calc-video-wave"`
	ChatsFolderEnabled                bool                  `json:"chats-folder-enabled"`
	PollsInP2GChats                   int                   `json:"polls-in-p2g-chats"`
	ChatInviteLinkPermissionsEnabled  bool                  `json:"chat-invite-link-permissions-enabled"`
	Cfs                               bool                  `json:"cfs"`
	CallEnableIceRenomination         bool                  `json:"callEnableIceRenomination"`
	MentionsEnabled                   bool                  `json:"mentions-enabled"`
	DoubleTapReaction                 string                `json:"double-tap-reaction"`
	PollTTL                           TTL                   `json:"poll-ttl"`
	ImageQuality                      float64               `json:"image-quality"`
	RenameSettingsToProfile           bool                  `json:"rename-settings-to-profile"`
	MaxAudioLength                    int                   `json:"max-audio-length"`
	SettingsEntryBanners              []SettingEntryBanner  `json:"settings-entry-banners"`
	CallsHotkeys                      bool                  `json:"calls-hotkeys"`
	InviteShort                       string                `json:"invite-short"`
	NotifTypingPresence               bool                  `json:"notif-typing-presence"`
	OfficialOrg                       bool                  `json:"official-org"`
	WebProfileDeletion                bool                  `json:"web-profile-deletion"`
	ChannelsEnabled                   bool                  `json:"channels-enabled"`
	EnableUnknownContactBottomSheet   int                   `json:"enable-unknown-contact-bottom-sheet"`
	UnsafeFilesAlert                  bool                  `json:"unsafe-files-alert"`
	MinImageSideSize                  int                   `json:"min-image-side-size"`
	AccountNicknameEnabled            bool                  `json:"account-nickname-enabled"`
	NickMinLength                     int                   `json:"nick-min-length"`
	InformerEnabled                   bool                  `json:"informer-enabled"`
	Stub                              string                `json:"stub"`
	FileUploadEnabled                 bool                  `json:"file-upload-enabled"`
	InviteLink                        string                `json:"invite-link"`
	DraftsSyncEnabled                 bool                  `json:"drafts-sync-enabled"`
	SettingsBusiness                  string                `json:"settings-business"`
	MentionsEntityNamesLimit          int                   `json:"mentions_entity_names_limit"`
	ChannelsComplaintEnabled          bool                  `json:"channels-complaint-enabled"`
	ReactionsEnabled                  bool                  `json:"reactions-enabled"`
	March826Theme                     bool                  `json:"march-8-26-theme"`
	WebmStickersEnabled               bool                  `json:"webm-stickers-enabled"`
	DeleteMsgFysLargeChatDisabled     bool                  `json:"delete-msg-fys-large-chat-disabled"`
	YMap                              YMap                  `json:"y-map"`
	RenderPolls                       bool                  `json:"render-polls"`
	FileUploadUnsupportedTypes        []string              `json:"file-upload-unsupported-types"`
	FaCreationDisabled                bool                  `json:"2fa-creation-disabled"`
	NickMaxLength                     int                   `json:"nick-max-length"`
	ShowWarningLinks                  bool                  `json:"show-warning-links"`
	ContentLevelAccess                bool                  `json:"content-level-access"`
	MaxFavoriteChats                  int                   `json:"max-favorite-chats"`
	AuthorVisibilityForwardEnabled    bool                  `json:"author-visibility-forward-enabled"`
	HasPhone                          bool                  `json:"has-phone"`
	SavedMessagesEnabled              bool                  `json:"saved-messages-enabled"`
	ReactPermission                   int                   `json:"react-permission"`
	RenameProfileToSettings           bool                  `json:"rename-profile-to-settings"`
	MaxReadmarks                      int                   `json:"max-readmarks"`
	StickerSuggestion                 []string              `json:"sticker-suggestion"`
	Creation2FaConfig                 TwoFactorConfig       `json:"creation-2fa-config"`
	ChatsPreloadPeriod                int                   `json:"chats-preload-period"`
	AudioPlayCmd                      bool                  `json:"audio-play-cmd"`
	DeleteMessageFromReply            bool                  `json:"delete-message-from-reply"`
	EnableFiltersForFolders           bool                  `json:"enable-filters-for-folders"`
	GroupCallPartLimit                int                   `json:"group-call-part-limit"`
	MaxFavoriteStickers               int                   `json:"max-favorite-stickers"`
	CallsVideoZoom                    bool                  `json:"calls-video-zoom"`
	FileUploadMaxSize                 int64                 `json:"file-upload-max-size"`
	MaxDescriptionLength              int                   `json:"max-description-length"`
	ChatsPageSize                     int                   `json:"chats-page-size"`
	MarkdownEnabled                   bool                  `json:"markdown-enabled"`
	WhiteListLinks                    []string              `json:"white-list-links"`
	SessionsEndDisabled               bool                  `json:"sessions-end-disabled"`
	EnableVideoMessagesTranscription  bool                  `json:"enable-video-messages-transcription"`
	ReactionsMenu                     []string              `json:"reactions-menu"`
	UniqueFavorites                   bool                  `json:"unique-favorites"`
	AudioPlayOpus                     bool                  `json:"audio-play-opus"`
	FamilyProtectionBotid             int                   `json:"family-protection-botid"`
	CallsSdkMapping                   CallsSDKMapping       `json:"calls-sdk-mapping"`
	StatSessionBackgroundThreshold    int                   `json:"stat-session-background-threshold"`
	CallsFullscreenMode               bool                  `json:"calls-fullscreen-mode"`
	ImageHeight                       int                   `json:"image-height"`
	ScheduledPostsEnabled             bool                  `json:"scheduled-posts-enabled"`
	SearchWebAppsShowcase             SearchWebappsShowcase `json:"search-webapps-showcase"`
	SavedMessagesAliases              []string              `json:"saved-messages-aliases"`
	SafeModeEnabled                   bool                  `json:"safe-mode-enabled"`
}

type CallRate struct {
	Limit    int `json:"limit"`
	Duration int `json:"duration"`
	SdkLimit int `json:"sdk-limit"`
	Delay    int `json:"delay"`
}

type ReactionSettings struct {
	Count       int   `json:"count"`
	IsActive    bool  `json:"isActive"`
	Included    bool  `json:"included"`
	ReactionIDs []any `json:"reactionIds"`
}

type TTL struct {
	Channel int `json:"channel"`
	Bigchat int `json:"bigchat"`
	Chat    int `json:"chat"`
}

type SettingEntryBanner struct {
	Icon  string `json:"icon"`
	Title string `json:"title"`
	Appid int    `json:"appid"`
}

type YMap struct {
	Tile     string `json:"tile"`
	Geocoder string `json:"geocoder"`
	Static   string `json:"static"`
}

type TwoFactorConfig struct {
	PassMinLen int  `json:"pass_min_len"`
	PassMaxLen int  `json:"pass_max_len"`
	HintMaxLen int  `json:"hint_max_len"`
	Enabled    bool `json:"enabled"`
}

type CallsSDKMapping struct {
	Off bool `json:"off"`
}

type SearchWebappsShowcase struct {
	Items []SearchWebAppsItem `json:"items"`
}

type SearchWebAppsItem struct {
	Icon  string `json:"icon"`
	Name  string `json:"name"`
	ID    int    `json:"id"`
	Title string `json:"title"`
}
