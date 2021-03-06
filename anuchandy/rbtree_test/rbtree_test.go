package main

import (
         "testing"
         "anuchandy/rbtree"
        )

func TestInsert(t *testing.T) {
  rbTree := getTestRBTree(t)
  bd := validate(t, rbTree)
  
  if bd != 3 {
    t.Error ("Expected black depth to be 3 but got", bd)
  }
  
  // Below we ensure nodes are aligned in the same way the 
  // algorithm is supposed to align the resulting tree.

  // Level 0

  // rbTree.RBRoot.Node.Color
  assertBlack(t, rbTree.RBRoot.Node)
  // rbTree.RBRoot.Node.Key
  assertKey(t, rbTree.RBRoot.Node, 11)

  
  // Level 1

  // rbTree.RBRoot.Node.Left.Color
  assertBlack(t, rbTree.RBRoot.Node.Left)
  // rbTree.RBRoot.Node.Left.Key
  assertKey(t, rbTree.RBRoot.Node.Left, 7)
  // rbTree.RBRoot.Node.Right.Color
  assertBlack(t, rbTree.RBRoot.Node.Right)
  // rbTree.RBRoot.Node.Left.Key
  assertKey(t, rbTree.RBRoot.Node.Right, 13)
  
  // Level 2

  // rbTree.RBRoot.Node.Left.Left.Color
  assertRed(t, rbTree.RBRoot.Node.Left.Left)
  // rbTree.RBRoot.Node.Left.Left.Key
  assertKey(t, rbTree.RBRoot.Node.Left.Left, 5)
  // rbTree.RBRoot.Node.Left.Right.Color
  assertRed(t, rbTree.RBRoot.Node.Left.Right)
  // rbTree.RBRoot.Node.Left.Right.Key
  assertKey(t, rbTree.RBRoot.Node.Left.Right, 9)
  // rbTree.RBRoot.Node.Right.Left.Color
  assertBlack(t, rbTree.RBRoot.Node.Right.Left)
  // rbTree.RBRoot.Node.Right.Left.Key
  assertKey(t, rbTree.RBRoot.Node.Right.Left, 12)
  // rbTree.RBRoot.Node.Right.Right.Color
  assertBlack(t, rbTree.RBRoot.Node.Right.Right)
  // rbTree.RBRoot.Node.Right.Right.Key
  assertKey(t, rbTree.RBRoot.Node.Right.Right, 15)
  
  // Level 3

  // rbTree.RBRoot.Node.Right.Right.Color
  assertBlack(t, rbTree.RBRoot.Node.Left.Left.Left)
  // rbTree.RBRoot.Node.Left.Left.Left.Key
  assertKey(t, rbTree.RBRoot.Node.Left.Left.Left, 4)
  // rbTree.RBRoot.Node.Left.Left.Right.Color
  assertBlack(t, rbTree.RBRoot.Node.Left.Left.Right)
  // rbTree.RBRoot.Node.Left.Left.Right.Key
  assertKey(t, rbTree.RBRoot.Node.Left.Left.Right, 6)
  // rbTree.RBRoot.Node.Left.Right.Left.Color
  assertBlack(t, rbTree.RBRoot.Node.Left.Right.Left)
  // rbTree.RBRoot.Node.Left.Right.Left.Key
  assertKey(t, rbTree.RBRoot.Node.Left.Right.Left, 8)
  // rbTree.RBRoot.Node.Left.Right.Right.Color
  assertBlack(t, rbTree.RBRoot.Node.Left.Right.Right)
  // rbTree.RBRoot.Node.Left.Right.Left.Key
  assertKey(t, rbTree.RBRoot.Node.Left.Right.Right, 10)
  // rbTree.RBRoot.Node.Right.Right.Left.Color
  assertRed(t, rbTree.RBRoot.Node.Right.Right.Left)
  // rbTree.RBRoot.Node.Right.Right.Left.Key
  assertKey(t, rbTree.RBRoot.Node.Right.Right.Left, 14)
  // rbTree.RBRoot.Node.Right.Right.Right.Color
  assertRed(t, rbTree.RBRoot.Node.Right.Right.Right)
  // rbTree.RBRoot.Node.Right.Right.Right.Key
  assertKey(t, rbTree.RBRoot.Node.Right.Right.Right, 16)
  // rbTree.RBRoot.Node.Left.Left.Left.Left.Color
  assertRed(t, rbTree.RBRoot.Node.Left.Left.Left.Left)
  // rbTree.RBRoot.Node.Left.Left.Left.Left.Key
  assertKey(t, rbTree.RBRoot.Node.Left.Left.Left.Left, 3)
  
  // Validate leaves
  // rbTree.RBRoot.Node.Left.Left.Left.Left
  assertLeaf(t, rbTree.RBRoot.Node.Left.Left.Left.Left)
  // rbTree.RBRoot.Node.Left.Left.Right
  assertLeaf(t, rbTree.RBRoot.Node.Left.Left.Right)
  // rbTree.RBRoot.Node.Left.Right.Left
  assertLeaf(t, rbTree.RBRoot.Node.Left.Right.Left)
  // rbTree.RBRoot.Node.Left.Right.Right
  assertLeaf(t, rbTree.RBRoot.Node.Left.Right.Right)
  // rbTree.RBRoot.Node.Right.Right.Left
  assertLeaf(t, rbTree.RBRoot.Node.Right.Right.Left)
  // rbTree.RBRoot.Node.Right.Right.Right
  assertLeaf(t, rbTree.RBRoot.Node.Right.Right.Right)
  
  // Validate single child
  if rbTree.RBRoot.Node.Left.Left.Left.Right != nil {
    t.Error ("RBRoot.Node.Left.Left.Left's (node 4) Right child must be nil", )
  }
}

func TestDelete1(t *testing.T) {
  rbTree := getTestRBTree1(t)
  //                  50B
  //                 /    \
  //                /      \
  //               /        \
  //             20R        70B
  //             /  \      /  \
  //            /    \    /    \
  //          10B    30B 60B   75B
  //          /  \   / \
  //        05B 13B 25R 35B
  //                / \
  //              24B 26B

  // If we delete 75B the case is 2.2.2
  // 1. Parent is BLACK node 'a' [70B]
  // 2. If the child of 'a' is BLACK 'b' [60B]
  // 3. If 'b' has no RED child [60B has no RED child]

  // After the deletion of 75B tree looks like below [intermediate state]
  
  //                  50B
  //                 /    \
  //                /      \
  //               /        \
  //             20R        70B => h + 1 to h
  //             /  \      /  \
  //            /    \    /    \
  //          10B    30B 60R   /\
  //          /  \   / \      /_ \  => h to h - 1
  //        05B 13B 25R 35B
  //                / \
  //              24B 26B

  // In this case BLACK depth of external nodes of 70B reduced by one  (3 to 2) for other
  // parts of the tree it is still 3
  // This became case 2.1.1

  // 1. Parent is BLACK node 'a' [50B]
  // 2. 'a' has a RED child 'b' [20R]
  // 3. Right child of 'b' which must be BLACK 'c' [30B]
  // 4. 'c' has a RED child (right or left) d [25R]

  // 'd' is left child (So there is 2 rotations)


  node75 := rbTree.SearchNode(75)
  assertKey(t, node75, 75)
  assertLeaf(t, node75)
  rbTree.DeleteNode(node75)

  //     Rotate 'b' left [ First rotation ]
  //                   50B
  //                   /  \
  //                  /    \
  //                 /      \
  //                30B      70B
  //                / \      / \
  //               /   \    /   \
  //             20R   35B 60B  75B
  //             /  \    
  //            /    \   
  //          10B    25B
  //          /  \   / \
  //        05B 13B 24B 26B

  //     Rotate 'a' right [ Second and final rotation ]
  //                  30B
  //                 /    \
  //                /      \
  //               /        \
  //             20R        50B
  //             /  \      /  \  
  //            /    \    /    \
  //          10B    25B 35B   70B
  //          /  \   / \       /
  //        05B 13B 24B 26B   60B
  
    

  validate(t, rbTree)
  // Level 0
  assertKey(t, rbTree.RBRoot.Node, 30)
  assertBlack(t, rbTree.RBRoot.Node)

  // Level 1
  assertKey(t, rbTree.RBRoot.Node.Left, 20)
  assertRed(t, rbTree.RBRoot.Node.Left)
  assertKey(t, rbTree.RBRoot.Node.Right, 50)
  assertBlack(t, rbTree.RBRoot.Node.Right)

  // Level 2
  assertKey(t, rbTree.RBRoot.Node.Left.Left, 10)
  assertBlack(t, rbTree.RBRoot.Node.Left.Left)
  assertKey(t, rbTree.RBRoot.Node.Left.Right, 25)
  assertBlack(t, rbTree.RBRoot.Node.Left.Right)
  assertKey(t, rbTree.RBRoot.Node.Right.Left, 35)
  assertBlack(t, rbTree.RBRoot.Node.Right.Left)
  assertKey(t, rbTree.RBRoot.Node.Right.Right, 70)
  assertBlack(t, rbTree.RBRoot.Node.Right.Right)

  // Level 3
  assertKey(t, rbTree.RBRoot.Node.Left.Left.Left, 5)
  assertBlack(t, rbTree.RBRoot.Node.Left.Left.Left)
  assertKey(t, rbTree.RBRoot.Node.Left.Left.Right, 13)
  assertBlack(t, rbTree.RBRoot.Node.Left.Left.Right)
  assertKey(t, rbTree.RBRoot.Node.Left.Right.Left, 24)
  assertBlack(t, rbTree.RBRoot.Node.Left.Right.Left)
  assertKey(t, rbTree.RBRoot.Node.Left.Right.Right, 26)
  assertBlack(t, rbTree.RBRoot.Node.Left.Right.Right)

  assertKey(t, rbTree.RBRoot.Node.Right.Right.Left, 60)
  assertRed(t, rbTree.RBRoot.Node.Right.Right.Left)

  // Assert leaves
  assertLeaf(t, rbTree.RBRoot.Node.Left.Left.Left)   // 5
  assertLeaf(t, rbTree.RBRoot.Node.Left.Left.Right)  // 13
  assertLeaf(t, rbTree.RBRoot.Node.Left.Right.Left)  // 24
  assertLeaf(t, rbTree.RBRoot.Node.Left.Right.Right) // 26
  assertLeaf(t, rbTree.RBRoot.Node.Right.Left)       // 35
  assertLeaf(t, rbTree.RBRoot.Node.Right.Right.Left) // 60

  // Assert parents
  assertKey(t, rbTree.RBRoot.Node.Left.Parent, 30)
  assertKey(t, rbTree.RBRoot.Node.Right.Parent, 30)

  assertKey(t, rbTree.RBRoot.Node.Left.Left.Parent, 20)
  assertKey(t, rbTree.RBRoot.Node.Left.Right.Parent, 20)
  assertKey(t, rbTree.RBRoot.Node.Right.Left.Parent, 50)
  assertKey(t, rbTree.RBRoot.Node.Right.Right.Parent, 50)

  assertKey(t, rbTree.RBRoot.Node.Left.Left.Left.Parent, 10)
  assertKey(t, rbTree.RBRoot.Node.Left.Left.Right.Parent, 10)
  assertKey(t, rbTree.RBRoot.Node.Left.Right.Left.Parent, 25)
  assertKey(t, rbTree.RBRoot.Node.Left.Right.Right.Parent, 25)

  assertKey(t, rbTree.RBRoot.Node.Right.Right.Left.Parent, 70)
}

