// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// BookingAppointment undocumented
type BookingAppointment struct {
	// Entity is the base model of BookingAppointment
	Entity

	ODataType string `json:"@odata.type,omitempty"`
	// AdditionalInformation undocumented
	AdditionalInformation *string `json:"additionalInformation,omitempty"`
	// AnonymousJoinWebURL undocumented
	AnonymousJoinWebURL *string `json:"anonymousJoinWebUrl,omitempty"`
	// CustomerEmailAddress undocumented
	CustomerEmailAddress *string `json:"customerEmailAddress,omitempty"`
	// CustomerID undocumented
	CustomerID *string `json:"customerId,omitempty"`
	// CustomerLocation undocumented
	CustomerLocation *Location `json:"customerLocation,omitempty"`
	// CustomerName undocumented
	CustomerName *string `json:"customerName,omitempty"`
	// CustomerNotes undocumented
	CustomerNotes *string `json:"customerNotes,omitempty"`
	// CustomerPhone undocumented
	CustomerPhone *string `json:"customerPhone,omitempty"`
	// Customers undocumented
	Customers []BookingCustomerInformationBase `json:"customers,omitempty"`
	// CustomerTimeZone undocumented
	CustomerTimeZone *string `json:"customerTimeZone,omitempty"`
	// Duration undocumented
	Duration *Duration `json:"duration,omitempty"`
	// End undocumented
	End *DateTimeTimeZone `json:"end,omitempty"`
	// FilledAttendeesCount undocumented
	FilledAttendeesCount *int `json:"filledAttendeesCount,omitempty"`
	// InvoiceAmount undocumented
	InvoiceAmount *float64 `json:"invoiceAmount,omitempty"`
	// InvoiceDate undocumented
	InvoiceDate *DateTimeTimeZone `json:"invoiceDate,omitempty"`
	// InvoiceID undocumented
	InvoiceID *string `json:"invoiceId,omitempty"`
	// InvoiceStatus undocumented
	InvoiceStatus *BookingInvoiceStatus `json:"invoiceStatus,omitempty"`
	// InvoiceURL undocumented
	InvoiceURL *string `json:"invoiceUrl,omitempty"`
	// IsLocationOnline undocumented
	IsLocationOnline *bool `json:"isLocationOnline,omitempty"`
	// JoinWebURL undocumented
	JoinWebURL *string `json:"joinWebUrl,omitempty"`
	// MaximumAttendeesCount undocumented
	MaximumAttendeesCount *int `json:"maximumAttendeesCount,omitempty"`
	// OnlineMeetingURL undocumented
	OnlineMeetingURL *string `json:"onlineMeetingUrl,omitempty"`
	// OptOutOfCustomerEmail undocumented
	OptOutOfCustomerEmail *bool `json:"optOutOfCustomerEmail,omitempty"`
	// PostBuffer undocumented
	PostBuffer *Duration `json:"postBuffer,omitempty"`
	// PreBuffer undocumented
	PreBuffer *Duration `json:"preBuffer,omitempty"`
	// Price undocumented
	Price *float64 `json:"price,omitempty"`
	// PriceType undocumented
	PriceType *BookingPriceType `json:"priceType,omitempty"`
	// Reminders undocumented
	Reminders []BookingReminder `json:"reminders,omitempty"`
	// SelfServiceAppointmentID undocumented
	SelfServiceAppointmentID *string `json:"selfServiceAppointmentId,omitempty"`
	// ServiceID undocumented
	ServiceID *string `json:"serviceId,omitempty"`
	// ServiceLocation undocumented
	ServiceLocation *Location `json:"serviceLocation,omitempty"`
	// ServiceName undocumented
	ServiceName *string `json:"serviceName,omitempty"`
	// ServiceNotes undocumented
	ServiceNotes *string `json:"serviceNotes,omitempty"`
	// SmsNotificationsEnabled undocumented
	SmsNotificationsEnabled *bool `json:"smsNotificationsEnabled,omitempty"`
	// StaffMemberIDs undocumented
	StaffMemberIDs []string `json:"staffMemberIds,omitempty"`
	// Start undocumented
	Start *DateTimeTimeZone `json:"start,omitempty"`
}

