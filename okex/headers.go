package okex

const (
	HeaderAccessKey        = "OK-ACCESS-KEY"        //The APIKey as a String
	HeaderAccessPassphrase = "OK-ACCESS-PASSPHRASE" //The passphrase you specified when creating the APIKey
	HeaderAccessSign       = "OK-ACCESS-SIGN"       //The Base64-encoded signature (see Signing Messages subsection for details)
	HeaderAccessTimestamp  = "OK-ACCESS-TIMESTAMP"  //The timestamp of your request.e.g : 2020-12-08T09:08:57.715Z
	HeaderSimulatedTrading = "x-simulated-trading"  //`x-simulated-trading: 1` needs to be added to the header of the Demo Trading request
	HeaderMethodGet        = "GET"
	HeaderMethodPost       = "POST"
)
