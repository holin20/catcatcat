package ezgo

func IsErr(err error) bool {
	return err != nil
}

func IsOk(err error) bool {
	return err == nil
}
