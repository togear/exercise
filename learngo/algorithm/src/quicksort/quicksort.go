package quicksort

func swap(A []int, i,j int) {
	temp := A[j];
	A[j] = A[i]
	A[i] = temp;
}

func partition(A []int, start,end int) int {
	i := start;
	j := end;
	pivot := A[start];

	for i<j {
		for i <=j && A[i] <= pivot {
			i++;
		}
		for i <=j  && A[j] >= pivot {
			j--;
		}
		if(i < j) {
			swap(A,i,j);
		}
	}
	swap(A,start,j);
	return j;
}

func quickSort(A []int,s,e int)[]int {
	if(s < e ) {
		t := partition(A,s,e)
		quickSort(A,s,t-1)
		quickSort(A,t+1,e)
	}
	return A;
}

func QuickSort(unsorted []int) []int {
	return quickSort(unsorted, 0 ,len(unsorted)-1)
}  
