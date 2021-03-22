package msgresolution

type adaptorUnit interface {
	isAdaptorUnit()
}

type adaptorUnits []adaptorUnit

func (au adaptorUnits) tryMerge() adaptorUnits {
	var mergedUnits []adaptorUnit
	for i := 0; i < len(au); i += 1 {
		if au[i] == nil {
			continue
		}
		if nau, ok := au[i].(*nestedAdaptorUnit); ok {
			for j := i + 1; j < len(au); j += 1 {
				if onau, ok := au[j].(*nestedAdaptorUnit); ok {
					if nau.merge(onau) {
						au[j] = nil
					}
				}
			}
		}
		mergedUnits = append(mergedUnits, au[i])
	}
	return mergedUnits
}

type nestedAdaptorUnit struct {
	name       string
	nestedUnit []adaptorUnit
}

func (nau *nestedAdaptorUnit) merge(other *nestedAdaptorUnit) bool {
	if nau.name != other.name {
		return false
	}
onuLoop:
	for _, onu := range other.nestedUnit {
		if onu, ok := onu.(*nestedAdaptorUnit); ok {
			for _, nu := range nau.nestedUnit {
				if nu, ok := nu.(*nestedAdaptorUnit); ok {
					if nu.merge(onu) {
						continue onuLoop
					}
				}
			}

		}
		nau.nestedUnit = append(nau.nestedUnit, onu)
	}
	return true
}
