package data

type JSONData struct {
	TRANSMEDIA struct {
		DeskripsiPerusahaan string `json:"deskripsiPerusahaan"`
		AlamatPerusahaan    string `json:"alamatPerusahaan"`
		KONTAK              struct {
			PHONE string `json:"PHONE"`
			FAX   string `json:"FAX"`
		} `json:"KONTAK"`
		SosialMedia struct {
			Twitter   string `json:"twitter"`
			Facebook  string `json:"facebook"`
			Youtube   string `json:"youtube"`
			Instagram string `json:"instagram"`
		} `json:"sosial_media"`
	} `json:"TRANSMEDIA"`
}
