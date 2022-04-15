package com.schmeister.devprep.LinkedLists;

import com.schmeister.devprep.util.Print;

public class Rotate {
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

	public void rotateList(int count) {

		Node rearHead = null;

		Node current = head;
		Node prev = null;

		rearHead = current;

		for (int i = 0; i < count && (current != null); i++) {
			prev = current;
			current = current.next;
		}
		if (current != null) {
			prev.next = null;
			head = current;
			while (current != null && current.next != null) {
				current = current.next;
			}
			current.next = rearHead;
		}
	}

	public static void rotate() {
		Rotate fm = new Rotate();
		for (int x = 5; x >= 0; x--) {
			fm.push(x);
		}
		fm.rotateList(2);
		Print.print("Rotate", fm.toString());
	}
}
