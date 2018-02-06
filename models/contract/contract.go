package contract

/********

	Contract bean

 */
type Contract struct {
	ContractId      int64
	ContractName    string
	ContractContent string
	Photo           string
	StartTime       string
	EndTime         string
	InsertTime string
	UpdateTime string
	IsActive   bool
}
