package main

import "fmt"

type Node struct {
	property     int
	nextNode     *Node
	previousNode *Node
}

type LinkedList struct {
	headNode *Node
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

func (LinkedList *LinkedList) NodeBetweenValues(firstProperty, secondProperty int) *Node {
	var node *Node
	var nodeWith *Node

	for node = LinkedList.headNode; node != nil; node = node.nextNode {
		if node.previousNode != nil && node.nextNode != nil {
			if node.previousNode.property == firstProperty && node.nextNode.property == secondProperty {
				nodeWith = node
				break
			}
		}
	}
	return nodeWith
}

func (LinkedList *LinkedList) AddToHead(property int) {
	var node = &Node{}

	node.property = property
	node.nextNode = nil
	if LinkedList.headNode != nil {
		node.nextNode = LinkedList.headNode
		LinkedList.headNode.previousNode = node
	}
	LinkedList.headNode = node
}

func (LinkedList *LinkedList) AddAfter(nodeProperty, property int) {
	var node = &Node{}
	node.property = property
	node.nextNode = nil
	var nodeWith *Node
	nodeWith = LinkedList.NodeWithValue(nodeProperty)
	if nodeWith != nil {
		node.nextNode = nodeWith.nextNode
		node.previousNode = nodeWith
		nodeWith.nextNode = node
	}
}

func (LinkedList *LinkedList) AddToEnd(property int) {
	var node = &Node{}
	node.property = property
	node.nextNode = nil
	var lastNode *Node
	lastNode = LinkedList.LastNode()
	if lastNode != nil {
		lastNode.nextNode = node
		node.previousNode = lastNode
	}
}

func (LinkedList *LinkedList) IterateList() {
	var node *Node

	for node = LinkedList.headNode; node != nil; node = node.nextNode {
		fmt.Println(node.property)
	}
}

func main() {
	var linkedList LinkedList
	linkedList = LinkedList{}
	linkedList.AddToHead(1)
	linkedList.AddToEnd(3)
	linkedList.AddToEnd(6)
	linkedList.AddAfter(3, 8)
	linkedList.IterateList()
	fmt.Println("node between: ", linkedList.NodeBetweenValues(3, 6).property)

}
