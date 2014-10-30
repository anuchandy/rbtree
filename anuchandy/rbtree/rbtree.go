package rbtree

// Type whose instances can hold reference to a function 
// which  can perform comparison of two keys and returns 
// an integer indicating whether one key is less than, 
// equal to, or greater  than the other.
//
type Comparer func(key1 interface{}, key2 interface{}) int

// Structure represents a node in RB tree.
//
type RBNode struct {
  Color bool // true: BLACK false:RED
  Left, Right, Parent  *RBNode
  Key interface{}
  Value interface{}
}

// Structure representing the root of the RB tree.
//
type RBRoot struct {
  Node *RBNode
}

// Structure representing the RedBlack Tree.
//
type RBTree struct {
  RBRoot *RBRoot
  Comparer Comparer
}

// Sets the color of the node as BLACK.
//
func (node *RBNode) setBlack() {
  node.Color = true
}

// Sets the color of the node as RED.
//
func (node *RBNode) setRed() {
  node.Color = false
}

// Returns TRUE if the node is BLACK else FALSE.
//
func (node *RBNode) IsBlack() bool {
  return node.Color
}

// Returns TRUE if the node is RED else FALSE.
//
func (node *RBNode) IsRed() bool {
  return !node.Color
}

// Returns TRUE if node is left child of its parent.
//
func (node *RBNode) isLeftChild() bool {
  return node.Parent.Left == node
}

// Returns TRUE if node is right child of its parent.
//
func (node *RBNode) isRightChild() bool {
  return node.Parent.Right == node
}

// Sets parent as node's parent
//
func (node *RBNode) setParent(parent *RBNode) {
  node.Parent = parent
}

// Rotate tree to left with node as root.
//
func (node *RBNode) Rotateleft(root *RBRoot) {
  right := node.Right
  parent := node.Parent
  
  node.Right = right.Left
  if node.Right != nil {
    node.Right.setParent(node)
  }
  
  right.Left = node
  right.setParent(parent)
  if parent != nil {
    if node.isLeftChild() {
      parent.Left = right
    } else {
      parent.Right = right
    }
  } else {
    root.Node = right
  }
  
  node.setParent(right)
}

// Rotate tree to right with node as root.
//
func (node *RBNode) Rotateright(root *RBRoot) {
  left := node.Left
  parent := node.Parent
  
  node.Left = left.Right
  if node.Left != nil {
    node.Left.setParent(node)
  }
  
  left.Right = node
  left.setParent(parent)
  if parent != nil {
    if node.isLeftChild() {
      parent.Left = left
    } else {
      parent.Right = left
    }
  } else {
    root.Node = left
  }
  
  node.setParent(left)
}

// Create and initialize a new RBTree instance and 
// return it.
//
func CreateNewRBTree(comparer Comparer) *RBTree {
  newTree := new(RBTree)
  newTree.RBRoot = &RBRoot { Node:nil }
  newTree.Comparer = comparer
  return newTree
}

// Create and initialize a new RBNode instance and 
// return it.
//
func CreateNewRBNode(key interface{}, value interface{}) *RBNode {
  newNode := new(RBNode)
  newNode.Key = key
  newNode.Value = value
  newNode.Color = false
  return newNode
}

// Search for RBNode with the given give and return the node
// if exists else nil.
//
func (rbTree *RBTree) SearchNode(key interface{}) *RBNode {
  var node *RBNode
  var curr *RBNode
  
  curr = rbTree.RBRoot.Node
  for ; curr != nil; {
    node = curr
    c := rbTree.Comparer(key, node.Key)
    if c > 0 {
      curr = curr.Right
    } else if c < 0 {
      curr = curr.Left
    } else {
      return node
    }
  }

  return nil
}

