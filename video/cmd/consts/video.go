package consts

var (
	ETCDAddr         = ":2379"
	VideoServiceHost = ""
	VideoServicePort = ":8200"
	VideoServiceName = "videoservice"

	DSN = "root:123456dx@tcp(192.168.222.166:3306)/video_db?charset=utf8&parseTime=True&loc=Local"

	// 存储服务
	VideoBucketName = "video"

	CoverBucketName = "cover"

	Location = "us-east-1"

	Endpoint        = "192.168.222.166:9000"
	AccessKeyID     = "admin"
	SecretAccessKey = "admin123456"
	UseSSL          = false

	RelationServiceName = "relationservice"
	UserServiceName     = "userservice"
	ActionServiceName   = "actionservice"
)
