package processes

import (
	"encoding/json"
	"go_code/MultiusersChatRoom/common/message"
	"go_code/MultiusersChatRoom/server/utils"
	"log"
	"net"
)

type SmsProcess struct {
}

func (this *SmsProcess) SendGroupMes(mes *message.Message) (err error) {

	var smsMes message.SmsMes

	err = json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		log.Println("SmsProcess json.Unmarshal Fail err=", err.Error())
		return
	}

	data, err := json.Marshal(mes)
	if err != nil {
		log.Println("SmsProcess json.Marshal Fail err=", err.Error())
		return
	}

	for id, up := range userMgr.onlineUsers {

		if id == smsMes.UserId {
			continue
		}

		err = this.SendMesToEachOnlineUser(data, up.Conn)
	}

	return

}

func (this *SmsProcess) SendMesToEachOnlineUser(data []byte, conn net.Conn) (err error) {

	tf := &utils.Transfer{
		Conn: conn,
	}

	err = tf.WritePkg(data)
	if err != nil {
		log.Println("SendMesToEachOnlineUser writePkg Fail err=", err.Error())

	}
	return
}
