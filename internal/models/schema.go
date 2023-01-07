package models

const (
	AccountCreateOrUpdateSubject = "com.account.create.or.update"
	AccountCreateOrUpdateAvro    = `{
		"type":"record",
		"name":"Account_Create_Or_Update",
		"namespace":"com.account.create.or.update",
		"fields":[
			 {
				"name":"email",
				"type":"string"
			 },
			 {
				"name":"full_number",
				"type":"string"
			 },
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
				"name":"name",
				"type":"string"
			 },
			 {
				"name":"public_place",
				"type":"string"
			 },
			 {
				"name":"status",
				"type":"string"
			 },
			 {
				"name":"zip_code",
				"type":"string"
			 }		   
		]
	 }`
	AccountDeleteSubject = "com.account.delete"
	AccountDeleteAvro    = `{
		"type":"record",
		"name":"Account_Delete",
		"namespace":"com.account.delete",
		"fields":[
			{
				"name":"email",
				"type":"string"
			 }	   
		]
	 }`
	AccountGetSubject = "com.account.get"
	AccountGetAvro    = `{
		"type":"record",
		"name":"Account_Get",
		"namespace":"com.account.get",
		"fields":[
			{
				"name":"email",
				"type":"string"
			 }	   
		]
	 }`
	AccountResponseSubject = "com.account.response"
	AccountResponseAvro    = `{
		"type":"record",
		"name":"Account_Response",
		"namespace":"com.account.response",
		"fields":[
			 {
				"name":"email",
				"type":"string"
			 },
			 {
				"name":"full_number",
				"type":"string"
			 },
			 {
				"name":"alias",
				"type":"string"
			 },
			 {
				"name":"city",
				"type":"string"
			 },
			 {
				"name":"date_time",
				"type":"string"
			 },
			 {
				"name":"district",
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
				"name":"status",
				"type":"string"
			 },
			 {
				"name":"zip_code",
				"type":"string"
			 }		   
		]
	 }`
)

type AccountCreateOrUpdateEvent struct {
	Email       string `json:"email" avro:"email"`
	FullNumber  string `json:"full_number" avro:"full_number"`
	Alias       string `json:"alias" avro:"alias"`
	City        string `json:"city" avro:"city"`
	District    string `json:"district" avro:"district"`
	Name        string `json:"name" avro:"name"`
	PublicPlace string `json:"public_place" avro:"public_place"`
	Status      string `json:"status" avro:"status"`
	ZipCode     string `json:"zip_code" avro:"zip_code"`
}

type AccountDeleteEvent struct {
	Email string `json:"email" avro:"email"`
}

type AccountGetEvent struct {
	Email string `json:"email" avro:"email"`
}

type AccountGetResponseEvent struct {
	Email       string `json:"email" avro:"email"`
	FullNumber  string `json:"full_number" avro:"full_number"`
	Alias       string `json:"alias" avro:"alias"`
	City        string `json:"city" avro:"city"`
	DateTime    string `json:"date_time" avro:"date_time"`
	District    string `json:"district" avro:"district"`
	Name        string `json:"name" avro:"name"`
	PublicPlace string `json:"public_place" avro:"public_place"`
	Status      string `json:"status" avro:"status"`
	ZipCode     string `json:"zip_code" avro:"zip_code"`
}
