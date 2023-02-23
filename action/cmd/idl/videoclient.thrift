namespace go video


# 视频点赞请求
struct FavoriteCountReq {
        1: required i64 video_id; // 用户id
        2: i64 count;
}
# 视频点赞响应
struct FavoriteCountResp {
        1: required i32 status_code; // 状态码，0-成功，其他值-失败
        2: required string status_msg; // 返回状态描述
        3: required i64 author_id;
}




# 视频评论请求
struct CommentCountReq {
        1: required i64 video_id; // 视频id
        2: i64 count;
}
# 视频评论响应
struct CommentCountResp {
        1: required i32 status_code; // 状态码，0-成功，其他值-失败
        2: required string status_msg; // 返回状态描述
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






# 视频列表请求
struct VideoListReq {
        1: required list<i64> video_ids; // 视频id
}
# 视频列表响应
struct VideoListResp {
        1: required i32 status_code; // 状态码，0-成功，其他值-失败
        2: required string status_msg; // 返回状态描述
        3: required list<Video> video_list;
}


service VideoService{
    FavoriteCountResp FavoriteCountMethod(1: FavoriteCountReq request);
    CommentCountResp CommentCountMethod(1: CommentCountReq request);
    VideoListResp VideoListMethod(1: VideoListReq request);
}