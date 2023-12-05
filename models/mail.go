package models

type Mail struct {
	Message_ID                string   `json:"message_id"`
	Date                      string   `json:"date"`
	From                      string   `json:"from"`
	To                        []string `json:"to"`
	Subject                   string   `json:"subject"`
	Cc                        []string `json:"cc"`
	Mime_version              string   `json:"mime_version"`
	Content_Type              string   `json:"content_type"`
	Content_Transfer_Encoding string   `json:"content_transfer_encoding"`
	Bcc                       []string `json:"bcc"`
	X_from                    string   `json:"x_from"`
	X_to                      []string `json:"x_to"`
	X_cc                      []string `json:"x_cc"`
	X_bcc                     []string `json:"x_bcc"`
	X_folder                  string   `json:"x_folder"`
	X_origin                  string   `json:"x_origin"`
	X_filename                string   `json:"x_filename"`
	Content                   string   `json:"content"`
}