func TestDelete2(t *testing.T) {
  rbTree := getTestRBTree2(t)
  //
  //                  50B
  //                 /    \
  //                /      \
  //               /        \
  //             20R        70B
  //             /  \      /  \
  //            /    \    /    \
  //          10B    30B 60B   75B
  //          /  \   / \
  //        05B 13B 25B 35R 
  //                    / \
  //                   34B 36B
  
  // If we delete 60B the case is 2.2.2
  // 1. Parent is BLACK node 'a' [70B]
  // 2. If the other child of 'a' is BLACK 'b' [75B]
  // 3. If 'b' has no RED child [75B has no RED child]

  // After the deletion of 60B tree looks like below [intermediate state]
  
  //
  //                  50B
  //                 /    \
  //                /      \
  //               /        \
  //             20R        70B
  //             /  \      /\   \
  //            /    \    /_ \   \
  //           /      \   h->h-1  \
  //          10B     30B         75R
  //          /  \    / \
  //        05B 13B  25B 35R 
  //                     / \
  //                    34B 36B
  
  // In this case BLACK depth of external nodes of 70B reduced by one (3 to 2) for other
  // parts of the tree it is still 3
  // This became case 2.1.1

  // 1. Parent is BLACK node 'a' [50B]
  // 2. The other child of 'a' is RED, child 'b' [20R]
  // 3. Right child of 'b' which must be BLACK 'c' [30B]
  // 4. 'c' has a RED child (right or left) d [35R]

  // 'd' is right child (So there is 3 rotations)

  node60 := rbTree.SearchNode(60)
  assertKey(t, node60, 60)
  assertLeaf(t, node60)
  rbTree.DeleteNode(node60)

  //  Rotate 'c' (30B) left [ First rotation ]
  //
  //                  50B
  //                 /    \
  //                /      \
  //               /        \
  //             20R        70B
  //             /  \      /\   \
  //            /    \    /_ \   \
  //           /      \   h->h-1  \
  //          10B     35B         75R
  //          /  \    / \
  //        05B 13B  30B 36B 
  //                 / \
  //               25B 34B
  
  //  Rotate 'b' (20R) left [ Second rotation ]

  //
  //                  50B
  //                 /    \
  //                /      \
  //               /        \
  //             35R        70B
  //             /  \      /\   \
  //            /    \    /_ \   \
  //           /      \   h->h-1  \
  //          20R    36B         75R
  //          / \
  //         /   \
  //        /     \
  //       10B     30B
  //       / \     /  \
  //      /   \   /    \
  //     05B 13B  25B 34B 
  //                 
  //               
  
  //  Rotate 'a' (50B) right [ Third and last rotation ]
  
  //
  //                35R     
  //               /   \     
  //              /     \    
  //             /       \  
  //          20R         50B
  //          / \          /\
  //         /   \        /  \
  //        /     \      /    \
  //       10B     30B  36B   70B
  //       / \     /  \         \
  //      /   \   /    \         \
  //     05B 13B  25B 34B        75R
  
  validate(t, rbTree)
  // Level 0
  assertKey(t, rbTree.RBRoot.Node, 35)
  assertBlack(t, rbTree.RBRoot.Node)

  // Level 1
  assertKey(t, rbTree.RBRoot.Node.Left, 20)
  assertRed(t, rbTree.RBRoot.Node.Left)
  assertKey(t, rbTree.RBRoot.Node.Right, 50)
  assertBlack(t, rbTree.RBRoot.Node.Right)

  // Level 2
  assertKey(t, rbTree.RBRoot.Node.Left.Left, 10)
  assertBlack(t, rbTree.RBRoot.Node.Left.Left)
  assertKey(t, rbTree.RBRoot.Node.Left.Right, 30)
  assertBlack(t, rbTree.RBRoot.Node.Left.Right)
  assertKey(t, rbTree.RBRoot.Node.Right.Left, 36)
  assertBlack(t, rbTree.RBRoot.Node.Right.Left)
  assertKey(t, rbTree.RBRoot.Node.Right.Right, 70)
  assertBlack(t, rbTree.RBRoot.Node.Right.Right)

  // Level 3
  assertKey(t, rbTree.RBRoot.Node.Left.Left.Left, 5)
  assertBlack(t, rbTree.RBRoot.Node.Left.Left.Left)
  assertKey(t, rbTree.RBRoot.Node.Left.Left.Right, 13)
  assertBlack(t, rbTree.RBRoot.Node.Left.Left.Right)
  assertKey(t, rbTree.RBRoot.Node.Left.Right.Left, 25)
  assertBlack(t, rbTree.RBRoot.Node.Left.Right.Left)
  assertKey(t, rbTree.RBRoot.Node.Left.Right.Right, 34)
  assertBlack(t, rbTree.RBRoot.Node.Left.Right.Right)

  assertKey(t, rbTree.RBRoot.Node.Right.Right.Right, 75)
  assertRed(t, rbTree.RBRoot.Node.Right.Right.Right)

  // Assert leaves
  assertLeaf(t, rbTree.RBRoot.Node.Left.Left.Left)    // 5
  assertLeaf(t, rbTree.RBRoot.Node.Left.Left.Right)   // 13
  assertLeaf(t, rbTree.RBRoot.Node.Left.Right.Left)   // 25
  assertLeaf(t, rbTree.RBRoot.Node.Left.Right.Right)  // 34
  assertLeaf(t, rbTree.RBRoot.Node.Right.Left)        // 36
  assertLeaf(t, rbTree.RBRoot.Node.Right.Right.Right) // 75

  // Assert parents
  assertKey(t, rbTree.RBRoot.Node.Left.Parent, 35)
  assertKey(t, rbTree.RBRoot.Node.Right.Parent, 35)

  assertKey(t, rbTree.RBRoot.Node.Left.Left.Parent, 20)
  assertKey(t, rbTree.RBRoot.Node.Left.Right.Parent, 20)
  assertKey(t, rbTree.RBRoot.Node.Right.Left.Parent, 50)
  assertKey(t, rbTree.RBRoot.Node.Right.Right.Parent, 50)

  assertKey(t, rbTree.RBRoot.Node.Left.Left.Left.Parent, 10)
  assertKey(t, rbTree.RBRoot.Node.Left.Left.Right.Parent, 10)
  assertKey(t, rbTree.RBRoot.Node.Left.Right.Left.Parent, 30)
  assertKey(t, rbTree.RBRoot.Node.Left.Right.Right.Parent, 30)

  assertKey(t, rbTree.RBRoot.Node.Right.Right.Right.Parent, 70)
}

func TestDelete3(t *testing.T) {
  rbTree := getTestRBTree3(t)
  //
  //                  50B
  //                 /    \
  //                /      \
  //               /        \
  //             20B        70B
  //             /  \      /  \
  //            /    \    /    \
  //          10B    30R 60B   75B
  //          /  \   / \
  //        05R 13R 25B 35B
  
  // If we delete 60B the case is 2.2.2
  // 1. Parent is BLACK node 'a' [70B]
  // 2. If the other child of 'a' is BLACK 'b' [75B]
  // 3. If 'b' has no RED child [75B has no RED child]

  // After the deletion of 60B tree looks like below [intermediate state]
  
  //
  //                  50B
  //                 /    \
  //                /      \
  //               /        \
  //             20B         70B
  //             /  \        /  \
  //            /    \      /    \
  //          10B    30R   / \   75R
  //          / \    / \  / _ \
  //         /   \  /   \  h -> h-1
  //        05R 13R 25B 35B
  
  // In this case BLACK depth of external nodes of 70B reduced by one (3 to 2) for other
  // parts of the tree it is still 3
  // This became case 2.2.1

  // 1. Parent is BLACK node 'a' [50B]
  // 2. The other child of 'a' is BLACK, child 'b' [20R]
  // 3.One child (or both), here right child of 'b' is RED 'c' [30R]
  
  // 'c' is RED right child - so there is two rotations

  node60 := rbTree.SearchNode(60)
  assertKey(t, node60, 60)
  assertLeaf(t, node60)
  rbTree.DeleteNode(node60)
  validate(t, rbTree)

  //  Rotate 'b' (20B) left [ First rotation ]
  //                  50B
  //                  / \
  //                 /   \
  //                /     \
  //              30B     70B      
  //              / \       \
  //             /   \       \
  //            20B  35B     75R
  //            /  \  
  //           /    \ 
  //          10B    25B
  //          / \
  //         /   \
  //        05R 13R

  //  Rotate 'a' (50B) right [ Second and last rotation ]
  //                  30B
  //                 /   \
  //                /     \
  //               /       \
  //             20B       50B
  //             /  \      /  \
  //            /    \    /    \
  //          10B    25B  35B  70B
  //          / \                \
  //         /   \                \
  //        05R 13R               75R

  // Level 0
  assertKey(t, rbTree.RBRoot.Node, 30)
  assertBlack(t, rbTree.RBRoot.Node)

  // Level 1
  assertKey(t, rbTree.RBRoot.Node.Left, 20)
  assertBlack(t, rbTree.RBRoot.Node.Left)
  assertKey(t, rbTree.RBRoot.Node.Right, 50)
  assertBlack(t, rbTree.RBRoot.Node.Right)

  // Level 2
  assertKey(t, rbTree.RBRoot.Node.Left.Left, 10)
  assertBlack(t, rbTree.RBRoot.Node.Left.Left)
  assertKey(t, rbTree.RBRoot.Node.Left.Right, 25)
  assertBlack(t, rbTree.RBRoot.Node.Left.Right)
  assertKey(t, rbTree.RBRoot.Node.Right.Left, 35)
  assertBlack(t, rbTree.RBRoot.Node.Right.Left)
  assertKey(t, rbTree.RBRoot.Node.Right.Right, 70)
  assertBlack(t, rbTree.RBRoot.Node.Right.Right)

  // Level 3
  assertKey(t, rbTree.RBRoot.Node.Left.Left.Left, 5)
  assertRed(t, rbTree.RBRoot.Node.Left.Left.Left)
  assertKey(t, rbTree.RBRoot.Node.Left.Left.Right, 13)
  assertRed(t, rbTree.RBRoot.Node.Left.Left.Right)
  assertKey(t, rbTree.RBRoot.Node.Right.Right.Right, 75)
  assertRed(t, rbTree.RBRoot.Node.Right.Right.Right)

  // Assert leaves
  assertLeaf(t, rbTree.RBRoot.Node.Left.Left.Left)    // 5
  assertLeaf(t, rbTree.RBRoot.Node.Left.Left.Right)   // 13
  assertLeaf(t, rbTree.RBRoot.Node.Left.Right)        // 25
  assertLeaf(t, rbTree.RBRoot.Node.Right.Left)        // 35
  assertLeaf(t, rbTree.RBRoot.Node.Right.Right.Right) // 75

  // Assert parents
  assertKey(t, rbTree.RBRoot.Node.Left.Parent, 30)
  assertKey(t, rbTree.RBRoot.Node.Right.Parent, 30)
  assertKey(t, rbTree.RBRoot.Node.Left.Left.Parent, 20)
  assertKey(t, rbTree.RBRoot.Node.Left.Right.Parent, 20)
  assertKey(t, rbTree.RBRoot.Node.Right.Left.Parent, 50)
  assertKey(t, rbTree.RBRoot.Node.Right.Right.Parent, 50)

  assertKey(t, rbTree.RBRoot.Node.Left.Left.Left.Parent, 10)
  assertKey(t, rbTree.RBRoot.Node.Left.Left.Right.Parent, 10)
  assertKey(t, rbTree.RBRoot.Node.Right.Right.Right.Parent, 70)
}

