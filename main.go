package main

// You are Given -
// 1. User preference matrix
// 2. A user-id.
// 3. A value k for ex - 3

// ratings = [[2, 0, 5, 0, 4, 3], -- user_id = 0  2 is the rating that user_id 0 gave item_id 0
//            [0, 4, 0, 4, 4, 0], -- user_id = 1
//            [0, 1, 0, 0, 0, 4],
//            [4, 0, 4, 0, 4, 2],
//            [3, 5, 0, 1, 0, 0]]

// Rows are users - Assume User Id is row index, Columns are items, assume item id is column index.
// If there is a 0 rating, assume that user has not rated that item yet.
// Ratings are from 1 to 5.

// similarity([2,0.....], [0,4....])  --> 2 * 0 + 0 * 4 + 5 * 0 + 0 * 4 ....

// You need to -
// Create a similarity function as well.
// 1. Find the top `k` user-ids who are most `similar` to the user-id in question based on the similarity function.

import (
	"fmt"
	"sort"
)

type UserSimilarity struct {
	Index int
	Score int
}

func similarity(userRatings1, userRatings2 []int) int {
	var result int

	for index := range userRatings1 {
		result += userRatings1[index] * userRatings2[index]
	}

	return result
}

func makeSimilarityArray(userId int, usersRatings [][]int, k int) []int {
	var userSimilarities []UserSimilarity

	for index := range usersRatings {
		if index == userId {
			continue
		}

		userSimilarities = append(userSimilarities, UserSimilarity{
			Index: index,
			Score: similarity(usersRatings[userId], usersRatings[index]),
		})
	}

	sort.Slice(userSimilarities, func(i, j int) bool {
		return userSimilarities[i].Score > userSimilarities[j].Score
	})

	fmt.Println("")
	resultDescription := fmt.Sprintf("Top users (id + similarity score) with most `similar` ratings to user with id %d:", userId)
	fmt.Println(resultDescription)
	fmt.Println(userSimilarities)
	fmt.Println("")

	topSimilarUsers := make([]int, 0, len(userSimilarities))

	for index, item := range userSimilarities {
		topSimilarUsers = append(topSimilarUsers, item.Index)

		if index+1 >= k {
			break
		}
	}

	return topSimilarUsers
}

func performComparison(userId int, ratings [][]int, k int) {
	similarRatingsUserIds := makeSimilarityArray(userId, ratings, k)
	resultDescription := fmt.Sprintf("IDs of top %d users with most `similar` ratings to user with id %d:", k, userId)

	fmt.Println(resultDescription)
	fmt.Println(similarRatingsUserIds)
	fmt.Println("")
}

func main() {
	testRatings := [][]int{
		{0, 4, 0, 4, 4, 0},
		{0, 1, 0, 0, 0, 4},
		{4, 0, 4, 0, 4, 2},
		{3, 5, 0, 1, 0, 0},
	}

	performComparison(3, testRatings, 2)
}
