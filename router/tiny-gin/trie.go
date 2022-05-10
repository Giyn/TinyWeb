/*
-------------------------------------
# @Time    : 2022/5/10 20:45:03
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : trie.go
# @Software: GoLand
-------------------------------------
*/

package tiny_gin

import "strings"

type node struct {
	pattern  string // 待匹配路由，如 /p/:lang
	part     string // 路由中的一部分，如 :lang
	children []*node
	isWild   bool // 是否精确匹配，part 含有 : 或 * 时为 true
}

// insert 插入结点
func (n *node) insert(pattern string, parts []string, height int) {
	if len(parts) == height {
		n.pattern = pattern
		return
	}
	part := parts[height]
	child := n.matchChild(part)
	if child == nil {
		child = &node{part: part, isWild: part[0] == ':' || part[0] == '*'}
		n.children = append(n.children, child)
	}
	child.insert(pattern, parts, height+1)
}

// search 查询结点
func (n *node) search(parts []string, height int) *node {
	if len(parts) == height || strings.HasSuffix(n.part, "*") {
		if n.pattern == "" {
			return nil
		}
		return n
	}
	part := parts[height]
	children := n.matchChildren(part)
	for _, child := range children {
		ans := child.search(parts, height+1)
		if ans != nil {
			return ans
		}
	}
	return nil
}

// matchChild 第一个匹配成功的结点，用于插入
func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		if child.part == part || child.isWild {
			return child
		}
	}
	return nil
}

// matchChildren 所有匹配成功的结点，用于查找
func (n *node) matchChildren(part string) (ans []*node) {
	for _, child := range n.children {
		if child.part == part || child.isWild {
			ans = append(ans, child)
		}
	}
	return
}
