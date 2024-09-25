package model

type Resource struct {
    Type   string            `json:"type"`
    Name   string            `json:"name"`
    Config map[string]string `json:"config"`
}

