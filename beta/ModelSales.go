
// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph


// SalesCreditMemo undocumented
type SalesCreditMemo struct {
    // Entity is the base model of SalesCreditMemo
    Entity
    // BillingPostalAddress undocumented
    BillingPostalAddress *PostalAddressType `json:"billingPostalAddress,omitempty"`
    // BillToCustomerID undocumented
    BillToCustomerID *UUID `json:"billToCustomerId,omitempty"`
    // BillToCustomerNumber undocumented
    BillToCustomerNumber *string `json:"billToCustomerNumber,omitempty"`
    // BillToName undocumented
    BillToName *string `json:"billToName,omitempty"`
    // CreditMemoDate undocumented
    CreditMemoDate *Date `json:"creditMemoDate,omitempty"`
    // CurrencyCode undocumented
    CurrencyCode *string `json:"currencyCode,omitempty"`
    // CurrencyID undocumented
    CurrencyID *UUID `json:"currencyId,omitempty"`
    // CustomerID undocumented
    CustomerID *UUID `json:"customerId,omitempty"`
    // CustomerName undocumented
    CustomerName *string `json:"customerName,omitempty"`
    // CustomerNumber undocumented
    CustomerNumber *string `json:"customerNumber,omitempty"`
    // DiscountAmount undocumented
    DiscountAmount *int `json:"discountAmount,omitempty"`
    // DiscountAppliedBeforeTax undocumented
    DiscountAppliedBeforeTax *bool `json:"discountAppliedBeforeTax,omitempty"`
    // DueDate undocumented
    DueDate *Date `json:"dueDate,omitempty"`
    // Email undocumented
    Email *string `json:"email,omitempty"`
    // ExternalDocumentNumber undocumented
    ExternalDocumentNumber *string `json:"externalDocumentNumber,omitempty"`
    // InvoiceID undocumented
    InvoiceID *UUID `json:"invoiceId,omitempty"`
    // InvoiceNumber undocumented
    InvoiceNumber *string `json:"invoiceNumber,omitempty"`
    // LastModifiedDateTime undocumented
    LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
    // Number undocumented
    Number *string `json:"number,omitempty"`
    // PaymentTermsID undocumented
    PaymentTermsID *UUID `json:"paymentTermsId,omitempty"`
    // PhoneNumber undocumented
    PhoneNumber *string `json:"phoneNumber,omitempty"`
    // PricesIncludeTax undocumented
    PricesIncludeTax *bool `json:"pricesIncludeTax,omitempty"`
    // Salesperson undocumented
    Salesperson *string `json:"salesperson,omitempty"`
    // SellingPostalAddress undocumented
    SellingPostalAddress *PostalAddressType `json:"sellingPostalAddress,omitempty"`
    // Status undocumented
    Status *string `json:"status,omitempty"`
    // TotalAmountExcludingTax undocumented
    TotalAmountExcludingTax *int `json:"totalAmountExcludingTax,omitempty"`
    // TotalAmountIncludingTax undocumented
    TotalAmountIncludingTax *int `json:"totalAmountIncludingTax,omitempty"`
    // TotalTaxAmount undocumented
    TotalTaxAmount *int `json:"totalTaxAmount,omitempty"`
    // Currency undocumented
    Currency *Currency `json:"currency,omitempty"`
    // Customer undocumented
    Customer *Customer `json:"customer,omitempty"`
    // PaymentTerm undocumented
    PaymentTerm *PaymentTerm `json:"paymentTerm,omitempty"`
    // SalesCreditMemoLines undocumented
    SalesCreditMemoLines []SalesCreditMemoLine `json:"salesCreditMemoLines,omitempty"`
}

// SalesCreditMemoLine undocumented
type SalesCreditMemoLine struct {
    // Entity is the base model of SalesCreditMemoLine
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
    // ShipmentDate undocumented
    ShipmentDate *Date `json:"shipmentDate,omitempty"`
    // TaxCode undocumented
    TaxCode *string `json:"taxCode,omitempty"`
    // TaxPercent undocumented
    TaxPercent *int `json:"taxPercent,omitempty"`
    // TotalTaxAmount undocumented
    TotalTaxAmount *int `json:"totalTaxAmount,omitempty"`
    // UnitOfMeasureID undocumented
    UnitOfMeasureID *UUID `json:"unitOfMeasureId,omitempty"`
    // UnitPrice undocumented
    UnitPrice *int `json:"unitPrice,omitempty"`
    // Account undocumented
    Account *Account `json:"account,omitempty"`
    // Item undocumented
    Item *Item `json:"item,omitempty"`
}

