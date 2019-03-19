package main

type Person struct {
    ID        int64   `xml:"id,omitempty"`
    Name      string   `xml:"name"`
	Password  string   `xml:"password"`
	Description string `xml:"description,omitempty"`
}

/*
type Address struct {
    City  string `json:"city,omitempty"`
    State string `json:"state,omitempty"`
}

var people []Person
*/