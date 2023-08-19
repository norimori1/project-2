Ads activity

Service is responsible for tracking and showing ads activity.

1) [POST] /ad/:ad_id/activate --- json {user_id, action_timestamp}
2) [POST] /ad/:ad_id/deactivate -- json {user_id,  action_timestamp}
3) [GET]  /ad/:ad_id/history
4) [GET] /user/:user_id/active_ads_count
