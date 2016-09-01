package msg

// MaliciousCategoryLong is a map from malicious category to long string names
var MaliciousCategoryLong = map[MaliciousCategory]string{
	MaliciousCategory_T_BOTS_4:    "Botnet",
	MaliciousCategory_T_CANDC_4:   "Command and Control Centers",
	MaliciousCategory_T_COMPR_4:   "Compromised & Links To Malware",
	MaliciousCategory_T_MALHOME_4: "Malware Call-Home",
	MaliciousCategory_T_MAL_4:     "Malware Distribution Point",
	MaliciousCategory_T_FRAUD_4:   "Phishing/Fraud",
	MaliciousCategory_T_SPAM_4:    "Spam URLs",
	MaliciousCategory_T_SPYWARE_4: "Spyware & Questionable Software",
	MaliciousCategory_T_UNKNOWN:   "Unknown Category",
}
