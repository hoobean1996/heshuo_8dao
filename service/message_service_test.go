package service_test

import (
	"encoding/xml"
	"testing"
	"wxcloudrun-golang/service"
)

func TestMain(t *testing.T) {
	message := &service.CommonMessage{}

	data := `<xml>
  <ToUserName><![CDATA[toUser]]></ToUserName>
  <FromUserName><![CDATA[FromUser]]></FromUserName>
  <CreateTime>123456789</CreateTime>
  <MsgType><![CDATA[event]]></MsgType>
  <Event><![CDATA[VIEW]]></Event>
  <EventKey><![CDATA[www.qq.com]]></EventKey>
</xml>`

	err := xml.Unmarshal([]byte(data), message)
	if err != nil {
		t.Error(err)
	}
}
