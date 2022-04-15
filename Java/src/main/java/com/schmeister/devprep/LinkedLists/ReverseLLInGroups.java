package com.schmeister.devprep.LinkedLists;

import com.schmeister.devprep.util.Print;

public class ReverseLLInGroups {
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
	
	private Node reverse(Node head, int k)
    {
        if(head == null)
          return null;
        
        Node current = head;
        Node next = null;
        Node prev = null;
 
        int count = 0;
 
        while (count < k && current != null) {
            next = current.next;
            current.next = prev;
            prev = current;
            current = next;
            count++;
        }
 
        if (next != null)
            head.next = reverse(next, k);
 
        return prev;
    }

	public static void reverseInGroups() {
		ReverseLLInGroups fm = new ReverseLLInGroups();
		for (int x = 0; x < 12; x++) {
			fm.push(x);
		}
		fm.head = fm.reverse(fm.head, 3);
		Print.print("ReverseInGroups", fm.toString());
	}
}