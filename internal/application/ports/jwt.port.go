package ports

type Token string

type IJWT interface {
	Generate(data interface{}) Token
	Verify(tokenString Token) (interface{}, error)
}
