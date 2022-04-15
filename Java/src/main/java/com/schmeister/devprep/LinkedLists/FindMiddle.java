package com.schmeister.devprep.LinkedLists;

import com.schmeister.devprep.util.Print;

public class FindMiddle {
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

	public int getMiddle() {
		Node node = head;
		Node nodeFast = head;
		if (head != null) {
			while (nodeFast != null && nodeFast.next != null) {
				node = node.next;
				nodeFast = nodeFast.next.next;
			}
			return node.data;
		}
		return 0;
	}

	@Override
	public String toString() {
		StringBuilder sb = new StringBuilder();
		
		Node node = head;
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

	public static void findMiddle() {
		FindMiddle fm = new FindMiddle();
		for (int x = 7; x >= 0; x--) {
			fm.push(x);
		}
		Print.print("FindMiddle", fm.toString());
		Print.print("FindMiddle", fm.getMiddle());
	}
}
