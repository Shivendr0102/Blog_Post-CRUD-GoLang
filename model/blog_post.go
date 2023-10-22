package model

type SearchPostRequest struct {
	SearchTerm *string `json:"searchTerm"`

	PageNumber int    `json:"pageNumber"`
	PageSize   int64  `json:"pageSize"`
	SortField  string `json:"sortField"`
	SortOrder  string `json:"sortOrder"`
}

type Address struct {
	House_No   *string `json:"house_no"`
	House_Name *string `json:"house_name"`
	Street     *string `json:"street"`
	City       *string `json:"city"`
	State      *string `json:"state"`
	Country    *string `json:"Country"`
	PinCode    *int    `json:"pincode"`
}

type Degree struct {
	SerialNo   *string `json:"serialnumber"`
	FieldName  *string `json:"fieldname"`
	StartDate  *string `json:"startdate"`
	EndDate    *string `json:"enddate"`
	Grade      *string `json:"grades"`
	College    *string `json:"college"`
	University *string `json:"univeristy"`
}

type Experience struct {
	Title            *string `json:"title"`
	Org_Name         *string `json:"org_name"`
	StartDate        *string `json:"startdate"`
	EndDate          *string `json:"enddate"`
	CurrentlyWorking bool    `json:"currentlyworking"`
}

type Author struct {
	AuthorId      *string `json:"author_id"`
	First_Name    *string `json:"firstname"`
	Last_Name     *string `json:"lastname"`
	Full_Name     *string `json:"fullname"`
	Date_Of_Birth *string `json:"dateofbirth"`
	Gender        *string `json:"gender"`
	Address
}

type BlogPost struct {
	Id           *string `json:"id"`
	Title        *string `json:"title"`
	Caption      *string `json:"caption"`
	Body         *string `json:"body"`
	Note         *string `json:"note"`
	CreatedDate  *string `json:"createddate"`
	ModifiedDate string  `json:"modifieddate"`
	Author
}

type PaginatedPostsResponse struct {
	Result              []*BlogPost
	PageNumber          int64  `json:"pageNumber"`
	PageSize            int64  `json:"pageSize"`
	TotalRecords        int64  `json:"totalRecords"`
	SelectedRecordCount int    `json:"selectedRecordCount"`
	SortField           string `json:"sortField"`
	SortOrder           string `json:"sortOrder"`
}
