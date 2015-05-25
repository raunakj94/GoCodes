package models

type  Steel struct {  
    ID      int
    Name    string
    Type    string
    Industry string
}


func (c Steel) Show(SteelID int) revel.Result {  
    var res models.Steel

    for _, Steel := range steels {
        if Steel.ID == SteelID {
            res = Steel
        }
    }

    if res.ID == 0 {
        return c.NotFound("Could not find Steel")
    }

    return c.RenderJson(res)
}
var Steel = []models.Beer{  
    models.Steel{1, "HArd Steel", "Ale", "Bierbrouwerij De Koningshoeven"},
    models.Steel{2, "Black Steel", "SOft steel", "TATA"},
    models.Steel{3, "tyt", "Stainless", "Sail"},
}

func (c Steel) List() revel.Result {  
    return c.RenderJson(Steel)
}


