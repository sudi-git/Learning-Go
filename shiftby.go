package sprint

func ShiftBy(r rune, step int) rune {

	postion := r - 'a'
//it counts the total steps taken and doesnt allow go beyond 26	
	newposition := (postion + rune(step))%26
//add the initial and new position to displays where it is
	return rune ('a'+newposition)


}

