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
  if rbTree.RBRoot.Node.Color != true {
    t.Error ("RBRoot.Node must be BLACK", )
  }
  
  if rbTree.RBRoot.Node.Key.(int) != 11 {
    t.Error ("RBRoot.Node key must be 11", )
  }

  // Level 1
  if rbTree.RBRoot.Node.Left.Color != true {
    t.Error ("RBRoot.Node.Left node must be BLACK", )
  }
 
  if rbTree.RBRoot.Node.Left.Key.(int) != 7 {
    t.Error ("RBRoot.Node.Left key must be 7", )
  }
  
  if rbTree.RBRoot.Node.Right.Color != true {
    t.Error ("RBRoot.Node.Right must be BLACK", )
  }
 
  if rbTree.RBRoot.Node.Right.Key.(int) != 13 {
    t.Error ("RBRoot.Node.Right key must be 13", )
  }

  // Level 2
  if rbTree.RBRoot.Node.Left.Left.Color != false {
    t.Error ("RBRoot.Node.Left.Left must be RED", )
  }
 
  if rbTree.RBRoot.Node.Left.Left.Key.(int) != 5 {
    t.Error ("RBRoot.Node.Left.Left key must be 5", )
  }
  
  if rbTree.RBRoot.Node.Left.Right.Color != false {
    t.Error ("BRoot.Node.Left.Right must be RED", )
  }
 
  if rbTree.RBRoot.Node.Left.Right.Key.(int) != 9 {
    t.Error ("BRoot.Node.Left.Right key must be 9", )
  }
  
  if rbTree.RBRoot.Node.Right.Left.Color != true {
    t.Error ("RBRoot.Node.Right.Left node must be BLACK", )
  }
 
  if rbTree.RBRoot.Node.Right.Left.Key.(int) != 12 {
    t.Error ("RBRoot.Node.Right.Left key must be 12", )
  }
  
  if rbTree.RBRoot.Node.Right.Right.Color != true {
    t.Error ("RBRoot.Node.Right.Right node must be BLACK", )
  }
 
  if rbTree.RBRoot.Node.Right.Right.Key.(int) != 15 {
    t.Error ("RBRoot.Node.Left.Right key must be 15", )
  }
  
  // Level 3
  if rbTree.RBRoot.Node.Left.Left.Left.Color != true {
    t.Error ("RBRoot.Node.Left.Left.Left must be BLACK", )
  }
 
  if rbTree.RBRoot.Node.Left.Left.Left.Key.(int) != 4 {
    t.Error ("RBRoot.Node.Left.Left.Left key must be 4", )
  }
  
  if rbTree.RBRoot.Node.Left.Left.Right.Color != true {
    t.Error ("RBRoot.Node.Left.Left.Right must be BLACK", )
  }  
 
  if rbTree.RBRoot.Node.Left.Left.Right.Key.(int) != 6 {
    t.Error ("RBRoot.Node.Left.Left.Right key must be 6", )
  }
  
  if rbTree.RBRoot.Node.Left.Right.Left.Color != true {
    t.Error ("RBRoot.Node.Left.Right.Left must be BLACK", )
  }
 
  if rbTree.RBRoot.Node.Left.Right.Left.Key.(int) != 8 {
    t.Error ("RBRoot.Node.Left.Right.Left key must be 8", )
  }
  
  if rbTree.RBRoot.Node.Left.Right.Right.Color != true {
    t.Error ("RBRoot.Node.Left.Right.Right must be BLACK", )
  }
 
  if rbTree.RBRoot.Node.Left.Right.Right.Key.(int) != 10 {
    t.Error ("RBRoot.Node.Left.Right.Right key must be 10", )
  }
  
  if rbTree.RBRoot.Node.Right.Right.Left.Color != false {
    t.Error ("RBRoot.Node.Right.Right.Left node must be RED", )
  }
  
  if rbTree.RBRoot.Node.Right.Right.Left.Key.(int) != 14 {
    t.Error ("RBRoot.Node.Right.Right.Left key must be 14", )
  }
  
  if rbTree.RBRoot.Node.Right.Right.Right.Color != false {
    t.Error ("RBRoot.Node.Right.Right.Right node must be RED", )
  }
  
  if rbTree.RBRoot.Node.Right.Right.Right.Key.(int) != 16 {
    t.Error ("RBRoot.Node.Right.Right.Right key must be 16", )
  }
  
  // Level 4
  if rbTree.RBRoot.Node.Left.Left.Left.Left.Color != false {
    t.Error ("RBRoot.Node.Left.Left.Left.Left must be RED", )
  }
 
  if rbTree.RBRoot.Node.Left.Left.Left.Left.Key.(int) != 3 {
    t.Error ("RBRoot.Node.Left.Left.Left.Left key must be 3", )
  }
  
  // Validate leaves
  if rbTree.RBRoot.Node.Left.Left.Left.Left.Left != nil {
    t.Error ("RBRoot.Node.Left.Left.Left.Left (node 3) must be leaf [Found left child]", )
  }
  
  if rbTree.RBRoot.Node.Left.Left.Left.Left.Right != nil {
    t.Error ("RBRoot.Node.Left.Left.Left.Left (node 3) must be leaf [Found right child]", )
  }
  
  if rbTree.RBRoot.Node.Left.Left.Right.Left != nil {
    t.Error ("rbTree.RBRoot.Node.Left.Left.Right (node 6) must be leaf [Found left child]", )
  }
  
  if rbTree.RBRoot.Node.Left.Left.Right.Right != nil {
    t.Error ("rbTree.RBRoot.Node.Left.Left.Right (node 6) must be leaf [Found right child]", )
  }
  
  if rbTree.RBRoot.Node.Left.Right.Left.Left != nil {
    t.Error ("rbTree.RBRoot.Node.Left.Right.Left (node 8) must be leaf [Found left child]", )
  }
  
  if rbTree.RBRoot.Node.Left.Right.Left.Right != nil {
    t.Error ("rbTree.RBRoot.Node.Left.Right.Left (node 8) must be leaf [Found right child]", )
  }
  
  if rbTree.RBRoot.Node.Left.Right.Right.Left != nil {
    t.Error ("rbTree.RBRoot.Node.Left.Right.Right (node 10) must be leaf [Found left child]", )
  }
  
  if rbTree.RBRoot.Node.Left.Right.Right.Right != nil {
    t.Error ("rbTree.RBRoot.Node.Left.Right.Right (node 10) must be leaf [Found right child]", )
  }
  
  if rbTree.RBRoot.Node.Right.Right.Left.Left != nil {
    t.Error ("rbTree.RBRoot.Node.Right.Right.Left (node 14) must be leaf [Found left child]", )
  }
  
  if rbTree.RBRoot.Node.Right.Right.Left.Right != nil {
    t.Error ("rbTree.RBRoot.Node.Right.Right.Left (node 14) must be leaf [Found right child]", )
  }
  
  if rbTree.RBRoot.Node.Right.Right.Right.Left != nil {
    t.Error ("rbTree.RBRoot.Node.Left.Right.Right (node 16) must be leaf [Found left child]", )
  }
  
  if rbTree.RBRoot.Node.Right.Right.Right.Right != nil {
    t.Error ("rbTree.RBRoot.Node.Left.Right.Right (node 16) must be leaf [Found right child]", )
  }
  
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

  // Level 1 [20R 70B]
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