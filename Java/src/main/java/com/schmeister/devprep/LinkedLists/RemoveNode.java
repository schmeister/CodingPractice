package com.schmeister.devprep.LinkedLists;

import com.schmeister.devprep.Print;

public class RemoveNode {
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

	public void push(int i) {
		Node n = new Node(i);
		n.next = head;
		head = n;
	}

	public void removeNode(int num) {

		Node n = head;
		int count = 0;
		while (count <= num && n != null) {
			if (count == num && n.next != null) {
				n.data = n.next.data;
				n.next = n.next.next;
			}
			count++;
			n = n.next;
		}
	}

	public static void removenode() {
		RemoveNode list1 = new RemoveNode();

		for (int x = 0; x < 10; x++) {
			list1.push(x);
		}
		// creating first list
		Print.print("RemoveNode", list1.toString());
		list1.removeNode(8);
		Print.print("RemoveNode", list1.toString());
	}
}