func NewBookingAppointment() (*BookingAppointment, error) {
	newBookingAppointment := &BookingAppointment{
		ODataType: "#microsoft.graph.BookingAppointment",
	}
	return newBookingAppointment, nil
}

// BookingBusiness undocumented
type BookingBusiness struct {
	// BookingNamedEntity is the base model of BookingBusiness
	BookingNamedEntity

	ODataType string `json:"@odata.type,omitempty"`
	// Address undocumented
	Address *PhysicalAddress `json:"address,omitempty"`
	// BusinessHours undocumented
	BusinessHours []BookingWorkHours `json:"businessHours,omitempty"`
	// BusinessType undocumented
	BusinessType *string `json:"businessType,omitempty"`
	// DefaultCurrencyIso undocumented
	DefaultCurrencyIso *string `json:"defaultCurrencyIso,omitempty"`
	// Email undocumented
	Email *string `json:"email,omitempty"`
	// IsPublished undocumented
	IsPublished *bool `json:"isPublished,omitempty"`
	// LanguageTag undocumented
	LanguageTag *string `json:"languageTag,omitempty"`
	// Phone undocumented
	Phone *string `json:"phone,omitempty"`
	// PublicURL undocumented
	PublicURL *string `json:"publicUrl,omitempty"`
	// SchedulingPolicy undocumented
	SchedulingPolicy *BookingSchedulingPolicy `json:"schedulingPolicy,omitempty"`
	// WebSiteURL undocumented
	WebSiteURL *string `json:"webSiteUrl,omitempty"`
	// Appointments undocumented
	Appointments []BookingAppointment `json:"appointments,omitempty"`
	// CalendarView undocumented
	CalendarView []BookingAppointment `json:"calendarView,omitempty"`
	// Customers undocumented
	Customers []BookingCustomer `json:"customers,omitempty"`
	// CustomQuestions undocumented
	CustomQuestions []BookingCustomQuestion `json:"customQuestions,omitempty"`
	// Services undocumented
	Services []BookingService `json:"services,omitempty"`
	// StaffMembers undocumented
	StaffMembers []BookingStaffMember `json:"staffMembers,omitempty"`
}

func NewBookingBusiness() (*BookingBusiness, error) {
	newBookingBusiness := &BookingBusiness{
		ODataType: "#microsoft.graph.BookingBusiness",
	}
	return newBookingBusiness, nil
}

// BookingCurrency undocumented
type BookingCurrency struct {
	// Entity is the base model of BookingCurrency
	Entity

	ODataType string `json:"@odata.type,omitempty"`
	// Symbol undocumented
	Symbol *string `json:"symbol,omitempty"`
}

func NewBookingCurrency() (*BookingCurrency, error) {
	newBookingCurrency := &BookingCurrency{
		ODataType: "#microsoft.graph.BookingCurrency",
	}
	return newBookingCurrency, nil
}

