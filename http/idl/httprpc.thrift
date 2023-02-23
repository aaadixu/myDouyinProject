namespace go httprpc

struct RegistReq {
    1: string username;
    2: string password;
}

struct RegistResp {
    1: i32 status_code;
    2: string status_msg;
    3: i64 user_id;
}


struct LoginReq {
    1: string username;
    2: string password;
}

struct LoginResp {
    1: i32 status_code;
    2: string status_msg;
    3: i64 user_id;
}

struct InfoReq {
    1: i64 user_id;
}

struct InfoResp {
    1: i32 status_code;
    2: string status_msg;
    3: User user;
}


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


service UserService {
    RegistResp RegistMethod(1: RegistReq request);
    LoginResp LoginMethod(1: LoginReq request) ;
    InfoResp InfoMethod(1: InfoReq request) ;
}






struct Video {
      1: i64 id; // 视频唯一标识
      2: User author; // 视频作者信息
      3: string play_url; // 视频播放地址
      4: string cover_url; // 视频封面地址
      5: i64 favorite_count; // 视频的点赞总数
      6: i64 comment_count; // 视频的评论总数
      7: bool is_favorite; // true-已点赞，false-未点赞
      8: string title; // 视频标题
}


struct FeedReq {

    1: i64 latest_time;
    2: i64 user_id;

}

struct FeedResp {

    1 : i32 status_code;
    2: string status_msg;
    3: list<Video> video_list;
    4: i64 next_time;

}



# 用户发布视频请求
struct PublishActionReq {
        1: required i64 user_id; // 用户鉴权token
        2: required binary data; // 视频数据
        3: required string title; // 视频标题
}
# 用户发布视频响应
struct PublishActionResp {
        1: required i32 status_code; // 状态码，0-成功，其他值-失败
        2: optional string status_msg; // 返回状态描述
}
# 视频发布列表请求
struct PublishListReq {
        1: required i64 user_id; // 用户id
}
# 视频发布列表响应
struct PublishListResp {
        1: required i32 status_code; // 状态码，0-成功，其他值-失败
        2: optional string status_msg; // 返回状态描述
        3: list<Video> video_list; // 用户发布的视频列表
}

service VideoService{
    FeedResp FeedMethod(1: FeedReq request);
    PublishActionResp PublishActionMethod(1: PublishActionReq request);
    PublishListResp PublishListMethod(1: PublishListReq request);
}





struct FavoriteActionReq {

    1: i64 user_id;

    2: i64 video_id;
    3: i32 action_type;
}

struct FavoriteActionResp {

    1 : i32 status_code;
    2: string status_msg;

}



# 登录用户的所有点赞视频
struct FavoriteListReq {
        1: required i64 user_id; // 用户鉴权token
}
# 登录用户的所有点赞视频响应
struct FavoriteListResp {
        1: required i32 status_code; // 状态码，0-成功，其他值-失败
        2: required string status_msg;
        3: required list<Video> video_list; // 返回状态描述
}


struct Comment  {
    1: required i64 id; // 视频评论id
    2: required User user; // 评论用户信息
    3: required string content; // 评论内容
    4: required string create_date; // 评论发布日期，格式 mm-dd

}


# 评论操作请求
struct CommentActionReq {
        1: required i64 user_id; // 用户id
        2: required i64 video_id;
        3: required i32 action_type;
        4: optional string comment_text;
        5: optional i64 comment_id;
}
# 评论操作响应
struct CommentActionResp {
        1: required i32 status_code; // 状态码，0-成功，其他值-失败
        2: required string status_msg; // 返回状态描述
        3: required Comment comment; // 评论成功返回评论内容，不需要重新拉取整个列表
}



# 视频评论列表请求
struct CommentListReq {
        1: required i64 video_id; // 用户鉴权token
}
# 视频评论列表响应
struct CommentListResp {
        1: required i32 status_code; // 状态码，0-成功，其他值-失败
        2: required string status_msg;
        3: required list<Comment> comment_list; // 返回状态描述
}

service ActionService{
    FavoriteActionResp FavoriteActionMethod(1: FavoriteActionReq request);
    FavoriteListResp FavoriteListMethod(1: FavoriteListReq request);
    CommentActionResp CommentActionMethod(1: CommentActionReq request);
    CommentListResp CommentListMethod(1: CommentListReq request);
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
    12: optional string message; // 发送的消息
    13: required i64 msg_type; // 消息类型，接收或发送

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





service RelationService {

    RelationActionResp RelationActionMethod(1: RelationActionReq request);
    FollowListResp FollowListMethod(1: FollowListReq request);
    FollowerListResp FollowerListMethod(1: FollowerListReq request);
    FriendListResp FriendListMethod(1: FriendListReq request);
}




struct ChatReq {
    1: i64 user_id;
    2: i64 to_user_id;
    3: i64 pre_msg_time;
}

struct ChatResp {
    1: i32 status_code;
    2: string status_msg;
    3: list<Message> message_list;
}


struct Message {

    1: i64 id; // 消息id
    2: i64 to_user_id; // 该消息接收者的id
    3: i64 from_user_id; // 该消息发送者的id
    4: string content; // 消息内容
    5: string create_time; // 消息创建时间
}


struct ActionReq {
    1: i64 user_id;
    2: i64 to_user_id;
    3: i32 action_type;
    4: string content;
}

struct ActionResp {
    1: i32 status_code;
    2: string status_msg;
}




service ChatService {
    ChatResp ChatMethod(1: ChatReq request);
    ActionResp ActionMethod(1: ActionReq request);
}