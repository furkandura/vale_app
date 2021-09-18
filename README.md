# Vale Uygulaması (API)

Vale firmalarını düşünerek kodlanmış apiler bulunmaktadır. Örneğin; müşteri ekleme, müşteriye araç plakası tanımlama, park kayıtları oluşturma gibi apiler bulunmaktadır. Dileyen kişiler front kısmını veya mobil uygulamasını yazarak pratik yapabilir.


# Ufak Bir Dökümantasyon

####  Firma Kayıt
	// Endpoint => api/company/register
	// Method => POST
	
    type  CompanyRegisterRequest  struct {
    
	    FullName string  `json:"full_name" validate:"required"`
    
	    CompanyName string  `json:"company_name" validate:"required"`
    
	    Phone string  `json:"phone" validate:"required"`
    
	    Password string  `json:"password" validate:"required,min=6,max=16"`
    
	    Email *string  `json:"email" validate:"omitempty,email"`
    
    }
     
    
   #### Firma Giriş

    // Endpoint => api/company/login
    // Method => POST
    
    type  CompanyLoginRequest  struct {
    
	    Phone string  `json:"phone" validate:"required"`
    
	    Password string  `json:"password" validate:"required"`
    
    }

#### Firma Bilgilerini Güncelle

    // Endpoint => api/company/update
    
    // Method => POST
    
    type  CompanyUpdateRequest  struct {
    
	    FullName string  `json:"full_name" validate:"required"`
    
	    Phone string  `json:"phone" validate:"required"`
    
	    Password string  `json:"password" validate:"required,min=6,max=16"`
    
	    Email *string  `json:"email" validate:"omitempty,email"`
    
    }

#### Park Kayıtlarını Listele

    // Endpoint => api/parking/all
    // Method => GET

#### Park Kaydı Oluştur

    // Endpoint => api/parking/create
    
    // Method => POST
    
    type  ParkingCreateRequest  struct {
    
	    CustomerId int  `json:"customer_id" validate:"required"`
    
	    Type int8  `json:"type" validate:"required"`
    
	    Plate string  `json:"plate" validate:"required"`
    
	    DateOfReceipt time.Time `json:"date_of_receipt" validate:"required"`
    
	    Amount *float64  `json:"amount"`
    
	    Note *string  `json:"note"`
    
    }
#### Park Kaydını Güncelle

    // Endpoint => api/parking/update
    
    // Method => POST

    type  ParkingUpdateRequest  struct {
    
	    ParkingId int  `json:"parking_id" validate:"required"`
    
	    CustomerId int  `json:"customer_id" validate:"required"`
    
	    Type int8  `json:"type" validate:"required"`
    
	    Plate string  `json:"plate" validate:"required"`
    
	    DateOfReceipt time.Time `json:"date_of_receipt" validate:"required"`
    
	    DateOfDelivery *time.Time `json:"date_of_delivery"`
    
	    Amount *float64  `json:"amount"`
    
	    Note *string  `json:"note"`
    
    }
#### Park Kaydını Sil

    // Endpoint => api/parking/delete/{id}
    // Method => GET
#### Müşterileri Listele

    // Endpoint => api/customer/all
    // Method => POST
#### Müşteri Oluştur

    // Endpoint => api/customer/create
    
    // Method => POST
    
 
    type  CustomerCreateRequest  struct {
    
	    FullName string  `json:"full_name" validate:"required"`
    
	    Phone *string  `json:"phone"`
    
	    Vehicles string  `json:"vehicles" validate:"required"`
    
	    Note *string  `json:"note"`
    
    }
#### Müşteri Güncelle

    // Endpoint => api/customer/update
    
    // Method => POST

    type  CustomerUpdateRequest  struct {
    
	    CustomerId int  `json:"customer_id" validate:"required"`
    
	    FullName string  `json:"full_name" validate:"required"`
    
	    Phone *string  `json:"phone"`
    
	    Vehicles string  `json:"vehicles" validate:"required"`
    
	    Note *string  `json:"note"`
    
    }
#### Müşteri Sil

    // Endpoint => api/customer/delete/{id}
    // Method => GET
#### Müşterinin Araç Plakalarını Göster

    // Endpoint => api/customer/get/vehicles/{id}
    // Method => GET