func TestDelete4(t *testing.T) {
  rbTree := getTestRBTree4(t)
  //
  //                  50B
  //                 /    \
  //                /      \
  //               /        \
  //             20B        70B
  //             /  \      /  \
  //            /    \    /    \
  //          10R    30B 60B   75B
  //          /  \   / \
  //        05B 13B 25R 35R

  // If we delete 75B the case is 2.2.2
  // 1. Parent is BLACK node 'a' [70B]
  // 2. If the other child of 'a' is BLACK 'b' [60B]
  // 3. If 'b' has no RED child [60B has no RED child]

  // After the deletion of 75B tree looks like below [intermediate state]

  //                  50B
  //                 /    \
  //                /      \
  //               /        \
  //             20B        70B
  //             /  \      /  \
  //            /    \    /   /\
  //          10R    30B 60R / _\
  //          /  \   / \     h to h-1
  //        05B 13B 25R 35R

  // In this case BLACK depth of external nodes of 70B reduced by one (3 to 2) for other
  // parts of the tree it is still 3
  // This became case 2.2.1

  // 1. Parent is BLACK node 'a' [50B]
  // 2. The other child of 'a' is BLACK, child 'b' [20B]
  // 3. One child (or both), here left child of 'b' is RED 'c' [10R]

  // 'c' is RED left child - so there is one rotation
  node75 := rbTree.SearchNode(75)
  assertKey(t, node75, 75)
  assertLeaf(t, node75)
  rbTree.DeleteNode(node75)
  validate(t, rbTree)

  //  Rotate 'a' (50B) right [ First and last rotation ]
  //                  20B
  //                 /    \
  //                /      \
  //               /        \
  //             10B        50B
  //             /  \       /    \
  //            /    \     /      \
  //          05B    13B  30B      70B
  //                     /  \      /  \
  //                    /    \    /   /\
  //                   25R   35R 60R /_ \ h to h-1

  // Level 0
  assertKey(t, rbTree.RBRoot.Node, 20)
  assertBlack(t, rbTree.RBRoot.Node)

  // Level 1
  assertKey(t, rbTree.RBRoot.Node.Left, 10)
  assertBlack(t, rbTree.RBRoot.Node.Left)
  assertKey(t, rbTree.RBRoot.Node.Right, 50)
  assertBlack(t, rbTree.RBRoot.Node.Right)

  // Level 2
  assertKey(t, rbTree.RBRoot.Node.Left.Left, 5)
  assertBlack(t, rbTree.RBRoot.Node.Left.Left)
  assertKey(t, rbTree.RBRoot.Node.Left.Right, 13)
  assertBlack(t, rbTree.RBRoot.Node.Left.Right)
  assertKey(t, rbTree.RBRoot.Node.Right.Left, 30)
  assertBlack(t, rbTree.RBRoot.Node.Right.Left)
  assertKey(t, rbTree.RBRoot.Node.Right.Right, 70)
  assertBlack(t, rbTree.RBRoot.Node.Right.Right)

  // Level 3
  assertKey(t, rbTree.RBRoot.Node.Right.Left.Left, 25)
  assertRed(t, rbTree.RBRoot.Node.Right.Left.Left)
  assertKey(t, rbTree.RBRoot.Node.Right.Left.Right, 35)
  assertRed(t, rbTree.RBRoot.Node.Right.Left.Right)
  assertKey(t, rbTree.RBRoot.Node.Right.Right.Left, 60)
  assertRed(t, rbTree.RBRoot.Node.Right.Right.Left)

  // Assert leaves
  assertLeaf(t, rbTree.RBRoot.Node.Left.Left)       // 5
  assertLeaf(t, rbTree.RBRoot.Node.Left.Right)      // 13
  assertLeaf(t, rbTree.RBRoot.Node.Right.Left.Left) // 25
  assertLeaf(t, rbTree.RBRoot.Node.Right.Left.Right)// 35
  assertLeaf(t, rbTree.RBRoot.Node.Right.Right.Left)// 60

  // Assert parents
  assertKey(t, rbTree.RBRoot.Node.Left.Parent, 20)
  assertKey(t, rbTree.RBRoot.Node.Right.Parent, 20)
  assertKey(t, rbTree.RBRoot.Node.Left.Left.Parent, 10)
  assertKey(t, rbTree.RBRoot.Node.Left.Right.Parent, 10)
  assertKey(t, rbTree.RBRoot.Node.Right.Left.Parent, 50)
  assertKey(t, rbTree.RBRoot.Node.Right.Right.Parent, 50)

  assertKey(t, rbTree.RBRoot.Node.Right.Left.Left.Parent, 30)
  assertKey(t, rbTree.RBRoot.Node.Right.Left.Right.Parent, 30)
  assertKey(t, rbTree.RBRoot.Node.Right.Right.Left.Parent, 70)
}

func TestDelet5(t *testing.T) {
  rbTree := getTestRBTree5(t)
  //
  //                  50B
  //                 /    \
  //                /      \
  //               /        \
  //             20R        70B
  //             /  \      /  \
  //            /    \    /    \
  //          10B    30B 60B   75B
  //          /  \   / \
  //        05B 13B 25B 35B

  // If we delete 75B the case is 2.2.2
  // 1. Parent is BLACK node 'a' [70B]
  // 2. If the other child of 'a' is BLACK 'b' [60B]
  // 3. If 'b' has no RED child [60B has no RED child]

  // After the deletion of 75B tree looks like below [intermediate state]

  //                  50B
  //                 /    \
  //                /      \
  //               /        \
  //             20R        70B
  //             /  \      /  \
  //            /    \    /   /\
  //          10B    30B 60R / _\
  //          /  \   / \     h to h-1
  //        05B 13B 25B 35B

  // In this case BLACK depth of external nodes of 70B reduced by one (3 to 2) for other
  // parts of the tree it is still 3
  // This became case 2.1.2

  // 1. Parent is BLACK node 'a' [50B]
  // 2. The other child of 'a' is RED, child 'b' [20R]
  // 3. The right child of 'b' must be BLACK 'c' [30B]
  // 4. If 'c' has no RED child

  // 'c' has no RED child - so there is one rotation
  node75 := rbTree.SearchNode(75)
  assertKey(t, node75, 75)
  assertLeaf(t, node75)
  rbTree.DeleteNode(node75)
  validate(t, rbTree)

  //  Rotate 'a' (50B) right [ First and last rotation ]
  //             20B
  //             /   \
  //            /     \
  //          10B      50B
  //          / \      / \
  //         /   \    /   \
  //        05B 13B  30R   70B
  //                 / \   /
  //               25B 35B 60R

  // Level 0
  assertKey(t, rbTree.RBRoot.Node, 20)
  assertBlack(t, rbTree.RBRoot.Node)

  // Level 1
  assertKey(t, rbTree.RBRoot.Node.Left, 10)
  assertBlack(t, rbTree.RBRoot.Node.Left)
  assertKey(t, rbTree.RBRoot.Node.Right, 50)
  assertBlack(t, rbTree.RBRoot.Node.Right)

  // Level 2
  assertKey(t, rbTree.RBRoot.Node.Left.Left, 5)
  assertBlack(t, rbTree.RBRoot.Node.Left.Left)
  assertKey(t, rbTree.RBRoot.Node.Left.Right, 13)
  assertBlack(t, rbTree.RBRoot.Node.Left.Right)
  assertKey(t, rbTree.RBRoot.Node.Right.Left, 30)
  assertRed(t, rbTree.RBRoot.Node.Right.Left)
  assertKey(t, rbTree.RBRoot.Node.Right.Right, 70)
  assertBlack(t, rbTree.RBRoot.Node.Right.Right)

  // Level 3
  assertKey(t, rbTree.RBRoot.Node.Right.Left.Left, 25)
  assertBlack(t, rbTree.RBRoot.Node.Right.Left.Left)
  assertKey(t, rbTree.RBRoot.Node.Right.Left.Right, 35)
  assertBlack(t, rbTree.RBRoot.Node.Right.Left.Right)
  assertKey(t, rbTree.RBRoot.Node.Right.Right.Left, 60)
  assertRed(t, rbTree.RBRoot.Node.Right.Right.Left)

  // Assert leaves
  assertLeaf(t, rbTree.RBRoot.Node.Left.Left)       // 5
  assertLeaf(t, rbTree.RBRoot.Node.Left.Right)      // 13
  assertLeaf(t, rbTree.RBRoot.Node.Right.Left.Left) // 25
  assertLeaf(t, rbTree.RBRoot.Node.Right.Left.Right)// 35
  assertLeaf(t, rbTree.RBRoot.Node.Right.Right.Left)// 60

  // Assert parents
  assertKey(t, rbTree.RBRoot.Node.Left.Parent, 20)
  assertKey(t, rbTree.RBRoot.Node.Right.Parent, 20)
  assertKey(t, rbTree.RBRoot.Node.Left.Left.Parent, 10)
  assertKey(t, rbTree.RBRoot.Node.Left.Right.Parent, 10)
  assertKey(t, rbTree.RBRoot.Node.Right.Left.Parent, 50)
  assertKey(t, rbTree.RBRoot.Node.Right.Right.Parent, 50)

  assertKey(t, rbTree.RBRoot.Node.Right.Left.Left.Parent, 30)
  assertKey(t, rbTree.RBRoot.Node.Right.Left.Right.Parent, 30)
  assertKey(t, rbTree.RBRoot.Node.Right.Right.Left.Parent, 70)
}

