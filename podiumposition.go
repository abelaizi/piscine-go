package piscine

func PodiumPosition(podium [][]string) [][]string {
	for i := 0; i < len(podium); i++ {
		for j := 0; j < len(podium)-1; j++ {
			if podium[i][0] < podium[j][0] {
				podium[i][0], podium[j][0] = podium[j][0], podium[i][0]
			}
		}
	}
	return podium
}
