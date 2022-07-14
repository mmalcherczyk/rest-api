package models

type User struct {
        ID          string    'json:"id"'
        Name        string    'json:"name"'
        Lastname    string    'json:"lastname"'
        PhoneNumber string    'json:"phonenumber"'
        PIN         int       'json:"pin"'

}