// BookingCustomQuestion undocumented
type BookingCustomQuestion struct {
	// Entity is the base model of BookingCustomQuestion
	Entity

	ODataType string `json:"@odata.type,omitempty"`
	// AnswerInputType undocumented
	AnswerInputType *AnswerInputType `json:"answerInputType,omitempty"`
	// AnswerOptions undocumented
	AnswerOptions []string `json:"answerOptions,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
}

func NewBookingCustomQuestion() (*BookingCustomQuestion, error) {
	newBookingCustomQuestion := &BookingCustomQuestion{
		ODataType: "#microsoft.graph.BookingCustomQuestion",
	}
	return newBookingCustomQuestion, nil
}

// BookingCustomer undocumented
type BookingCustomer struct {
	// BookingPerson is the base model of BookingCustomer
	BookingPerson

	ODataType string `json:"@odata.type,omitempty"`
	// Addresses undocumented
	Addresses []PhysicalAddress `json:"addresses,omitempty"`
	// Phones undocumented
	Phones []Phone `json:"phones,omitempty"`
}

func NewBookingCustomer() (*BookingCustomer, error) {
	newBookingCustomer := &BookingCustomer{
		ODataType: "#microsoft.graph.BookingCustomer",
	}
	return newBookingCustomer, nil
}

// BookingCustomerInformation undocumented
type BookingCustomerInformation struct {
	// BookingCustomerInformationBase is the base model of BookingCustomerInformation
	BookingCustomerInformationBase

	ODataType string `json:"@odata.type,omitempty"`
	// CustomerID undocumented
	CustomerID *string `json:"customerId,omitempty"`
	// CustomQuestionAnswers undocumented
	CustomQuestionAnswers []BookingQuestionAnswer `json:"customQuestionAnswers,omitempty"`
	// EmailAddress undocumented
	EmailAddress *string `json:"emailAddress,omitempty"`
	// Location undocumented
	Location *Location `json:"location,omitempty"`
	// Name undocumented
	Name *string `json:"name,omitempty"`
	// Notes undocumented
	Notes *string `json:"notes,omitempty"`
	// Phone undocumented
	Phone *string `json:"phone,omitempty"`
	// SmsNotificationsEnabled undocumented
	SmsNotificationsEnabled *bool `json:"smsNotificationsEnabled,omitempty"`
	// TimeZone undocumented
	TimeZone *string `json:"timeZone,omitempty"`
}

func NewBookingCustomerInformation() (*BookingCustomerInformation, error) {
	newBookingCustomerInformation := &BookingCustomerInformation{
		ODataType: "#microsoft.graph.BookingCustomerInformation",
	}
	return newBookingCustomerInformation, nil
}

// BookingCustomerInformationBase undocumented
type BookingCustomerInformationBase struct {
	// Object is the base model of BookingCustomerInformationBase
	Object

	ODataType string `json:"@odata.type,omitempty"`
}

func NewBookingCustomerInformationBase() (*BookingCustomerInformationBase, error) {
	newBookingCustomerInformationBase := &BookingCustomerInformationBase{
		ODataType: "#microsoft.graph.BookingCustomerInformationBase",
	}
	return newBookingCustomerInformationBase, nil
}

// BookingNamedEntity undocumented
type BookingNamedEntity struct {
	// Entity is the base model of BookingNamedEntity
	Entity

	ODataType string `json:"@odata.type,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
}

func NewBookingNamedEntity() (*BookingNamedEntity, error) {
	newBookingNamedEntity := &BookingNamedEntity{
		ODataType: "#microsoft.graph.BookingNamedEntity",
	}
	return newBookingNamedEntity, nil
}

// BookingPerson undocumented
type BookingPerson struct {
	// BookingNamedEntity is the base model of BookingPerson
	BookingNamedEntity

	ODataType string `json:"@odata.type,omitempty"`
	// EmailAddress undocumented
	EmailAddress *string `json:"emailAddress,omitempty"`
}

func NewBookingPerson() (*BookingPerson, error) {
	newBookingPerson := &BookingPerson{
		ODataType: "#microsoft.graph.BookingPerson",
	}
	return newBookingPerson, nil
}

// BookingQuestionAnswer undocumented
type BookingQuestionAnswer struct {
	// Object is the base model of BookingQuestionAnswer
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// Answer undocumented
	Answer *string `json:"answer,omitempty"`
	// AnswerInputType undocumented
	AnswerInputType *AnswerInputType `json:"answerInputType,omitempty"`
	// AnswerOptions undocumented
	AnswerOptions []string `json:"answerOptions,omitempty"`
	// IsRequired undocumented
	IsRequired *bool `json:"isRequired,omitempty"`
	// Question undocumented
	Question *string `json:"question,omitempty"`
	// QuestionID undocumented
	QuestionID *string `json:"questionId,omitempty"`
	// SelectedOptions undocumented
	SelectedOptions []string `json:"selectedOptions,omitempty"`
}

