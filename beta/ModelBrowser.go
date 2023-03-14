
// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph


// BrowserSharedCookie undocumented
type BrowserSharedCookie struct {
    // Entity is the base model of BrowserSharedCookie
    Entity
    // Comment undocumented
    Comment *string `json:"comment,omitempty"`
    // CreatedDateTime undocumented
    CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
    // DeletedDateTime undocumented
    DeletedDateTime *time.Time `json:"deletedDateTime,omitempty"`
    // DisplayName undocumented
    DisplayName *string `json:"displayName,omitempty"`
    // History undocumented
    History []BrowserSharedCookieHistory `json:"history,omitempty"`
    // HostOnly undocumented
    HostOnly *bool `json:"hostOnly,omitempty"`
    // HostOrDomain undocumented
    HostOrDomain *string `json:"hostOrDomain,omitempty"`
    // LastModifiedBy undocumented
    LastModifiedBy *IdentitySet `json:"lastModifiedBy,omitempty"`
    // LastModifiedDateTime undocumented
    LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
    // Path undocumented
    Path *string `json:"path,omitempty"`
    // SourceEnvironment undocumented
    SourceEnvironment *BrowserSharedCookieSourceEnvironment `json:"sourceEnvironment,omitempty"`
    // Status undocumented
    Status *BrowserSharedCookieStatus `json:"status,omitempty"`
}

// BrowserSharedCookieHistory undocumented
type BrowserSharedCookieHistory struct {
    // Object is the base model of BrowserSharedCookieHistory
    Object
    // Comment undocumented
    Comment *string `json:"comment,omitempty"`
    // DisplayName undocumented
    DisplayName *string `json:"displayName,omitempty"`
    // HostOnly undocumented
    HostOnly *bool `json:"hostOnly,omitempty"`
    // HostOrDomain undocumented
    HostOrDomain *string `json:"hostOrDomain,omitempty"`
    // LastModifiedBy undocumented
    LastModifiedBy *IdentitySet `json:"lastModifiedBy,omitempty"`
    // Path undocumented
    Path *string `json:"path,omitempty"`
    // PublishedDateTime undocumented
    PublishedDateTime *time.Time `json:"publishedDateTime,omitempty"`
    // SourceEnvironment undocumented
    SourceEnvironment *BrowserSharedCookieSourceEnvironment `json:"sourceEnvironment,omitempty"`
}

// BrowserSite undocumented
type BrowserSite struct {
    // Entity is the base model of BrowserSite
    Entity
    // AllowRedirect undocumented
    AllowRedirect *bool `json:"allowRedirect,omitempty"`
    // Comment undocumented
    Comment *string `json:"comment,omitempty"`
    // CompatibilityMode undocumented
    CompatibilityMode *BrowserSiteCompatibilityMode `json:"compatibilityMode,omitempty"`
    // CreatedDateTime undocumented
    CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
    // DeletedDateTime undocumented
    DeletedDateTime *time.Time `json:"deletedDateTime,omitempty"`
    // History undocumented
    History []BrowserSiteHistory `json:"history,omitempty"`
    // LastModifiedBy undocumented
    LastModifiedBy *IdentitySet `json:"lastModifiedBy,omitempty"`
    // LastModifiedDateTime undocumented
    LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
    // MergeType undocumented
    MergeType *BrowserSiteMergeType `json:"mergeType,omitempty"`
    // Status undocumented
    Status *BrowserSiteStatus `json:"status,omitempty"`
    // TargetEnvironment undocumented
    TargetEnvironment *BrowserSiteTargetEnvironment `json:"targetEnvironment,omitempty"`
    // WebURL undocumented
    WebURL *string `json:"webUrl,omitempty"`
}

// BrowserSiteHistory undocumented
type BrowserSiteHistory struct {
    // Object is the base model of BrowserSiteHistory
    Object
    // AllowRedirect undocumented
    AllowRedirect *bool `json:"allowRedirect,omitempty"`
    // Comment undocumented
    Comment *string `json:"comment,omitempty"`
    // CompatibilityMode undocumented
    CompatibilityMode *BrowserSiteCompatibilityMode `json:"compatibilityMode,omitempty"`
    // LastModifiedBy undocumented
    LastModifiedBy *IdentitySet `json:"lastModifiedBy,omitempty"`
    // MergeType undocumented
    MergeType *BrowserSiteMergeType `json:"mergeType,omitempty"`
    // PublishedDateTime undocumented
    PublishedDateTime *time.Time `json:"publishedDateTime,omitempty"`
    // TargetEnvironment undocumented
    TargetEnvironment *BrowserSiteTargetEnvironment `json:"targetEnvironment,omitempty"`
}

// BrowserSiteList undocumented
type BrowserSiteList struct {
    // Entity is the base model of BrowserSiteList
    Entity
    // Description undocumented
    Description *string `json:"description,omitempty"`
    // DisplayName undocumented
    DisplayName *string `json:"displayName,omitempty"`
    // LastModifiedBy undocumented
    LastModifiedBy *IdentitySet `json:"lastModifiedBy,omitempty"`
    // LastModifiedDateTime undocumented
    LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
    // PublishedBy undocumented
    PublishedBy *IdentitySet `json:"publishedBy,omitempty"`
    // PublishedDateTime undocumented
    PublishedDateTime *time.Time `json:"publishedDateTime,omitempty"`
    // Revision undocumented
    Revision *string `json:"revision,omitempty"`
    // Status undocumented
    Status *BrowserSiteListStatus `json:"status,omitempty"`
    // SharedCookies undocumented
    SharedCookies []BrowserSharedCookie `json:"sharedCookies,omitempty"`
    // Sites undocumented
    Sites []BrowserSite `json:"sites,omitempty"`
}