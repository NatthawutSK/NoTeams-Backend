package config

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

func LoadConfig(path string) IConfig {
	envMap, err := godotenv.Read(path)
	if err != nil {
		log.Fatalf("load dotenv failed: %v", err)
	}

	return &config{
		app: &app{
			host: envMap["APP_HOST"],
			port: func() int {
				p, err := strconv.Atoi(envMap["APP_PORT"])
				if err != nil {
					log.Fatalf("load port failed: %v", err)
				}
				return p
			}(),
			name:    envMap["APP_NAME"],
			version: envMap["APP_VERSION"],
			readTimeout: func() time.Duration {
				t, err := strconv.Atoi(envMap["APP_READ_TIMEOUT"])
				if err != nil {
					log.Fatalf("load read timeout failed: %v", err)
				}
				return time.Duration(int64(t) * int64(math.Pow10(9)))
			}(),
			writeTimeout: func() time.Duration {
				t, err := strconv.Atoi(envMap["APP_WRITE_TIMEOUT"])
				if err != nil {
					log.Fatalf("load write timeout failed: %v", err)
				}
				return time.Duration(int64(t) * int64(math.Pow10(9)))
			}(),
			bodyLimit: func() int {
				b, err := strconv.Atoi(envMap["APP_BODY_LIMIT"])
				if err != nil {
					log.Fatalf("load body limit failed: %v", err)
				}
				return b
			}(),
			fileUploadLimit: func() int {
				b, err := strconv.Atoi(envMap["APP_FILE_UPLOAD_LIMIT"])
				if err != nil {
					log.Fatalf("load file limit failed: %v", err)
				}
				return b
			}(),
		},
		db: &db{
			host: envMap["DB_HOST"],
			port: func() int {
				p, err := strconv.Atoi(envMap["DB_PORT"])
				if err != nil {
					log.Fatalf("load db port failed: %v", err)
				}
				return p
			}(),
			protocol: envMap["DB_PROTOCOL"],
			username: envMap["DB_USERNAME"],
			password: envMap["DB_PASSWORD"],
			database: envMap["DB_DATABASE"],
			sslMode:  envMap["DB_SSL_MODE"],
			maxConnections: func() int {
				m, err := strconv.Atoi(envMap["DB_MAX_CONNECTIONS"])
				if err != nil {
					log.Fatalf("load db max connections failed: %v", err)
				}
				return m
			}(),
		},
		jwt: &jwt{
			secertKey: envMap["JWT_SECRET_KEY"],
			accessExpiresAt: func() int {
				t, err := strconv.Atoi(envMap["JWT_ACCESS_EXPIRES"])
				if err != nil {
					log.Fatalf("load access expires at failed: %v", err)
				}
				return t
			}(),
			refreshExpiresAt: func() int {
				t, err := strconv.Atoi(envMap["JWT_REFRESH_EXPIRES"])
				if err != nil {
					log.Fatalf("load refresh expires at failed: %v", err)
				}
				return t
			}(),
		},
		s3: &s3{
			s3AccessKey: envMap["S3_ACCESS_KEY"],
			s3SecretKey: envMap["S3_SECRET_KEY"],
			s3Bucket:    envMap["S3_BUCKET"],
			s3Region:    envMap["S3_REGION"],
			s3Session:   envMap["S3_SESSION_TOKEN"],
		},
	}
}

type IConfig interface {
	App() IAppConfig
	Db() IDbConfig
	Jwt() IJwtConfig
	S3() IS3Config
}

type config struct {
	app *app
	db  *db
	jwt *jwt
	s3  *s3
}

type IAppConfig interface {
	Url() string // host:port
	Name() string
	Version() string
	ReadTimeout() time.Duration
	WriteTimeout() time.Duration
	BodyLimit() int
	FileLimit() int
	Host() string
	Port() int
}

type app struct {
	host            string
	port            int
	name            string
	version         string
	readTimeout     time.Duration
	writeTimeout    time.Duration
	bodyLimit       int //bytes
	fileUploadLimit int //bytes
}

func (c *config) App() IAppConfig {
	return c.app
}

func (a *app) FileLimit() int {
	return a.fileUploadLimit
}

func (a *app) Url() string {
	return fmt.Sprintf("%s:%d", a.host, a.port)
}
func (a *app) Name() string                { return a.name }
func (a *app) Version() string             { return a.version }
func (a *app) ReadTimeout() time.Duration  { return a.readTimeout }
func (a *app) WriteTimeout() time.Duration { return a.writeTimeout }
func (a *app) BodyLimit() int              { return a.bodyLimit }
func (a *app) FileUploadLimit() int        { return a.fileUploadLimit }
func (a *app) Host() string                { return a.host }
func (a *app) Port() int                   { return a.port }

type IDbConfig interface {
	Url() string
	MaxOpenConns() int
}

type db struct {
	host           string
	port           int
	protocol       string
	username       string
	password       string
	database       string
	sslMode        string
	maxConnections int
}

func (c *config) Db() IDbConfig {
	return c.db
}
func (d *db) Url() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		d.host,
		d.port,
		d.username,
		d.password,
		d.database,
		d.sslMode,
	)
}
func (d *db) MaxOpenConns() int { return d.maxConnections }

type IJwtConfig interface {
	SecretKey() []byte
	AccessExpiresAt() int
	RefreshExpiresAt() int
	SetJwtAccessExpires(t int)
	SetJwtRefreshExpires(t int)
}

type jwt struct {
	secertKey        string
	accessExpiresAt  int //sec
	refreshExpiresAt int //sec
}

func (c *config) Jwt() IJwtConfig {
	return c.jwt
}
func (j *jwt) SecretKey() []byte          { return []byte(j.secertKey) }
func (j *jwt) AccessExpiresAt() int       { return j.accessExpiresAt }
func (j *jwt) RefreshExpiresAt() int      { return j.refreshExpiresAt }
func (j *jwt) SetJwtAccessExpires(t int)  { j.accessExpiresAt = t }
func (j *jwt) SetJwtRefreshExpires(t int) { j.refreshExpiresAt = t }

type s3 struct {
	s3AccessKey string
	s3SecretKey string
	s3Bucket    string
	s3Region    string
	s3Session   string
}

type IS3Config interface {
	S3AccessKey() string
	S3SecretKey() string
	S3Bucket() string
	S3Region() string
	S3Session() string
}

func (s *s3) S3AccessKey() string { return s.s3AccessKey }
func (s *s3) S3SecretKey() string { return s.s3SecretKey }
func (s *s3) S3Bucket() string    { return s.s3Bucket }
func (s *s3) S3Region() string    { return s.s3Region }
func (s *s3) S3Session() string   { return s.s3Session }

func (c *config) S3() IS3Config {
	return c.s3
}
