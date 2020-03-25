package main
import "fmt"

const (
	GroupMaxNumber     = 1024
	GroupWLANMaxNumber = 4
	TimerMaxNumber     = 3
)
type GroupInfo struct {
	Next *GroupNode
	Id uint16
}
type MacNode struct {
	Next *MacNode
	Addr string
}
type BWNode struct {
	Next *BWNode
	Group [GroupMaxNumber]GroupInfo
	Name string /* the name is only */
	List *MacNode /* mac maybe less */
}
type WLANNode struct {
	Next *WLANNode
	Group [GroupMaxNumber]GroupInfo
	SSID string
	Passwd string
}
type DateNode struct {
	Start string
	Stop string
	Week string
}
type TimerNode struct {
	Next *TimerNode
	Group [GroupMaxNumber]GroupInfo
	Name string
	Timer [TimerMaxNumber]DateNode
}
type APNode struct {
	Next *APNode
	Mac string
	IPAddr string
}
type GroupNode struct {
	Next *GroupNode
	Id uint16 /* less then 65535 groups */
	Name string /* group name */
	ApList *APNode
	WLANList [GroupWLANMaxNumber] *WLANNode
	WLANCount uint8
	BWFlag uint8 /*  0, disable 1, black 2, white */
	BWList *BWNode
	LEDTimerFlag uint8 /* 0, disable 1,enable */
	Timer *TimerNode
}

var GroupList GroupNode /* the first node */
//var WLANList *WLANNode
//var TimerList *TimerNode
//var BWList *BWNode
func InitGroupList(head *GroupNode) {
	head.Name = "default"
}
func AddGroupNode(head *GroupNode, id uint16, name string) {
	var temp *GroupNode
	temp = new(GroupNode)
	temp.Id = id
	temp.Name = name
	temp.Next = head.Next
	head.Next = temp
	fmt.Printf("%d %s add\n", temp.Id, temp.Name)
}
func  SearchGroupNodeById(head *GroupNode, id uint16) (temp *GroupNode, pre *GroupNode) {
	var prev  *GroupNode
	prev = nil
	for temp := head; temp != nil; temp = temp.Next {
		if temp.Id == id {
			fmt.Printf("You have found the id %d group\n", id)
			return temp, prev
		}
		prev = temp
	}
	return nil, nil
}
func DelGroupNode(head *GroupNode, id uint16) int{
	temp, prev := SearchGroupNodeById(head, id)
	if nil == temp {
		fmt.Printf("The id %d group is not found\n", id)
		return -1
	}
	if nil == prev {
		head = head.Next
	} else {
		prev.Next = temp.Next
	}
	fmt.Printf("The id %d group del\n", id)
	return 0
}
func main() {
	AddGroupNode(&GroupList, 1, "default")
	AddGroupNode(&GroupList, 2, "hello")
	AddGroupNode(&GroupList, 3, "world")
	SearchGroupNodeById(&GroupList, 1)
	DelGroupNode(&GroupList, 1)
	SearchGroupNodeById(&GroupList, 0)
	fmt.Printf("hanfushun hello world\n")
}