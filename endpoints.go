package merit

type apiEndpoint string

const (

	// Sales invoices
	// epGetListOfSalesInvoices apiEndpoint = iota
	// epGetSalesInvoiceDetails
	// epDeleteInvoice
	// epCreateSalesInvoice

	// Sales offers
	epGetListOfSalesOffers apiEndpoint = "v2/getoffers"
	epCreateSalesOffer     apiEndpoint = "v2/sendoffer"
	epSetOfferStatus       apiEndpoint = "v2/setofferstatus"
	// epCreateInvoiceFromOffer
	// epGetSalesOfferDetails
	// epUpdateSalesOffer

	// Get price
	// epGetPrice

	// Recurring Invoices
	// epCreateRecurringInvoice
	// epSendIndicationValues
	// epGetReccurringInvoicesClientAddressList
	// epGetReccurringInvoicesList
	// epGetReccurringInvoiceDetails

	// Purchase invoices
	// epGetListOfPurchaseInvoices
	// epGetPurchaseInvoiceDetails
	// epDeletePurchaseInvoice
	// epCreatePurchaseInvoice

	// Inventory movements
	epGetListOfLocations          apiEndpoint = "v2/getlocations"
	epGetListOfInventoryMovements apiEndpoint = "v2/getinvmovements"
	// epCreateInventoryMovement

	// Payments
	// epGetListOfPayments
	// epGetListOfPaymentTypes
	// epCreatePaymentOfSalesInvoice
	// epCreatePaymentOfPurchaseInvoice
	// epCreatePaymentOfSalesOffer
	// epDeletePayment

	// General ledger transactions
	// epCreateGeneralLedgerTransaction
	// epGetListOfGeneralLedgerTransactions
	// epGetGeneralLedgerTransactionDetails
	// epGetGeneralLedgerTransactionFullDetails

	// Tax list
	epGetListOfTaxes apiEndpoint = "v1/gettaxes"

	// Send tax
	// epSendTax

	// Customers
	epGetListOfCustomers apiEndpoint = "v1/getcustomers"
	// epCreateCustomer
	// epUpdateCustomer

	// Vendors
	// epGetListOfVendors
	// epCreateVendor
	// epUpdateVendor

	// Accounts list
	// epGetAccountsList

	// Project List
	// epGetProjectList

	// Cost centers list
	// epGetCostCentersList

	// Dimensions
	epGetListOfDimensions apiEndpoint = "v2/getdimensions"
	// epCreateDimensions
	// epAddDimensionValues

	// Departments List
	// epGetDepartmentsList

	// Unit of measure list
	epGetUnitOfMeasureList apiEndpoint = "v1/getunits"

	// Banks list
	// epGetBanksList

	// Financial years
	// epGetFinancialYearsList

	// Items
	epGetItems      apiEndpoint = "v1/getitems"
	epGetItemGroups apiEndpoint = "v2/getitemgroups"
	epSendItems     apiEndpoint = "v2/senditems"
	// epAddItemGroups
	// epAddItemGroups
	// epUpdateItem

	// Reports
	// epCustomerDebtsReport
	// epCustomerPaymentReport
	// epStatementOfProfitOrLoss
	// epStatementOfFinancialPosition
	epGetInventoryReport apiEndpoint = "v2/getinventoryreport"
	// epSalesReport
	epGetPurchaseReport apiEndpoint = "v2/getpurchrep"
)