// Delete the node from the RB tree ensure resulting tree 
// is RB tree.
//
func (rbTree *RBTree) DeleteNode(node *RBNode) {
  root := rbTree.RBRoot
  var parent *RBNode
  if node.Left != nil && node.Right != nil {
    // Find in-order successor [The left most child in the 
    // right sub-tree of node]
    successor := node.Right
    for ; successor.Left != nil ; {
       successor = successor.Left
    }

    // node's parent relinquish node and accept successor as 
    // child
    if node.Parent != nil {
      if node.isLeftChild() {
        node.Parent.Left = successor        
      } else {
        node.Parent.Right = successor
      }
    } else {
      root.Node = successor
    }

    // Adopt node's left child
    successor.Left = node.Left
    successor.Left.setParent(successor)
    
    parent = successor.Parent
    // successor's right child is either leaf node or external node
    child := successor.Right
    if node != parent { //successor is not immediate right child of 
      // node. Since successor is not immediate right child of node, 
      // it will be left child of it's parent. let successor's parent 
      // look after successor's right child
      parent.Left = child
      if child != nil {
        child.setParent(parent)
      }

      // Adopt node's right child  
      successor.Right = node.Right
      successor.Right.setParent(successor)
    } else {
      parent = successor
    }

    // successor accept it's new parent as node's parent
    successor.setParent(node.Parent)
    // successor now replaces 'node', so color must be same as node's 
    // color
    oldColor := successor.Color
    successor.Color = node.Color
    
    if child != nil {
      // Successor had a leaf right child which must be RED. Successor 
      // was originally BLACK in this case. Set child's color to BLACK 
      // to maintain black height
      child.Color = true
    } else if oldColor {
      // Successor was a leaf node and was black. The black height got 
      // reduced by 1 due to  the removal of this successor, fix this
      rbTree.fixBlackDepth(parent)
    }
  } else {
    // node is either a leaf or a non-leaf node with only one child
    var child *RBNode
    if node.Right != nil {
      child = node.Right
    } else {
      child = node.Left
    }

    parent = node.Parent
    if parent != nil {
      if parent.Left == node {
        parent.Left = child
      } else {
        parent.Right = child
      }
    } else {
      root.Node = child
    }

    if child != nil {
      // Child is originally RED and node is BLACK. Set child to 
      // BLACK to maintain black height.
      child.setParent(parent)
      child.Color = true
    } else if node.Color {
      rbTree.fixBlackDepth(parent)
    }
  }
}

// Put the given value into the RB tree with KEY as key, if key 
// already exists then replace the value.
//
func (rbTree *RBTree) PutValue(key interface{}, value interface{}) (*RBNode, interface{}) {
  var prev *RBNode
  var pCurr **RBNode
  
  prev = nil
  pCurr = &rbTree.RBRoot.Node
  for ; *pCurr != nil ;  {
    prev = *pCurr
    c := rbTree.Comparer(key, prev.Key)
    if c > 0 {
      pCurr = &((*(pCurr)).Right)
    } else if c < 0 {
      pCurr = &((*(pCurr)).Left)
    } else {
      oldVal := (*(pCurr)).Value
      (*(pCurr)).Value = value
      return *pCurr, oldVal
    }
  }

  *pCurr = CreateNewRBNode(key, value)
  (*(pCurr)).Parent = prev
  rbTree.fixDoubleRed(*pCurr)
  return *pCurr, nil
}

