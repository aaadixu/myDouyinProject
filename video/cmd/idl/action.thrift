namespace go action


struct IsUserFavoriteVideoReq {
    1: i64 video_id;
    2: i64 user_id;
}
struct IsUserFavoriteVideoResp {
    1: i32 status_code;
    2: string status_msg;
    3: bool is_favorite;
}


service ActionService{
    IsUserFavoriteVideoResp IsUserFavoriteVideoMethod(1: IsUserFavoriteVideoReq request);
}