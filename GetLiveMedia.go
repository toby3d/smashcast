package hitGox

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"

	just "github.com/toby3d/hitGox/tools"
	f "github.com/valyala/fasthttp"
)

type (
	LiveMedia struct {
		AuthToken  string       `json:"authToken"`
		Livestream []Livestream `json:"livestream"`
		MediaType  string       `json:"media_type"`
		Request    struct {
			This string `json:"this"`
		} `json:"request"`
	}

	Livestream struct {
		CategoryChannels   string `json:"category_channels"`
		CategoryID         string `json:"category_id"`
		CategoryLogoLarge  string `json:"category_logo_large"`
		CategoryLogoSmall  string `json:"category_logo_small"`
		CategoryMediaCount string `json:"category_media_count"`
		CategoryName       string `json:"category_name"`
		CategoryNameShort  string `json:"category_name_short"`
		CategorySeoKey     string `json:"category_seo_key"`
		CategoryUpdated    string `json:"category_updated"`
		CategoryViewers    string `json:"category_viewers"`
		Channel            struct {
			ChannelLink     string `json:"channel_link"`
			Followers       string `json:"followers"`
			LivestreamCount string `json:"livestream_count"`
			MediaIsLive     string `json:"media_is_live"`
			MediaLiveSince  string `json:"media_live_since"`
			PartnerType     string `json:"partner_type"`
			Recordings      string `json:"recordings"`
			Teams           string `json:"teams"`
			TwitterAccount  string `json:"twitter_account"`
			TwitterEnabled  string `json:"twitter_enabled"`
			UserBetaProfile string `json:"user_beta_profile"`
			UserCover       string `json:"user_cover"`
			UserID          string `json:"user_id"`
			UserLogo        string `json:"user_logo"`
			UserLogoSmall   string `json:"user_logo_small"`
			UserMediaID     string `json:"user_media_id"`
			UserName        string `json:"user_name"`
			UserPartner     string `json:"user_partner"`
			UserStatus      string `json:"user_status"`
			Videos          string `json:"videos"`
		} `json:"channel"`
		Following              bool      `json:"following"`
		MediaBgImage           string    `json:"media_bg_image"`
		MediaCategoryID        string    `json:"media_category_id"`
		MediaChatChannel       string    `json:"media_chat_channel"`
		MediaChatEnabled       string    `json:"media_chat_enabled"`
		MediaCountries         []string  `json:"media_countries"`
		MediaDailyViews        string    `json:"media_daily_views"`
		MediaDateAdded         string    `json:"media_date_added"`
		MediaDateUpdated       string    `json:"media_date_updated"`
		MediaDeleted           string    `json:"media_deleted"`
		MediaDescription       string    `json:"media_description"`
		MediaDescriptionMD     string    `json:"media_description_md"`
		MediaDisplayName       string    `json:"media_display_name"`
		MediaDownloadLink      string    `json:"media_download_link"`
		MediaDuration          string    `json:"media_duration"`
		MediaDurationFormat    string    `json:"media_duration_format"`
		MediaDvr               string    `json:"media_dvr"`
		MediaFeatured          string    `json:"media_featured"`
		MediaFeaturedCountries string    `json:"media_featured_countries"`
		MediaFeaturedForced    string    `json:"media_featured_forced"`
		MediaFeaturedWeight    string    `json:"media_featured_weight"`
		MediaFile              string    `json:"media_file"`
		MediaHidden            string    `json:"media_hidden"`
		MediaHost              string    `json:"media_host"`
		MediaHostedID          string    `json:"media_hosted_id"`
		MediaID                string    `json:"media_id"`
		MediaIngestHost        string    `json:"media_ingest_host"`
		MediaIsLive            string    `json:"media_is_live"`
		MediaIsSpherical       bool      `json:"media_is_spherical"`
		MediaKey               string    `json:"media_key"`
		MediaLiveDelay         string    `json:"media_live_delay"`
		MediaLiveNotified      string    `json:"media_live_notified"`
		MediaLiveSince         string    `json:"media_live_since"`
		MediaMature            string    `json:"media_mature"`
		MediaMonthlyViews      string    `json:"media_monthly_views"`
		MediaName              string    `json:"media_name"`
		MediaNotifyUsers       string    `json:"media_notify_users"`
		MediaOfflineID         string    `json:"media_offline_id"`
		MediaPasswordProtected string    `json:"media_password_protected"`
		MediaPlaying           string    `json:"media_playing"`
		MediaPrivacy           string    `json:"media_privacy"`
		MediaProcessing        string    `json:"media_processing"`
		MediaProfiles          string    `json:"media_profiles"`
		MediaRecSession        string    `json:"media_rec_session"`
		MediaRecording         string    `json:"media_recording"`
		MediaRelay             string    `json:"media_relay"`
		MediaRepairSource      string    `json:"media_repair_source"`
		MediaStartInSec        string    `json:"media_start_in_sec"`
		MediaStatus            string    `json:"media_status"`
		MediaTags              string    `json:"media_tags"`
		MediaThumbnail         string    `json:"media_thumbnail"`
		MediaThumbnailLarge    string    `json:"media_thumbnail_large"`
		MediaTitle             string    `json:"media_title"`
		MediaTotalViews        string    `json:"media_total_views"`
		MediaTranscoding       string    `json:"media_transcoding"`
		MediaTypeID            string    `json:"media_type_id"`
		MediaUploaded          string    `json:"media_uploaded"`
		MediaUserID            string    `json:"media_user_id"`
		MediaUserName          string    `json:"media_user_name"`
		MediaViews             string    `json:"media_views"`
		MediaViewsDaily        string    `json:"media_views_daily"`
		MediaViewsMonthly      string    `json:"media_views_monthly"`
		MediaViewsWeekly       string    `json:"media_views_weekly"`
		MediaWeeklyViews       string    `json:"media_weekly_views"`
		MediaYtUpload          string    `json:"media_yt_upload"`
		MediaYtUploadStatus    string    `json:"media_yt_upload_status"`
		Subscribed             bool      `json:"subscribed"`
		TeamName               string    `json:"team_name"`
		UserBanned             string    `json:"user_banned"`
		MediaHostedMedia       LiveMedia `json:"media_hosted_media"`
	}
)

func GetLiveMedia(channel, authToken string, showHidden, fast bool) (*LiveMedia, error) {
	var args f.Args
	args.Add("authToken", authToken)
	if showHidden {
		args.Add("showHidden", strconv.FormatBool(showHidden))
	}
	if fast {
		args.Add("fast", strconv.FormatBool(fast))
	}

	url := fmt.Sprintf(APIEndpoint, fmt.Sprint("media/live/", channel))
	resp, err := just.GET(url, &args)
	if err != nil {
		return nil, err
	}

	var obj LiveMedia
	json.NewDecoder(bytes.NewReader(resp)).Decode(&obj)

	return &obj, nil
}
