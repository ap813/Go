package main

import "fmt"

type node struct {
	num int
	next *node
}

func insertTail(head *node, n int) *node {

	if head == nil {
		head = new(node);
		head.num = n;
		return head;
	}

	head.next = insertTail(head.next, n);

	return head;
}

func insertHead(head *node, n int) *node{

	root := new(node);

	root.num = n;

	root.next = head;

	return root;
}

func deleteTail(head *node) *node {
	if(head.next == nil) {
		head = nil;
		return nil;
	}

	head.next = deleteTail(head.next);

	return head;
}

func deleteHead(head *node) *node {

	temp := head.next;

	head = nil;

	return temp;
}


func printList(head *node) {

	if head == nil {
		return;
	}

	fmt.Println(head.num);

	printList(head.next);
}

func main() {

	root := insertHead(nil, 8);
	root = insertHead(root, 3);

	printList(root);

	root = deleteHead(root);

	printList(root);
}
