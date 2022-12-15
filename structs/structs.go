package structs

type Topik struct {
	Id         int64  `json:"id"`
	Judul      string `json:"judul"`
	Pertanyaan string `json:"pertanyaan"`
	Periode    string `json:"-"`
	StartDate  string `json:"start_date"`
	EndDate    string `json:"end_date"`
	CreatedBy  string `json:"created_by"`
	CreatedAt  string `json:"-"`
	UpdateAt   string `json:"-"`
	Deleted    string `json:"-"`
	IsRelease  string `json:"is_realese"`
	Target     string `json:"target"`
}

type Konten struct {
	Id         int64  `json:"-"`
	IdTopik    string `json:"id_topik"`
	Judul      string `json:"judul"`
	Keterangan string `json:"keterangan"`
	ImageSrc   string `json:"image_src"`
	Deleted    int    `json:"-"`
	Voting     string
}

type Respon struct {
	Id        int64  `json:"-"`
	IdTopik   string `json:"id_topik"`
	Response  string `json:"id_konten"`
	CreatedBy string `json:"-"`
	CreatedAt string `json:"-"`
	UpdateAt  string `json:"-"`
}

type Target struct {
	Id          int64  `json:"-"`
	Target      string `json:"target"`
	Description string `json:"description"`
	CreatedAt   string `json:"-"`
	UpdateAt    string `json:"-"`
}

type Listener struct {
	Id       int64  `json:"-"`
	Username string `json:"username"`
	Target   string `json:"target"`
}

type User struct {
	Id        int64  `json:"-"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	FullName  string `json:"full_name"`
	Gender    string `json:"gender"`
	Level     string `json:"level"`
	CreatedAt string `json:"-"`
	UpdateAt  string `json:"-"`
}
type Polling struct {
	Topik  Topik
	Konten Konten
}

type Account struct {
	Username string
	Password string
	Email    string
	Level    string
}

// =======================

type ValidationKonten struct {
	IdTopik    string `json:"id_topik" binding:"required"`
	Judul      string `json:"judul" binding:"required"`
	Keterangan string `json:"keterangan"`
	ImageSrc   string `json:"image_src" binding:"required"`
}

type ValidateUser struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
	FullName string `json:"full_name" binding:"required"`
	Gender   string `json:"gender" binding:"required"`
	Level    string `json:"level" binding:"required"`
}
