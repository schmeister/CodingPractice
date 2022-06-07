package com.schmeister.google;

import java.util.PriorityQueue;

class KthLargest {
    private static int k;
    private PriorityQueue<Integer> heap;
    
    public KthLargest(int k, int[] nums) {
        this.k = k;
        heap = new PriorityQueue<>();
        
        for (int num: nums) {
            heap.offer(num);
        }
        
        while (heap.size() > k) {
            heap.poll();
        }
    }
    
    public int add(int val) {
        heap.offer(val);
        if (heap.size() > k) {
            heap.poll();
        }

        return heap.peek();
    }
    
	public static void testApp() {
		System.out.println("Done");
		int k = 3;
		int[] arr = { 4, 5, 8, 2 };
		KthLargest kthLargest = new KthLargest(3, arr);
		System.out.printf("%s\t%d\n",kthLargest.heap.toString(), kthLargest.add(3)); // returns 4
		System.out.printf("%s\t%d\n",kthLargest.heap.toString(), kthLargest.add(5)); // returns 5
		System.out.printf("%s\t%d\n",kthLargest.heap.toString(), kthLargest.add(10)); // returns 5
		System.out.printf("%s\t%d\n",kthLargest.heap.toString(), kthLargest.add(9)); // returns 8
		System.out.printf("%s\t%d\n",kthLargest.heap.toString(), kthLargest.add(4)); // returns 8
	}

}
