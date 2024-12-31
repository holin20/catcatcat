package ezgo

func IsErr(err error) bool {
	if c, ok := err.(*cause); ok {
		return c != nil
	}
	return err != nil
}

func IsOk(err error) bool {
	return err == nil
}