// SalesInvoice undocumented
type SalesInvoice struct {
    // Entity is the base model of SalesInvoice
    Entity
    // BillingPostalAddress undocumented
    BillingPostalAddress *PostalAddressType `json:"billingPostalAddress,omitempty"`
    // BillToCustomerID undocumented
    BillToCustomerID *UUID `json:"billToCustomerId,omitempty"`
    // BillToCustomerNumber undocumented
    BillToCustomerNumber *string `json:"billToCustomerNumber,omitempty"`
    // BillToName undocumented
    BillToName *string `json:"billToName,omitempty"`
    // CurrencyCode undocumented
    CurrencyCode *string `json:"currencyCode,omitempty"`
    // CurrencyID undocumented
    CurrencyID *UUID `json:"currencyId,omitempty"`
    // CustomerID undocumented
    CustomerID *UUID `json:"customerId,omitempty"`
    // CustomerName undocumented
    CustomerName *string `json:"customerName,omitempty"`
    // CustomerNumber undocumented
    CustomerNumber *string `json:"customerNumber,omitempty"`
    // CustomerPurchaseOrderReference undocumented
    CustomerPurchaseOrderReference *string `json:"customerPurchaseOrderReference,omitempty"`
    // DiscountAmount undocumented
    DiscountAmount *int `json:"discountAmount,omitempty"`
    // DiscountAppliedBeforeTax undocumented
    DiscountAppliedBeforeTax *bool `json:"discountAppliedBeforeTax,omitempty"`
    // DueDate undocumented
    DueDate *Date `json:"dueDate,omitempty"`
    // Email undocumented
    Email *string `json:"email,omitempty"`
    // ExternalDocumentNumber undocumented
    ExternalDocumentNumber *string `json:"externalDocumentNumber,omitempty"`
    // InvoiceDate undocumented
    InvoiceDate *Date `json:"invoiceDate,omitempty"`
    // LastModifiedDateTime undocumented
    LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
    // Number undocumented
    Number *string `json:"number,omitempty"`
    // OrderID undocumented
    OrderID *UUID `json:"orderId,omitempty"`
    // OrderNumber undocumented
    OrderNumber *string `json:"orderNumber,omitempty"`
    // PaymentTermsID undocumented
    PaymentTermsID *UUID `json:"paymentTermsId,omitempty"`
    // PhoneNumber undocumented
    PhoneNumber *string `json:"phoneNumber,omitempty"`
    // PricesIncludeTax undocumented
    PricesIncludeTax *bool `json:"pricesIncludeTax,omitempty"`
    // Salesperson undocumented
    Salesperson *string `json:"salesperson,omitempty"`
    // SellingPostalAddress undocumented
    SellingPostalAddress *PostalAddressType `json:"sellingPostalAddress,omitempty"`
    // ShipmentMethodID undocumented
    ShipmentMethodID *UUID `json:"shipmentMethodId,omitempty"`
    // ShippingPostalAddress undocumented
    ShippingPostalAddress *PostalAddressType `json:"shippingPostalAddress,omitempty"`
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
    // Currency undocumented
    Currency *Currency `json:"currency,omitempty"`
    // Customer undocumented
    Customer *Customer `json:"customer,omitempty"`
    // PaymentTerm undocumented
    PaymentTerm *PaymentTerm `json:"paymentTerm,omitempty"`
    // SalesInvoiceLines undocumented
    SalesInvoiceLines []SalesInvoiceLine `json:"salesInvoiceLines,omitempty"`
    // ShipmentMethod undocumented
    ShipmentMethod *ShipmentMethod `json:"shipmentMethod,omitempty"`
}

// SalesInvoiceLine undocumented
type SalesInvoiceLine struct {
    // Entity is the base model of SalesInvoiceLine
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
    // ShipmentDate undocumented
    ShipmentDate *Date `json:"shipmentDate,omitempty"`
    // TaxCode undocumented
    TaxCode *string `json:"taxCode,omitempty"`
    // TaxPercent undocumented
    TaxPercent *int `json:"taxPercent,omitempty"`
    // TotalTaxAmount undocumented
    TotalTaxAmount *int `json:"totalTaxAmount,omitempty"`
    // UnitOfMeasureID undocumented
    UnitOfMeasureID *UUID `json:"unitOfMeasureId,omitempty"`
    // UnitPrice undocumented
    UnitPrice *int `json:"unitPrice,omitempty"`
    // Account undocumented
    Account *Account `json:"account,omitempty"`
    // Item undocumented
    Item *Item `json:"item,omitempty"`
}

