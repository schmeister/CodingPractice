package com.schmeister.devprep.LinkedLists;

import com.schmeister.devprep.util.Print;

public class IntersectingList {
	class Node {
		int data;
		Node next;

		Node(int d) {
			data = d;
			next = null;
		}
	}

	Node head1 = null;
	Node head2 = null;

	@Override
	public String toString() {

		String one = loop(head1);
		String two = loop(head2);

		return one + "  :  " + two;
	}

	private String loop(Node node) {
		StringBuilder sb = new StringBuilder();
		if (node != null) {
			while (node != null) {
				sb.append(node.data);
				if (node.next != null)
					sb.append(" -> ");

				node = node.next;
			}
		}
		return sb.toString();
	}

	public Node push(Node head, int i) {
		Node node = new Node(i);
		node.next = head;
		head = node;
		return head;
	}

	public Node append(Node head, int i) {
		Node node = new Node(i);
		if (head == null)
			return node;

		Node cur = head;
		while (cur.next != null) {
			cur = cur.next;
		}
		cur.next = node;

		return head;
	}

	public int getIntersecting() {
		Node cur1 = head1;
		Node cur2=  head2;
		int retVal = -1;
		
		while (cur1.next != null && cur2.next != null) {
			if (cur1.data < cur2.data)
				cur1 = cur1.next;
			if (cur1.data > cur2.data)
				cur2 = cur2.next;
			if (cur1.data == cur2.data) {
				retVal = cur1.data;
				break;
			}
		}
		return retVal;
	}

	public static void intersectingList() {
		IntersectingList fm = new IntersectingList();
		for (int x = 0; x < 30; x += 7) {
			fm.head1 = fm.append(fm.head1, x);
		}
		for (int x = 1; x < 35; x += 4) {
			fm.head2 = fm.append(fm.head2, x);
		}
		Print.print("Intersecting",fm.toString());
		Print.print("Intersecting",fm.getIntersecting());
	}
}