package wxxx

import (
	"wxxx/mm_pb"
)

var Log = &logger{}

const shortUrl = "http://szshort.weixin.qq.com"

type client struct {
	shortLink       HttpCli
	clientVersion   int32
	loginRsaVersion uint32
	osType          string
}

func newClient(shortLink HttpCli) *client {
	cli := &client{
		shortLink:       shortLink,
		clientVersion:   0x17000523,
		loginRsaVersion: 174,
		osType:          "iPad iPhone OS9.3.3",
	}
	return cli
}

func (this *client) GetLoginQrCode(deviceId string) (*GetLoginQrCodeResponse, error) {
	var deviceIdBytes []byte
	if deviceId == "" {
		deviceId = this.newRandomDeviceId()
		deviceIdBytes = Md5Bytes(deviceId)
		deviceIdBytes[0] = 0x49
	} else {
		deviceIdBytes = []byte(deviceId)
	}
	//generate a random aes key
	//rsaLen := uint32(2048)
	aesKey := this.newRandomAesKey()
	req := &mm_pb.GetLoginQRCodeRequest{
		BaseRequest:     this.newBasRequest(aesKey.Key, deviceIdBytes, 0, 0),
		Aes:             aesKey,
		Opcode:          new(uint32),
		DeviceName:      nil,
		UserName:        nil,
		ExtDevLoginType: new(uint32),
		HardwareExtra:   nil,
		SoftType:        nil,
		//Rsa: &mm_pb.RSAPem{
		//	Len: rsaLen,
		//	Pem: kRsaPubKeyModules,
		//},
	}
	resp := &mm_pb.GetLoginQRCodeResponse{}
	pkg, err := this.shortLinkExecute(shortUrl,
		req,
		mm_pb.CGI_TYPE_GETLOGINQRCODE, mm_pb.CGI_URL_GETLOGINQRCODE,
		7, aesKey.Key, nil, 0, resp)
	if err != nil {
		return nil, err
	}
	//fmt.Println(resp)
	return &GetLoginQrCodeResponse{
		QRCodeId:      *resp.Uuid,
		Uin:           pkg.uin,
		AesKey:        resp.AESKey.Key,
		Cookies:       pkg.cookie,
		QRCodeContent: resp.QRCode.Src,
	}, nil
}

func (this *client) CheckLoginQrCode(qrCodeId string, aesKey []byte) (*CheckLoginQrCodeResponse, error) {
	panic("implement me")
}