func TestDelet6(t *testing.T) {
  rbTree := getTestRBTree6(t)
  //
  //                  50B
  //                 /    \
  //                /      \
  //               /        \
  //             20R        70B
  //             /  \      /  \
  //            /    \    /    \
  //          10B    30B 60B   75B
  //          /  \   / \
  //        05B 13R 25B 35B
  //            / \
  //           /   \
  //          11B  14B

  // If we delete 25B the case is 2.2.2
  // 1. Parent is BLACK node 'a' [30B]
  // 2. If the other child of 'a' is BLACK 'b' [35B]
  // 3. If 'b' has no RED child [35B has no RED child]

  // After the deletion of 25B tree looks like below [intermediate state]

  //                  50B
  //                 /    \
  //                /      \
  //               /        \
  //             20R         70B
  //             /  \        /  \
  //            /    \      /    \
  //          10B     30B   60B   75B
  //          /  \    / \
  //        05B 13R  /\  35R
  //            / \ /_ \
  //           /   \ h ->h-1
  //          11B  14B

  // In this case BLACK depth of external nodes of 30B reduced by one (3 to 2) for other
  // parts of the tree it is still 3
  // This became case 1.1

  // 1. Parent is RED node 'a' [20B]
  // 2. The other child of 'b' must be RED, child 'b' [10B]
  // 3. 'b' has one RED child (can have 0, 1, or 2 here 1) the right child of 'b' is RED 'c' [13R]

  // There are two rotations
  node25 := rbTree.SearchNode(25)
  assertKey(t, node25, 25)
  assertLeaf(t, node25)
  rbTree.DeleteNode(node25)
  validate(t, rbTree)

  //  Rotate 'b' (10B) left [ First rotation ]
  //                  50B
  //                 /    \
  //                /      \
  //               /        \
  //             20B         70B
  //             /  \        /  \
  //            /    \      /    \
  //          13R     30B   60B   75B
  //          /  \    / \
  //        10B 14B  /\  35R
  //        / \     /_ \
  //       /   \    h ->h-1
  //      05B  11B

  //  Rotate 'a' (20B) right [ Second and last rotation ]
  //                  50B
  //                 /    \
  //                /      \
  //               /        \
  //             13R         70B
  //             /  \        /  \
  //            /    \      /    \
  //          10B     20B   60B   75B
  //          /  \    / \
  //        05B 11B  14B  30B
  //                      / \
  //                     /\  35R
  //                    /_ \
  //                   h ->h-1

  // Level 0 (no change)
  assertKey(t, rbTree.RBRoot.Node, 50)
  assertBlack(t, rbTree.RBRoot.Node)

  // Level 1
  assertKey(t, rbTree.RBRoot.Node.Left, 13)
  assertRed(t, rbTree.RBRoot.Node.Left)
  assertKey(t, rbTree.RBRoot.Node.Right,70) // (no change)
  assertBlack(t, rbTree.RBRoot.Node.Right)

  // Level 2
  assertKey(t, rbTree.RBRoot.Node.Left.Left, 10)
  assertBlack(t, rbTree.RBRoot.Node.Left.Left)
  assertKey(t, rbTree.RBRoot.Node.Left.Right, 20)
  assertBlack(t, rbTree.RBRoot.Node.Left.Right)
  assertKey(t, rbTree.RBRoot.Node.Right.Left,60)
  assertBlack(t, rbTree.RBRoot.Node.Right.Left)
  assertKey(t, rbTree.RBRoot.Node.Right.Right, 75)
  assertBlack(t, rbTree.RBRoot.Node.Right.Right)

  // Level 3
  assertKey(t, rbTree.RBRoot.Node.Left.Left.Left, 5)
  assertBlack(t, rbTree.RBRoot.Node.Left.Left.Left)
  assertKey(t, rbTree.RBRoot.Node.Left.Left.Right, 11)
  assertBlack(t, rbTree.RBRoot.Node.Left.Left.Right)
  assertKey(t, rbTree.RBRoot.Node.Left.Right.Left, 14)
  assertBlack(t, rbTree.RBRoot.Node.Left.Right.Left)
  assertKey(t, rbTree.RBRoot.Node.Left.Right.Right, 30)
  assertBlack(t, rbTree.RBRoot.Node.Left.Right.Right)

  // Level 4
  assertKey(t, rbTree.RBRoot.Node.Left.Right.Right.Right, 35)
  assertRed(t, rbTree.RBRoot.Node.Left.Right.Right.Right)

  // Assert leaves
  assertLeaf(t, rbTree.RBRoot.Node.Left.Left.Left)         // 5
  assertLeaf(t, rbTree.RBRoot.Node.Left.Left.Right)        // 11
  assertLeaf(t, rbTree.RBRoot.Node.Left.Right.Left)        // 14
  assertLeaf(t, rbTree.RBRoot.Node.Left.Right.Right.Right) // 35
  assertLeaf(t, rbTree.RBRoot.Node.Right.Left)             // 60
  assertLeaf(t, rbTree.RBRoot.Node.Right.Right)            // 75

  // Assert parents
  assertKey(t, rbTree.RBRoot.Node.Left.Parent, 50)
  assertKey(t, rbTree.RBRoot.Node.Right.Parent, 50)
  assertKey(t, rbTree.RBRoot.Node.Left.Left.Parent, 13)
  assertKey(t, rbTree.RBRoot.Node.Left.Right.Parent, 13)
  assertKey(t, rbTree.RBRoot.Node.Right.Left.Parent, 70)
  assertKey(t, rbTree.RBRoot.Node.Right.Right.Parent, 70)

  assertKey(t, rbTree.RBRoot.Node.Left.Left.Left.Parent, 10)
  assertKey(t, rbTree.RBRoot.Node.Left.Left.Right.Parent, 10)
  assertKey(t, rbTree.RBRoot.Node.Left.Right.Left.Parent, 20)
  assertKey(t, rbTree.RBRoot.Node.Left.Right.Right.Parent, 20)

  assertKey(t, rbTree.RBRoot.Node.Left.Right.Right.Right.Parent, 30)
 }

func TestDelet7(t *testing.T) {
  rbTree := getTestRBTree7(t)
  //
  //                  50B
  //                 /    \
  //                /      \
  //               /        \
  //             20R        70B
  //             /  \      /  \
  //            /    \    /    \
  //          10B    30B 60B   75B
  //          /  \   / \
  //        05R 13B 25B 35B
  //        / \
  //       /   \
  //     04B   06B

  // If we delete 25B the case is 2.2.2
  // 1. Parent is BLACK node 'a' [30B]
  // 2. If the other child of 'a' is BLACK 'b' [35B]
  // 3. If 'b' has no RED child [35B has no RED child]

  // After the deletion of 25B tree looks like below [intermediate state]

  //                  50B
  //                 /    \
  //                /      \
  //               /        \
  //             20R         70B
  //             /  \        /  \
  //            /    \      /    \
  //          10B     30B   60B   75B
  //          /  \    / \
  //        05R 13B  /\  35R
  //        / \     /_ \
  //       /   \    h ->h-1
  //     04B  06B

  // In this case BLACK depth of external nodes of 30B reduced by one (3 to 2) for other
  // parts of the tree it is still 3
  // This became case 1.1

  // 1. Parent is RED node 'a' [20B]
  // 2. The other child of 'b' must be RED, child 'b' [10B]
  // 3. 'b' has one RED child (can have 0, 1, or 2 here 1) the left child of 'b' is RED 'c' [05R]

  // There is one rotation
  node25 := rbTree.SearchNode(25)
  assertKey(t, node25, 25)
  assertLeaf(t, node25)
  rbTree.DeleteNode(node25)
  validate(t, rbTree)

  // Rotate 'a' (20R) right [First and last rotation]
  //                  50B
  //                 /    \
  //                /      \
  //               /        \
  //             10R         70B
  //             /  \        /  \
  //            /    \      /    \
  //          05B     20B   60B   75B
  //          /  \    /  \
  //        04B 06B  13B  30B
  //                        \
  //                        35R

  // Level 0 (no change)
  assertKey(t, rbTree.RBRoot.Node, 50)
  assertBlack(t, rbTree.RBRoot.Node)

  // Level 1
  assertKey(t, rbTree.RBRoot.Node.Left, 10)
  assertRed(t, rbTree.RBRoot.Node.Left)
  assertKey(t, rbTree.RBRoot.Node.Right,70) // (no change)
  assertBlack(t, rbTree.RBRoot.Node.Right)

  // Level 2
  assertKey(t, rbTree.RBRoot.Node.Left.Left, 5)
  assertBlack(t, rbTree.RBRoot.Node.Left.Left)
  assertKey(t, rbTree.RBRoot.Node.Left.Right, 20)
  assertBlack(t, rbTree.RBRoot.Node.Left.Right)
  assertKey(t, rbTree.RBRoot.Node.Right.Left,60)
  assertBlack(t, rbTree.RBRoot.Node.Right.Left)
  assertKey(t, rbTree.RBRoot.Node.Right.Right, 75)
  assertBlack(t, rbTree.RBRoot.Node.Right.Right)

  // Level 3
  assertKey(t, rbTree.RBRoot.Node.Left.Left.Left, 4)
  assertBlack(t, rbTree.RBRoot.Node.Left.Left.Left)
  assertKey(t, rbTree.RBRoot.Node.Left.Left.Right, 6)
  assertBlack(t, rbTree.RBRoot.Node.Left.Left.Right)
  assertKey(t, rbTree.RBRoot.Node.Left.Right.Left, 13)
  assertBlack(t, rbTree.RBRoot.Node.Left.Right.Left)
  assertKey(t, rbTree.RBRoot.Node.Left.Right.Right, 30)
  assertBlack(t, rbTree.RBRoot.Node.Left.Right.Right)

  // Level 4
  assertKey(t, rbTree.RBRoot.Node.Left.Right.Right.Right, 35)
  assertRed(t, rbTree.RBRoot.Node.Left.Right.Right.Right)

  // Assert leaves
  assertLeaf(t, rbTree.RBRoot.Node.Left.Left.Left)         // 4
  assertLeaf(t, rbTree.RBRoot.Node.Left.Left.Right)        // 6
  assertLeaf(t, rbTree.RBRoot.Node.Left.Right.Left)        // 13
  assertLeaf(t, rbTree.RBRoot.Node.Left.Right.Right.Right) // 35
  assertLeaf(t, rbTree.RBRoot.Node.Right.Left)             // 60
  assertLeaf(t, rbTree.RBRoot.Node.Right.Right)            // 75

  assertKey(t, rbTree.RBRoot.Node.Left.Parent, 50)
  assertKey(t, rbTree.RBRoot.Node.Right.Parent, 50)
  assertKey(t, rbTree.RBRoot.Node.Left.Left.Parent, 10)
  assertKey(t, rbTree.RBRoot.Node.Left.Right.Parent, 10)
  assertKey(t, rbTree.RBRoot.Node.Right.Left.Parent, 70)
  assertKey(t, rbTree.RBRoot.Node.Right.Right.Parent, 70)
  assertKey(t, rbTree.RBRoot.Node.Left.Left.Left.Parent, 5)
  assertKey(t, rbTree.RBRoot.Node.Left.Left.Right.Parent, 5)
  assertKey(t, rbTree.RBRoot.Node.Left.Right.Left.Parent, 20)
  assertKey(t, rbTree.RBRoot.Node.Left.Right.Right.Parent, 20)

  assertKey(t, rbTree.RBRoot.Node.Left.Right.Right.Right.Parent, 30)
 }

