package services

import (
	"log"

	"serasa/models"

	"gopkg.in/mgo.v2/bson"

	mgo "gopkg.in/mgo.v2"
)

type NegativationService struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	COLLECTION = "negativation"
)

func (p *NegativationService) Connect() {
	session, err := mgo.Dial(p.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(p.Database)
}

func (p *NegativationService) List(filterNegativation models.FilterNegativation) ([]models.Negativation, error) {
	var negativations []models.Negativation

	filter := bson.M{}
	if filterNegativation.CompanyName != nil {
		filter = bson.M{"companyName": bson.RegEx{*filterNegativation.CompanyName, "i"}}
	}

  if filterNegativation.CompanyDocument != nil {
		filter = bson.M{"companyDocument": bson.RegEx{*filterNegativation.CompanyDocument, "i"}}
	}

	if filterNegativation.CustomerDocument != nil {
		filter = bson.M{"customerDocument": bson.RegEx{*filterNegativation.CustomerDocument, "i"}}
	}

	err := db.C(COLLECTION).Find(filter).All(&negativations)
	return negativations, err
}

func (p *NegativationService) Create(negativations []models.Negativation) error {
  for _, value := range negativations {
    err := db.C(COLLECTION).Insert(&value)
    if err != nil {
      return err
    }
  }
	return nil
}
