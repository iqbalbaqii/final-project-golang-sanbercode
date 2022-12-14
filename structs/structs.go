package structs

type Topik struct {
	Id         int64  `json:"-"`
	Judul      string `json:"judul"`
	Pertanyaan string `json:"pertanyaan"`
	Periode    string `json:"periode"`
	StartDate  string `json:"start_date"`
	EndDate    string `json:"end_date"`
	CreatedBy  string `json:"-"`
	CreatedAt  string `json:"-"`
	UpdateAt   string `json:"-"`
	Deleted    int    `json:"-"`
	IsRelease  int    `json:"-"`
	Target     string `json:"target"`
}

type Konten struct {
	Id         int64  `json:"-"`
	IdTopik    int64  `json:"id_topik"`
	Judul      string `json:"judul"`
	Keterangan string `json:"keterangan"`
	ImageSrc   string `json:"image_src"`
	Deleted    int    `json:"-"`
}

type Respon struct {
	Id        int64  `json:"-"`
	IdTopik   int64  `json:"id_topik"`
	Response  int64  `json:"response"`
	CreatedBy string `json:"created_by"`
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
	Level     int    `json:"level"`
	CreatedAt string `json:"-"`
	UpdateAt  string `json:"-"`
}
type Polling struct {
	Topik  Topik
	Konten Konten
}
