package models

type ResponseErrorCustom struct {
	ResponseCode    string      `json:"responseCode"`
	ResponseMessage interface{} `json:"responseMessage"`
}

type ServerConfig struct {
	AppName             string `env:"APP_NAME"`
	AppPort             string `env:"APP_PORT"`
	ServiceHost         string `env:"SERVICE_HOST"`
	Protocol            string `env:"PROTOCOL_SERVER"`
	JWTSecret           string `env:"SECRET"`
	DBUsername          string `env:"DB_USERNAME"`
	DBPassword          string `env:"DB_PASSWORD"`
	DBHost              string `env:"DB_HOST"`
	DBPort              string `env:"DB_PORT"`
	DBName              string `env:"DB_NAME"`
	DriverName          string `env:"DRIVER_NAME"`
	DBUserSqlServer     string `env:"DB_USER_SQLSERVER"`
	DBPasswordSqlServer string `env:"DB_PASSWORD_SQLSERVER"`
	DBHostSqlServer     string `env:"DB_HOST_SQLSERVER"`
	DBPortSqlServer     string `env:"DB_PORT_SQLSERVER"`
	DBNameSqlServer     string `env:"DB_NAME_SQLSERVER"`
	JSONPathFile        string `env:"JSON_PATHFILE,required"`
}

type JwtModel struct {
	Token string `json:"token"`
}

type ResponseCustom struct {
	ResponseCode    string      `json:"responseCode"`
	ResponseMessage string      `json:"responseMessage"`
	Omzet           int         `json:"omzet,omitempty"`
	Data            interface{} `json:"data"`
}

type ResponseDataPagination struct {
	TotalData     int      `json:"total_data"`
	DataPerPage   int      `json:"data_per_page"`
	Page          int      `json:"page"`
	TotalPage     int      `json:"total_page"`
	NumberCurrent int      `json:"number_current"`
	NumberEnd     int      `json:"number_end"`
	NextUrlPage   string   `json:"next_url_page"`
	PrevUrlPage   string   `json:"prev_url_page"`
	Omzet         int      `json:"omzet"`
	Data          []Report `json:"data_detail"`
}

type ResponseCustomErr struct {
	ResponseCode    string `json:"responseCode"`
	ResponseMessage string `json:"responseMessage"`
}

type ErrMeta struct {
	FieldErr string
}
