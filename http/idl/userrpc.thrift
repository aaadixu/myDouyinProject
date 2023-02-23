namespace go user

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


struct AddWorkNumReq {
    1: i64 user_id;
}

struct AddWorkNumResp {
        1: i32 status_code;
        2: string status_msg;
}










struct TotalFavoritedReq {
    1: i64 user_id;
    2: i8 count;
}

struct TotalFavoritedResp {
        1: i32 status_code;
        2: string status_msg;
}

struct FavoriteCountReq {
    1: i64 user_id;
    2: i8 count;
}

struct FavoriteCountResp {
        1: i32 status_code;
        2: string status_msg;
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
    RegistResp RegistMethod(1: RegistReq request);
    LoginResp LoginMethod(1: LoginReq request) ;
    InfoResp InfoMethod(1: InfoReq request) ;
    AddWorkNumResp AddWorkNumMethod(1: AddWorkNumReq request);
    FavoriteCountResp FavoriteCountMethod(1: FavoriteCountReq request);
    TotalFavoritedResp TotalFavoritedMethod(1: TotalFavoritedReq request);
    FollowCountResp FollowCountMethod(1: FollowCountReq request);
    FollowerCountResp FollowerCountMethod(1: FollowerCountReq request);

    InfosResp InfosMethod(1: InfosReq request) ;
}



