package ports

type IHasher interface {
	Hash(password string) (string, error)
	Compare(hashPassoword, password string) (bool, error)
}
