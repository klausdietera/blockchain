package models

type Account struct {
	PublicKey       string
	SecondPublicKey string
	Balance         int64
	Delegate        *Delegate // fix
	Votes           []string  //
	Referral        *Account  //
	Stakes          string    //
}
