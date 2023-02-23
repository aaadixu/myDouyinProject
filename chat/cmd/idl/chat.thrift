namespace go chat

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



