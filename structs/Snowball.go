package structs

import (
	"log"
)

type Snowball struct {
	Response struct {
		Num           string `json:"num"`
		ReportType    int    `json:"report_type"`
		PartnerId     string `json:"partnerid"`
		Datetime      string `json:"datetime"`
		Status        int    `json:"status"`
		OverflowRatio int    `json:"overflow_ratio"`
	} `json:"response"`

	Credits []struct {
		CredID         int     `json:"cred_id"`
		CredType       int     `json:"cred_type"`
		CredActive     int     `json:"cred_active"`
		CredSum        int     `json:"cred_sum"`
		CredSumDebt    int     `json:"cred_sum_debt"`
		CredCurrency   string  `json:"cred_currency"`
		CredUpdate     string  `json:"cred_update"`
		CredPercent    int     `json:"cred_percent"`
		CredFullCost   float64 `json:"cred_full_cost"`
		CredAnnuity    int     `json:"cred_annuity"`
		CredSource     string  `json:"cred_source"`
		CredDate       string  `json:"cred_date"`
		CredEnddate    string  `json:"cred_enddate"`
		CredDatePayout string  `json:"cred_date_payout"`
	} `json:"credits"`

	Collapsed struct {
		Short struct {
			Count int     `json:"count"`
			Sum   float64 `json:"sum"`
		} `json:"short"`

		Overflow struct {
			Count int     `json:"count"`
			Sum   float64 `json:"sum"`
		} `json:"overflow"`
	} `json:"collapsed"`

	overdue Agregate `json:"overdue"`
}

type Agregate struct {
	Count int     `json:"count"`
	Sum   float64 `json:"sum"`
}

func (sb Snowball) Generate() float64 {
	/**
	var a, b *int64
	if true { // супер условие
	   var fooA = int64(100)
	   a = &fooA
	}
	if true { // супер условие
	   var fooB = int64(-100)
	   b = &fooB
	}
	foo(a, b)
	*/
	var a, b *int64
	*a = 100
	*b = -100
	log.Println(generateSum(true, a, b))
	*a = 200
	*b = 100
	log.Println(generateSum(false, a, b))
	*a = -100
	*b = -200
	log.Println(generateSum(false, a, b))
	res, _ := generateSum(true, nil, nil)
	return res
}

func (ag Agregate) Generate(limit int) {
	ag.Count = 1
	ag.Sum = 1
}
