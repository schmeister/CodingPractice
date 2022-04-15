package com.schmeister.devprep.LinkedLists;

import com.schmeister.devprep.util.Print;

public class PairwiseSwap {
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

	public void swap() {
		Node temp = head;
		Node newHead = null;

		while (temp != null && temp.next != null) 
		{
			/* Swap the data */
			int k = temp.data;
			temp.data = temp.next.data;
			temp.next.data = k;
			temp = temp.next.next;
		}
		head = newHead;
	}

	public static void pairwiseSwap() {
		PairwiseSwap list1 = new PairwiseSwap();

		Node temp = null;
		for (int x = 0; x <= 25; x += 5) {
			temp = list1.append(x);
		}

		Print.print("PairwiseSwap", list1.toString());
		list1.swap();
		Print.print("PairwiseSwap", list1.toString());
	}
}