func NewBookingQuestionAnswer() (*BookingQuestionAnswer, error) {
	newBookingQuestionAnswer := &BookingQuestionAnswer{
		ODataType: "#microsoft.graph.BookingQuestionAnswer",
	}
	return newBookingQuestionAnswer, nil
}

// BookingQuestionAssignment undocumented
type BookingQuestionAssignment struct {
	// Object is the base model of BookingQuestionAssignment
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// IsRequired undocumented
	IsRequired *bool `json:"isRequired,omitempty"`
	// QuestionID undocumented
	QuestionID *string `json:"questionId,omitempty"`
}

func NewBookingQuestionAssignment() (*BookingQuestionAssignment, error) {
	newBookingQuestionAssignment := &BookingQuestionAssignment{
		ODataType: "#microsoft.graph.BookingQuestionAssignment",
	}
	return newBookingQuestionAssignment, nil
}

// BookingReminder undocumented
type BookingReminder struct {
	// Object is the base model of BookingReminder
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// Message undocumented
	Message *string `json:"message,omitempty"`
	// Offset undocumented
	Offset *Duration `json:"offset,omitempty"`
	// Recipients undocumented
	Recipients *BookingReminderRecipients `json:"recipients,omitempty"`
}

func NewBookingReminder() (*BookingReminder, error) {
	newBookingReminder := &BookingReminder{
		ODataType: "#microsoft.graph.BookingReminder",
	}
	return newBookingReminder, nil
}

// BookingSchedulingPolicy undocumented
type BookingSchedulingPolicy struct {
	// Object is the base model of BookingSchedulingPolicy
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// AllowStaffSelection undocumented
	AllowStaffSelection *bool `json:"allowStaffSelection,omitempty"`
	// MaximumAdvance undocumented
	MaximumAdvance *Duration `json:"maximumAdvance,omitempty"`
	// MinimumLeadTime undocumented
	MinimumLeadTime *Duration `json:"minimumLeadTime,omitempty"`
	// SendConfirmationsToOwner undocumented
	SendConfirmationsToOwner *bool `json:"sendConfirmationsToOwner,omitempty"`
	// TimeSlotInterval undocumented
	TimeSlotInterval *Duration `json:"timeSlotInterval,omitempty"`
}

func NewBookingSchedulingPolicy() (*BookingSchedulingPolicy, error) {
	newBookingSchedulingPolicy := &BookingSchedulingPolicy{
		ODataType: "#microsoft.graph.BookingSchedulingPolicy",
	}
	return newBookingSchedulingPolicy, nil
}

// BookingService undocumented
type BookingService struct {
	// BookingNamedEntity is the base model of BookingService
	BookingNamedEntity

	ODataType string `json:"@odata.type,omitempty"`
	// AdditionalInformation undocumented
	AdditionalInformation *string `json:"additionalInformation,omitempty"`
	// CustomQuestions undocumented
	CustomQuestions []BookingQuestionAssignment `json:"customQuestions,omitempty"`
	// DefaultDuration undocumented
	DefaultDuration *Duration `json:"defaultDuration,omitempty"`
	// DefaultLocation undocumented
	DefaultLocation *Location `json:"defaultLocation,omitempty"`
	// DefaultPrice undocumented
	DefaultPrice *float64 `json:"defaultPrice,omitempty"`
	// DefaultPriceType undocumented
	DefaultPriceType *BookingPriceType `json:"defaultPriceType,omitempty"`
	// DefaultReminders undocumented
	DefaultReminders []BookingReminder `json:"defaultReminders,omitempty"`
	// Description undocumented
	Description *string `json:"description,omitempty"`
	// IsAnonymousJoinEnabled undocumented
	IsAnonymousJoinEnabled *bool `json:"isAnonymousJoinEnabled,omitempty"`
	// IsHiddenFromCustomers undocumented
	IsHiddenFromCustomers *bool `json:"isHiddenFromCustomers,omitempty"`
	// IsLocationOnline undocumented
	IsLocationOnline *bool `json:"isLocationOnline,omitempty"`
	// LanguageTag undocumented
	LanguageTag *string `json:"languageTag,omitempty"`
	// MaximumAttendeesCount undocumented
	MaximumAttendeesCount *int `json:"maximumAttendeesCount,omitempty"`
	// Notes undocumented
	Notes *string `json:"notes,omitempty"`
	// PostBuffer undocumented
	PostBuffer *Duration `json:"postBuffer,omitempty"`
	// PreBuffer undocumented
	PreBuffer *Duration `json:"preBuffer,omitempty"`
	// SchedulingPolicy undocumented
	SchedulingPolicy *BookingSchedulingPolicy `json:"schedulingPolicy,omitempty"`
	// SmsNotificationsEnabled undocumented
	SmsNotificationsEnabled *bool `json:"smsNotificationsEnabled,omitempty"`
	// StaffMemberIDs undocumented
	StaffMemberIDs []string `json:"staffMemberIds,omitempty"`
	// WebURL undocumented
	WebURL *string `json:"webUrl,omitempty"`
}

