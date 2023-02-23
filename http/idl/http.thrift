namespace go http

struct RegistReq {
    1: string Name (api.body="username",api.query="username"); // 添加 api 注解为方便进行参数绑定
    2: string Password (api.body="password",api.query="password",api.form="password"); // 添加 api 注解为方便进行参数绑定
}

struct RegistResp {
    1: i32 status_code;
    2: string status_msg;
    3: i64 user_id;
    4: string token;
}


struct LoginReq {
    1: string Name (api.body="username",api.query="username",api.form="username"); // 添加 api 注解为方便进行参数绑定
    2: string Password (api.body="password",api.query="password",api.form="password"); // 添加 api 注解为方便进行参数绑定
}

struct LoginResp {
    1: i32 status_code;
    2: string status_msg;
    3: i64 user_id;
    4: string token;
}

struct InfoReq {
    1: i64 user_id;
    2: string token;
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
    RegistResp RegistMethod(1: RegistReq request) (api.post="/douyin/user/register/");
    LoginResp LoginMethod(1: LoginReq request) (api.post="/douyin/user/login/");
    InfoResp InfoMethod(1: InfoReq request) (api.get="/douyin/user/");
}



struct FeedReq {

    1: i64 latest_time;
    2: string token;

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

struct FeedResp {

    1 : i32 status_code;
    2: string status_msg;
    3: list<Video> video_list;
    4: i64 next_time;

}



# 用户发布视频请求
struct PublishActionReq {
        1: required string token; // 用户鉴权token
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
        2: required string token; // 用户鉴权token
}
# 视频发布列表响应
struct PublishListResp {
        1: required i32 status_code; // 状态码，0-成功，其他值-失败
        2: optional string status_msg; // 返回状态描述
        3: list<Video> video_list; // 用户发布的视频列表
}

service VideoService{
    FeedResp FeedMethod(1: FeedReq request) (api.get="/douyin/feed/");
    PublishActionResp PublishActionMethod(1: PublishActionReq request) (api.post="/douyin/publish/action/");
    PublishListResp PublishListMethod(1: PublishListReq request) (api.get="/douyin/publish/list/");
}



# 用户的点赞请求
struct FavoriteActionReq {
        1: required string token; // 用户鉴权token
        2: required i64 video_id; // 视频id
        3: required i32 action_type; // 1-点赞，2-取消点赞
}
# 用户的点赞响应
struct FavoriteActionResp {
        1: required i32 status_code; // 状态码，0-成功，其他值-失败
        2: optional string status_msg; // 返回状态描述
}
# 用户的所有点赞视频请求
struct FavoriteListReq {
        1: required i64 user_id; // 用户id
        2: required string token; // 用户鉴权token
}
# 用户的所有点赞视频响应
struct FavoriteListResp {
        1: required i32 status_code; // 状态码，0-成功，其他值-失败
        2: optional string status_msg; // 返回状态描述
        3: list<Video> video_list; // 用户点赞视频列表
}

# 用户评论请求
struct CommentActionReq {
      1: required string token; // 用户鉴权token
      2: required i64 video_id; // 视频id
      3: required i32 action_type; // 1-发布评论，2-删除评论
      4: optional string comment_text; // 用户填写的评论内容，在action_type=1的时候使用
      5: optional i64 comment_id; // 要删除的评论id，在action_type=2的时候使用
}
# 用户评论响应
struct CommentActionResp {
      1: required i32 status_code; // 状态码，0-成功，其他值-失败
      2: optional string status_msg; // 返回状态描述
      3: optional Comment comment; // 评论成功返回评论内容，不需要重新拉取整个列表
}
# 用户评论
struct Comment {
      1: required i64 id; // 视频评论id
      2: required User user; // 评论用户信息
      3: required string content; // 评论内容
      4: required string create_date; // 评论发布日期，格式 mm-dd
}
# 用户评论列表请求
struct CommentListReq {
      1: required string token; // 用户鉴权token
      2: required i64 video_id; // 视频id
}
# 用户评论列表响应
struct CommentListResp {
      1: required i32 status_code; // 状态码，0-成功，其他值-失败
      2: optional string status_msg; // 返回状态描述
      3: list<Comment> comment_list; // 评论列表
}

# 互动服务
service ActionService {
    CommentActionResp CommentActionMethod(1: CommentActionReq request) (api.post="/douyin/comment/action/");
    CommentListResp CommentListMethod(1: CommentListReq request) (api.get="/douyin/comment/list/");
    FavoriteActionResp FavoriteActionMethod(1: FavoriteActionReq request) (api.post="/douyin/favorite/action/");
    FavoriteListResp FavoriteListMethod(1: FavoriteListReq request) (api.get="/douyin/favorite/list/");
}



# 社交接口
# 登录用户对其他用户进行关注或取消关注请求
struct RelationActionReq {
      1: required string token; // 用户鉴权token
      2: required i64 to_user_id; // 对方用户id
      3: required i32 action_type; // 1-关注，2-取消关注
}
# 登录用户对其他用户进行关注或取消关注响应
struct RelationActionResp {
      1: required i32 status_code; // 状态码，0-成功，其他值-失败
      2: optional string status_msg // 返回状态描述
}
# 登录用户关注的所有用户列表请求
struct RelationFollowListReq {
      1: required i64 user_id; // 用户id
      2: required string token; // 用户鉴权token
}
# 登录用户关注的所有用户列表响应
struct RelationFollowListResp {
      1: required i32 status_code; // 状态码，0-成功，其他值-失败
      2: optional string status_msg; // 返回状态描述
      3: list<User> user_list; // 用户信息列表
}
# 关注登录用户的粉丝列表请求
struct RelationFollowerListReq {
      1: required i64 user_id; // 用户id
      2: required string token; // 用户鉴权token
}
# 关注登录用户的粉丝列表响应
struct RelationFollowerListResp {
      1: required i32 status_code; // 状态码，0-成功，其他值-失败
      2: optional string status_msg; // 返回状态描述
      3: list<User> user_list; // 用户列表
}
# 关注登录用户的好友列表请求
struct RelationFriendListReq {
      1: required i64 user_id; // 用户id
      2: required string token; // 用户鉴权token
}
# 关注登录用户的好友列表响应
struct RelationFriendListResp {
      1: required i32 status_code; // 状态码，0-成功，其他值-失败
      2: optional string status_msg; // 返回状态描述
      3: list<FriendUser> user_list; // 用户列表
}
struct FriendUser{
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
    12: optional string message; // 和该好友的最新聊天消息
    13: required i64 msgType; // message消息的类型，0 => 当前请求用户接收的消息， 1 => 当前请求用户发送的消息
}
# -------------------------------
# 关系服务
service RelationService {
    RelationActionResp RelationActionMethod(1: RelationActionReq request) (api.post="/douyin/relation/action/");
    RelationFollowListResp RelationFollowListMethod(1: RelationFollowListReq request) (api.get="/douyin/relation/follow/list/");
    RelationFollowerListResp RelationFollowerListMethod(1: RelationFollowerListReq request) (api.get="/douyin/relation/follower/list/");
    RelationFriendListResp RelationFriendListMethod(1: RelationFriendListReq request) (api.get="/douyin/relation/friend/list/");
}



# 消息接收请求
struct MessageChatReq {
      1: required string token; // 用户鉴权token
      2: required i64 to_user_id; // 对方用户id
      3: required i64 pre_msg_time; //上次最新消息的时间（新增字段-apk更新中）
}
# 消息接收响应
struct MessageChatResp {
      1: required i32 status_code; // 状态码，0-成功，其他值-失败
      2: optional string status_msg; // 返回状态描述
      3: list<Message> message_list; // 消息列表
}
# 消息发送请求
struct MessageSendReq {
      1: required string token; // 用户鉴权token
      2: required i64 to_user_id; // 对方用户id
      3: required i32 action_type; // 1-发送消息
      4: required string content; // 消息内容
}
# 消息发送响应
struct MessageSendResp {
      1: required i32 status_code; // 状态码，0-成功，其他值-失败
      2: optional string status_msg;  // 返回状态描述
}
# 消息
struct Message {
      1: required i64 id; // 消息id
      2: required i64 to_user_id; // 该消息接收者的id
      3: required i64 from_user_id; // 该消息发送者的id
      4: required string content; // 消息内容
      5: optional string create_time; // 消息创建时间
}
# -------------------------------
# 消息服务
service ChatService {
    MessageChatResp MessageChatMethod(1: MessageChatReq request) (api.get="/douyin/message/chat/");
    MessageSendResp MessageSendMethod(1: MessageSendReq request) (api.post="/douyin/message/action/");
}