/*
 * @Author: ph4ntom
 * @Date: 2021-03-08 18:19:04
 * @LastEditors: ph4ntom
 * @LastEditTime: 2021-03-20 13:43:53
 */
package protocol

import (
	"Stowaway/crypto"
	"net"
)

const (
	HI = iota
	UUID
	UUIDRET
	MYINFO
	MYMEMO
	SHELLREQ
	SHELLRES
	SHELLCOMMAND
	SHELLRESULT
	LISTENREQ
	LISTENRES
	SSHREQ
	SSHRES
	SSHCOMMAND
	SSHRESULT
)

const ADMIN_UUID = "IAMADMINXD"
const TEMP_UUID = "IAMNEWHERE"
const TEMP_ROUTE = "THEREISNOROUTE"

type Message interface {
	ConstructHeader()
	ConstructData(Header, interface{})
	ConstructSuffix()
	DeconstructHeader()
	DeconstructData() (Header, interface{}, error)
	DeconstructSuffix()
	SendMessage()
}

/**
 * @description:
 * @param {Message} message
 * @param {Header} header
 * @param {interface{}} mess
 * @return {*}
 */
func ConstructMessage(message Message, header Header, mess interface{}) {
	message.ConstructData(header, mess)
	message.ConstructHeader()
	message.ConstructSuffix()
}

/**
 * @description: See function name
 * @param {Message} message
 * @return {*}
 */
func DestructMessage(message Message) (Header, interface{}, error) {
	message.DeconstructHeader()
	header, mess, err := message.DeconstructData()
	message.DeconstructSuffix()
	return header, mess, err
}

type Header struct {
	Sender      string // sender and accepter are both 10bytes
	Accepter    string
	MessageType uint16
	RouteLen    uint32
	Route       string
	DataLen     uint64
}

type HIMess struct {
	GreetingLen uint16
	Greeting    string
	IsAdmin     uint16
}

type UUIDMess struct {
	UUIDLen uint16
	UUID    string
}

type UUIDRetMess struct {
	OK uint16
}

type MyInfo struct {
	UsernameLen uint64
	Username    string
	HostnameLen uint64
	Hostname    string
}

type MyMemo struct {
	MemoLen uint64
	Memo    string
}

type ShellReq struct {
	Start uint16
}

type ShellRes struct {
	OK uint16
}

type ShellCommand struct {
	CommandLen uint64
	Command    string
}

type ShellResult struct {
	ResultLen uint64
	Result    string
}

type ListenReq struct {
	AddrLen uint64
	Addr    string
}

type ListenRes struct {
	OK uint16
}

type SSHReq struct {
	Method         uint16
	AddrLen        uint64
	Addr           string
	UsernameLen    uint64
	Username       string
	PasswordLen    uint64
	Password       string
	CertificateLen uint64
	Certificate    []byte
}

type SSHRes struct {
	OK uint16
}

type SSHCommand struct {
	CommandLen uint64
	Command    string
}

type SSHResult struct {
	ResultLen uint64
	Result    string
}

/**
 * @description: The struct containing  essential components to use "PrepareAndDecideWhichSProto" or "PrepareAndDecideWhichRProto"
 * @param {*}
 * @return {*}
 */
type MessageComponent struct {
	ID     string
	Conn   net.Conn
	Secret string
}

/**
 * @description: Decide which transmission protocol you want to use for sending message
 * @param {net.Conn} conn
 * @return {*}
 */
func PrepareAndDecideWhichSProto(conn net.Conn, secret string, uuid string) Message {
	// Now only apply tcp raw
	// TODO: HTTP
	tMessage := new(TCPMessage)
	tMessage.Conn = conn
	tMessage.ID = uuid
	tMessage.CryptoSecret, _ = crypto.KeyPadding([]byte(secret))
	return tMessage
}

/**
 * @description: Decide which transmission protocol you want to use for receving message
 * @param {net.Conn} conn
 * @return {*}
 */
func PrepareAndDecideWhichRProto(conn net.Conn, secret string, uuid string) Message {
	// Now only apply tcp raw
	// TODO: HTTP
	tMessage := new(TCPMessage)
	tMessage.Conn = conn
	tMessage.ID = uuid
	tMessage.CryptoSecret, _ = crypto.KeyPadding([]byte(secret))
	return tMessage
}