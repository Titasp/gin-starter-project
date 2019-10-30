package models

import (
	"github.com/titasp/gin-starter-project/database"
)

type Certificate struct {
	ID   string `json:"certificate_id,omitempty"`
	Data []byte `json:"certificate_data"`
}

func (c Certificate) GetByID(userId string, certId string) (*Certificate, error) {
	db := database.GetDB()
	var cert *Certificate

	tx, err := db.Begin(false)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	userBucket := tx.Bucket([]byte(userId))
	if userBucket == nil {
		return nil, nil
	}

	err = userBucket.ForEach(func(orgId, v []byte) error {
		if orgBucket := userBucket.Bucket([]byte(orgId)); orgBucket != nil {
			certData := orgBucket.Get([]byte(certId))
			if certData != nil {
				cert = &Certificate{
					ID:   certId,
					Data: certData,
				}
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	// Commit the transaction.
	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return cert, nil
}
