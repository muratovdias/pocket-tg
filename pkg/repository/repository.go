package repository

type Bucket string

const (
	AccessToken  Bucket = "access_tokens"
	RequestToken Bucket = "request_tokens"
)

type TokenRepo interface {
	Save(chatID int64, toke string, bucket Bucket) error
	Get(chatID int64, bucket Bucket) (string, error)
}

//type Repository struct {
//	TokenRepo
//}
//
//func NewRepository(boltRepo *boltdb.Repo) *Repository {
//	return &Repository{TokenRepo: boltRepo}
//}
