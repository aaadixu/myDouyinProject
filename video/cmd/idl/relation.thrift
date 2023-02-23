namespace go relation


struct IsFollowingReq {
    1: i64 user_id;
    2: i64 to_user_id;
}

struct IsFollowingResp {

    1: i32 status_code;
    2: string status_msg;
    3: i8 following_type; // 0表示未知错误，1表示已关注，2表示未关注

}



service RelationService {
    IsFollowingResp IsFollowingMethod(1: IsFollowingReq request);
}