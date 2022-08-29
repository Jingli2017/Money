package main

import (
	"fmt"
	"strings"
)

/*
Each user completed 1 purchase.
completed_purchase_user_ids = [
"3123122444","234111110", "8321125440", "99911063"]

And our ops team provided us with some raw log data from our ad server showing every time a
user clicked on one of our ads:
ad_clicks = [
"IP_Address,Time,Ad_Text",
"122.121.0.1,2016-11-03 11:41:19,Buy wool coats for your pets",
"96.3.199.11,2016-10-15 20:18:31,2017 Pet Mittens",
"122.121.0.250,2016-11-01 06:13:13,The Best Hollywood Coats",
"82.1.106.8,2016-11-12 23:05:14,Buy wool coats for your pets",
"92.130.6.144,2017-01-01 03:18:55,Buy wool coats for your pets",
"92.130.6.145,2017-01-01 03:18:55,2017 Pet Mittens",
]

The client also sent over the IP addresses of all their users.

all_user_ips = [
"User_ID,IP_Address",
"2339985511,122.121.0.155",
"234111110,122.121.0.1",
"3123122444,92.130.6.145",
"39471289472,2001:0db8:ac10:fe01:0000:0000:0000:0000",
"8321125440,82.1.106.8",
"99911063,92.130.6.144"
]
*/

func main() {
	completedIds := []string{"3123122444","234111110", "8321125440", "99911063"}
	allUserIps := []string{"2339985511,122.121.0.155",
		"234111110,122.121.0.1",
		"3123122444,92.130.6.145",
		"39471289472,2001:0db8:ac10:fe01:0000:0000:0000:0000",
		"8321125440,82.1.106.8",
		"99911063,92.130.6.144"}
	clicks := []string{
		"122.121.0.1,2016-11-03 11:41:19,Buy wool coats for your pets",
		"96.3.199.11,2016-10-15 20:18:31,2017 Pet Mittens",
		"122.121.0.250,2016-11-01 06:13:13,The Best Hollywood Coats",
		"82.1.106.8,2016-11-12 23:05:14,Buy wool coats for your pets",
		"92.130.6.144,2017-01-01 03:18:55,Buy wool coats for your pets",
		"92.130.6.145,2017-01-01 03:18:55,2017 Pet Mittens",
	}

	advertise(completedIds, allUserIps, clicks)
}

func advertise(completedIds []string, allUserIps []string, adClicks []string) {
	userIdToIp := make(map[string]string)
	for _, userIp := range allUserIps {
		split := strings.Split(userIp, ",")
		userId := split[0]
		ip := split[1]
		userIdToIp[userId] = ip
	}

	completedIps := make(map[string]bool)
	for _, id := range completedIds{
		completedIps[userIdToIp[id]] = true
	}

	adCount := make(map[string]int)
	purchaseCount := make(map[string]int)
	for _, click := range adClicks {
		split := strings.Split(click, ",")
		ip := split[0]
		ad := split[2]
		adCount[ad]++
		if _, ok := completedIps[ip]; ok {
			purchaseCount[ad]++
		}
	}

	//output
	for ad, count := range adCount {
		fmt.Printf("%d of %d, %s", purchaseCount[ad], count, ad)
	}
}
