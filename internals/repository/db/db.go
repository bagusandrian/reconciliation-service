package db

//go:generate mockery --name=Handler --filename=mock_handler.go --inpackage
type DB interface {
	GetDummyQuery()
}
