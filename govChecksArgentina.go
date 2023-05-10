package metamap

type (

	// ArgentinaDNI
	ArgentinaDNIRequest struct {
		// Document number from either a National ID or a driver license.
		DocumentNumber string `json:"documentNumber"`
		// "2023-05-10"
		DateOfBirth string `json:"dateOfBirth"`
		// M for male, F for female.
		Gender string `json:"gender"`
		// Document issue date.
		DateOfIssue string `json:"dateOfIssue"`
		CallbackUrl string `json:"callbackUrl"`
		// Use Metadata to add internal references for your outputs/webhooks.
		Metadata map[string]interface{} `json:"metadata"`
	}

	ArgentinaDNIResponse struct {
		Error interface{} `json:"error"`
		Data  struct {
			DateOfIssue    string `json:"dateOfIssue"`
			DocumentNumber string `json:"documentNumber"`
		} `json:"data"`
	}

	// ArgentinaRENAPER
	ArgentinaRENAPERRequest struct {
		// Document number from either a National ID or a driver license.
		DocumentNumber string `json:"documentNumber"`
		// "2023-05-10"
		DateOfBirth string `json:"dateOfBirth"`
		// M for male, F for female.
		Gender string `json:"gender"`
		// Document issue date.
		DateOfIssue string `json:"dateOfIssue"`
		CallbackUrl string `json:"callbackUrl"`
		// Use Metadata to add internal references for your outputs/webhooks.
		Metadata map[string]interface{} `json:"metadata"`
	}

	ArgentinaRENAPERResponse struct {
		Error interface{} `json:"error"`
		Data  struct {
			DateOfBirth  string      `json:"dateOfBirth"`
			FullName     string      `json:"fullName"`
			DniNumber    string      `json:"dniNumber"`
			Cuit         string      `json:"cuit"`
			PhoneNumbers []string    `json:"phoneNumbers"`
			Deceased     interface{} `json:"deceased"`
		} `json:"data"`
	}

	// ArgentinaRENAPERPremium
	ArgentinaRENAPERPremiumRequest struct {
		// Document number from either a National ID or a driver license.
		DocumentNumber string `json:"documentNumber"`
		// E.g "2023-05-10"
		DateOfBirth string `json:"dateOfBirth"`
		// Date of document issue E.g "2023-05-10"
		IssueDate string `json:"issueDate"`
		// M for male, F for female.
		Gender      string `json:"gender"`
		CallbackUrl string `json:"callbackUrl"`
		// Use Metadata to add internal references for your outputs/webhooks.
		Metadata map[string]interface{} `json:"metadata"`
	}

	ArgentinaRENAPERPremiumResponse struct {
		Error interface{} `json:"error"`
		Data  struct {
			IdMainIssue          string `json:"idMainIssue"`
			IdReprint            string `json:"idReprint"`
			CopyCode             string `json:"copyCode"`
			IssueDate            string `json:"issueDate"`
			ExpiryDate           string `json:"expiryDate"`
			LastName             string `json:"lastName"`
			FirstName            string `json:"firstName"`
			DateOfBirth          string `json:"dateOfBirth"`
			StreetName           string `json:"streetName"`
			StreetNumber         string `json:"streetNumber"`
			Floor                string `json:"floor"`
			Apartment            string `json:"apartment"`
			ZipCode              string `json:"zipCode"`
			Neighborhood         string `json:"neighborhood"`
			Monobloc             string `json:"monobloc"`
			City                 string `json:"city"`
			Town                 string `json:"town"`
			State                string `json:"state"`
			Country              string `json:"country"`
			DeathNotice          string `json:"deathNotice"`
			DateOfDeath          string `json:"dateOfDeath"`
			CitizenId            string `json:"citizenId"`
			Nationality          string `json:"nationality"`
			PlaceOfBirth         string `json:"placeOfBirth"`
			Taxid                string `json:"taxid"`
			Gender               string `json:"gender"`
			NationalId           string `json:"nationalId"`
			NationalIdValidation string `json:"nationalIdValidation"`
			IssueDateValidation  string `json:"issueDateValidation"`
			GenderValidation     string `json:"genderValidation"`
		} `json:"data"`
	}
)

/*
ArgentinaDNI verify a user's DNI (Documento Nacional de Identidad) based on card number and issue date.

This method takes in the ArgentinaDNIRequest{} struct as a parameter.

MetaMap connects with the Argentinian National Direction of Migration (Dirección Nacional de Migraciones) site to validate that the
National Identity Document (Documento Nacional de Identidad / DNI) number present in the ID card exists and its date of issue matches the last one issued to the user.
*/
func (c *Client) ArgentinaDNI(req ArgentinaDNIRequest) (*ArgentinaDNIResponse, error) {
	url := "govchecks/v1/ar/dni"
	method := MethodPOST
	var response ArgentinaDNIResponse
	c.IsBasic = false
	if err := c.newRequest(method, url, req, response); err != nil {
		return &ArgentinaDNIResponse{}, err
	}

	return &response, nil
}

/*
ArgentinaRENAPER verify a user's DNI (Documento Nacional de Identidad) number and identity.

NOTE: Use the ArgentinaDNI() to validate the submitted DNI card.

This method takes in the ArgentinaRENAPERRequest{} struct as a parameter.

MetaMap connects with the Argentinian National Registry of Persons (Registro Nacional de la Persona / RENAPER) to validate that the
National Identity Document (Documento Nacional de Identidad / DNI) number present in the ID card exists and its owner matches the data obtained from it.
*/
func (c *Client) ArgentinaRENAPER(req ArgentinaRENAPERRequest) (*ArgentinaRENAPERResponse, error) {
	url := "govchecks/v1/ar/renaper"
	method := MethodPOST
	var response ArgentinaRENAPERResponse
	c.IsBasic = false
	if err := c.newRequest(method, url, req, response); err != nil {
		return &ArgentinaRENAPERResponse{}, err
	}

	return &response, nil
}

/*
ArgentinaRENAPERPremium verify a user's DNI (Documento Nacional de Identidad) number and identity.

This method takes in the ArgentinaRENAPERPremiumRequest{} struct as a parameter.

NOTE: You do not need to validate the submitted DNI card parameters via ArgentinaDNI() as this check is implemented in here.

MetaMap connects to the Argentinian National Registry of Persons (Registro Nacional de la Persona / RENAPER) to validate that
the DNI (Documento Nacional de Identidad / National Identity Document) exists in government DB a date of issue and copy of the card match the last issued to the user.
This version complements the response with additional data points compared to the basic Renaper DB.
*/
func (c *Client) ArgentinaRENAPERPremium(req ArgentinaRENAPERPremiumRequest) (*ArgentinaRENAPERPremiumResponse, error) {
	url := "govchecks/v1/ar/renaper-premium"
	method := MethodPOST
	var response ArgentinaRENAPERPremiumResponse
	c.IsBasic = false
	if err := c.newRequest(method, url, req, response); err != nil {
		return &ArgentinaRENAPERPremiumResponse{}, err
	}

	return &response, nil
}
