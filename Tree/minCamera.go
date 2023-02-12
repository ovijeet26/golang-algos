package tree

//Leet code link -> https://leetcode.com/problems/binary-tree-cameras/
func minCameraCover(root *TreeNode) int {

	if root.Left == nil && root.Right == nil {

		return 1
	}

	cameraCount := 0

	_, monitor := dfsCamera(root, &cameraCount)

	if !monitor {
		cameraCount++
	}

	return cameraCount
}

// isCamera and isMonitored
func dfsCamera(node *TreeNode, cameraCount *int) (bool, bool) {

	// If we hit a leaf, node, then we pass that it's not a camera, also it's not monitored.
	if node == nil {

		return false, true
	}

	cameraLeft, monitorLeft := dfsCamera(node.Left, cameraCount)
	cameraRight, monitorRight := dfsCamera(node.Right, cameraCount)

	currentCamera, currentMonitor := false, false
	if cameraLeft || cameraRight {

		currentCamera = false
		currentMonitor = true
	}

	if monitorLeft == false || monitorRight == false {

		*cameraCount++
		currentCamera = true
		currentMonitor = true
	}

	return currentCamera, currentMonitor
}
