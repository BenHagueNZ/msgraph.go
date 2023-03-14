
// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph


// PurchaseInvoice undocumented
type PurchaseInvoice struct {
    // Entity is the base model of PurchaseInvoice
    Entity
    // BuyFromAddress undocumented
    BuyFromAddress *PostalAddressType `json:"buyFromAddress,omitempty"`
    // CurrencyCode undocumented
    CurrencyCode *string `json:"currencyCode,omitempty"`
    // CurrencyID undocumented
    CurrencyID *UUID `json:"currencyId,omitempty"`
    // DiscountAmount undocumented
    DiscountAmount *int `json:"discountAmount,omitempty"`
    // DiscountAppliedBeforeTax undocumented
    DiscountAppliedBeforeTax *bool `json:"discountAppliedBeforeTax,omitempty"`
    // DueDate undocumented
    DueDate *Date `json:"dueDate,omitempty"`
    // InvoiceDate undocumented
    InvoiceDate *Date `json:"invoiceDate,omitempty"`
    // LastModifiedDateTime undocumented
    LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
    // Number undocumented
    Number *string `json:"number,omitempty"`
    // PayToAddress undocumented
    PayToAddress *PostalAddressType `json:"payToAddress,omitempty"`
    // PayToContact undocumented
    PayToContact *string `json:"payToContact,omitempty"`
    // PayToName undocumented
    PayToName *string `json:"payToName,omitempty"`
    // PayToVendorID undocumented
    PayToVendorID *UUID `json:"payToVendorId,omitempty"`
    // PayToVendorNumber undocumented
    PayToVendorNumber *string `json:"payToVendorNumber,omitempty"`
    // PricesIncludeTax undocumented
    PricesIncludeTax *bool `json:"pricesIncludeTax,omitempty"`
    // ShipToAddress undocumented
    ShipToAddress *PostalAddressType `json:"shipToAddress,omitempty"`
    // ShipToContact undocumented
    ShipToContact *string `json:"shipToContact,omitempty"`
    // ShipToName undocumented
    ShipToName *string `json:"shipToName,omitempty"`
    // Status undocumented
    Status *string `json:"status,omitempty"`
    // TotalAmountExcludingTax undocumented
    TotalAmountExcludingTax *int `json:"totalAmountExcludingTax,omitempty"`
    // TotalAmountIncludingTax undocumented
    TotalAmountIncludingTax *int `json:"totalAmountIncludingTax,omitempty"`
    // TotalTaxAmount undocumented
    TotalTaxAmount *int `json:"totalTaxAmount,omitempty"`
    // VendorID undocumented
    VendorID *UUID `json:"vendorId,omitempty"`
    // VendorInvoiceNumber undocumented
    VendorInvoiceNumber *string `json:"vendorInvoiceNumber,omitempty"`
    // VendorName undocumented
    VendorName *string `json:"vendorName,omitempty"`
    // VendorNumber undocumented
    VendorNumber *string `json:"vendorNumber,omitempty"`
    // Currency undocumented
    Currency *Currency `json:"currency,omitempty"`
    // PurchaseInvoiceLines undocumented
    PurchaseInvoiceLines []PurchaseInvoiceLine `json:"purchaseInvoiceLines,omitempty"`
    // Vendor undocumented
    Vendor *Vendor `json:"vendor,omitempty"`
}

// PurchaseInvoiceLine undocumented
type PurchaseInvoiceLine struct {
    // Entity is the base model of PurchaseInvoiceLine
    Entity
    // AccountID undocumented
    AccountID *UUID `json:"accountId,omitempty"`
    // AmountExcludingTax undocumented
    AmountExcludingTax *int `json:"amountExcludingTax,omitempty"`
    // AmountIncludingTax undocumented
    AmountIncludingTax *int `json:"amountIncludingTax,omitempty"`
    // Description undocumented
    Description *string `json:"description,omitempty"`
    // DiscountAmount undocumented
    DiscountAmount *int `json:"discountAmount,omitempty"`
    // DiscountAppliedBeforeTax undocumented
    DiscountAppliedBeforeTax *bool `json:"discountAppliedBeforeTax,omitempty"`
    // DiscountPercent undocumented
    DiscountPercent *int `json:"discountPercent,omitempty"`
    // DocumentID undocumented
    DocumentID *UUID `json:"documentId,omitempty"`
    // ExpectedReceiptDate undocumented
    ExpectedReceiptDate *Date `json:"expectedReceiptDate,omitempty"`
    // InvoiceDiscountAllocation undocumented
    InvoiceDiscountAllocation *int `json:"invoiceDiscountAllocation,omitempty"`
    // ItemID undocumented
    ItemID *UUID `json:"itemId,omitempty"`
    // LineType undocumented
    LineType *string `json:"lineType,omitempty"`
    // NetAmount undocumented
    NetAmount *int `json:"netAmount,omitempty"`
    // NetAmountIncludingTax undocumented
    NetAmountIncludingTax *int `json:"netAmountIncludingTax,omitempty"`
    // NetTaxAmount undocumented
    NetTaxAmount *int `json:"netTaxAmount,omitempty"`
    // Quantity undocumented
    Quantity *int `json:"quantity,omitempty"`
    // Sequence undocumented
    Sequence *int `json:"sequence,omitempty"`
    // TaxCode undocumented
    TaxCode *string `json:"taxCode,omitempty"`
    // TaxPercent undocumented
    TaxPercent *int `json:"taxPercent,omitempty"`
    // TotalTaxAmount undocumented
    TotalTaxAmount *int `json:"totalTaxAmount,omitempty"`
    // UnitCost undocumented
    UnitCost *int `json:"unitCost,omitempty"`
    // Account undocumented
    Account *Account `json:"account,omitempty"`
    // Item undocumented
    Item *Item `json:"item,omitempty"`
}
