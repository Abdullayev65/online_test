package utill

func AnyNotNil(all ...interface{}) bool {
	for _, a := range all {
		if a != nil {
			return true
		}
	}
	return false
}
