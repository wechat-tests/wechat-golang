package wxxx

func NewClient(shortLink HttpCli) Client {
	return newClient(shortLink)
}

type (
	Client interface {
		// get login qr code
		GetLoginQrCode(deviceId string) (*GetLoginQrCodeResponse, error)
		// check login qr code state
		CheckLoginQrCode(qrCodeId string, aesKey []byte) (*CheckLoginQrCodeResponse, error)
	}

	GetLoginQrCodeResponse struct {
		QRCodeId      string `json:"qrCodeId"`
		Uin           uint32 `json:"uin"`
		AesKey        []byte `json:"aesKey"`
		Cookies       []byte `json:"cookies"`
		QRCodeContent []byte `json:"qrCodeContent"`
	}
	CheckLoginQrCodeResponse struct {
	}
)
