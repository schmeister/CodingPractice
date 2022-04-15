package com.schmeister.devprep.Trees;

import java.util.ArrayList;
import java.util.Iterator;

class Node {
	int data;
	ArrayList<Node> children;

	public Node(int data) {
		this.data = data;
		this.children = new ArrayList<Node>();
	}
}

public class Nary {
	Node root;
	void traverse(Node root) {
		if (root == null)
			return;
		
		System.out.printf("%d ",root.data);
		
		Iterator<Node> iter = root.children.iterator();
		while (iter.hasNext()) {
			Node n = iter.next();
			traverse(n);
		}
	}
	
	public static void nary() {
		Nary nary = new Nary();
		nary.root = new Node(0);
		
		nary.root.children.add(new Node(1));
		nary.root.children.add(new Node(2));
		nary.root.children.add(new Node(3));
		
		nary.root.children.get(0).children.add(new Node(11));
		
		System.out.printf("Nary preorder: ");
		nary.traverse(nary.root);
		System.out.printf("\n");
		
	}
}
