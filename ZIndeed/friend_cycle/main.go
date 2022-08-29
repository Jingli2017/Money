package main

import (
	"fmt"
	"strings"
)

/*
Sample Output:
Output:
1: 2, 3
2: 1
3: 1, 4
4: 3
6: None
*/

/*
Sample Input:
employees = [
  "1, Bill, Engineer",
  "2, Joe, HR",
  "3, Sally, Engineer",
  "4, Richard, Business",
  "6, Tom, Engineer"
]

friendships = [
  "1, 2",
  "1, 3",
  "3, 4"
]
*/
func main() {
	friendShips := []string{"1,2", "1,3", "3,4"}
	employees := []string{
		"1,Bill,Engineer",
		"2,Joe,HR",
		"3,Sally,Engineer",
		"4,Richard,Business",
		"6,Tom,Engineer",
	}

	res := findFriends(employees, friendShips)
	fmt.Println(res)

	res_2 := findFriendsDep(res, employees)
	fmt.Println(res_2)
}

func findFriends(employees []string, friendShips []string) map[string][]string {
	friendsMap := make(map[string][]string)
	for _, friends := range friendShips {
		split := strings.Split(friends, ",")
		src := split[0]
		dst := split[1]
		friendsMap[src] = append(friendsMap[src], dst)
		friendsMap[dst] = append(friendsMap[dst], src)
	}
	// cache up
	for _, e := range employees {
		split := strings.Split(e, ",")
		id := split[0]
		if _, ok := friendsMap[id]; !ok {
			friendsMap[id] = nil
		}
	}

	return friendsMap
}

//Pt 2.Now for each department count the number of employees that have a friend in another department
/*
Sample Output:
Output:
"Engineer: 2 of 3"
"HR: 1 of 1"
"Business: 1 of 1"
*/
func findFriendsDep(friendsMap map[string][]string, employees []string) []string {
	depToIds := make(map[string][]string)
	idToDep := make(map[string]string)
	for _, e := range employees {
		split := strings.Split(e, ",")
		id := split[0]
		dep := split[2]
		depToIds[dep] = append(depToIds[dep], id)
		idToDep[id] = dep
	}

	depToOutsideFriends := make(map[string]int)
	for dep, ids := range depToIds {
		for _, id := range ids {
			for _, friend := range friendsMap[id] {
				if idToDep[friend] != dep {
					depToOutsideFriends[dep]++
					break
				}
			}
		}
	}

	output := make([]string, 0)
	for dep := range depToIds {
		output = append(output, fmt.Sprintf("%s: %d of %d", dep,  depToOutsideFriends[dep], len(depToIds[dep])))
	}

	return output
}
