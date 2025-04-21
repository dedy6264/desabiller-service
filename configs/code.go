package configs

var (
	VALIDATE_ERROR_CODE = "34"
	DB_ERROR            = "81"
	DB_NOT_FOUND        = "82"

	//trx
	SUCCESS_MSG           = "SUCCESS"
	PENDING_MSG           = "PENDING"
	FAILED_MSG            = "FAILED"
	BILLER_DISRUPTION_MSG = "BILLER DISRUPTION"
	BILL_NOTFOUND_MSG     = "BILL NOT FOUND"
	BIIL_PAID_MSG         = "SUDAH TERBAYAR"
	UNDEFINED_MSG         = "UNDEFINED"

	SUCCESS_CODE         = "00"
	PENDING_CODE         = "05"
	FAILED_CODE          = "09"
	INQUIRY_SUCCESS_CODE = "10"
	INQUIRY_PENDING_CODE = "15"
	INQUIRY_FAILED_CODE  = "19"
	//status code worker/helper
	WORKER_SUCCESS_CODE = "000"
	WORKER_PENDING_CODE = "500"

	WORKER_FAILED_CODE      = "900"
	WORKER_INVALID_PARAM    = "910"
	WORKER_CREDENTIAL_ERROR = "920"
	WORKER_VALIDATION_ERROR = "930"
	WORKER_SYSTEM_ERROR     = "940"
	WORKER_UNDEFINED_ERROR  = "950"
	// WORKER_VALIDATE_SYSTEM   = "090"
	BILLER_DISRUPTION         = "960"
	WORKER_BILL_NOTFOUND_CODE = "970"
)

// pending{
// 	02	BILL UNPAID	Failed	Your bill is unpaid, only reaching inquiry status. Please finish your payment first.

// }
// failed{
// 	WORKER_FAILED_CODE
// 	"06"	TRANSACTION NOT FOUND	Failed	There is no transaction with your inputted ref_id. Please check again your inputted ref_id to find your transaction.
// "07"	FAILED	Failed	Your current transaction has failed. Please try again.
// "13"	CUSTOMER NUMBER BLOCKED	Failed	Your customer number (customer_id) has been blocked. You can change your customer number (customer_id) or contact our Customer Service.
// "18",	NUMBER NOT AVAILABLE	Failed	You can see all available E-SIM number by using E-SIM List API.
// "20",	CODE NOT FOUND	Failed	Your inputted product_code isn’t in the database. Check again your product_code, you can check product_code list by using Pricelist API.
// "21",	NUMBER EXPIRED	Failed	Your phone number (customer_id) is expired. You can try other phone number.
// "132",	PRODUCT CODE NOT ELIGIBLE DUE TO SUBSCRIBER LOCATION	Failed	Your inputted product_code isn’t eligible due to subscriber location. Please try again with different product_code.
// 	"106",	PRODUCT IS TEMPORARILY OUT OF SERVICE	Failed	The product_code that you pick is in non-active status. You can retry your transaction with another product_code that has active status.
//new
// "09",	INQUIRY FAILED	Failed	Your inquiry process failed. Please try to do the inquiry again.
// "30",	PAYMENT HAVE TO BE DONE VIA COUNTER / PDAM	Failed	Your inputted payment has to be done at the counter. Please do your transaction at the counter.
// "33",	TRANSACTION CAN'T BE PROCESS, PLEASE TRY AGAIN LATER	Failed	Transaction can't be process, please try again later.
// "37",	PAYMENT FAILED	Failed	Your payment failed, please try again later.
// "38",	PAYMENT FAILED, PLEASE DO ANOTHER REQUEST	Failed	Your payment failed, please try again with a new request.
// "91",	DATABASE CONNECTION ERROR	Failed	There is an error on the database connection. Please try again later.
// "92",	GENERAL ERROR	Failed	The received response code is undefined yet. Please contact our Customer Service.
// "105",	MISC ERROR / BILLER SYSTEM ERROR	Failed	There is an error from the supplier. Please try again later.
//endnew

