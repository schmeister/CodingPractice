package com.schmeister.devprep.LinkedLists;

import com.schmeister.devprep.Print;

public class Sort {
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

	public void sort() {
		Node zeros = null;
		Node ones = null;
		Node twos = null;

		while (head != null) {
			Node cur = head;
			head = head.next;
			
			if (cur.data % 3 == 0) {
				cur.next = zeros;
				zeros = cur;
			}
			else if (cur.data % 3 == 1) {
				cur.next = ones;
				ones = cur;
			}
			else if (cur.data % 3 == 2) {
				cur.next = twos;
				twos = cur;
			}
		}
		Print.print("Sort (z)", zeros.toString());
		Print.print("Sort (o)", ones.toString());
		Print.print("Sort (t)", twos.toString());
	}

	public static void sortLinkedList() {
		Sort list1 = new Sort();

		for (int x = 0; x < 10; x++) {
			list1.push(x % 3);
		}
		// creating first list
		Print.print("Sort", list1.toString());
		list1.sort();
	}
}