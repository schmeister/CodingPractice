package com.schmeister.devprep.LinkedLists;

import com.schmeister.devprep.Print;

public class Palindrome {
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

	Node left;

	boolean isPalindrome(Node right) {
		left = head;
		// Stop recursion when right becomes null
		if (right == null)
			return true;

		// If sub-list is not palindrome then no need to
		// check for the current left and right, return
		// false
		boolean isp = isPalindrome(right.next);
		if (isp == false)
			return false;

		// Check values at current left and right
		boolean isp1 = (right.data == left.data);

		left = left.next;

		// Move left to next node;
		return isp1;
	}

	boolean isPal() {
		return isPalindrome(head);
	}

	public static void palindrome() {
		Palindrome list1 = new Palindrome();
		Palindrome list2 = new Palindrome();

		// creating first list
		list1.append(7);
		list1.append(5);
		list1.append(9);
		list1.append(4);
		list1.append(6);
		Print.print("Palindrome", list1.toString());
		Print.print("Palindrome", list1.isPal());

		list2.append(7);
		list2.append(5);
		list2.append(9);
		list2.append(5);
		list2.append(7);
		Print.print("Palindrome", list2.toString());
		Print.print("Palindrome", list2.isPal());
	}
}