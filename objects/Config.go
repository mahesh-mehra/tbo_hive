package objects

var ConfigObj = &Config{}

type KafkaTopicsStruct struct {
	Chats string `json:"chats"`
}

type RedisStruct struct {
	Host     string `json:"host"`
	DB       int    `json:"db"`
	Password string `json:"password"`
}

type TwillioStruct struct {
	Contact string `json:"contact"`
	Url     string `json:"url"`
	Sid     string `json:"sid"`
	Key     string `json:"key"`
}

type Config struct {
	Scylla          string            `json:"scylla"`
	ScyllaNamespace string            `json:"scylla_namespace"`
	ScyllaUsername  string            `json:"scylla_username"`
	ScyllaPassword  string            `json:"scylla_password"`
	KafkaBrokers    string            `json:"kafka_brokers"`
	KafkaTopics     KafkaTopicsStruct `json:"kafka_topics"`
	Http            string            `json:"http"`
	Redis           RedisStruct       `json:"redis"`
	Twillio         TwillioStruct     `json:"twillio"`
	SecretKey       string            `json:"secret_key"`
	LocalPath       LocalPathStruct   `json:"local_path"`
}

type LocalPathStruct struct {
	Images        string `json:"images"`
	Videos        string `json:"videos"`
	ProfilePhotos string `json:"profile_photos"`
}
