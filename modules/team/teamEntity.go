package team

type CreateTeamRes struct {
	TeamId     string `json:"team_id" db:"team_id"`
	TeamName   string `json:"team_name" db:"team_name"`
	TeamPoster string `json:"team_poster" db:"team_poster"`
}

type GetTeamByIdRes struct {
	TeamId     string `json:"team_id" db:"team_id"`
	TeamName   string `json:"team_name" db:"team_name"`
	TeamPoster string `json:"team_poster" db:"team_poster"`
}

type JoinTeamRes struct {
	TeamId     string `json:"team_id" db:"team_id"`
	TeamName   string `json:"team_name" db:"team_name"`
	TeamPoster string `json:"team_poster" db:"team_poster"`
}

type GetTeamByUserIdRes struct {
	TeamId     string `json:"team_id" db:"team_id"`
	TeamName   string `json:"team_name" db:"team_name"`
	TeamPoster string `json:"team_poster" db:"team_poster"`
}