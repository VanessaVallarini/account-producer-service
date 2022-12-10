package models

const (
	AccountCreateSubject = "com.account.create"
	AccountCreateAvro    = `{
		"type":"record",
		"name":"Account_Create",
		"namespace":"com.account.create",
		"fields":[
			 {
				"name":"alias",
				"type":"string"
			 },
			 {
				"name":"city",
				"type":"string"
			 },
			 {
				"name":"district",
				"type":"string"
			 },
			 {
				"name":"email",
				"type":"string"
			 },
			 {
				"name":"full_number",
				"type":"string"
			 },
			 {
				"name":"name",
				"type":"string"
			 },
			 {
				"name":"public_place",
				"type":"string"
			 },
			 {
				"name":"zip_code",
				"type":"string"
			 }		   
		]
	 }`
)

type AccountCreateEvent struct {
	Alias       string `json:"alias" avro:"alias"`
	City        string `json:"city" avro:"city"`
	District    string `json:"district" avro:"district"`
	Email       string `json:"email" avro:"email"`
	FullNumber  string `json:"full_number" avro:"full_number"`
	Name        string `json:"name" avro:"name"`
	PublicPlace string `json:"public_place" avro:"public_place"`
	ZipCode     string `json:"zip_code" avro:"zip_code"`
}
