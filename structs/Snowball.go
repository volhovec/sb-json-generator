package structs

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

	Overdue struct {
		Count int     `json:"count"`
		Sum   float64 `json:"sum"`
	} `json:"overdue"`
}

func (sb Snowball) Generate() string {

	return "Другой"
}
