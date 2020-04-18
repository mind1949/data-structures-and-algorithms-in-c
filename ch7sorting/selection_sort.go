package sort

// SelectionSort 选择排序
func SelectionSort(s []int) {
	n := len(s)                // c0*1
	for i := 0; i < n-1; i++ { // c1*n
		minIdx := i                  // c2*1
		for j := i + 1; j < n; j++ { // c3*(n+n-1+n-2+...+2)
			if s[j] < s[minIdx] { // c4*(n-1+n-2+n-3+...+1)
				minIdx = j // c5*(n-1+n-2+n-3+...+1)
			}
		}
		s[i], s[minIdx] = s[minIdx], s[i] // c6*(n-1)
	}
}

/*
证明:
* 循环不变式: s[:i+1]为已排序, 且s[:]中的元素为排序前的元素
* 初始化:
> 循环在第一次迭代之前,它为真
第一次贴点之前i=0, s[0:i]为s[0:1], 仅有一个元素, 所以循环不变式为真
* 保持:
> 若循环在某次迭代之前它为真, 那么下次迭代之前它仍为真
在进行下一次迭代过程中, 选择排序做的工作就是从s[i+1:n]中找到比s[i]小, 且为s[i+1:n]中最小的数互换位置, 这样s[0:i+1]依然为有序的, 且s[:]中的元素为排序前的元素, 所以循环不变式依然为真
* 终止:
当循环终止时, i=n-1, 又根据`保持`中的分析, 此时s[0:i+1]为有序, 也就是s[0:n]为有序, 并且s[:]中元素为排序前的元素, 所以可以证明当循环终止时输入的数据集合已经排序完成.
*/

/*
分析:
假设每一行的运行时间为常量时间ci,
所以总时间可以用如下公式表示:
T(n)= c0+c1*n+c2+c3*(n+n-1+n-2+...+2)+c4*(n-1+n-2+n-3+...+1)+c5*(n-1+n-2+n-3+...+1)+c6*(n-1)
	= c0+c1*n+c2+c3(n+2)*(n-1)/2+c4*(n)*(n-1)/2+c5*(n)*(n-1)/2+c6*n-c6
	= c0+c1*n+c2+(c3/2)*(n^2+n-2)+(c4/2)*(n^2-n)+(c5/2)*(n^2-n)+c6*n-c6
	= (c3/2+c4/2+c5/2)*n^2 + (c1+c3/2-c4/2-c5/2+c6)*n + (c0+c2-c3-c6)
	= a*n^2+b*n+c
最坏情况与最好情况的量级军事`n^2`, 用符号表示是`Θ(n^2)`
*/
