package main

import "fmt"

type node struct {
	left *node
	right * node
	num int
}

func insert(root *node, num int) *node {

	if root == nil {
		root := new(node);
		root.num = num;
		return root;
	}

	if root.num > num {
		root.left = insert(root.left,num);
	} else if root.num < num {
		root.right = insert(root.right,num);
	}

	return root;
}

func search(root *node, num int) bool {

	if root == nil {
		return false;
	}

	if root.num == num {
		return true;
	}

	return search(root.left, num) || search(root.right, num);
}

func delete(root *node, num int) *node {

	if root == nil {
		return nil;
	}

	if root.num == num {
		if root.left != nil && root.right != nil {
			// Find lowest number in Right Subtree
			// and replace deleted node
			temp := findLowest(root.right);

			root.num = temp.num;

			root.right = delete(root.right, temp.num);

		} else if root.left == nil {
			// Replace Root with the right
			root = root.right;
		} else if root.right == nil {
			// Replace Root with the left
			root = root.left;
		} else {
			// Delete the node
			root = nil;
		}

		return root;
	}

	root.left = delete(root.left, num);
	root.right = delete(root.right, num);
	return root;
}

func findLowest(root *node) *node{
	if root.left == nil {
		return root;
	}

	return findLowest(root.left);
}

func inorder(root *node) {

	if root == nil {
		return;
	}

	inorder(root.left);
	fmt.Println(root.num);
	inorder(root.right);
}

func preorder(root *node) {

	if root == nil {
		 return;
	}

	fmt.Println(root.num);
	preorder(root.left);
	preorder(root.right);
}

func postorder(root *node) {

	if root == nil {
		return;
	}

	postorder(root.left);
	postorder(root.right);
	fmt.Println(root.num);
}

func main() {

	root := insert(nil, 11);
	root = insert(root, 9);
	root = insert(root, 10);
	root = insert(root, 5);
	root = insert(root, 14);
	root = insert(root, 12);

	root = delete(root, 9);

	inorder(root);
}
