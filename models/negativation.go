package models

type Negativation struct {
	CompanyDocument   string        `bson:"companyDocument" json:"companyDocument"`
	CompanyName       string        `bson:"companyName" json:"companyName"`
	CustomerDocument  string        `bson:"customerDocument" json:"customerDocument"`
	Value             int           `bson:"value" json:"value"`
	Contract          string        `bson:"contract" json:"contract"`
	DebtDate          string        `bson:"debtDate" json:"debtDate"`
	InclusionDate     string        `bson:"inclusionDate" json:"inclusionDate"`

}

type FilterNegativation struct {
	CompanyName      *string
	CustomerDocument *string
}
