package method

type FormatFuc func(str string) (format string)

func Format(fuc FormatFuc, str string) {
	fuc(str)
}
