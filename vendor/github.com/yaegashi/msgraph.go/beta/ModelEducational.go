// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// EducationalActivity undocumented
type EducationalActivity struct {
	// ItemFacet is the base model of EducationalActivity
	ItemFacet
	// CompletionMonthYear undocumented
	CompletionMonthYear *Date `json:"completionMonthYear,omitempty"`
	// EndMonthYear undocumented
	EndMonthYear *Date `json:"endMonthYear,omitempty"`
	// Institution undocumented
	Institution *InstitutionData `json:"institution,omitempty"`
	// Program undocumented
	Program *EducationalActivityDetail `json:"program,omitempty"`
	// StartMonthYear undocumented
	StartMonthYear *Date `json:"startMonthYear,omitempty"`
}

// EducationalActivityDetail undocumented
type EducationalActivityDetail struct {
	// Object is the base model of EducationalActivityDetail
	Object
	// Abbreviation undocumented
	Abbreviation *string `json:"abbreviation,omitempty"`
	// Activities undocumented
	Activities *string `json:"activities,omitempty"`
	// Awards undocumented
	Awards *string `json:"awards,omitempty"`
	// Description undocumented
	Description *string `json:"description,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// FieldsOfStudy undocumented
	FieldsOfStudy *string `json:"fieldsOfStudy,omitempty"`
	// Grade undocumented
	Grade *string `json:"grade,omitempty"`
	// Notes undocumented
	Notes *string `json:"notes,omitempty"`
	// WebURL undocumented
	WebURL *string `json:"webUrl,omitempty"`
}
