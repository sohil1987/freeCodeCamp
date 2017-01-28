package main

import (
	"fmt"
	"strings"
	"sync"
)

func parallelLetFreq() {
	var wg sync.WaitGroup
	wg.Add(3)
	go count(euro, &wg)
	go count(dutch, &wg)
	go count(us, &wg)
	wg.Wait()

}

func count(str string, wg *sync.WaitGroup) {
	defer wg.Done()
	counter := make(map[string]int)
	for _, letter := range strings.Split(str, "") {
		fmt.Sprintln(letter)
		letter = strings.ToLower(letter)
		_, ok := counter[letter]
		if ok {
			counter[letter]++
		} else {
			counter[letter] = 1
		}
	}
	fmt.Println("Letters --> ", counter)
}

var (
	euro = `Freude schöner Götterfunken
Tochter aus Elysium,
Wir betreten feuertrunken,
Himmlische, dein Heiligtum!
Deine Zauber binden wieder
Was die Mode streng geteilt;
Alle Menschen werden Brüder,
Wo dein sanfter Flügel weilt.`
	dutch = `Wilhelmus van Nassouwe
ben ik, van Duitsen bloed,
den vaderland getrouwe
blijf ik tot in den dood.
Een Prinse van Oranje
ben ik, vrij, onverveerd,
den Koning van Hispanje
heb ik altijd geëerd.`
	us = `O say can you see by the dawn's early light,
What so proudly we hailed at the twilight's last gleaming,
Whose broad stripes and bright stars through the perilous fight,
O'er the ramparts we watched, were so gallantly streaming?
And the rockets' red glare, the bombs bursting in air,
Gave proof through the night that our flag was still there;
O say does that star-spangled banner yet wave,
O'er the land of the free and the home of the brave?`
)
