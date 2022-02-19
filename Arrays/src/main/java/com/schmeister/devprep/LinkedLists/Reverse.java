package com.schmeister.devprep.LinkedLists;

import com.schmeister.devprep.Print;

public class Reverse {
	class Node {
		int data;
		Node next;

		Node(int d) {
			data = d;
			next = null;
		}
	}

	Node head = null;

	public void push(int i) {
		Node node = new Node(i);
		node.next = head;
		head = node;
	}

	@Override
	public String toString() {
		StringBuilder sb = new StringBuilder();
		
		Node node = head;
		Node rev = null;
		if (head != null) {
			while (node != null) {
				sb.append(node.data);
				if (node.next != null)
					sb.append(" -> ");
				
				node = node.next;
			}
		}
		return sb.toString();
	}

	public void reverseList() {
		Node prev = null;
		Node current = head;
		Node next = null;

		while (current != null) {
			next = current.next;
			current.next = prev;
			prev = current;
			current = next;
		}
		head = prev;
	}

	public static void reverse() {
		Reverse fm = new Reverse();
		for (int x = 7; x >= 0; x--) {
			fm.push(x);
		}
		fm.reverseList();
		Print.print("Reverse", fm.toString());
	}
}
