package config

type DB struct {
	DB_HOST     string
	DB_USER     string
	DB_PASSWORD string
	DB_PORT     string
	DB_NAME     string
}

type FRONT_END_URL struct {
	FRONT_END_URL_PROD string
	FRONT_END_URL_DEV  string
}

type SECRETS struct {
	JWT_SECRET string
}

type GOOGLE struct {
	IMAGE_BUCKET_NAME string
	URL_BUCKET_NAME   string
}

type AVAILABLE_LANGUAGES struct {
	EN_US string
	PT_BR string
	FR_FR string
	ES_ES string
	ZH_CN string
}

type TELEMETRY struct {
	END_POINT    string
	SERVICE_NAME string
}

type SPOTIFYAUTHSERVICE struct {
	CLIENT_ID     string
	CLIENT_SECRET string
	REDIRECT_URI  string
}

var DB_POSTGRES_CONTAINER = DB{
	DB_HOST:     "postgres_reading",
	DB_USER:     "postgres",
	DB_PASSWORD: "postgres",
	DB_PORT:     "5432",
	DB_NAME:     "postgres",
}

var DB_POSTGRES_LOCAL = DB{
	DB_HOST:     "localhost",
	DB_USER:     "postgres",
	DB_PASSWORD: "postgres",
	DB_PORT:     "5432",
	DB_NAME:     "postgres",
}

var DB_NEON = DB{
	DB_HOST:     "",
	DB_USER:     "",
	DB_PASSWORD: "",
	DB_NAME:     "",
}

var SECRETS_VAR = SECRETS{
	JWT_SECRET: "c779ff9ab1405e4d1bda57e0d41b13dd674c30f874da93c8dbf4f5fa470aa128",
}

var FRONT_END_URL_VAR = FRONT_END_URL{
	FRONT_END_URL_DEV:  "http://localhost:4200",
	FRONT_END_URL_PROD: "http://localhost:4200",
}

var GOOGLE_VAR = GOOGLE{
	IMAGE_BUCKET_NAME: "you-choose",
	URL_BUCKET_NAME:   "https://storage.googleapis.com/",
}

var AVAILABLE_LANGUAGES_VAR = AVAILABLE_LANGUAGES{
	EN_US: "en-US",
	PT_BR: "pt-BR",
	FR_FR: "fr-FR",
	ES_ES: "es-ES",
	ZH_CN: "zh-CN",
}

var TELEMETRY_VAR = TELEMETRY{
	END_POINT:    "localhost:4317",
	SERVICE_NAME: "you-choose",
}

var SPOTIFY_AUTH_SERVICE = SPOTIFYAUTHSERVICE{
	CLIENT_ID:     "83dc318f5c3c4ca0b4bfb6c7f7333631",
	CLIENT_SECRET: "ff5a6136396649419f1af87d08954e43",
	REDIRECT_URI:  "http://localhost:8080/callback",
}
