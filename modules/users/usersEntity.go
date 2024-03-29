package users

type User struct {
	UserId   string `db:"user_id" json:"user_id"`
	Email    string `db:"email" json:"email"`
	Username string `db:"username" json:"username"`
	Dob      string `db:"dob" json:"dob"`
	Phone    string `db:"phone" json:"phone"`
	Avatar   string `db:"avatar" json:"avatar"`
	Bio      string `db:"bio" json:"bio"`
}

type UserCredentialCheck struct {
	UserId   string `db:"user_id" json:"user_id"`
	Email    string `db:"email" json:"email"`
	Password string `db:"password" json:"password"`
	Username string `db:"username" json:"username"`
	Dob      string `db:"dob" json:"dob"`
	Phone    string `db:"phone" json:"phone"`
	Bio      string `db:"bio" json:"bio"`
	Avatar   string `db:"avatar" json:"avatar"`
}

type UserPassport struct {
	User  *User      `json:"user"`
	Token *UserToken `json:"token"`
}

type UserToken struct {
	OauthId      string `db:"oauth_id" json:"oauth_id"`
	AccessToken  string `db:"access_token" json:"access_token"`
	RefreshToken string `db:"refresh_token" json:"refresh_token"`
}

type UserClaims struct {
	Id string `json:"id" db:"id"`
}

type Oauth struct {
	OauthId string `db:"oauth_id" json:"oauth_id"`
	UserId  string `db:"user_id" json:"user_id"`
}

type FindMember struct {
	UserId   string `db:"user_id" json:"user_id"`
	Username string `db:"username" json:"username"`
	Avatar   string `db:"avatar" json:"avatar"`
	Email    string `db:"email" json:"email"`
}

type TeamsByUserIdRes struct {
	TeamId     string `db:"team_id" json:"team_id"`
	TeamName   string `db:"team_name" json:"team_name"`
	Role       string `db:"role" json:"role"`
	TeamPoster string `db:"team_poster" json:"team_poster"`
}
