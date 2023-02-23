namespace go user
struct User {
    1: i64 id;
    2: string name;
    3: i64 follow_count;
    4: i64 follower_count;
    5: bool is_follow;
    6: string avatar; //用户头像
    7: string background_image; //用户个人页顶部大图
    8: string signature; //个人简介
    9: i64 total_favorited; //获赞数量
    10: i64 work_count; //作品数量
    11: i64 favorite_count; //点赞数量
}

struct FollowCountReq{
    1: i64 user_id;
    2: i64 count;
}
struct FollowCountResp{
    1: i32 status_code;
    2: string status_msg;
}

struct FollowerCountReq{
    1: i64 user_id;
    2: i64 count;
}
struct FollowerCountResp{
    1: i32 status_code;
    2: string status_msg;
}


struct InfosReq {
    1: list<i64> user_ids;
}

struct InfosResp {
    1: i32 status_code;
    2: string status_msg;
    3: list<User> users;
}




service UserService {
    FollowCountResp FollowCountMethod(1: FollowCountReq request);
    FollowerCountResp FollowerCountMethod(1: FollowerCountReq request);

     InfosResp InfosMethod(1: InfosReq request) ;
}