func TestDelet8(t *testing.T) {
  rbTree := getTestRBTree8(t)
  //
  //                  50B
  //                 /    \
  //                /      \
  //               /        \
  //             20R        70B
  //             /  \      /  \
  //            /    \    /    \
  //          10B    30B 60B   75B
  //          /  \   / \
  //        05B 13B 25B 35B

  // If we delete 25B the case is 2.2.2
  // 1. Parent is BLACK node 'a' [30B]
  // 2. If the other child of 'a' is BLACK 'b' [35B]
  // 3. If 'b' has no RED child [35B has no RED child]

  // After the deletion of 25B tree looks like below [intermediate state]

  //                  50B
  //                 /    \
  //                /      \
  //               /        \
  //             20R         70B
  //             /  \        /  \
  //            /    \      /    \
  //          10B     30B   60B   75B
  //          /  \    / \
  //        05B 13B  /\  35R
  //                /_ \
  //               h ->h-1

  // In this case BLACK depth of external nodes of 30B reduced by one (3 to 2) for other
  // parts of the tree it is still 3
  // This became case 1.2

  // 1. Parent is RED node 'a' [20B]
  // 2. The other child of 'b' must be BLACK, child 'b' [10B]
  // 3. 'b' has no RED child (both 05 and 13 are BLACK)

  // There is one rotation
  node25 := rbTree.SearchNode(25)
  assertKey(t, node25, 25)
  assertLeaf(t, node25)
  rbTree.DeleteNode(node25)
  validate(t, rbTree)

  // Simple recolouring will fix ( 20R -> 20B, 10B -> 10R)
  //                  50B
  //                 /    \
  //                /      \
  //               /        \
  //             20B         70B
  //             /  \        /  \
  //            /    \      /    \
  //          10R     30B   60B   75B
  //          /  \    / \
  //        05B 13B  /\  35R
  //                /_ \
  //               h ->h-1

  // Level 0 (no change)
  assertKey(t, rbTree.RBRoot.Node, 50)
  assertBlack(t, rbTree.RBRoot.Node)

  // Level 1
  assertKey(t, rbTree.RBRoot.Node.Left, 20)
  assertBlack(t, rbTree.RBRoot.Node.Left)
  assertKey(t, rbTree.RBRoot.Node.Right,70) // (no change)
  assertBlack(t, rbTree.RBRoot.Node.Right)

  // Level 2
  assertKey(t, rbTree.RBRoot.Node.Left.Left, 10)
  assertRed(t, rbTree.RBRoot.Node.Left.Left)
  assertKey(t, rbTree.RBRoot.Node.Left.Right, 30)
  assertBlack(t, rbTree.RBRoot.Node.Left.Right)
  assertKey(t, rbTree.RBRoot.Node.Right.Left,60)
  assertBlack(t, rbTree.RBRoot.Node.Right.Left)
  assertKey(t, rbTree.RBRoot.Node.Right.Right, 75)
  assertBlack(t, rbTree.RBRoot.Node.Right.Right)

  // Level 3
  assertKey(t, rbTree.RBRoot.Node.Left.Left.Left, 5)
  assertBlack(t, rbTree.RBRoot.Node.Left.Left.Left)
  assertKey(t, rbTree.RBRoot.Node.Left.Left.Right, 13)
  assertBlack(t, rbTree.RBRoot.Node.Left.Left.Right)
  assertKey(t, rbTree.RBRoot.Node.Left.Right.Right, 35)
  assertRed(t, rbTree.RBRoot.Node.Left.Right.Right)

  //                  50B
  //                 /    \
  //                /      \
  //               /        \
  //             20B         70B
  //             /  \        /  \
  //            /    \      /    \
  //          10R     30B   60B   75B
  //          /  \    / \
  //        05B 13B  /\  35R
  //                /_ \
  //               h ->h-1

  // Assert leaves
  assertLeaf(t, rbTree.RBRoot.Node.Left.Left.Left)          // 5
  assertLeaf(t, rbTree.RBRoot.Node.Left.Left.Right)         // 13
  assertLeaf(t, rbTree.RBRoot.Node.Left.Right.Right)        // 35
  assertLeaf(t, rbTree.RBRoot.Node.Right.Left)              // 60
  assertLeaf(t, rbTree.RBRoot.Node.Right.Right)             // 75

  // Assert parents
  assertKey(t, rbTree.RBRoot.Node.Left.Parent, 50)
  assertKey(t, rbTree.RBRoot.Node.Right.Parent, 50)
  assertKey(t, rbTree.RBRoot.Node.Left.Left.Parent, 20)
  assertKey(t, rbTree.RBRoot.Node.Left.Right.Parent, 20)
  assertKey(t, rbTree.RBRoot.Node.Right.Left.Parent, 70)
  assertKey(t, rbTree.RBRoot.Node.Right.Right.Parent, 70)
  assertKey(t, rbTree.RBRoot.Node.Left.Left.Left.Parent, 10)
  assertKey(t, rbTree.RBRoot.Node.Left.Left.Right.Parent, 10)
  assertKey(t, rbTree.RBRoot.Node.Left.Right.Right.Parent, 30)
 }

//function 'getTestRBTree1' returns below tree
//
//                  50B
//                 /    \
//                /      \
//               /        \
//             20R        70B
//             /  \      /  \
//            /    \    /    \
//          10B    30B 60B   75B
//          /  \   / \
//        05B 13B 25R 35B
//                / \
//              24B 26B
func getTestRBTree1(t *testing.T) *rbtree.RBTree {
  // Level 0 [50B]
  rbTree := rbtree.CreateNewRBTree(IntKeyComparator)
  rbTree.RBRoot.Node = rbtree.CreateNewRBNode(50, 50)
  rbTree.RBRoot.Node.Color = true
  rbTree.RBRoot.Node.Parent = nil

  // Level 1 [20R 70B]
  rbTree.RBRoot.Node.Left = rbtree.CreateNewRBNode(20, 20)
  rbTree.RBRoot.Node.Left.Parent = rbTree.RBRoot.Node
  rbTree.RBRoot.Node.Right = rbtree.CreateNewRBNode(70, 70)
  rbTree.RBRoot.Node.Right.Parent = rbTree.RBRoot.Node
  rbTree.RBRoot.Node.Right.Color = true
  
  // Level 2 [10B 30B 60B 75B]
  rbTree.RBRoot.Node.Left.Left  = rbtree.CreateNewRBNode(10, 10)
  rbTree.RBRoot.Node.Left.Left.Color = true
  rbTree.RBRoot.Node.Left.Left.Parent = rbTree.RBRoot.Node.Left
  rbTree.RBRoot.Node.Left.Right  = rbtree.CreateNewRBNode(30, 30)
  rbTree.RBRoot.Node.Left.Right.Color = true
  rbTree.RBRoot.Node.Left.Right.Parent = rbTree.RBRoot.Node.Left
  rbTree.RBRoot.Node.Right.Left = rbtree.CreateNewRBNode(60, 60)
  rbTree.RBRoot.Node.Right.Left.Color = true
  rbTree.RBRoot.Node.Right.Left.Parent = rbTree.RBRoot.Node.Right
  rbTree.RBRoot.Node.Right.Right = rbtree.CreateNewRBNode(75, 75)
  rbTree.RBRoot.Node.Right.Right.Color = true
  rbTree.RBRoot.Node.Right.Right.Parent = rbTree.RBRoot.Node.Right

  // Level 3 [5B 13B 25R 35B]
  rbTree.RBRoot.Node.Left.Left.Left = rbtree.CreateNewRBNode(5, 5)
  rbTree.RBRoot.Node.Left.Left.Left.Color = true
  rbTree.RBRoot.Node.Left.Left.Left.Parent = rbTree.RBRoot.Node.Left.Left
  rbTree.RBRoot.Node.Left.Left.Right = rbtree.CreateNewRBNode(13, 13)
  rbTree.RBRoot.Node.Left.Left.Right.Color = true
  rbTree.RBRoot.Node.Left.Left.Right.Parent = rbTree.RBRoot.Node.Left.Left
  rbTree.RBRoot.Node.Left.Right.Left = rbtree.CreateNewRBNode(25, 25)
  rbTree.RBRoot.Node.Left.Right.Left.Parent = rbTree.RBRoot.Node.Left.Right
  rbTree.RBRoot.Node.Left.Right.Right = rbtree.CreateNewRBNode(35, 35)
  rbTree.RBRoot.Node.Left.Right.Right.Color = true
  rbTree.RBRoot.Node.Left.Right.Right.Parent = rbTree.RBRoot.Node.Left.Right
  
  // Level 4 [24B 26B]
  rbTree.RBRoot.Node.Left.Right.Left.Left = rbtree.CreateNewRBNode(24, 24)
  rbTree.RBRoot.Node.Left.Right.Left.Left.Color = true
  rbTree.RBRoot.Node.Left.Right.Left.Left.Parent = rbTree.RBRoot.Node.Left.Right.Left
  rbTree.RBRoot.Node.Left.Right.Left.Right = rbtree.CreateNewRBNode(26, 26)
  rbTree.RBRoot.Node.Left.Right.Left.Right.Color = true
  rbTree.RBRoot.Node.Left.Right.Left.Right.Parent = rbTree.RBRoot.Node.Left.Right.Left
  
  validate(t, rbTree)
  return rbTree
}