// SalesOrder undocumented
type SalesOrder struct {
    // Entity is the base model of SalesOrder
    Entity
    // BillingPostalAddress undocumented
    BillingPostalAddress *PostalAddressType `json:"billingPostalAddress,omitempty"`
    // BillToCustomerID undocumented
    BillToCustomerID *UUID `json:"billToCustomerId,omitempty"`
    // BillToCustomerNumber undocumented
    BillToCustomerNumber *string `json:"billToCustomerNumber,omitempty"`
    // BillToName undocumented
    BillToName *string `json:"billToName,omitempty"`
    // CurrencyCode undocumented
    CurrencyCode *string `json:"currencyCode,omitempty"`
    // CurrencyID undocumented
    CurrencyID *UUID `json:"currencyId,omitempty"`
    // CustomerID undocumented
    CustomerID *UUID `json:"customerId,omitempty"`
    // CustomerName undocumented
    CustomerName *string `json:"customerName,omitempty"`
    // CustomerNumber undocumented
    CustomerNumber *string `json:"customerNumber,omitempty"`
    // DiscountAmount undocumented
    DiscountAmount *int `json:"discountAmount,omitempty"`
    // DiscountAppliedBeforeTax undocumented
    DiscountAppliedBeforeTax *bool `json:"discountAppliedBeforeTax,omitempty"`
    // Email undocumented
    Email *string `json:"email,omitempty"`
    // ExternalDocumentNumber undocumented
    ExternalDocumentNumber *string `json:"externalDocumentNumber,omitempty"`
    // FullyShipped undocumented
    FullyShipped *bool `json:"fullyShipped,omitempty"`
    // LastModifiedDateTime undocumented
    LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
    // Number undocumented
    Number *string `json:"number,omitempty"`
    // OrderDate undocumented
    OrderDate *Date `json:"orderDate,omitempty"`
    // PartialShipping undocumented
    PartialShipping *bool `json:"partialShipping,omitempty"`
    // PaymentTermsID undocumented
    PaymentTermsID *UUID `json:"paymentTermsId,omitempty"`
    // PhoneNumber undocumented
    PhoneNumber *string `json:"phoneNumber,omitempty"`
    // PricesIncludeTax undocumented
    PricesIncludeTax *bool `json:"pricesIncludeTax,omitempty"`
    // RequestedDeliveryDate undocumented
    RequestedDeliveryDate *Date `json:"requestedDeliveryDate,omitempty"`
    // Salesperson undocumented
    Salesperson *string `json:"salesperson,omitempty"`
    // SellingPostalAddress undocumented
    SellingPostalAddress *PostalAddressType `json:"sellingPostalAddress,omitempty"`
    // ShippingPostalAddress undocumented
    ShippingPostalAddress *PostalAddressType `json:"shippingPostalAddress,omitempty"`
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
    // Currency undocumented
    Currency *Currency `json:"currency,omitempty"`
    // Customer undocumented
    Customer *Customer `json:"customer,omitempty"`
    // PaymentTerm undocumented
    PaymentTerm *PaymentTerm `json:"paymentTerm,omitempty"`
    // SalesOrderLines undocumented
    SalesOrderLines []SalesOrderLine `json:"salesOrderLines,omitempty"`
}

// SalesOrderLine undocumented
type SalesOrderLine struct {
    // Entity is the base model of SalesOrderLine
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
    // InvoiceDiscountAllocation undocumented
    InvoiceDiscountAllocation *int `json:"invoiceDiscountAllocation,omitempty"`
    // InvoicedQuantity undocumented
    InvoicedQuantity *int `json:"invoicedQuantity,omitempty"`
    // InvoiceQuantity undocumented
    InvoiceQuantity *int `json:"invoiceQuantity,omitempty"`
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
    // ShipmentDate undocumented
    ShipmentDate *Date `json:"shipmentDate,omitempty"`
    // ShippedQuantity undocumented
    ShippedQuantity *int `json:"shippedQuantity,omitempty"`
    // ShipQuantity undocumented
    ShipQuantity *int `json:"shipQuantity,omitempty"`
    // TaxCode undocumented
    TaxCode *string `json:"taxCode,omitempty"`
    // TaxPercent undocumented
    TaxPercent *int `json:"taxPercent,omitempty"`
    // TotalTaxAmount undocumented
    TotalTaxAmount *int `json:"totalTaxAmount,omitempty"`
    // UnitOfMeasureID undocumented
    UnitOfMeasureID *UUID `json:"unitOfMeasureId,omitempty"`
    // UnitPrice undocumented
    UnitPrice *int `json:"unitPrice,omitempty"`
    // Account undocumented
    Account *Account `json:"account,omitempty"`
    // Item undocumented
    Item *Item `json:"item,omitempty"`
}

