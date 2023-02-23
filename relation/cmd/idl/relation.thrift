namespace go relation


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

struct FriendUser {
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
    12: string message; // 发送的消息
    13: i64 msg_type; // 消息类型，接收或发送



}


# 用户关注操作请求
struct RelationActionReq {

    1: i64 user_id;
    2: i64 to_user_id;
    3: i32 action_type;
}
# 用户关注操作响应
struct RelationActionResp {

    1: i32 status_code;
    2: string status_msg;
}

# 用户关注列表
struct FollowListReq{
    1: i64 user_id;
}
struct FollowListResp{
    1: i32 status_code;
    2: string status_msg;
    3: list<User> user_list;
}


// 用户粉丝列表
struct FollowerListReq{
    1: i64 user_id;
}
struct FollowerListResp{
    1: i32 status_code;
    2: string status_msg;
    3: list<User> user_list;
}


// 用户好友列表
struct FriendListReq{
    1: i64 user_id;
}
struct FriendListResp{
    1: i32 status_code;
    2: string status_msg;
    3: list<FriendUser> user_list;
}




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
    RelationActionResp RelationActionMethod(1: RelationActionReq request);
    FollowListResp FollowListMethod(1: FollowListReq request);
    FollowerListResp FollowerListMethod(1: FollowerListReq request);
    FriendListResp FriendListMethod(1: FriendListReq request);
    IsFollowingResp IsFollowingMethod(1: IsFollowingReq request);
}