// function 'getTestRBTree2' returns below tree
//
//                  50B
//                 /    \
//                /      \
//               /        \
//             20R        70B
//             /  \      /  \
//            /    \    /    \
//          10B    30B 60B   75B
//          /  \   / \
//        05B 13B 25B 35R 
//                    / \
//                   34B 36B
func getTestRBTree2(t *testing.T) *rbtree.RBTree {
  // Level 0 [50B]
  rbTree := rbtree.CreateNewRBTree(IntKeyComparator)
  rbTree.RBRoot.Node = rbtree.CreateNewRBNode(50, 50)
  rbTree.RBRoot.Node.Color = true
  rbTree.RBRoot.Node.Parent = nil

  // Level 1 [20R 70B]
  rbTree.RBRoot.Node.Left = rbtree.CreateNewRBNode(20, 20)
  rbTree.RBRoot.Node.Left.Parent = rbTree.RBRoot.Node
  rbTree.RBRoot.Node.Right = rbtree.CreateNewRBNode(70, 70)
  rbTree.RBRoot.Node.Right.Parent = rbTree.RBRoot.Node
  rbTree.RBRoot.Node.Right.Color = true
  
  // Level 2 [10B 30B 60B 75B]
  rbTree.RBRoot.Node.Left.Left  = rbtree.CreateNewRBNode(10, 10)
  rbTree.RBRoot.Node.Left.Left.Color = true
  rbTree.RBRoot.Node.Left.Left.Parent = rbTree.RBRoot.Node.Left
  rbTree.RBRoot.Node.Left.Right  = rbtree.CreateNewRBNode(30, 30)
  rbTree.RBRoot.Node.Left.Right.Color = true
  rbTree.RBRoot.Node.Left.Right.Parent = rbTree.RBRoot.Node.Left
  rbTree.RBRoot.Node.Right.Left = rbtree.CreateNewRBNode(60, 60)
  rbTree.RBRoot.Node.Right.Left.Color = true
  rbTree.RBRoot.Node.Right.Left.Parent = rbTree.RBRoot.Node.Right
  rbTree.RBRoot.Node.Right.Right = rbtree.CreateNewRBNode(75, 75)
  rbTree.RBRoot.Node.Right.Right.Color = true
  rbTree.RBRoot.Node.Right.Right.Parent = rbTree.RBRoot.Node.Right

  // Level 3 [5B 13B 25R 35B]
  rbTree.RBRoot.Node.Left.Left.Left = rbtree.CreateNewRBNode(5, 5)
  rbTree.RBRoot.Node.Left.Left.Left.Color = true
  rbTree.RBRoot.Node.Left.Left.Left.Parent = rbTree.RBRoot.Node.Left.Left
  rbTree.RBRoot.Node.Left.Left.Right = rbtree.CreateNewRBNode(13, 13)
  rbTree.RBRoot.Node.Left.Left.Right.Color = true
  rbTree.RBRoot.Node.Left.Left.Right.Parent = rbTree.RBRoot.Node.Left.Left
  rbTree.RBRoot.Node.Left.Right.Left = rbtree.CreateNewRBNode(25, 25)
  rbTree.RBRoot.Node.Left.Right.Left.Color = true
  rbTree.RBRoot.Node.Left.Right.Left.Parent = rbTree.RBRoot.Node.Left.Right
  rbTree.RBRoot.Node.Left.Right.Right = rbtree.CreateNewRBNode(35, 35)
  rbTree.RBRoot.Node.Left.Right.Right.Parent = rbTree.RBRoot.Node.Left.Right

  // Level 4 [34B 36B]
  rbTree.RBRoot.Node.Left.Right.Right.Left = rbtree.CreateNewRBNode(34, 34)
  rbTree.RBRoot.Node.Left.Right.Right.Left.Color = true
  rbTree.RBRoot.Node.Left.Right.Right.Left.Parent = rbTree.RBRoot.Node.Left.Right.Right
  rbTree.RBRoot.Node.Left.Right.Right.Right = rbtree.CreateNewRBNode(36, 36)
  rbTree.RBRoot.Node.Left.Right.Right.Right.Color = true
  rbTree.RBRoot.Node.Left.Right.Right.Right.Parent = rbTree.RBRoot.Node.Left.Right.Right
  
  validate(t, rbTree)
  return rbTree
}

//function 'getTestRBTree3' returns below tree
//
//                  50B
//                 /    \
//                /      \
//               /        \
//             20B        70B
//             /  \      /  \
//            /    \    /    \
//          10B    30R 60B   75B
//          /  \   / \
//        05R 13R 25B 35B
//
func getTestRBTree3(t *testing.T) *rbtree.RBTree {
  // Level 0 [50B]
  rbTree := rbtree.CreateNewRBTree(IntKeyComparator)
  rbTree.RBRoot.Node = rbtree.CreateNewRBNode(50, 50)
  rbTree.RBRoot.Node.Color = true
  rbTree.RBRoot.Node.Parent = nil

  // Level 1 [20R 70B]
  rbTree.RBRoot.Node.Left = rbtree.CreateNewRBNode(20, 20)
  rbTree.RBRoot.Node.Left.Color = true
  rbTree.RBRoot.Node.Left.Parent = rbTree.RBRoot.Node
  rbTree.RBRoot.Node.Right = rbtree.CreateNewRBNode(70, 70)
  rbTree.RBRoot.Node.Right.Color = true
  rbTree.RBRoot.Node.Right.Parent = rbTree.RBRoot.Node
  
  // Level 2 [10B 30R 60B 75B]
  rbTree.RBRoot.Node.Left.Left  = rbtree.CreateNewRBNode(10, 10)
  rbTree.RBRoot.Node.Left.Left.Color = true
  rbTree.RBRoot.Node.Left.Left.Parent = rbTree.RBRoot.Node.Left
  rbTree.RBRoot.Node.Left.Right  = rbtree.CreateNewRBNode(30, 30)
  rbTree.RBRoot.Node.Left.Right.Parent = rbTree.RBRoot.Node.Left
  rbTree.RBRoot.Node.Right.Left = rbtree.CreateNewRBNode(60, 60)
  rbTree.RBRoot.Node.Right.Left.Color = true
  rbTree.RBRoot.Node.Right.Left.Parent = rbTree.RBRoot.Node.Right
  rbTree.RBRoot.Node.Right.Right = rbtree.CreateNewRBNode(75, 75)
  rbTree.RBRoot.Node.Right.Right.Color = true
  rbTree.RBRoot.Node.Right.Right.Parent = rbTree.RBRoot.Node.Right

  // Level 3 [5R 13R 25B 35B]
  rbTree.RBRoot.Node.Left.Left.Left = rbtree.CreateNewRBNode(5, 5)
  rbTree.RBRoot.Node.Left.Left.Left.Parent = rbTree.RBRoot.Node.Left.Left
  rbTree.RBRoot.Node.Left.Left.Right = rbtree.CreateNewRBNode(13, 13)
  rbTree.RBRoot.Node.Left.Left.Right.Parent = rbTree.RBRoot.Node.Left.Left
  rbTree.RBRoot.Node.Left.Right.Left = rbtree.CreateNewRBNode(25, 25)
  rbTree.RBRoot.Node.Left.Right.Left.Color = true
  rbTree.RBRoot.Node.Left.Right.Left.Parent = rbTree.RBRoot.Node.Left.Right
  rbTree.RBRoot.Node.Left.Right.Right = rbtree.CreateNewRBNode(35, 35)
  rbTree.RBRoot.Node.Left.Right.Right.Color = true
  rbTree.RBRoot.Node.Left.Right.Right.Parent = rbTree.RBRoot.Node.Left.Right
  
  validate(t, rbTree)
  return rbTree
}

