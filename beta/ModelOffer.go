
// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph


// OfferShiftRequestObject undocumented
type OfferShiftRequestObject struct {
    // ScheduleChangeRequest is the base model of OfferShiftRequestObject
    ScheduleChangeRequest
    // RecipientActionDateTime undocumented
    RecipientActionDateTime *time.Time `json:"recipientActionDateTime,omitempty"`
    // RecipientActionMessage undocumented
    RecipientActionMessage *string `json:"recipientActionMessage,omitempty"`
    // RecipientUserID undocumented
    RecipientUserID *string `json:"recipientUserId,omitempty"`
    // SenderShiftID undocumented
    SenderShiftID *string `json:"senderShiftId,omitempty"`
}