// SalesQuote undocumented
type SalesQuote struct {
    // Entity is the base model of SalesQuote
    Entity
    // AcceptedDate undocumented
    AcceptedDate *Date `json:"acceptedDate,omitempty"`
    // BillingPostalAddress undocumented
    BillingPostalAddress *PostalAddressType `json:"billingPostalAddress,omitempty"`
    // BillToCustomerID undocumented
    BillToCustomerID *UUID `json:"billToCustomerId,omitempty"`
    // BillToCustomerNumber undocumented
    BillToCustomerNumber *string `json:"billToCustomerNumber,omitempty"`
    // BillToName undocumented
    BillToName *string `json:"billToName,omitempty"`
    // CurrencyCode undocumented
    CurrencyCode *string `json:"currencyCode,omitempty"`
    // CurrencyID undocumented
    CurrencyID *UUID `json:"currencyId,omitempty"`
    // CustomerID undocumented
    CustomerID *UUID `json:"customerId,omitempty"`
    // CustomerName undocumented
    CustomerName *string `json:"customerName,omitempty"`
    // CustomerNumber undocumented
    CustomerNumber *string `json:"customerNumber,omitempty"`
    // DiscountAmount undocumented
    DiscountAmount *int `json:"discountAmount,omitempty"`
    // DocumentDate undocumented
    DocumentDate *Date `json:"documentDate,omitempty"`
    // DueDate undocumented
    DueDate *Date `json:"dueDate,omitempty"`
    // Email undocumented
    Email *string `json:"email,omitempty"`
    // ExternalDocumentNumber undocumented
    ExternalDocumentNumber *string `json:"externalDocumentNumber,omitempty"`
    // LastModifiedDateTime undocumented
    LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
    // Number undocumented
    Number *string `json:"number,omitempty"`
    // PaymentTermsID undocumented
    PaymentTermsID *UUID `json:"paymentTermsId,omitempty"`
    // PhoneNumber undocumented
    PhoneNumber *string `json:"phoneNumber,omitempty"`
    // Salesperson undocumented
    Salesperson *string `json:"salesperson,omitempty"`
    // SellingPostalAddress undocumented
    SellingPostalAddress *PostalAddressType `json:"sellingPostalAddress,omitempty"`
    // SentDate undocumented
    SentDate *time.Time `json:"sentDate,omitempty"`
    // ShipmentMethodID undocumented
    ShipmentMethodID *UUID `json:"shipmentMethodId,omitempty"`
    // ShippingPostalAddress undocumented
    ShippingPostalAddress *PostalAddressType `json:"shippingPostalAddress,omitempty"`
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
    // ValidUntilDate undocumented
    ValidUntilDate *Date `json:"validUntilDate,omitempty"`
    // Currency undocumented
    Currency *Currency `json:"currency,omitempty"`
    // Customer undocumented
    Customer *Customer `json:"customer,omitempty"`
    // PaymentTerm undocumented
    PaymentTerm *PaymentTerm `json:"paymentTerm,omitempty"`
    // SalesQuoteLines undocumented
    SalesQuoteLines []SalesQuoteLine `json:"salesQuoteLines,omitempty"`
    // ShipmentMethod undocumented
    ShipmentMethod *ShipmentMethod `json:"shipmentMethod,omitempty"`
}

// SalesQuoteLine undocumented
type SalesQuoteLine struct {
    // Entity is the base model of SalesQuoteLine
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
    // UnitOfMeasureID undocumented
    UnitOfMeasureID *UUID `json:"unitOfMeasureId,omitempty"`
    // UnitPrice undocumented
    UnitPrice *int `json:"unitPrice,omitempty"`
    // Account undocumented
    Account *Account `json:"account,omitempty"`
    // Item undocumented
    Item *Item `json:"item,omitempty"`
}
