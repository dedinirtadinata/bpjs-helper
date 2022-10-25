package models

type TicketModel struct {
	Status       string `json:"status"`
	Message      string `json:"message"`
	AntrianNomor string `json:"antrian_nomor"`
	ServiceName  string `json:"service_name"`
	LeftQueue    int    `json:"left_queue"`
}
