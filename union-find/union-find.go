package unionfind

type UnionFind interface {
	Union(p, q int)
	Connected(p, q int) bool
}
