package pqutil

func Concate(a string, b string, separator string) string {
	if len(a) > 0 && len(b) > 0 {
		return a + separator + b
	}
	if len(a) > 0 {
		return a
	}
	return b
}

//func FromInt64(i int64) string {
//	return strconv.FormatInt(int64(i), 10)
//}

//func ToInt64(s string) (int64, base.Error) {
//	i, e := strconv.ParseInt(s, 10, 64)
//	if e != nil {
//		return i, liberror.CantConvertStringToInt64(s)
//	}
//	return i, nil
//}
//
//func FromInt(i int) string {
//	return strconv.Itoa(i)
//}
//
//func ToInt(s string) (int, base.Error) {
//	i, e := strconv.Atoi(s)
//	if e != nil {
//		return i, liberror.CantConvertStringToInt64(s)
//	}
//	return i, nil
//}

//func JoinInt64(list []int64, separator string) string {
//	return JoinInt64WithConvertMethod(list, separator, func(elem int64) string {
//		return FromInt64(elem)
//	})
//}
//
//func JoinInt64WithConvertMethod(list []int64, separator string, convert func(int64) string) string {
//	var buffer bytes.Buffer
//	for i, elem := range list {
//		if i != 0 {
//			buffer.WriteString(separator)
//		}
//		buffer.WriteString(convert(elem))
//	}
//	return buffer.String()
//}
