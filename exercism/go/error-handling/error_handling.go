package erratum

func Use(opener ResourceOpener, input string) (err error) {
	var resource Resource

	//  while ! opened
	opened := false
	for !opened {
		resource, err = opener()
		if err != nil {
			// if Transient, keep trying.
			if _, ok := err.(TransientError); ok {
				continue
			}
			// return real error
			return
		}
		opened = true
	}

	defer resource.Close()
	defer func() {
		if r := recover(); r != nil {
			ferr, ok := r.(FrobError)
			if ok {
				resource.Defrob(ferr.defrobTag)
			}
			err = r.(error)
		}
	}()

	resource.Frob(input)
	return nil
}