// WORKER_VALIDATION_ERROR
//new
// "01",	INVOICE HAS BEEN PAID	Failed	The invoice with your inputted data has already been paid. You don’t need to pay it again, or you can check for your inputted customer_id.
// "03",	INVALID REF ID	Failed	Your inputted ref_id isn’t valid. Please follow the correct format for ref_id (alpha_num only without space). Try again with a valid ref_id.
// "04",	BILLING ID EXPIRED	Failed	Your reference ID (ref_id) is expired. Please retry your inquiry request with a different reference ID (ref_id).
// "08",	BILLING ID BLOCKED	Failed	The customer ID for your inputted product code is blocked by IAK. Please contact our Customer Service.
// "11",	DUPLICATE REF ID	Failed	The ref_id that you’ve inputted is already been inputted, try again with another ref_id.
// "31",	TRANSACTION REJECTED DUE TO EXCEEDING MAXIMAL TOTAL BILL ALLOWED, MAXIMAL TOTAL BILL Rp50,000,000	Failed	Your current transaction is exceeding the maximum total bill allowed, maximum total bill is Rp50.000.000.
// "32",	TRANSACTION FAILED, PLEASE PAY BILL OF ALL PERIOD	Failed	Your current transaction isn’t covering all periods. Please try again to pay with all periods.
// "34",	BILL HAS BEEN PAID	Failed	Your bill for your current transaction has been paid. Please try again for another transaction.
// "35",	TRANSACTION REJECTED DUE TO ANOTHER UNPAID ARREAR	Failed	Your transaction id failed. Please pay for your other arrear first, then try again your transaction.
// "36",	EXCEEDING DUE DATE, PLEASE PAY IN THE COUNTER / PDAM	Failed	Your current transaction is exceeding the due date. Please pay the transaction at the counter.
// "40",	TRANSACTION REJECTED DUE TO ALL OR ONE OF THE ARREAR/INVOICE HAS BEEN PAID	Failed	Your transaction has already been paid. You can try again with another transaction.
// "41",	CAN'T BE PAID IN COUNTER, PLEASE PAY TO THE CORRESPONDING COMPANY	Failed	Your current transaction cannot be paid in the counter. You can pay your transaction to the corresponding company.
// "42",	PAYMENT REQUEST HAVEN'T BEEN RECEIVED	Failed	Your current transaction is still in the inquiry process. Please continue your payment process first.
// "100",	INVALID SIGNATURE	Failed	Your sign field doesn’t contain the right key for your current request. Please check again your sign value.
// "101",	INVALID COMMAND	Failed	The command that you’ve inputted is not a valid command, try checking your commands field for typos or try another command.
// "103",	TIMEOUT	Failed	Your current request exceeds the timeout limit. You can try to request it again.
//endnew
// "14",	INCORRECT DESTINATION NUMBER	Failed	Your customer_id that you’ve inputted isn’t a valid phone number. Please check again your customer_id.
// "16",	NUMBER NOT MATCH WITH OPERATOR	Failed	Your phone number (customer_id) that you’ve inputted doesn’t match with your desired operator (product_code). Please check again your phone number or change your operator.
// "19",	NUMBER IS ALREADY IN USE	Failed	Please select other E-SIM number.
// "131",	TOP UP REGION BLOCKED FOR PLAYER	Failed	Your current destination number top up request is blocked in that region. Please try again with a different destination number.
// "141",	INVALID USER ID / ZONE ID / SERVER ID / ROLENAME	Failed	Your inputted user ID / Zone ID / Server ID / Role name isn’t valid. Please try again with another user ID / Zone ID / Server ID / Role name. You can check on Inquiry Game Server.
// "142",	INVALID USER ID	Failed	Your current destination number (user id) top up request is invalid. Please try again with a different destination number or try checking for typos in your field.
// "206",	THIS DESTINATION NUMBER HAS BEEN BLOCKED	Failed	The customer_id that you inputted is blocked or not in whitelist. You can unblock it by remove customer number blacklist in API Security menu blacklist (https://developer.iak.id/end-user-blacklist) or add customer number whitelist in API Security menu whitelist (https://developer.iak.id/end-user-whitelist) on developer.iak.id.

// WORKER_INVALID_PARAM
// "203",	NUMBER IS TOO LONG	Failed	Your inputted customer ID is too long. Please check again your customer ID.
// 	"205",	WRONG COMMAND	Failed	The command that you’ve inputted is not a valid command, try check your commands field for typos or try another command.
// 	"107",	ERROR IN XML FORMAT	Failed	The body format of your request isn’t correct or there is an error in your body (required, ajax error, etc). Please use the correct JSON or XML format corresponding to your request to API. You can see the required body request for each request in the API Documentation.
//new
// "93"	INVALID AMOUNT	Failed	The amount inputted isn’t valid. Please check again your inputted amount.
//endnew

// WORKER_CREDENTIAL_ERROR
// 	"102",	INVALID IP ADDRESS	Failed	Your IP address isn’t allowed to make a transaction. You can add your IP address to your allowed IP address list in https://developer.iak.id/prod-setting.

// system
// 	"12",	BALANCE MAXIMUM LIMIT EXCEEDED	Failed	-
// 	"204",	WRONG AUTHENTICATION	Failed	Your sign (signature) field doesn’t contain the right key for your current request. Please check again your sign value.
// 	"17",	INSUFFICIENT DEPOSIT	Failed	Your current deposit is lower than the product_price you want to buy. You can add more money into your deposit by doing top up on iak.id deposit menu, or if you are in development mode, you can add your development deposit by clicking the + (plus) sign on development deposit menu (https://developer.iak.id/sandbox-report).
// 	"110",	SYSTEM UNDER MAINTENANCE	Failed	The system is currently under maintenance, you can try again later.
// 	"202",	MAXIMUM 1 NUMBER 1 TIME IN 1 DAY	Failed	You can only top up to a phone number once in a day (based on your developer setting). If you want to allow more than one top up to a phone number, you can go to https://developer.iak.id/ then choose API Setting menu, you can turn on “Allow multiple transactions for the same number” in development or production settings.
// 	"207",	MAXIMUM 1 NUMBER WITH ANY CODE 1 TIME IN 1 DAY	Failed	You’ve already done a transaction today. Please do another transaction tomorrow, or disable the high restriction setting in https://developer.iak.id/prod-setting.
// 	"121",	MONTHLY TOP UP LIMIT EXCEEDED	Failed	This response code applies to OVO products.
// 	"117",	PAGE NOT FOUND	Failed	The API URL that you want to hit is not found. Try checking your request URL for any typos or try other API URLs.
// 	"10",	REACH TOP UP LIMIT USING SAME DESTINATION NUMBER IN 1 DAY	Failed	Your current destination number top up request is reaching the limit on that day. Please try again tomorrow.
// 201	UNDEFINED RESPONSE CODE	Pending	The received response code is undefined yet. Please contact our Customer Service.
// 05	UNDEFINED ERROR
//new
// "94",	SERVICE HAS EXPIRED	Failed	-
// "108",	SORRY, YOUR ID CAN'T BE USED FOR THIS PRODUCT TRANSACTION	Failed	The customer ID that you’ve inputted can’t be used for this product transaction. Please try again with another customer ID or another product.
// "109",	SYSTEM CUT OFF	Failed	PLN product code cannot receive a request at 11PM until 1AM (GMT +7). Please try again when the service is available.
//endnew

//other
// }
