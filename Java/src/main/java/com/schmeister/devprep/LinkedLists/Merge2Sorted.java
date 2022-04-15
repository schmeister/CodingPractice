package com.schmeister.devprep.LinkedLists;

import com.schmeister.devprep.util.Print;

public class Merge2Sorted {
	class Node {
		int data;
		Node next;
		boolean visited = false;

		Node(int d) {
			data = d;
			next = null;
		}

		@Override
		public String toString() {
			StringBuilder sb = new StringBuilder();

			Node n = this;
			while (n != null) {
				sb.append(n.data);
				if (n.next != null)
					sb.append(" -> ");
				n = n.next;
			}
			return sb.toString();
		}
	}

	Node head = null;

	@Override
	public String toString() {
		int maxLen = 19;
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

	static Node merge(Node h1, Node h2) {
		if (h1 == null) {
			return h2;
		}
		if (h2 == null) {
			return h1;
		}

		// start with the linked list
		// whose head data is the least
		if (h1.data < h2.data) {
			h1.next = merge(h1.next, h2);
			return h1;
		} else {
			h2.next = merge(h1, h2.next);
			return h2;
		}
	}

	public static void merge2Sorted() {
		Merge2Sorted list1 = new Merge2Sorted();
		Merge2Sorted list2 = new Merge2Sorted();

		Node temp = null;
		for (int x = 0; x <= 25; x += 5) {
			temp = list1.append(x);
		}
		for (int x = 1; x <= 40; x += 3) {
			temp = list2.append(x);
		}

		Print.print("Merge2Sorted (1)", list1.toString());
		Print.print("Merge2Sorted (2)", list2.toString());
		list1.head = list1.merge(list1.head, list2.head);
		Print.print("Merge2Sorted", list1.toString());
	}
}