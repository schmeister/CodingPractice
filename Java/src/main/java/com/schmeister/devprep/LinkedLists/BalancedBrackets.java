package com.schmeister.devprep.LinkedLists;

import com.schmeister.devprep.util.Print;

public class BalancedBrackets {
	class Node {
		char data;
		Node next;
		boolean visited = false;

		Node(char c) {
			data = c;
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

	public void push(char c) {
		Node node = new Node(c);
		node.next = head;
		head = node;
	}

	public char pop() {
		Node n = head;
		head = head.next;
		return n.data;
	}

	public boolean balancedBrackets(String str) {

		int len = str.length();
		String pushes = "({[<";
		String pops = ")}]>";
		for (int x = 0; x < len; x++) {
			char c = str.charAt(x);
			if (pushes.indexOf(c) >= 0) {
				push(c);

			} else if (pops.indexOf(c) >= 0) {
				if (head == null)
					return false;

				char p = pop();

				boolean matched = false;
				if (c == ')')
					matched = p == '(';
				else if (c == '}')
					matched = p == '{';
				else if (c == ']')
					matched = p == '[';
				else if (c == '>')
					matched = p == '<';

				if (!matched)
					return false;
			}
		}
		if (head != null)
			return false;
		return true;
	}

	public static void balancedBrackets() {
		BalancedBrackets bb = new BalancedBrackets();
//		bb.balancedBrackets("{}[]{}[]<>");
		Print.print("BalancedBrackets", bb.balancedBrackets("({[][][]{}})"));

	}
}