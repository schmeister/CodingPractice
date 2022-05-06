package erratum

func Use(opener ResourceOpener, input string) (err error) {
    var res Resource
    opened := false
    
    for !opened {
    	res, err = opener()
        if err != nil {
            if _, ok := err.(TransientError); ok {
                continue
            }
        	return
        }
    	opened = true
    }
	defer res.Close()
	defer func() {
        if r := recover(); r != nil {
        	ferr, ok := r.(FrobError)
            if ok {
            	res.Defrob(ferr.defrobTag)
            }
        	err = r.(error)
        }
    }()
    
    res.Frob(input)
    return
}