func NewBookingService() (*BookingService, error) {
	newBookingService := &BookingService{
		ODataType: "#microsoft.graph.BookingService",
	}
	return newBookingService, nil
}

// BookingStaffMember undocumented
type BookingStaffMember struct {
	// BookingPerson is the base model of BookingStaffMember
	BookingPerson

	ODataType string `json:"@odata.type,omitempty"`
	// AvailabilityIsAffectedByPersonalCalendar undocumented
	AvailabilityIsAffectedByPersonalCalendar *bool `json:"availabilityIsAffectedByPersonalCalendar,omitempty"`
	// ColorIndex undocumented
	ColorIndex *int `json:"colorIndex,omitempty"`
	// IsEmailNotificationEnabled undocumented
	IsEmailNotificationEnabled *bool `json:"isEmailNotificationEnabled,omitempty"`
	// MembershipStatus undocumented
	MembershipStatus *BookingStaffMembershipStatus `json:"membershipStatus,omitempty"`
	// Role undocumented
	Role *BookingStaffRole `json:"role,omitempty"`
	// TimeZone undocumented
	TimeZone *string `json:"timeZone,omitempty"`
	// UseBusinessHours undocumented
	UseBusinessHours *bool `json:"useBusinessHours,omitempty"`
	// WorkingHours undocumented
	WorkingHours []BookingWorkHours `json:"workingHours,omitempty"`
}

func NewBookingStaffMember() (*BookingStaffMember, error) {
	newBookingStaffMember := &BookingStaffMember{
		ODataType: "#microsoft.graph.BookingStaffMember",
	}
	return newBookingStaffMember, nil
}

// BookingWorkHours undocumented
type BookingWorkHours struct {
	// Object is the base model of BookingWorkHours
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// Day undocumented
	Day *DayOfWeek `json:"day,omitempty"`
	// TimeSlots undocumented
	TimeSlots []BookingWorkTimeSlot `json:"timeSlots,omitempty"`
}

func NewBookingWorkHours() (*BookingWorkHours, error) {
	newBookingWorkHours := &BookingWorkHours{
		ODataType: "#microsoft.graph.BookingWorkHours",
	}
	return newBookingWorkHours, nil
}

// BookingWorkTimeSlot undocumented
type BookingWorkTimeSlot struct {
	// Object is the base model of BookingWorkTimeSlot
	Object

	ODataType string `json:"@odata.type,omitempty"`
	// End undocumented
	End *TimeOfDay `json:"end,omitempty"`
	// Start undocumented
	Start *TimeOfDay `json:"start,omitempty"`
}

func NewBookingWorkTimeSlot() (*BookingWorkTimeSlot, error) {
	newBookingWorkTimeSlot := &BookingWorkTimeSlot{
		ODataType: "#microsoft.graph.BookingWorkTimeSlot",
	}
	return newBookingWorkTimeSlot, nil
}