//function 'getTestRBTree4' returns below tree
//
//                  50B
//                 /    \
//                /      \
//               /        \
//             20B        70B
//             /  \      /  \
//            /    \    /    \
//          10R    30B 60B   75B
//          /  \   / \
//        05B 13B 25R 35R
//
func getTestRBTree4(t *testing.T) *rbtree.RBTree {
  // Level 0 [50B]
  rbTree := rbtree.CreateNewRBTree(IntKeyComparator)
  rbTree.RBRoot.Node = rbtree.CreateNewRBNode(50, 50)
  rbTree.RBRoot.Node.Color = true
  rbTree.RBRoot.Node.Parent = nil

  // Level 1 [20B 70B]
  rbTree.RBRoot.Node.Left = rbtree.CreateNewRBNode(20, 20)
  rbTree.RBRoot.Node.Left.Color = true
  rbTree.RBRoot.Node.Left.Parent = rbTree.RBRoot.Node
  rbTree.RBRoot.Node.Right = rbtree.CreateNewRBNode(70, 70)
  rbTree.RBRoot.Node.Right.Color = true
  rbTree.RBRoot.Node.Right.Parent = rbTree.RBRoot.Node
  
  // Level 2 [10R 30B 60B 75B]
  rbTree.RBRoot.Node.Left.Left  = rbtree.CreateNewRBNode(10, 10)
  rbTree.RBRoot.Node.Left.Left.Parent = rbTree.RBRoot.Node.Left
  rbTree.RBRoot.Node.Left.Right  = rbtree.CreateNewRBNode(30, 30)
  rbTree.RBRoot.Node.Left.Right.Color = true
  rbTree.RBRoot.Node.Left.Right.Parent = rbTree.RBRoot.Node.Left
  rbTree.RBRoot.Node.Right.Left = rbtree.CreateNewRBNode(60, 60)
  rbTree.RBRoot.Node.Right.Left.Color = true
  rbTree.RBRoot.Node.Right.Left.Parent = rbTree.RBRoot.Node.Right
  rbTree.RBRoot.Node.Right.Right = rbtree.CreateNewRBNode(75, 75)
  rbTree.RBRoot.Node.Right.Right.Color = true
  rbTree.RBRoot.Node.Right.Right.Parent = rbTree.RBRoot.Node.Right

  // Level 3 [5B 13B 25R 35R]
  rbTree.RBRoot.Node.Left.Left.Left = rbtree.CreateNewRBNode(5, 5)
  rbTree.RBRoot.Node.Left.Left.Left.Color = true
  rbTree.RBRoot.Node.Left.Left.Left.Parent = rbTree.RBRoot.Node.Left.Left
  rbTree.RBRoot.Node.Left.Left.Right = rbtree.CreateNewRBNode(13, 13)
  rbTree.RBRoot.Node.Left.Left.Right.Color = true
  rbTree.RBRoot.Node.Left.Left.Right.Parent = rbTree.RBRoot.Node.Left.Left
  rbTree.RBRoot.Node.Left.Right.Left = rbtree.CreateNewRBNode(25, 25)
  rbTree.RBRoot.Node.Left.Right.Left.Parent = rbTree.RBRoot.Node.Left.Right
  rbTree.RBRoot.Node.Left.Right.Right = rbtree.CreateNewRBNode(35, 35)
  rbTree.RBRoot.Node.Left.Right.Right.Parent = rbTree.RBRoot.Node.Left.Right
  
  validate(t, rbTree)
  return rbTree
}

//function 'getTestRBTree5' returns below tree
//
//                  50B
//                 /    \
//                /      \
//               /        \
//             20R        70B
//             /  \      /  \
//            /    \    /    \
//          10B    30B 60B   75B
//          /  \   / \
//        05B 13B 25B 35B
//
func getTestRBTree5(t *testing.T) *rbtree.RBTree {
  // Level 0 [50B]
  rbTree := rbtree.CreateNewRBTree(IntKeyComparator)
  rbTree.RBRoot.Node = rbtree.CreateNewRBNode(50, 50)
  rbTree.RBRoot.Node.Color = true
  rbTree.RBRoot.Node.Parent = nil

  // Level 1 [20R 70B]
  rbTree.RBRoot.Node.Left = rbtree.CreateNewRBNode(20, 20)
  rbTree.RBRoot.Node.Left.Parent = rbTree.RBRoot.Node
  rbTree.RBRoot.Node.Right = rbtree.CreateNewRBNode(70, 70)
  rbTree.RBRoot.Node.Right.Color = true
  rbTree.RBRoot.Node.Right.Parent = rbTree.RBRoot.Node

  // Level 2 [10B 30B 60B 75B]
  rbTree.RBRoot.Node.Left.Left  = rbtree.CreateNewRBNode(10, 10)
  rbTree.RBRoot.Node.Left.Left.Color = true
  rbTree.RBRoot.Node.Left.Left.Parent = rbTree.RBRoot.Node.Left
  rbTree.RBRoot.Node.Left.Right  = rbtree.CreateNewRBNode(30, 30)
  rbTree.RBRoot.Node.Left.Right.Color = true
  rbTree.RBRoot.Node.Left.Right.Parent = rbTree.RBRoot.Node.Left
  rbTree.RBRoot.Node.Right.Left = rbtree.CreateNewRBNode(60, 60)
  rbTree.RBRoot.Node.Right.Left.Color = true
  rbTree.RBRoot.Node.Right.Left.Parent = rbTree.RBRoot.Node.Right
  rbTree.RBRoot.Node.Right.Right = rbtree.CreateNewRBNode(75, 75)
  rbTree.RBRoot.Node.Right.Right.Color = true
  rbTree.RBRoot.Node.Right.Right.Parent = rbTree.RBRoot.Node.Right

  // Level 3 [5B 13B 25B 35B]
  rbTree.RBRoot.Node.Left.Left.Left = rbtree.CreateNewRBNode(5, 5)
  rbTree.RBRoot.Node.Left.Left.Left.Color = true
  rbTree.RBRoot.Node.Left.Left.Left.Parent = rbTree.RBRoot.Node.Left.Left
  rbTree.RBRoot.Node.Left.Left.Right = rbtree.CreateNewRBNode(13, 13)
  rbTree.RBRoot.Node.Left.Left.Right.Color = true
  rbTree.RBRoot.Node.Left.Left.Right.Parent = rbTree.RBRoot.Node.Left.Left
  rbTree.RBRoot.Node.Left.Right.Left = rbtree.CreateNewRBNode(25, 25)
  rbTree.RBRoot.Node.Left.Right.Left.Color = true
  rbTree.RBRoot.Node.Left.Right.Left.Parent = rbTree.RBRoot.Node.Left.Right
  rbTree.RBRoot.Node.Left.Right.Right = rbtree.CreateNewRBNode(35, 35)
  rbTree.RBRoot.Node.Left.Right.Right.Color = true
  rbTree.RBRoot.Node.Left.Right.Right.Parent = rbTree.RBRoot.Node.Left.Right

  validate(t, rbTree)
  return rbTree
}

//function 'getTestRBTree6' returns below tree
//
//                  50B
//                 /    \
//                /      \
//               /        \
//             20R        70B
//             /  \      /  \
//            /    \    /    \
//          10B    30B 60B   75B
//          /  \   / \
//        05B 13R 25B 35B
//            / \
//           /   \
//          11B  14B
func getTestRBTree6(t *testing.T) *rbtree.RBTree {
  // Level 0 [50B]
  rbTree := rbtree.CreateNewRBTree(IntKeyComparator)
  rbTree.RBRoot.Node = rbtree.CreateNewRBNode(50, 50)
  rbTree.RBRoot.Node.Color = true
  rbTree.RBRoot.Node.Parent = nil

  // Level 1 [20R 70B]
  rbTree.RBRoot.Node.Left = rbtree.CreateNewRBNode(20, 20)
  rbTree.RBRoot.Node.Left.Parent = rbTree.RBRoot.Node
  rbTree.RBRoot.Node.Right = rbtree.CreateNewRBNode(70, 70)
  rbTree.RBRoot.Node.Right.Color = true
  rbTree.RBRoot.Node.Right.Parent = rbTree.RBRoot.Node

  // Level 2 [10B 30B 60B 75B]
  rbTree.RBRoot.Node.Left.Left  = rbtree.CreateNewRBNode(10, 10)
  rbTree.RBRoot.Node.Left.Left.Color = true
  rbTree.RBRoot.Node.Left.Left.Parent = rbTree.RBRoot.Node.Left
  rbTree.RBRoot.Node.Left.Right  = rbtree.CreateNewRBNode(30, 30)
  rbTree.RBRoot.Node.Left.Right.Color = true
  rbTree.RBRoot.Node.Left.Right.Parent = rbTree.RBRoot.Node.Left
  rbTree.RBRoot.Node.Right.Left = rbtree.CreateNewRBNode(60, 60)
  rbTree.RBRoot.Node.Right.Left.Color = true
  rbTree.RBRoot.Node.Right.Left.Parent = rbTree.RBRoot.Node.Right
  rbTree.RBRoot.Node.Right.Right = rbtree.CreateNewRBNode(75, 75)
  rbTree.RBRoot.Node.Right.Right.Color = true
  rbTree.RBRoot.Node.Right.Right.Parent = rbTree.RBRoot.Node.Right

  // Level 3 [5B 13R 25B 35B]
  rbTree.RBRoot.Node.Left.Left.Left = rbtree.CreateNewRBNode(5, 5)
  rbTree.RBRoot.Node.Left.Left.Left.Color = true
  rbTree.RBRoot.Node.Left.Left.Left.Parent = rbTree.RBRoot.Node.Left.Left
  rbTree.RBRoot.Node.Left.Left.Right = rbtree.CreateNewRBNode(13, 13)
  rbTree.RBRoot.Node.Left.Left.Right.Parent = rbTree.RBRoot.Node.Left.Left
  rbTree.RBRoot.Node.Left.Right.Left = rbtree.CreateNewRBNode(25, 25)
  rbTree.RBRoot.Node.Left.Right.Left.Color = true
  rbTree.RBRoot.Node.Left.Right.Left.Parent = rbTree.RBRoot.Node.Left.Right
  rbTree.RBRoot.Node.Left.Right.Right = rbtree.CreateNewRBNode(35, 35)
  rbTree.RBRoot.Node.Left.Right.Right.Color = true
  rbTree.RBRoot.Node.Left.Right.Right.Parent = rbTree.RBRoot.Node.Left.Right

  // Level 4 [11B 14B]
  rbTree.RBRoot.Node.Left.Left.Right.Left = rbtree.CreateNewRBNode(11, 11)
  rbTree.RBRoot.Node.Left.Left.Right.Left.Color = true
  rbTree.RBRoot.Node.Left.Left.Right.Left.Parent = rbTree.RBRoot.Node.Left.Left.Right
  rbTree.RBRoot.Node.Left.Left.Right.Right = rbtree.CreateNewRBNode(14, 14)
  rbTree.RBRoot.Node.Left.Left.Right.Right.Color = true
  rbTree.RBRoot.Node.Left.Left.Right.Right.Parent = rbTree.RBRoot.Node.Left.Left.Right

  validate(t, rbTree)
  return rbTree
}

