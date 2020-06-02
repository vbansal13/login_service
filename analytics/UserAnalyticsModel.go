package analytics

//UserAnalyticsModel describes analytics message model
type UserAnalyticsModel struct {
	Username        string `json:"username"`
	LoginSuccessful bool   `json:"login_successful"`
	Message         string `json:"message"`
}
