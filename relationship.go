package anaconda

import (
	"net/url"
)

type RelationshipResponse struct {
	Relationship Relationship `json:"relationship"`
}
type Relationship struct {
	Target Target `json:"target"`
	Source Source `json:"source"`
}
type Target struct {
	Id          int64  `json:"id"`
	Id_str      string `json:"id_str"`
	Screen_name string `json:"screen_name"`
	Following   bool   `json:"following"`
	Followed_by bool   `json:"followed_by"`
}
type Source struct {
	Id                    int64  `json:"id"`
	Id_str                string `json:"id_str"`
	Screen_name           string `json:"screen_name"`
	Following             bool   `json:"following"`
	Followed_by           bool   `json:"followed_by"`
	Can_dm                bool   `json:"can_dm"`
	Blocking              bool   `json:"blocking"`
	Blocked_By            bool   `json:"blocked_by"`
	Muting                bool   `json:"muting"`
	Marked_spam           bool   `json:"marked_spam"`
	All_replies           bool   `json:"all_replies"`
	Want_retweets         bool   `json:"want_retweets"`
	Notifications_enabled bool   `json:"notifications_enabled"`
}

func (a TwitterApi) GetFriendshipsShow(v url.Values) (relationshipResponse RelationshipResponse, err error) {
	response_ch := make(chan response)
	a.queryQueue <- query{a.baseUrl + "/friendships/show.json", v, &relationshipResponse, _GET, response_ch}
	return relationshipResponse, (<-response_ch).err
}
