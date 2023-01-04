package main

import (
	"fmt"
)

type Node struct {
	property int
	nextNode *Node
}

type LinkedList struct {
	headNode *Node
}

func (LinkedList *LinkedList) AddToHead(property int) {
	var node = Node{}
	node.property = property
	if node.nextNode != nil {
		node.nextNode = LinkedList.headNode
	}

	LinkedList.headNode = &node
}

func (LinkedList *LinkedList) IterateList() {
	var node *Node

	for node = LinkedList.headNode; node != nil; node = node.nextNode {
		fmt.Println(node.property)
	}
}

func (LinkedList *LinkedList) LastNode() *Node {
	var node *Node
	var lastNode *Node
	for node = LinkedList.headNode; node != nil; node = node.nextNode {
		if node.nextNode == nil {
			lastNode = node
		}
	}
	return lastNode
}

func (LinkedList *LinkedList) AddToEnd(property int) {
	var node = &Node{}
	node.property = property
	node.nextNode = nil
	var lastNode *Node
	lastNode = LinkedList.LastNode()
	if lastNode != nil {
		lastNode.nextNode = node
	}
}

func (LinkedList *LinkedList) NodeWithValue(property int) *Node {
	var node *Node
	var nodeWith *Node

	for node = LinkedList.headNode; node != nil; node = node.nextNode {
		if node.property == property {
			nodeWith = node
			break
		}
	}
	return nodeWith
}

func (LinkedList *LinkedList) AddAfter(nodeProperty, property int) {
	var node = &Node{}
	node.property = property
	node.nextNode = nil
	var nodeWith *Node
	nodeWith = LinkedList.NodeWithValue(nodeProperty)
	if nodeWith != nil {
		node.nextNode = nodeWith.nextNode
		nodeWith.nextNode = node
	}
}

func main() {

	var linkedList LinkedList
	linkedList = LinkedList{}
	linkedList.AddToHead(1)
	linkedList.AddToHead(3)
	linkedList.AddToEnd(4)
	linkedList.AddToEnd(5)
	linkedList.AddAfter(3, 7)
	// fmt.Println(linkedList.headNode.property)
	linkedList.IterateList()

}
