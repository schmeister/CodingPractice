package com.schmeister.devprep.LinkedLists;

import com.schmeister.devprep.Print;

public class Queue {
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
	public Node remove()
	{
		Node node = head;
		head = head.next;
		
		node.next = null;
		
		return node;
	}

	public static void queue() {
		Queue list1 = new Queue();

		// creating first list
		list1.append(7);
		list1.append(5);
		list1.append(9);
		list1.append(4);
		list1.append(6);
		Print.print("Queue", list1.toString());
		Print.print("Queue", list1.remove().toString());
		Print.print("Queue", list1.toString());
	}
}