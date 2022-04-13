package com.schmeister.devprep.LinkedLists;

import com.schmeister.devprep.Print;

public class AddAsNumbers {
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

	public void add(Node head2) {
		Node list1 = head;
		Node list2 = head2;
		int carry = 0;

		while (list1 != null && list2 != null) {
			int tmp = (list1.data + list2.data + carry) % 10;
			carry = (list1.data + list2.data + carry) / 10;
			list1.data = tmp;

			list1 = list1.next;
			list2 = list2.next;
		}
		while (list1 != null) {
			int tmp = (list1.data + carry) % 10;
			carry = (list1.data + carry) / 10;
			list1.data = tmp;

			list1 = list1.next;
		}
		while (list2 != null) {
			int tmp = (list2.data + carry) % 10;
			carry = (list2.data + carry) / 10;
			list2.data = tmp;

			list2 = list2.next;
		}
	}

	public static void addAsNumbers() {
		AddAsNumbers list1 = new AddAsNumbers();
		AddAsNumbers list2 = new AddAsNumbers();

		// creating first list
		list1.append(7);
		list1.append(5);
		list1.append(9);
		list1.append(4);
		list1.append(6);
		Print.print("addAsNumbers", list1.toString());

		// creating second list
		list2.append(8);
		list2.append(4);
		Print.print("addAsNumbers", list2.toString());

		list1.add(list2.head);
		Print.print("addAsNumbers", list1.toString());
	}
}