//function 'getTestRBTree7' returns below tree
//
//                  50B
//                 /    \
//                /      \
//               /        \
//             20R        70B
//             /  \      /  \
//            /    \    /    \
//          10B    30B 60B   75B
//          /  \   / \
//        05R 13B 25B 35B
//        / \
//       /   \
//     04B   06B
func getTestRBTree7(t *testing.T) *rbtree.RBTree {
  // Level 0 [50B]
  rbTree := rbtree.CreateNewRBTree(IntKeyComparator)
  rbTree.RBRoot.Node = rbtree.CreateNewRBNode(50, 50)
  rbTree.RBRoot.Node.Color = true
  rbTree.RBRoot.Node.Parent = nil

  // Level 1 [20R 70B]
  rbTree.RBRoot.Node.Left = rbtree.CreateNewRBNode(20, 20)
  rbTree.RBRoot.Node.Left.Parent = rbTree.RBRoot.Node
  rbTree.RBRoot.Node.Right = rbtree.CreateNewRBNode(70, 70)
  rbTree.RBRoot.Node.Right.Color = true
  rbTree.RBRoot.Node.Right.Parent = rbTree.RBRoot.Node

  // Level 2 [10B 30B 60B 75B]
  rbTree.RBRoot.Node.Left.Left  = rbtree.CreateNewRBNode(10, 10)
  rbTree.RBRoot.Node.Left.Left.Color = true
  rbTree.RBRoot.Node.Left.Left.Parent = rbTree.RBRoot.Node.Left
  rbTree.RBRoot.Node.Left.Right  = rbtree.CreateNewRBNode(30, 30)
  rbTree.RBRoot.Node.Left.Right.Color = true
  rbTree.RBRoot.Node.Left.Right.Parent = rbTree.RBRoot.Node.Left
  rbTree.RBRoot.Node.Right.Left = rbtree.CreateNewRBNode(60, 60)
  rbTree.RBRoot.Node.Right.Left.Color = true
  rbTree.RBRoot.Node.Right.Left.Parent = rbTree.RBRoot.Node.Right
  rbTree.RBRoot.Node.Right.Right = rbtree.CreateNewRBNode(75, 75)
  rbTree.RBRoot.Node.Right.Right.Color = true
  rbTree.RBRoot.Node.Right.Right.Parent = rbTree.RBRoot.Node.Right

  // Level 3 [5R 13B 25B 35B]
  rbTree.RBRoot.Node.Left.Left.Left = rbtree.CreateNewRBNode(5, 5)
  rbTree.RBRoot.Node.Left.Left.Left.Parent = rbTree.RBRoot.Node.Left.Left
  rbTree.RBRoot.Node.Left.Left.Right = rbtree.CreateNewRBNode(13, 13)
  rbTree.RBRoot.Node.Left.Left.Right.Color = true
  rbTree.RBRoot.Node.Left.Left.Right.Parent = rbTree.RBRoot.Node.Left.Left
  rbTree.RBRoot.Node.Left.Right.Left = rbtree.CreateNewRBNode(25, 25)
  rbTree.RBRoot.Node.Left.Right.Left.Color = true
  rbTree.RBRoot.Node.Left.Right.Left.Parent = rbTree.RBRoot.Node.Left.Right
  rbTree.RBRoot.Node.Left.Right.Right = rbtree.CreateNewRBNode(35, 35)
  rbTree.RBRoot.Node.Left.Right.Right.Color = true
  rbTree.RBRoot.Node.Left.Right.Right.Parent = rbTree.RBRoot.Node.Left.Right

  // Level 4 [04B 06B]
  rbTree.RBRoot.Node.Left.Left.Left.Left = rbtree.CreateNewRBNode(4, 4)
  rbTree.RBRoot.Node.Left.Left.Left.Left.Color = true
  rbTree.RBRoot.Node.Left.Left.Left.Left.Parent = rbTree.RBRoot.Node.Left.Left.Left
  rbTree.RBRoot.Node.Left.Left.Left.Right = rbtree.CreateNewRBNode(6, 6)
  rbTree.RBRoot.Node.Left.Left.Left.Right.Color = true
  rbTree.RBRoot.Node.Left.Left.Left.Right.Parent = rbTree.RBRoot.Node.Left.Left.Left

  validate(t, rbTree)
  return rbTree
}

//function 'getTestRBTree7' returns below tree
//
//                  50B
//                 /    \
//                /      \
//               /        \
//             20R        70B
//             /  \      /  \
//            /    \    /    \
//          10B    30B 60B   75B
//          /  \   / \
//        05B 13B 25B 35B
//
func getTestRBTree8(t *testing.T) *rbtree.RBTree {
  // Level 0 [50B]
  rbTree := rbtree.CreateNewRBTree(IntKeyComparator)
  rbTree.RBRoot.Node = rbtree.CreateNewRBNode(50, 50)
  rbTree.RBRoot.Node.Color = true
  rbTree.RBRoot.Node.Parent = nil

  // Level 1 [20R 70B]
  rbTree.RBRoot.Node.Left = rbtree.CreateNewRBNode(20, 20)
  rbTree.RBRoot.Node.Left.Parent = rbTree.RBRoot.Node
  rbTree.RBRoot.Node.Right = rbtree.CreateNewRBNode(70, 70)
  rbTree.RBRoot.Node.Right.Color = true
  rbTree.RBRoot.Node.Right.Parent = rbTree.RBRoot.Node

  // Level 2 [10B 30B 60B 75B]
  rbTree.RBRoot.Node.Left.Left  = rbtree.CreateNewRBNode(10, 10)
  rbTree.RBRoot.Node.Left.Left.Color = true
  rbTree.RBRoot.Node.Left.Left.Parent = rbTree.RBRoot.Node.Left
  rbTree.RBRoot.Node.Left.Right  = rbtree.CreateNewRBNode(30, 30)
  rbTree.RBRoot.Node.Left.Right.Color = true
  rbTree.RBRoot.Node.Left.Right.Parent = rbTree.RBRoot.Node.Left
  rbTree.RBRoot.Node.Right.Left = rbtree.CreateNewRBNode(60, 60)
  rbTree.RBRoot.Node.Right.Left.Color = true
  rbTree.RBRoot.Node.Right.Left.Parent = rbTree.RBRoot.Node.Right
  rbTree.RBRoot.Node.Right.Right = rbtree.CreateNewRBNode(75, 75)
  rbTree.RBRoot.Node.Right.Right.Color = true
  rbTree.RBRoot.Node.Right.Right.Parent = rbTree.RBRoot.Node.Right

  // Level 3 [5B 13B 25B 35B]
  rbTree.RBRoot.Node.Left.Left.Left = rbtree.CreateNewRBNode(5, 5)
  rbTree.RBRoot.Node.Left.Left.Left.Color = true
  rbTree.RBRoot.Node.Left.Left.Left.Parent = rbTree.RBRoot.Node.Left.Left
  rbTree.RBRoot.Node.Left.Left.Right = rbtree.CreateNewRBNode(13, 13)
  rbTree.RBRoot.Node.Left.Left.Right.Color = true
  rbTree.RBRoot.Node.Left.Left.Right.Parent = rbTree.RBRoot.Node.Left.Left
  rbTree.RBRoot.Node.Left.Right.Left = rbtree.CreateNewRBNode(25, 25)
  rbTree.RBRoot.Node.Left.Right.Left.Color = true
  rbTree.RBRoot.Node.Left.Right.Left.Parent = rbTree.RBRoot.Node.Left.Right
  rbTree.RBRoot.Node.Left.Right.Right = rbtree.CreateNewRBNode(35, 35)
  rbTree.RBRoot.Node.Left.Right.Right.Color = true
  rbTree.RBRoot.Node.Left.Right.Right.Parent = rbTree.RBRoot.Node.Left.Right

  validate(t, rbTree)
  return rbTree
}

func assertKey(t *testing.T, node *rbtree.RBNode, key int) {
  nodeKey := node.Key.(int)
  if nodeKey != key {
    t.Error("Expected key ", key, " but got ", nodeKey, )
  }
}

func assertLeaf(t *testing.T, node *rbtree.RBNode) {
  if node.Left != nil || node.Right != nil {
    t.Error("Expected leaf node but found node with child", )
  }
}

func assertBlack(t *testing.T, node *rbtree.RBNode) {
  if !node.Color {
    t.Error("Expected BLACK node but got RED", )
  }
}

func assertRed(t *testing.T, node *rbtree.RBNode) {
  if node.Color {
    t.Error("Expected RED node but got BLACK", )
  }
}

// Gets a sample tree to be used by test cases
//
func getTestRBTree(t *testing.T) *rbtree.RBTree {
  var rbTree = rbtree.CreateNewRBTree(IntKeyComparator)
  entries := []int{ 10, 11, 12, 13, 14, 15, 16, 9, 8, 7, 6, 5, 4, 3 }  
  for i := 0; i < len(entries); i++ {
    rbTree.PutValue(entries[i], entries[i])
    validate(t, rbTree)
  }
  
  // The validate call in above FOR loop will ensure the tree remains
  // RBTree after adding each entry.
  return rbTree
}

func IntKeyComparator(key1 interface{}, key2 interface{}) int {
  intKey1 := key1.(int)
  intKey2 := key2.(int)
  if intKey1 == intKey2 {
    return 0
  } else if intKey1 > intKey2 {
    return 1
  }

  return -1
}

func validate(t *testing.T, rbTree *rbtree.RBTree) int {
  if rbTree.RBRoot.Node.IsRed() {
    t.Error("RBTree root cannot be RED", )
    return 0;
  }
   
  bh := validateIntern(t, rbTree, rbTree.RBRoot.Node)
  bh--
  t.Log(bh)
  return bh
}

func validateIntern(t *testing.T, rbTree *rbtree.RBTree, rbNode *rbtree.RBNode) int {
  if rbNode == nil {
    return 1  
  }

  lChild := rbNode.Left
  rChild := rbNode.Right

  if rbNode.IsRed() {
    if (lChild != nil && lChild.IsRed()) || (rChild != nil && rChild.IsRed()) {
      t.Error("Double RED found", )
      return 0;
   }
  }

  lbh := validateIntern(t, rbTree, lChild)
  rbh := validateIntern(t, rbTree, rChild)
  
  if (lChild != nil && (rbTree.Comparer(lChild.Key, rbNode.Key) >= 0)) || 
    (rChild != nil && (rbTree.Comparer(rChild.Key, rbNode.Key) <= 0)) {
    t.Error("BST violation", )
    return 0;
  }
  
  if lbh != 0 && rbh != 0 && lbh != rbh {
    t.Error("BLACK Depth violation", )
    return 0;
  }
  
  if lbh != 0 && rbh != 0 {
    if rbNode.IsRed() {
      return lbh
    } else {
      return lbh + 1
    }
  } else {
    return 0
  }
}