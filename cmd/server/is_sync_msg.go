package server

import "strconv"

func parseSyncMessage(msg []byte) (bool, int) {
	if len(msg) > 5 && string(msg[:5]) == "SYNC:" {
		msgIndx, err := strconv.Atoi(string(msg[5:]))
		if err != nil {
			return false, 0
		}
		return true, msgIndx
	}
	return false, 0
}
