package com.schmeister.devprep.LinkedLists;

import com.schmeister.devprep.Print;

public class NthFromEnd {
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

	public Node nthFromEnd(int nth) {
		Node node1 = head;
		Node node2 = head;
		int count = 0;

		while (node1.next != null) {
			count++;
			node1 = node1.next;
		}

		for (int x = 0; x <= (count - nth); x++) {
			node2 = node2.next;
		}
		return node2;
	}

	public static void nthFromEnd() {
		NthFromEnd fm = new NthFromEnd();
		Node temp = null;
		for (int x = 0; x <= 7; x++) {
			temp = fm.append(x);
		}
		int pos = 5;
		Print.print("NthFromEnd", fm.toString());
		Print.print("NthFromEnd (" + pos + ")", fm.nthFromEnd(pos).data);
	}
}