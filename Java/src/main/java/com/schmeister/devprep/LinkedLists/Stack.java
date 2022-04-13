package com.schmeister.devprep.LinkedLists;

import com.schmeister.devprep.Print;

public class Stack {
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

	public Node pop() {
		Node n = head;
		head = head.next;
		n.next = null;
		return n;
	}

	public static void stack() {
		Stack list1 = new Stack();

		// creating first list
		list1.push(7);
		list1.push(5);
		list1.push(9);
		list1.push(4);
		list1.push(6);
		Print.print("Stack", list1.toString());
		Print.print("Stack", list1.pop().toString());
		Print.print("Stack", list1.toString());
	}
}