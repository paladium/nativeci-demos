package main

func main() {
	const n = 35
	d := [n + 1]int{}
	d[0] = 1
	d[1] = 1
	for i := 2; i <= n; i++ {
		d[i] = d[i-1] + d[i-2]
	}
	print(d[n])
}
