package handler

import(
	"github.com/labstack/echo/v4"
	"github.com/google/uuid"

	"10xers/domain"
)

type(
	// this is where phone domain initiated
	// and can be used anywhere
	PhoneHandler struct{
		Phone domain.PhoneInterface
	}
)

func NewPhoneHandler(phoneDomain domain.PhoneInterface)*PhoneHandler{
	if phoneDomain == nil{
		panic("phone domain ")
	}

	return &PhoneHandler{
		Phone: phoneDomain,
	}
}


func (ph *PhoneHandler) CreatePhone(c echo.Context) error {
    var phone domain.Phone
    if err := c.Bind(&phone); err != nil {
        return err
    }
	phone.ID = uuid.New()
    if err := ph.Phone.CreatePhone(&phone) ; err != nil{
        return err
    }
    return c.JSON(200,"success creating new phone")
}

func (ph *PhoneHandler) SearchPhone(c echo.Context) error {
	name := c.QueryParam("name")
    phone, err := ph.Phone.SearchPhonesByName(name)
    if err != nil {
        return err
    }
    return c.JSON(200, phone)
}

func (ph *PhoneHandler) UpdatePhone(c echo.Context) error {
    id := c.Param("id")
    uuid,err := uuid.Parse(id)
	if err != nil {
		return err
	}
    // Check if the phone exists
    foundPhone, err := ph.Phone.GetPhoneByID(uuid)
    if err != nil {
        return echo.NewHTTPError(404, "Phone not found")
    }
    
    var phone domain.Phone
    if err := c.Bind(&phone); err != nil {
        return err
    }

    // handle case if one of phone property is not proper inserted
    if phone.Brand == ""{
        phone.Brand = foundPhone.Brand
    }
    if phone.Name == ""{
        phone.Name = foundPhone.Name
    }
    if phone.Price == 0{
        phone.Price = foundPhone.Price
    }
    if phone.StockQuantity == 0 {
        phone.StockQuantity = foundPhone.StockQuantity
    }

	phone.ID = uuid
    if err := ph.Phone.UpdatePhone(&phone) ; err != nil{
        return err
    }
    return c.JSON(200,"success update phone with id: "+id)
}

func (ph *PhoneHandler) DeletePhoneByID(c echo.Context) error {
    id := c.Param("id")
	uuid,err := uuid.Parse(id)
	if err != nil {
		return err
	}
    // Check if the phone exists
    _, err = ph.Phone.GetPhoneByID(uuid)
    if err != nil {
        return echo.NewHTTPError(404, "Phone not found")
    }
    
    // Check if the phone exists
    _, err = ph.Phone.GetPhoneByID(uuid)
    if err != nil {
        return echo.NewHTTPError(404, "Phone not found")
    }

    if err := ph.Phone.DeletePhoneByID(uuid) ; err != nil{
        return err
    }
    
    return c.JSON(200,"success delete phone with id: "+id)
}