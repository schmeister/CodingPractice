package com.schmeister.devprep.LinkedLists;

import com.schmeister.devprep.Print;

public class RemoveLoop {
	class Node {
		int data;
		Node next;
		boolean visited = false;

		Node(int d) {
			data = d;
			next = null;
		}
	}

	Node head = null;

	@Override
	public String toString() {
		int maxLen = 15;
		int len = 0;

		Node node = head;
		StringBuilder sb = new StringBuilder();
		if (node != null) {
			while (node != null) {
				sb.append(node.data);
				if (node.next != null)
					sb.append(" -> ");

				node = node.next;
				if (len++ >= maxLen)
					break;
			}
		}
		String one = sb.toString();

		return one;
	}

	public Node push(int i) {
		Node node = new Node(i);
		node.next = head;
		head = node;
		return head;
	}

	public Node append(int i) {
		Node node = new Node(i);
		if (head == null) {
			head = node;
			return node;
		}

		Node cur = head;
		while (cur.next != null) {
			cur = cur.next;
		}
		cur.next = node;

		return node;
	}

	public boolean hasLoop() {
		Node slow = head;
		Node fast = head.next;

		while (fast != null && fast.next != null) {
			if (slow.data == fast.data) {
				remove(slow);
				return true;
			}
			slow = slow.next;
			fast = fast.next.next;
		}
		return false;
	}

	public Node remove(Node lnode) {
		Node node1 = lnode;
		Node node2 = lnode;
		int nodesInLoop = 1;

		// count nodes in loop:
		while (node1.next.data != node2.data) {
			nodesInLoop++;
			node1 = node1.next;
		}

		// Advance node2 by size of loop
		node1 = head;
		node2 = head;
		for (int x = 0; x < nodesInLoop; x++) {
			node2 = node2.next;
		}

		// node2 will now complete loop at the same time node1 reaches loop
		while (node1.next != node2.next) {
			node1 = node1.next;
			node2 = node2.next;
		}
		node2.next = null;

		return node1;
	}

	public static void removeLoop() {
		RemoveLoop fm = new RemoveLoop();
		Node temp = null;
		for (int x = 0; x <= 7; x++) {
			temp = fm.append(x);
		}
		temp.next = fm.head.next.next.next.next;

		Print.print("RemoveLoop", fm.toString());
		Print.print("RemoveLoop (loop?)", fm.hasLoop());
		Print.print("RemoveLoop", fm.toString());
	}
}