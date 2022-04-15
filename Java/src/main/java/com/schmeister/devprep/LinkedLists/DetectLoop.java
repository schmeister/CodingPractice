package com.schmeister.devprep.LinkedLists;

import com.schmeister.devprep.util.Print;

public class DetectLoop {
	class Node {
		int data;
		Node next;
		boolean visited = false;

		Node(int d) {
			data = d;
			next = null;
		}
	}

	Node head = null;
	String slowString;
	String fastString;

	@Override
	public String toString() {
		int count = 0;
		Node node = head;
		StringBuilder sb = new StringBuilder();
		if (node != null) {
			while (node != null) {
				sb.append(node.data);
				if (node.next != null)
					sb.append(" -> ");

				node = node.next;
				if (count++ > 10)
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

	public boolean hasLoop() {
		StringBuilder ssb = new StringBuilder();
		StringBuilder fsb = new StringBuilder();
		boolean retVal = false;

		Node slow = head;
		Node fast = head.next;
		while (fast.next != null && fast.next.next != null) {
			ssb.append(slow.data + " ");
			fsb.append(fast.data + " ");

			if (slow.data == fast.data) {
				retVal = true;
				break;
			}

			slow = slow.next;
			fast = fast.next.next;
		}
		slowString = ssb.toString();
		fastString = fsb.toString();
		return retVal;
	}

	public static void detectLoop() {
		DetectLoop fm = new DetectLoop();
		fm.push(20);
		fm.push(4);
		fm.push(15);
		fm.push(14);
		fm.push(3);
		fm.push(10);
		fm.head.next.next.next.next.next.next = fm.head.next;

		boolean hasLoop = fm.hasLoop();
		Print.print("DetectLoop (slow)", fm.slowString);
		Print.print("DetectLoop (fast)", fm.fastString);
		Print.print("DetectLoop", hasLoop);
	}
}