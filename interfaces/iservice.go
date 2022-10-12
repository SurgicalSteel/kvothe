package interfaces

import "github.com/SurgicalSteel/kvothe/models"

type InterfaceKvotheService interface {
	GetSongQuoteByID(id int64) (*models.SongQuote, int, error)
	GetAllSongData() ([]models.SongQuote, int, error)
	BackfillRedis() error

	//mock data
	GetQuoteByIDMockData(rawID string) (string, error)
	AddNewQuoteMockData(q string) (string, error)

	GetPersonByIDMockData(rawID string) (string, error)
	AddNewPersonMockData(p string) (string, error)
}
