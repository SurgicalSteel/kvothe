package services

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/SurgicalSteel/kvothe/utils"
)

func (ks *KvotheService) GetQuoteByIDMockData(rawID string) (string, error) {
	rawID = strings.Trim(rawID, " ")
	id, err := strconv.Atoi(rawID)
	if err != nil {
		return "", err
	}

	if q, ok := quoteMap[id]; ok {
		return q, nil
	}
	return "", fmt.Errorf("Quote with id %d is not found!", id)
}

func (ks *KvotheService) GetPersonByIDMockData(rawID string) (string, error) {
	rawID = strings.Trim(rawID, " ")
	id, err := strconv.Atoi(rawID)
	if err != nil {
		return "", err
	}

	if q, ok := personMap[id]; ok {
		return q, nil
	}
	return "", fmt.Errorf("Person with id %d is not found!", id)
}

func (ks *KvotheService) AddNewQuoteMockData(q string) (string, error) {
	for _, tq := range quoteMap {
		if strings.EqualFold(tq, q) {
			return "", fmt.Errorf("Quote is already exist")
		}
	}

	maxID := 0
	for tid, _ := range quoteMap {
		maxID = utils.Maxint(maxID, tid)
	}
	maxID += 1
	quoteMap[maxID] = q

	return "Quote has been added", nil
}

func (ks *KvotheService) AddNewPersonMockData(p string) (string, error) {
	for _, tp := range personMap {
		if strings.EqualFold(tp, p) {
			return "", fmt.Errorf("Person is already exist")
		}
	}

	maxID := 0
	for tid, _ := range personMap {
		maxID = utils.Maxint(maxID, tid)
	}
	maxID += 1
	quoteMap[maxID] = p

	return "Person has been added", nil
}