// parent's left or right child [which was leaf] got deleted, it's 
// color was BLACK, so BLACK height property violated, fix it.
//
func (rbTree *RBTree) fixBlackDepth(parent *RBNode) {
  root := rbTree.RBRoot
  var child *RBNode = nil
  for ; parent != nil ; {
    if parent.Right == child {
      // Black depth of right sub-tree of parent got reduced by one
      var sibling *RBNode = parent.Left
      if parent.IsRed() {
        // parent is RED, the left child i.e. 'sibling' must present 
        // and it must be BLACK
        if sibling.Left != nil && sibling.Left.IsRed() {
          // Left child of sibling is RED
          sibling.setRed()
          sibling.Left.setBlack()
          parent.setBlack()
          parent.Rotateright(root)
          break;
        } else if sibling.Right != nil && sibling.Right.IsRed() {
          // Right child of sibling is RED
          sibling.Rotateleft(root)
          parent.setBlack()
          parent.Rotateright(root)
          break
        } else {
          // Both children of sibling is BLACK (note: we consider 
          // external nodes are BLACK)
          sibling.setRed()
          parent.setBlack()
          break;
        }
      } else {
        // parent is a BLACK, the left child i.e. 'sibling' must 
        // present it can be either RED or BLACK
        if sibling.IsRed() {
          // sibling (left child of parent) is RED

          // Grandchild of parent must be present and will be BLACK
          // because it's parent 'sibling' is RED
          var grandChild *RBNode = sibling.Right
          if grandChild.Left != nil && grandChild.Left.IsRed() {
            // Left child of grandchild is RED
            grandChild.Left.setBlack()
            sibling.Rotateleft(root)
            parent.Rotateright(root)
            break
          } else if grandChild.Right != nil && grandChild.Right.IsRed() {
            // Right child of grandchild is RED
            grandChild.Right.setBlack()
            grandChild.Rotateleft(root)
            sibling.Rotateleft(root)
            parent.Rotateright(root)
            break
          } else {
            // Both children of grandchild is BLACK (note: we consider 
            // external nodes are BLACK)
            grandChild.setRed()
            sibling.setBlack()
            parent.Rotateright(root)
            break
          }
        } else {
          // sibling (left child of parent) is BLACK
          if sibling.Left != nil && sibling.Left.IsRed() {
            // Left child of sibling is RED
            sibling.Left.setBlack()
            parent.Rotateright(root)
            break
          } else if sibling.Right != nil && sibling.Right.IsRed() {
            // Right child of sibling is RED
            sibling.Right.setBlack()
            sibling.Rotateleft(root)
            parent.Rotateright(root)
            break
          } else {
          // Both children of sibling is BLACK (note: we consider 
          // external nodes are BLACK)
            sibling.setRed()
            child = parent
            parent = parent.Parent
            continue
          }
        }
      }
    } else {
      // Back depth of left sub-tree of parent got reduced by one
      var sibling *RBNode = parent.Right
      if parent.IsRed() {
        if sibling.Right != nil && sibling.Right.IsRed() {
          sibling.setRed()
          sibling.Right.setBlack()
          parent.setBlack()
          parent.Rotateleft(root)
          break;
        } else if sibling.Left != nil && sibling.Left.IsRed() {
          sibling.Rotateright(root)
          parent.setBlack()
          parent.Rotateleft(root)
          break
        } else {
          sibling.setRed()
          parent.setBlack()
          break;
        }
      } else {
        if sibling.IsRed() {
          var grandChild *RBNode = sibling.Left
          if grandChild.Right != nil && grandChild.Right.IsRed() {
            grandChild.Right.setBlack()
            sibling.Rotateright(root)
            parent.Rotateleft(root)
            break
          } else if grandChild.Left != nil && grandChild.Left.IsRed() {
            grandChild.Left.setBlack()
            grandChild.Rotateright(root)
            sibling.Rotateright(root)
            parent.Rotateleft(root)
            break
          } else {
            grandChild.setRed()
            sibling.setBlack()
            parent.Rotateleft(root)
            break
          }
        } else {
          if sibling.Right != nil && sibling.Right.IsRed() {
            sibling.Right.setBlack()
            parent.Rotateleft(root)
            break
          } else if sibling.Left != nil && sibling.Left.IsRed() {
            sibling.Left.setBlack()
            sibling.Rotateright(root)
            parent.Rotateleft(root)
            break
          } else {
            sibling.setRed()
            child = parent
            parent = parent.Parent
            continue
          }
        }
      }
    }
  }
}

// The node inserted as a new node, double RED problem can occur
// if it's parent is RED, if so fix it.
//
func (rbTree *RBTree) fixDoubleRed(node *RBNode) {
  root := rbTree.RBRoot
  parent := node.Parent
  for ; parent != nil && parent.IsRed() ; {
    // grandparent must present and will be BLACK 
    // because parent is RED
    grandParent := parent.Parent
    if grandParent.Right == parent {
      // uncle may or may not present and can be 
      // RED or BLACK       
      uncle := grandParent.Left
      if uncle != nil && uncle.IsRed() {
        grandParent.setRed()
        uncle.setBlack()
        parent.setBlack()
        node = grandParent
        parent = node.Parent
      } else {
        // uncle is BLACK or external node
        if node == parent.Right {
          grandParent.setRed()
          parent.setBlack()
          // Single rotation
          grandParent.Rotateleft(root)
        } else {
          grandParent.setRed()
          node.setBlack()
          // Double rotation
          parent.Rotateright(root)
          grandParent.Rotateleft(root)
        }
      }
    } else {
      uncle := grandParent.Right
      if uncle != nil && uncle.IsRed() {
        grandParent.setRed()
        uncle.setBlack()
        parent.setBlack()
        node = grandParent
        parent = node.Parent
      } else {
        if node == parent.Left {
          grandParent.setRed()
          parent.setBlack()
          // Single rotation
          grandParent.Rotateright(root)
        } else {
          grandParent.setRed()
          node.setBlack()
          // Double rotation
          parent.Rotateleft(root)
          grandParent.Rotateright(root)
        }
      }
    }
  }
  
  root.Node.setBlack()
}