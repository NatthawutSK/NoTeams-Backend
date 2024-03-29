package files

type FileRes struct {
	FileName string `json:"filename"`
	Url      string `json:"url"`
}

type GetFilesTeamRes []*GetFilesTeam

type GetFilesTeam struct {
	FileId    string `json:"file_id" db:"file_id"`
	FileName  string `json:"file_name" db:"file_name"`
	FileUrl   string `json:"file_url" db:"file_url"`
	Username  string `json:"username" db:"username"`
	CreatedAt string `json:"created_at" db:"created_at"`
}

type FileTeamByIdRes struct {
	FileId    string `json:"file_id" db:"file_id"`
	FileName  string `json:"file_name" db:"file_name"`
	FileUrl   string `json:"file_url" db:"file_url"`
	Username  string `json:"username" db:"username"`
	CreatedAt string `json:"created_at" db:"created_at"`
}
