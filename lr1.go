package lr1

type Rule func(items []interface{}) (interface{}, error)

// Simplified LR(1) parser that applies replacement rules to the tail of list of items
// and replaces found sequences with parsed items.
func Parse(items []interface{}, rules ...Rule) ([]interface{}, error) {

	//Initially, we should have a stack of tokenized strings
	if len(items) == 0 {
		return items, nil
	}

	for {
		triggered := false

		//given the list of tokens and parsed items, we go backwards, from the end to beginning
		for b := len(items); b >= 0; b-- {

			// and calculate tha tail of the list, increasing the size of the tail on each iteration
			tail := items[b:len(items)]

			// each time we have some tail, consisting of some elements we apply rules to it
			for _, rule := range rules {

				// the rule we apply may produce a replacement for the tail
				replace, err := rule(tail)

				if err != nil {
					return nil, err
				}

				if replace == nil {
					continue
				}

				// if rule have produced the replacement -- replace the tail with it
				items = append(items[0:b], replace)
				triggered = true
			}
		}

		// if no rule have been triggered for the whole stack - no need to continue, just exit
		if !triggered {
			break
		}
	}

	return items, nil
}
